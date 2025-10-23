package main

import (
	"auctionhouse/internal/database"
	"auctionhouse/internal/handlers"
	"auctionhouse/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	app := fiber.New()

	//middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("api/auth/register", handlers.Register)
	app.Post("api/auth/login", handlers.Login)

	//Protected routes
	protected := app.Group("/api", middleware.JWTMiddleware())
	protected.Get("/auctions", handlers.GetAuctions)
	protected.Get("/auctions/open", handlers.GetOpenAuctions)

	log.Fatal(app.Listen(":3000"))

}
