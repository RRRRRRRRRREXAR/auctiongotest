package repository

import (
	"auctionhouse/internal/models"
	"database/sql"
)

type BidRepository struct {
	DB *sql.DB
}

func NewBidRepository(db *sql.DB) *BidRepository {
	return &BidRepository{DB: db}
}

func (bidRepository BidRepository) CreateBid(bid models.Bid) (*models.Bid, error) {
	result, err := bidRepository.DB.Exec("INSERT INTO bids (auction_id, user_id, amount) VALUES (?, ?, ?)",
		bid.AuctionID, bid.UserID, bid.Amount)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	bid.Id = int(id)
	return &bid, nil
}
