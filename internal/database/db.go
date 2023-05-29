package database

import (
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(mainURL, cacheAddr, statsAddr string) (*gorm.DB, *redis.Client, driver.Conn) {
	main := InitMain(mainURL)
	cache := InitCache(cacheAddr)
	stats := InitStats(statsAddr)

	return main, cache, stats
}

func InitMain(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func InitCache(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})

	return client
}

func InitStats(addr string) driver.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}