package repository

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
)

type User interface {
}

type Segment interface {
	CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error)
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
