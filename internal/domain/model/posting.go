package model

import "time"

type PostingStatus uint8

const (
	PostingStatus_CREATED    PostingStatus = 1
	PostingStatus_BUILDING   PostingStatus = 2
	PostingStatus_BUILT      PostingStatus = 3
	PostingStatus_DELIVERING PostingStatus = 4
	PostingStatus_DELIVERED  PostingStatus = 5
)

type Posting struct {
	ID            int64
	CartID        int64
	PositionsIDs  []int64
	PostingStatus PostingStatus
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
