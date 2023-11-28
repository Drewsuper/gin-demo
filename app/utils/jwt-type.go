package utils

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}
