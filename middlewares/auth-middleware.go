// Package middlewares
package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mehranhn/go_example_1/models"
)

type JWTConfig struct {
	SecretKey       string
	SigningMethod   string
	TokenExpiration time.Duration
	ContextKey      string
}

func DefaultJWTConfig(secretKey string) JWTConfig {
	return JWTConfig{
		SecretKey:       secretKey,
		SigningMethod:   "HS256",
		TokenExpiration: 24 * time.Hour,
		ContextKey:      "user",
	}
}

// i could implement a authentication using a access token and refresh token but this is a simple test so i'm going to skip it
func JWTAuthMiddleware(config JWTConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		parts := strings.Split(authHeader, " ")
		var tokenString string
		if len(parts) == 1 {
			tokenString = parts[0]
		} else if len(parts) == 2 && parts[0] != "Bearer" {
			tokenString = parts[1]
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)

		}

		token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "invalid_token",
				"message": "Invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			c.Locals(config.ContextKey, claims)
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "invalid_token",
			"message": "Invalid token claims",
		})
	}
}
