package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db}
}

func (r *SegmentPostgres) CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error) {
	var id, count int
	query := fmt.Sprintf("SELECT COUNT(name) FROM %s WHERE name = $1", segmentsTable)
	err := r.db.QueryRow(query, input.Name).Scan(&count)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("Сегмент с таким именем уже существует.")
	}
	query = fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", segmentsTable)
	err = r.db.QueryRow(query, input.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SegmentPostgres) DeleteSegment(input dynamic_segmentation.SegmentInfo) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", segmentsTable)
	_, err := r.db.Exec(query, input.Name)
	return err
}
