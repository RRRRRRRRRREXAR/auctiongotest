package handlers

import (
	"auctionhouse/internal/database"
	"auctionhouse/internal/models"
	"auctionhouse/internal/repository"

	"github.com/gofiber/fiber/v2"
)

// PlaceBid godoc
// @Summary Place a bid on an auction
// @Description Place a new bid on an existing auction
// @Tags bids
// @Accept json
// @Produce json
// @Param bid body models.Bid true "Bid details"
// @Success 200 {object} models.Bid
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string "Bid with same amount exists"
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /bids [post]
func PlaceBid(c *fiber.Ctx) error {
	var bidModel models.Bid
	if err := c.BodyParser(&bidModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	repo := repository.NewBidRepository(database.DB)
	existingBid, err := repo.FindBidByAuctionAndAmount(bidModel.AuctionId, bidModel.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check existing bids",
		})
	}
	if existingBid != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "A bid with the same amount already exists for this auction",
		})
	}

	placedBid, err := repo.CreateBid(bidModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to place bid",
		})
	}
	return c.JSON(placedBid)
}

// GetBidsByAuctionId godoc
// @Summary Get bids by auction ID
// @Description Retrieve all bids for a specific auction
// @Tags bids
// @Produce json
// @Param id path int true "Auction ID"
// @Success 200 {array} models.Bid
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /bids/{id} [get]
func GetBidsByAuctionId(c *fiber.Ctx) error {
	auctionID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid auction ID",
		})
	}
	repo := repository.NewBidRepository(database.DB)
	bids, err := repo.GetBidsByAuctionId(auctionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve bids",
		})
	}
	return c.JSON(bids)
}
