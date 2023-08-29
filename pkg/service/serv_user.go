package service

import (
	"errors"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
)

type ServUser struct {
	repo repository.User
}

func NewUserService(repo repository.User) *ServUser {
	return &ServUser{repo}
}

func (s *ServUser) AddToSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	if len(input.AddToSegments) == 0 {
		return nil
	}
	var cleanSegmentsInfo []dynamic_segmentation.SegmentInfo
	errorList := make([]error, 0)
	for _, segmentInfo := range input.AddToSegments {
		if segmentInfo.Name != "" {
			cleanSegmentsInfo = append(cleanSegmentsInfo, segmentInfo)
		} else {
			errorList = append(errorList, errors.New("Отсутствует имя сегмента"))
		}
	}
	input.AddToSegments = cleanSegmentsInfo
	errorList = append(errorList, s.repo.AddToSegments(input)...)
	return errorList
}

func (s *ServUser) DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	if len(input.DeleteFromSegments) == 0 {
		return nil
	}
	var cleanSegmentsInfo []dynamic_segmentation.SegmentInfo
	errorList := make([]error, 0)
	for _, segmentInfo := range input.DeleteFromSegments {
		if segmentInfo.Name != "" {
			cleanSegmentsInfo = append(cleanSegmentsInfo, segmentInfo)
		} else {
			errorList = append(errorList, errors.New("Отсутствует имя сегмента"))
		}
	}
	input.DeleteFromSegments = cleanSegmentsInfo
	errorList = append(errorList, s.repo.DeleteFromSegments(input)...)
	return errorList
}

func (s *ServUser) GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error) {
	return s.repo.GetActiveSegments(id)
}
