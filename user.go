package dynamic_segmentation

import "time"

type UserUpdatesInfo struct {
	UserId             int           `json:"user_id"`
	AddToSegments      []SegmentInfo `json:"add_to_segments"`
	DeleteFromSegments []SegmentInfo `json:"delete_from_segments"`
}

type DateInfo struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type ReportInfo struct {
	UserId      int       `json:"user_id" db:"user_id"`
	SegmentName string    `json:"segment_name" db:"segment_name"`
	Operation   string    `json:"operation" db:"operation"`
	Date        time.Time `json:"date" db:"date"`
}

type AutoDeletionInfo struct {
	UserId       int       `db:"user_id"`
	SegmentName  string    `db:"segment_name"`
	AddingTime   time.Time `db:"adding_time"`
	DeletionTime time.Time `db:"deletion_time"`
}
