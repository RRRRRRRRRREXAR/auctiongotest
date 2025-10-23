package handlers

import (
	"auctionhouse/internal/database"
	"auctionhouse/internal/models"
	"auctionhouse/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAuctions(c *fiber.Ctx) error {

	auctions := []fiber.Map{
		{"id": 1, "title": "Antique Vase", "price": 150.00},
		{"id": 2, "title": "Vintage Car", "price": 20000.00},
	}

	return c.JSON(auctions)
}

func GetOpenAuctions(c *fiber.Ctx) error {
	repo := repository.NewAuctionRepository(database.DB)
	auctions, err := repo.GetOpenAuctions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch auctions"})
	}

	return c.JSON(auctions)
}

func CreateAuction(c *fiber.Ctx, auctionModel models.Auction) error {
	repo := repository.NewAuctionRepository(database.DB)
	createdAuction, err := repo.CreateAuction(auctionModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create auction"})
	}
	return c.JSON(createdAuction)
}
