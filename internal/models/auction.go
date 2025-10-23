package models

type Auction struct {
	Id          int     `json:"id"`
	ItemName    string  `json:"item_name"`
	StartingBid float64 `json:"starting_bid"`
	Description string  `json:"description"`
	EndTime     string  `json:"end_time"`
	CreatedAt   string  `json:"created_at"`
}
