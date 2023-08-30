package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db}
}

func (s *SegmentPostgres) CreateSegment(input dynamic_segmentation.SegmentInfo) (int, error) {
	var id, count int
	query := fmt.Sprintf("SELECT COUNT(name) FROM %s WHERE name = $1", segmentsTable)
	err := s.db.QueryRow(query, input.Name).Scan(&count)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("Сегмент с таким именем уже существует.")
	}
	query = fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", segmentsTable)
	err = s.db.QueryRow(query, input.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SegmentPostgres) DeleteSegment(input dynamic_segmentation.SegmentInfo) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", segmentsTable)
	_, err := s.db.Exec(query, input.Name)
	return err
}

func (s *SegmentPostgres) CreateSegmentWihtPercent(percent int, input dynamic_segmentation.SegmentInfo) (int, error) {
	id, err := s.CreateSegment(input)
	var users []dynamic_segmentation.UserInfo
	percentStr := strconv.FormatFloat(float64(percent)/100, 'f', -1, 32)
	query := fmt.Sprintf("SELECT user_id FROM (SELECT DISTINCT user_id FROM %s) t ORDER BY RANDOM()"+
		" LIMIT (SELECT COUNT(DISTINCT user_id)* $1 from %s)", usersInSegmentsTable, usersInSegmentsTable)
	err = s.db.Select(&users, query, percentStr)
	if err != nil {
		return 0, err
	}
	for _, user := range users {
		err = s.AddOneToSement(user.UserId, input.Name)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
