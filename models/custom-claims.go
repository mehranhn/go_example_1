// Package models
package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
    UserID string `json:"user_id"`
    Phone  string `json:"phone"`
    jwt.RegisteredClaims
}
