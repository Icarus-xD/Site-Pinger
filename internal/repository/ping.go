package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Icarus-xD/SitePinger/internal/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PingRepo struct {
	db *gorm.DB
	cache *redis.Client
	ctx context.Context
}

func NewPingRepo(db *gorm.DB, cache *redis.Client) *PingRepo {
	return &PingRepo{
		db: db,
		cache: cache,
		ctx: context.Background(),
	}
}

func (r *PingRepo) GetAll() ([]model.SitePing, error) {
	var pings []model.SitePing

	if result := r.db.Find(&pings); result.Error != nil {
		return pings, result.Error
	}

	return pings, nil
}

func (r *PingRepo) GetById(id uuid.UUID) (model.SitePing, error) {
	var ping model.SitePing

	cached, err := r.cache.Get(r.ctx, id.String()).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cached), &ping)

		if err == nil {
			return ping, nil
		}
	}

	if err := r.db.First(&ping, id).Error; err != nil {
		return ping, err
	}
	
	marshaledPing, err := json.Marshal(ping)
	if err == nil {
		r.cache.SetNX(r.ctx, id.String(), marshaledPing, time.Hour * 24)
	}

	return ping, nil
}

func (r *PingRepo) GetMin() ([]model.SitePing, error) {
	var pings []model.SitePing
	var minPing model.SitePing

	if err := r.db.Order("ping").First(&minPing).Error; err != nil {
		return pings, err
	}

	if err := r.db.Where("ping = ?", minPing.Ping).Find(&pings).Error; err != nil {
		return pings, err
	}

	return pings, nil
}

func (r *PingRepo) GetMax() ([]model.SitePing, error) {
	var pings []model.SitePing
	var maxPing model.SitePing

	if err := r.db.Order("ping desc").First(&maxPing).Error; err != nil {
		return pings, err
	}

	if err := r.db.Where("ping = ?", maxPing.Ping).Find(&pings).Error; err != nil {
		return pings, err
	}

	return pings, nil
}

func (r *PingRepo) UpdateById(id uuid.UUID, newPing int64) {	
	var ping model.SitePing

	if result := r.db.First(&ping, id); result.Error != nil {
		log.Fatalln(result.Error)
	}

	if ping.Ping == newPing { return }

	ping.Ping = newPing
	r.db.Save(&ping)

	cacheTTL, err := r.cache.TTL(r.ctx, id.String()).Result()
	if err != nil {
		log.Println("error getting TTL", err)
		cacheTTL = time.Hour * 24
	}

	if cacheTTL == -2 { 
		return 
	}

	marshaledPing, err := json.Marshal(ping)
	if err != nil {
		return
	}

	r.cache.SetXX(r.ctx, id.String(), marshaledPing, cacheTTL)
}

func (r *PingRepo) InvalidateMinMaxCache() {
	maxPings, err := r.GetMax()
	if err == nil {
		marshaledPings, err := json.Marshal(maxPings)

		if err == nil {
			r.cache.Set(r.ctx, "max", marshaledPings, time.Hour * 24)
		}
	}

	minPings, err := r.GetMin()
	if err == nil {
		marshaledPings, err := json.Marshal(minPings)

		if err == nil {
			r.cache.Set(r.ctx, "min", marshaledPings, time.Hour * 24)
		}
	}
}