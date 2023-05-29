package repository

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Icarus-xD/SitePinger/internal/model"
)

const (
	endpointsStatsQuery string = "SELECT endpoint, COUNT(*) AS request_count FROM endpoints_stats GROUP BY endpoint"
)

type StatsRepo struct {
	db driver.Conn
	ctx context.Context
}

func NewStatsRepo(db driver.Conn) *StatsRepo {
	return &StatsRepo{
		db: db,
		ctx: context.Background(),
	}
}

func (r *StatsRepo) SaveStat(endpoint string) error {
	return r.db.AsyncInsert(r.ctx, fmt.Sprintf("INSERT INTO endpoints_stats (endpoint) VALUES (%s)", endpoint), false)
}

func (r *StatsRepo) GetEndpointsStats() ([]model.EndpointStats, error) {
	var stats []model.EndpointStats
	
	rows, err := r.db.Query(r.ctx, endpointsStatsQuery)
	if err != nil {
		return stats, err
	}

	err = rows.ScanStruct(&stats)
	if err != nil {
		return stats, err
	}

	return stats, nil
}