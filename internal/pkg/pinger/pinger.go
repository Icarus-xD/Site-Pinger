package pinger

import (
	"fmt"
	"log"
	"time"

	"github.com/Icarus-xD/SitePinger/internal/model"
	"github.com/go-ping/ping"
	"github.com/google/uuid"
)

type pingsRepo interface {
	GetAll() ([]model.SitePing, error)
	UpdateById(id uuid.UUID, newPing int64)
	InvalidateMinMaxCache()
}

type pingUpdate struct {
	id uuid.UUID
	ping int64
}

func RunPings(repo pingsRepo) {
	sitesPing, err := repo.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	sitesCount := len(sitesPing)

	updates := make(chan pingUpdate, sitesCount)
	defer close(updates)

	for _, sp := range sitesPing {
		go pingSite(sp.ID, sp.Site, updates)
	}

	invalidationCounter := sitesCount
	for update := range updates {
		repo.UpdateById(update.id, update.ping)
		invalidationCounter--

		if (invalidationCounter == 0) {
			repo.InvalidateMinMaxCache()

			invalidationCounter = sitesCount
		}
	}
}

func pingSite(id uuid.UUID, site string, updates chan <-pingUpdate) {

	pinger, err := ping.NewPinger(site)
	if err != nil {
		log.Println(err)
	}
	
	pinger.SetPrivileged(true)
	pinger.Size = 64
	pinger.Count = 4
	pinger.Timeout = time.Second * 5

	for {
		err = pinger.Run()
		if err != nil {
			log.Println(err)
		}

		stats := pinger.Statistics()

		updates <- pingUpdate{
			id: id,
			ping: stats.AvgRtt.Milliseconds(),
		}

		fmt.Printf("Ping to %s: %v\n", site, stats.AvgRtt.Milliseconds())

		time.Sleep(time.Minute)
	}
	
}