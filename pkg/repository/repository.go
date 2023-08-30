package repository

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
)

type User interface {
	AddToSegments(input dynamic_segmentation.UserUpdatesInfo) []error
	DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error
	GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error)
	GetReport(input dynamic_segmentation.DateInfo) (*sqlx.Rows, *sqlx.Rows, *sqlx.Rows, error)
}

type Segment interface {
	CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error)
	DeleteSegment(input dynamic_segmentation.SegmentInfo) error
	CreateSegmentWihtPercent(percent int, input dynamic_segmentation.SegmentInfo) (int, error)
}

type Background interface {
	DeleteExpirated() error
}

type Repository struct {
	User
	Segment
	Background
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:       NewUserPostgres(db),
		Segment:    NewSegmentPostgres(db),
		Background: NewBackgroundPostgres(db),
	}
}
