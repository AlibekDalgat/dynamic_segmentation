package service

import (
	"errors"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
)

type ServSegment struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *ServSegment {
	return &ServSegment{repo}
}

func (s *ServSegment) CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error) {
	if input.Name == "" {
		return 0, errors.New("Отсутствует имя сегмента")
	}
	return s.repo.CreateSegment(input)
}

func (s *ServSegment) DeleteSegment(input dynamic_segmentation.SegmentInfo) error {
	if input.Name == "" {
		return errors.New("Отсутствует имя сегмента")
	}
	return s.repo.DeleteSegment(input)
}
