package repository

import (
	"fmt"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type BackgroundPostgres struct {
	db *sqlx.DB
}

func NewBackgroundPostgres(db *sqlx.DB) *BackgroundPostgres {
	return &BackgroundPostgres{db}
}

func (b *BackgroundPostgres) DeleteExpirated() error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ttl < $1 RETURNING user_id, segment_name, adding_time, $1 as deletion_time",
		usersInSegmentsTable)
	rows, err := b.db.Queryx(query, time.Now().In(loc))
	if err != nil {
		return err
	}
	for rows.Next() {
		var row dynamic_segmentation.AutoDeletionInfo
		err = rows.StructScan(&row)
		if err != nil {
			return err
		}
		query = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, adding_time, deletion_time) values ($1, $2, $3, $4)",
			deletedUsersFromSegments)
		_, err = b.db.Exec(query, row.UserId, row.SegmentName, row.AddingTime, row.DeletionTime)
		if err != nil {
			return err
		}
		logrus.Infof("Срок нахождения пользователя %d в сегменте %s истёк", row.UserId, row.SegmentName)
	}
	return nil
}
