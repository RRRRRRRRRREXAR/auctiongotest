package middleware

import (
	"auctionhouse/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}
		tokenString := authHeader[len("Bearer "):]
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		c.Locals("userId", claims.UserId)
		c.Locals("email", claims.Email)
		return c.Next()
	}
}
