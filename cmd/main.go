package main

import (
	"log"

	"github.com/Icarus-xD/SitePinger/internal/config"
	"github.com/Icarus-xD/SitePinger/internal/database"
	"github.com/Icarus-xD/SitePinger/internal/pkg/pinger"
	"github.com/Icarus-xD/SitePinger/internal/repository"
	"github.com/Icarus-xD/SitePinger/internal/service"
	"github.com/Icarus-xD/SitePinger/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	mainDB, cacheDB, statsDB := database.Init(config.PostgresUrl, config.RedisAddr, config.ClickhouseAddr)

	sqlDB, err := mainDB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
	defer cacheDB.Close()
	defer statsDB.Close()

	router := gin.Default()

	pingRepo := repository.NewPingRepo(mainDB, cacheDB)
	pingService := service.NewPingService(pingRepo)

	statsRepo := repository.NewStatsRepo(statsDB)
	statsService := service.NewStatsServvice(statsRepo)

	handler := rest.NewHandler(pingService, statsService)
	handler.InitRouter(router)

	go pinger.RunPings(pingRepo)

	if err := router.Run(config.AppPort); err != nil {
		log.Fatalln(err)
	}
}