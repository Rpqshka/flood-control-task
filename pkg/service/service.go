package service

import (
	"context"
	"task/pkg/repository"
)

type Checker interface {
	SetParam(ctx context.Context) error
	GetParam() (int64, int64, error)
}

type Service struct {
	*FloodService
	Checker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		FloodService: NewFloodService(*repos.FloodMongo),
		Checker:      NewFloodService(*repos.FloodMongo),
	}
}
