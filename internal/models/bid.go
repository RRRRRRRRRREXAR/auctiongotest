package models

type Bid struct {
	Id        int     `json:"id"`
	Amount    float64 `json:"amount"`
	UserId    int     `json:"user_id"`
	AuctionId int     `json:"auction_id"`
}
