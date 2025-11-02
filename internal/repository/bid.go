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
		bid.AuctionId, bid.UserId, bid.Amount)
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

func (r *BidRepository) FindBidByAuctionAndAmount(auctionID int, amount float64) (*models.Bid, error) {
	var bid models.Bid
	err := r.DB.QueryRow("SELECT id, auction_id, user_id, amount FROM bids WHERE auction_id = ? AND amount = ?",
		auctionID, amount).Scan(&bid.Id, &bid.AuctionId, &bid.UserId, &bid.Amount)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &bid, nil
}

func (bidRepository BidRepository) GetBidsByAuctionId(auctionID int) ([]models.Bid, error) {
	rows, err := bidRepository.DB.Query("SELECT id, auction_id, user_id, amount FROM bids WHERE auction_id = ?", auctionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bids []models.Bid
	for rows.Next() {
		var bid models.Bid
		if err := rows.Scan(&bid.Id, &bid.AuctionId, &bid.UserId, &bid.Amount); err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}
