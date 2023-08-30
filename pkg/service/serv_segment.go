package service

import (
	"errors"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo}
}

func (s *SegmentService) CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error) {
	if input.Name == "" {
		return 0, errors.New("Отсутствует имя сегмента")
	}
	return s.repo.CreateSegment(input)
}

func (s *SegmentService) DeleteSegment(input dynamic_segmentation.SegmentInfo) error {
	if input.Name == "" {
		return errors.New("Отсутствует имя сегмента")
	}
	return s.repo.DeleteSegment(input)
}
