package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
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
		if err := u.db.QueryRow(query, input.User_id, input.AddToSegments[i].Name).Scan(&count); err != nil {
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
			errorList = append(errorList, errors.New(fmt.Sprintf("Сегмента %s несуществует", input.AddToSegments[i].Name)))
			continue
		}
		query = fmt.Sprintf("INSERT INTO %s (user_id, segment_name) values ($1, $2)", usersInSegmentsTable)
		if _, err := u.db.Exec(query, input.User_id, input.AddToSegments[i].Name); err != nil {
			errorList = append(errorList, err)
		}
	}
	return errorList
}

func (u *UserPostgres) DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	errorList := make([]error, 0)
	for i := 0; i < len(input.DeleteFromSegments); i++ {
		query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_name = $2", usersInSegmentsTable)
		if _, err := u.db.Exec(query, input.User_id, input.DeleteFromSegments[i].Name); err != nil {
			errorList = append(errorList, err)
		}
	}
	return errorList
}

func (u *UserPostgres) GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error) {
	var segmnets []dynamic_segmentation.SegmentInfo
	query := fmt.Sprintf("SELECT segment_name FROM %s WHERE user_id = $1", usersInSegmentsTable)
	err := u.db.Select(&segmnets, query, id)
	return segmnets, err
}
