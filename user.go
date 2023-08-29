package dynamic_segmentation

type UserUpdatesInfo struct {
	User_id            int           `json:"user_id"`
	AddToSegments      []SegmentInfo `json:"add_to_segments"`
	DeleteFromSegments []SegmentInfo `json:"delete_from_segments"`
}
