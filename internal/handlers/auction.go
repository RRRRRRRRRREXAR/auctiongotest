package handlers

import (
	"auctionhouse/internal/database"
	"auctionhouse/internal/models"
	"auctionhouse/internal/repository"

	"github.com/gofiber/fiber/v2"
)

// GetOpenAuctions godoc
// @Summary Get all open auctions
// @Description Retrieve all currently open auctions
// @Tags auctions
// @Produce json
// @Success 200 {array} models.Auction
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /auctions/open [get]
func GetOpenAuctions(c *fiber.Ctx) error {
	repo := repository.NewAuctionRepository(database.DB)
	auctions, err := repo.GetOpenAuctions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch auctions"})
	}

	return c.JSON(auctions)
}

// CreateAuction godoc
// @Summary Create a new auction
// @Description Create a new auction with the provided details
// @Tags auctions
// @Accept json
// @Produce json
// @Param auction body models.Auction true "Auction details"
// @Success 200 {object} models.Auction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /auctions [post]
func CreateAuction(c *fiber.Ctx) error {
	var auctionModel models.Auction
	if err := c.BodyParser(&auctionModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	repo := repository.NewAuctionRepository(database.DB)
	createdAuction, err := repo.CreateAuction(auctionModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create auction"})
	}
	return c.JSON(createdAuction)
}

// GetAuctionDetails godoc
// @Summary Get auction details
// @Description Retrieve details of a specific auction including all bids
// @Tags auctions
// @Produce json
// @Param id path int true "Auction ID"
// @Success 200 {object} map[string]interface{} "auction and bids"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /auctions/{id} [get]
func GetAuctionDetails(c *fiber.Ctx) error {
	auctionID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid auction ID"})
	}
	repo := repository.NewAuctionRepository(database.DB)
	bidRepo := repository.NewBidRepository(database.DB)
	auction, err := repo.GetAuctionById(auctionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch auction details"})
	}
	if auction == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Auction not found"})
	}
	bids, err := bidRepo.GetBidsByAuctionId(auctionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch bids"})
	}

	return c.JSON(fiber.Map{"auction": auction, "bids": bids})
}
