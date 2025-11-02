package main

import (
	"auctionhouse/internal/database"
	"auctionhouse/internal/handlers"
	"auctionhouse/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "auctionhouse/docs" // Import generated docs
)

// @title Auction House API
// @version 1.0
// @description API for auction house application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@auctionhouse.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	app := fiber.New()

	//middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("api/auth/register", handlers.Register)
	app.Post("api/auth/login", handlers.Login)

	//Protected routes
	protected := app.Group("/api", middleware.JWTMiddleware())
	protected.Get("/auctions/open", handlers.GetOpenAuctions)
	protected.Get("/auctions/:id", handlers.GetAuctionDetails)
	protected.Post("/auctions", handlers.CreateAuction)
	protected.Get("/auctions/:id/bids", handlers.GetBidsByAuctionId)

	protected.Post("/bids", handlers.PlaceBid)

	log.Fatal(app.Listen(":3000"))

}
