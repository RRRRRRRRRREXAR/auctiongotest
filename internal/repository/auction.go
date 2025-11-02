package repository

import (
	"auctionhouse/internal/models"
	"database/sql"
)

type AuctionRepository struct {
	DB *sql.DB
}

func NewAuctionRepository(db *sql.DB) *AuctionRepository {
	return &AuctionRepository{DB: db}
}

func (auctionRepository AuctionRepository) GetOpenAuctions() ([]models.Auction, error) {
	rows, err := auctionRepository.DB.Query("SELECT id, item_name, starting_bid, description, end_time, created_at FROM auctions WHERE end_time > datetime('now')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auctions []models.Auction
	for rows.Next() {
		var auction models.Auction
		if err := rows.Scan(&auction.Id, &auction.ItemName, &auction.StartingBid, &auction.Description, &auction.EndTime, &auction.CreatedAt); err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}
	return auctions, nil
}

func (auctionRepository AuctionRepository) CreateAuction(auction models.Auction) (*models.Auction, error) {
	result, err := auctionRepository.DB.Exec("INSERT INTO auctions (item_name, starting_bid, description, end_time) VALUES (?, ?, ?, ?)",
		auction.ItemName, auction.StartingBid, auction.Description, auction.EndTime)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	auction.Id = int(id)
	return &auction, nil
}

func (auctionRepository AuctionRepository) GetAuctionById(auctionID int) (*models.Auction, error) {
	var auction models.Auction
	err := auctionRepository.DB.QueryRow("SELECT id, item_name, starting_bid, description, end_time, created_at FROM auctions WHERE id = ?",
		auctionID).Scan(&auction.Id, &auction.ItemName, &auction.StartingBid, &auction.Description, &auction.EndTime, &auction.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &auction, nil
}
