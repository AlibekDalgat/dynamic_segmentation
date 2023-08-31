package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
	"time"
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
	tx, err := s.db.Begin()
	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time, deletion_time) SELECT user_id, segment_name, adding_time, $1 FROM %s WHERE segment_name = $2", deletedUsersFromSegments, usersInSegmentsTable)
	_, err = tx.Exec(query, time.Now().In(loc), input.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE name = $1", segmentsTable)
	_, err = s.db.Exec(query, input.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *SegmentPostgres) CreateSegmentWihtPercent(percent int, input dynamic_segmentation.SegmentInfo) (int, error) {
	id, err := s.CreateSegment(input)
	var users []dynamic_segmentation.UserInfo
	query := fmt.Sprintf("SELECT user_id FROM (SELECT DISTINCT user_id FROM %s) t ORDER BY RANDOM()"+
		" LIMIT (SELECT COUNT(DISTINCT user_id)*($1 /100.) from %s)", usersInSegmentsTable, usersInSegmentsTable)
	err = s.db.Select(&users, query, percent)
	if err != nil {
		return 0, err
	}
	for _, user := range users {
		err = s.AddOneToSement(user.UserId, input.Name)
		if err != nil {
			return 0, errors.New("Сегмент добавлен без пользователей. " + err.Error())
		}
	}
	return id, nil
}
