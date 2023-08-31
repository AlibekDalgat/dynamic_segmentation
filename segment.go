package dynamic_segmentation

import "time"

type SegmentInfo struct {
	Name string    `json:"name" binding:"required" db:"segment_name"`
	Ttl  time.Time `json:"ttl"`
}
