package models

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	Email string
	jwt.RegisteredClaims
}

type JwtConfig struct {
	Secret    string
	ExpiresIn int64
}
