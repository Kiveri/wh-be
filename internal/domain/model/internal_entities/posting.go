package internal_entities

type PostingStatus uint8

const (
	PostingStatus_CREATED    PostingStatus = 1
	PostingStatus_BUILDING   PostingStatus = 2
	PostingStatus_BUILT      PostingStatus = 3
	PostingStatus_DELIVERING PostingStatus = 4
	PostingStatus_DELIVERED  PostingStatus = 5
)

type Posting struct {
	ID           int64
	CartID       int64
	PositionsIDs []int64
	Status       PostingStatus
	IsActive     bool
}

func NewPosting(cartID int64) *Posting {
	return &Posting{
		CartID:   cartID,
		Status:   PostingStatus_CREATED,
		IsActive: true,
	}
}

func (p *Posting) SetCartAndPositionsIDs(cartID int64, positionsIDs []int64) {
	p.CartID = cartID
	p.PositionsIDs = positionsIDs
	p.Status = PostingStatus_BUILDING
}

func (p *Posting) ChangePostingStatus(status PostingStatus) {
	p.Status = status
}
