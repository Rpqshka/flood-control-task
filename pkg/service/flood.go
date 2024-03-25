package service

import (
	"context"
	"task/pkg/repository"
)

type FloodService struct {
	repo repository.FloodMongo
}

func NewFloodService(repo repository.FloodMongo) *FloodService {
	return &FloodService{repo: repo}
}

func (s *FloodService) Check(ctx context.Context, userID int64) (bool, error) {
	return s.repo.Check(ctx, userID)
}

func (s *FloodService) SetParam(ctx context.Context) error {
	return s.repo.SetParam(ctx)
}

func (s *FloodService) GetParam() (int64, int64, error) {
	return s.repo.GetParam()
}
