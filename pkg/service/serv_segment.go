package service

import "github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"

type ServSegment struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *ServSegment {
	return &ServSegment{repo}
}
