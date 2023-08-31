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
	return s.repo.CreateSegment(input)
}

func (s *SegmentService) DeleteSegment(input dynamic_segmentation.SegmentInfo) error {
	return s.repo.DeleteSegment(input)
}

func (s *SegmentService) CreateSegmentWihtPercent(percent int, input dynamic_segmentation.SegmentInfo) (int, error) {
	if percent <= 0 || percent > 100 {
		return 0, errors.New("Неправильно ввёдены проценты")
	}
	return s.repo.CreateSegmentWihtPercent(percent, input)
}
