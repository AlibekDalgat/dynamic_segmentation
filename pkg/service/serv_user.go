package service

import "github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"

type ServUser struct {
	repo repository.Segment
}

func NewUserService(repo repository.Segment) *ServUser {
	return &ServUser{repo}
}
