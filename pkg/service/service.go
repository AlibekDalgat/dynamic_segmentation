package service

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
)

type User interface {
	AddToSegments(input dynamic_segmentation.UserUpdatesInfo) []error
	DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error
	GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error)
}

type Segment interface {
	CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error)
	DeleteSegment(input dynamic_segmentation.SegmentInfo) error
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
