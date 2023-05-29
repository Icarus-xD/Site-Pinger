package service

import (
	"github.com/Icarus-xD/SitePinger/internal/model"
	"github.com/google/uuid"
)

type PingRepo interface {
	GetById(id uuid.UUID) (model.SitePing, error)
	GetMin() ([]model.SitePing, error)
	GetMax() ([]model.SitePing, error)
}

type PingService struct {
	repo PingRepo
}

func NewPingService(repo PingRepo) *PingService {
	return &PingService{
		repo: repo,
	}
}

func (s *PingService) GetById(id uuid.UUID) (model.SitePing, error) {
	return s.repo.GetById(id)
}

func (s *PingService) GetMin() ([]model.SitePing, error) {
	return s.repo.GetMin()
}

func (s *PingService) GetMax() ([]model.SitePing, error) {
	return s.repo.GetMax()
}