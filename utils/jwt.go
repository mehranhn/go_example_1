package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mehranhn/go_example_1/models"
)

func GenerateJWTToken(secretKey string, userID, phone string, expiration time.Duration) (string, error) {
    claims := models.CustomClaims{
        UserID: userID,
        Phone:  phone,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Subject:   userID,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

func ExtractClaimsFromContext(c *fiber.Ctx) (*models.CustomClaims, error) {
    claims, ok := c.Locals("user").(*models.CustomClaims)
    if !ok {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "No user claims found")
    }
    return claims, nil
}
