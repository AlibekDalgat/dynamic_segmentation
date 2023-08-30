package service

import "github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"

type BackgroundService struct {
	repo repository.Background
}

func NewBackgroundService(repo repository.Background) *BackgroundService {
	return &BackgroundService{repo}
}

func (b *BackgroundService) DeleteExpirated() error {
	return b.repo.DeleteExpirated()
}
