package service

import (
	"encoding/csv"
	"errors"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

func (u *UserService) AddToSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	if len(input.AddToSegments) == 0 {
		return nil
	}
	var cleanSegmentsInfo []dynamic_segmentation.SegmentInfo
	errorList := make([]error, 0)
	for _, segmentInfo := range input.AddToSegments {
		if segmentInfo.Name == "" {
			errorList = append(errorList, errors.New("Отсутствует имя сегмента"))
		} else if !segmentInfo.Ttl.IsZero() && segmentInfo.Ttl.Before(time.Now()) {
			errorList = append(errorList, errors.New("Заданый TTL меньше актуального времени"))
		} else {
			cleanSegmentsInfo = append(cleanSegmentsInfo, segmentInfo)
		}
	}
	input.AddToSegments = cleanSegmentsInfo
	errorList = append(errorList, u.repo.AddToSegments(input)...)
	return errorList
}

func (u *UserService) DeleteFromSegments(input dynamic_segmentation.UserUpdatesInfo) []error {
	if len(input.DeleteFromSegments) == 0 {
		return nil
	}
	var cleanSegmentsInfo []dynamic_segmentation.SegmentInfo
	errorList := make([]error, 0)
	for _, segmentInfo := range input.DeleteFromSegments {
		if segmentInfo.Name != "" {
			cleanSegmentsInfo = append(cleanSegmentsInfo, segmentInfo)
		} else {
			errorList = append(errorList, errors.New("Отсутствует имя сегмента"))
		}
	}
	input.DeleteFromSegments = cleanSegmentsInfo
	errorList = append(errorList, u.repo.DeleteFromSegments(input)...)
	return errorList
}

func (u *UserService) GetActiveSegments(id int) ([]dynamic_segmentation.SegmentInfo, error) {
	return u.repo.GetActiveSegments(id)
}

func (u *UserService) GetReport(input dynamic_segmentation.DateInfo) (*os.File, error) {
	if input.Month > 12 || input.Month < 1 {
		return nil, errors.New("Некорректный ввод месяца")
	}
	rowsDelAdd, rowsDelDel, rowsAct, err := u.repo.GetReport(input)
	if err != nil {
		return nil, err
	}
	defer rowsDelAdd.Close()
	defer rowsDelDel.Close()
	defer rowsAct.Close()

	file, err := os.Create("report.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"user_id", "segment_name", "operation", "date"}
	err = writer.Write(headers)
	if err != nil {
		return nil, err
	}
	operationAdding := "добавление"
	operationDeletion := "удаление"
	err = writeToFile(writer, rowsDelDel, operationDeletion)
	if err != nil {
		return nil, err
	}
	err = writeToFile(writer, rowsDelAdd, operationAdding)
	if err != nil {
		return nil, err
	}
	err = writeToFile(writer, rowsAct, operationAdding)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeToFile(writer *csv.Writer, rows *sqlx.Rows, operation string) error {
	for rows.Next() {
		var row dynamic_segmentation.ReportInfo
		row.Operation = operation
		err := rows.StructScan(&row)
		if err != nil {
			return err
		}
		formattedTime := row.Date.Format("2006-01-02 15:04:05")
		recored := []string{strconv.Itoa(row.UserId), row.SegmentName, operation, formattedTime}
		err = writer.Write(recored)
		if err != nil {
			return err
		}
	}
	return nil
}
