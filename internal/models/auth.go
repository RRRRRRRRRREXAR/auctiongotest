package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
