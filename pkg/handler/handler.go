package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}
