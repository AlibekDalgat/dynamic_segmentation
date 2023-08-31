package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db}
}

func (u *UserPostgres) AddToSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	errorList := make([]error, 0)
	for i := 0; i < len(input.AddToSegments); i++ {
		var count int
		query := fmt.Sprintf("SELECT COUNT(*) FROM %s "+
			"WHERE user_id = $1 AND segment_name = $2", usersInSegmentsTable)
		if err := u.db.QueryRow(query, input.UserId, input.AddToSegments[i].Name).Scan(&count); err != nil {
			errorList = append(errorList, err)
			continue
		}
		if count > 0 {
			continue
		}
		query = fmt.Sprintf("SELECT COUNT(name) FROM %s WHERE name = $1", segmentsTable)
		if err := u.db.QueryRow(query, input.AddToSegments[i].Name).Scan(&count); err != nil {
			errorList = append(errorList, err)
			continue
		}
		if count == 0 {
			errorList = append(errorList, errors.New(fmt.Sprintf("Сегмента %s несуществует",
				input.AddToSegments[i].Name)))
			continue
		}
		if !input.AddToSegments[i].Ttl.IsZero() {
			query = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time, ttl) values ($1, $2, $3, $4)",
				usersInSegmentsTable)
			if _, err := u.db.Exec(query, input.UserId, input.AddToSegments[i].Name, time.Now().In(loc), input.AddToSegments[i].Ttl); err != nil {
				errorList = append(errorList, err)
			}
		} else {
			query = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time) values ($1, $2, $3)",
				usersInSegmentsTable)
			if _, err := u.db.Exec(query, input.UserId, input.AddToSegments[i].Name, time.Now().In(loc)); err != nil {
				errorList = append(errorList, err)
			}
		}
	}
	return errorList
}

func (u *UserPostgres) DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	errorList := make([]error, 0)
	for i := 0; i < len(input.DeleteFromSegments); i++ {
		tx, err := u.db.Begin()
		if err != nil {
			errorList = append(errorList, err)
			return errorList
		}
		var adding_time, now_time time.Time
		query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_name = $2 "+
			"RETURNING adding_time", usersInSegmentsTable)
		if err = tx.QueryRow(query, input.UserId, input.DeleteFromSegments[i].Name).Scan(&adding_time); err != nil {
			tx.Rollback()
			errorList = append(errorList, err)
			continue
		}
		now_time = time.Now().In(loc)
		query = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time, deletion_time) values ($1, $2, $3, $4)",
			deletedUsersFromSegments)
		if _, err = tx.Exec(query, input.UserId, input.DeleteFromSegments[i].Name, adding_time, now_time); err != nil {
			tx.Rollback()
			errorList = append(errorList, err)
			continue
		}
		tx.Commit()
	}
	return errorList
}

func (u *UserPostgres) GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error) {
	var segmnets []dynamic_segmentation.SegmentInfo
	query := fmt.Sprintf("SELECT segment_name FROM %s WHERE user_id = $1", usersInSegmentsTable)
	err := u.db.Select(&segmnets, query, id)
	return segmnets, err
}

func (u *UserPostgres) GetReport(input dynamic_segmentation.DateInfo) (*sqlx.Rows, *sqlx.Rows, *sqlx.Rows, error) {
	//var rows []dynamic_segmentation.ReportInfo
	query := fmt.Sprintf("SELECT user_id, segment_name, adding_time as date FROM %s "+
		"WHERE DATE_PART('year', adding_time) = $1 AND DATE_PART('month', adding_time) = $2", deletedUsersFromSegments)
	rowsDelAdd, err := u.db.Queryx(query, input.Year, input.Month)
	if err != nil {
		return nil, nil, nil, err
	}
	query = fmt.Sprintf("SELECT user_id, segment_name, deletion_time as date FROM %s "+
		"WHERE DATE_PART('year', deletion_time) = $1 AND DATE_PART('month', deletion_time) = $2", deletedUsersFromSegments)
	rowsDelDel, err := u.db.Queryx(query, input.Year, input.Month)
	if err != nil {
		return nil, nil, nil, err
	}
	query = fmt.Sprintf("SELECT user_id, segment_name, adding_time as date FROM %s "+
		"WHERE DATE_PART('year', adding_time) = $1 AND DATE_PART('month', adding_time) = $2", usersInSegmentsTable)
	rowsAct, err := u.db.Queryx(query, input.Year, input.Month)
	if err != nil {
		return nil, nil, nil, err
	}
	return rowsDelAdd, rowsDelDel, rowsAct, nil
}

func (s *SegmentPostgres) AddOneToSement(userId int, segmentName string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time) values ($1, $2, $3)",
		usersInSegmentsTable)
	_, err := s.db.Exec(query, userId, segmentName, time.Now().In(loc))
	if err != nil {
		return err
	}
	return nil
}
