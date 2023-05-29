package service

import "github.com/Icarus-xD/SitePinger/internal/model"

type StatsRepo interface {
	SaveStat(endpoint string) error
	GetEndpointsStats() ([]model.EndpointStats, error)
}

type StatsService struct {
	repo StatsRepo
}

func NewStatsServvice(repo StatsRepo) *StatsService {
	return &StatsService{
		repo: repo,
	}
}

func (s *StatsService) SaveStat(endpoint string) error {
	return s.repo.SaveStat(endpoint)
}

func (s *StatsService) GetEndpointsStats() ([]model.EndpointStats, error) {
	return s.repo.GetEndpointsStats()
}