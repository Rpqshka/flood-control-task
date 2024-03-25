package service

import (
	"task/pkg/repository"
)

type Service struct {
	*FloodService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		FloodService: NewFloodService(*repos.FloodMongo),
	}
}
