package repository

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
