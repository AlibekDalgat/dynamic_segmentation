package service

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
)

type User interface {
}

type Segment interface {
	CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error)
}

type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repos),
		Segment: NewSegmentService(repos),
	}
}
