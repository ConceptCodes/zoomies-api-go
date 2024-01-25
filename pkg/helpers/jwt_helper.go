package helpers

import (
	"fmt"
	"log"
	"time"
	"zoomies-api-go/pkg/config"
	"zoomies-api-go/pkg/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint, email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.AppConfig.Jwt.ExpiresIn)).Unix(),
	})

	return claims.SignedString([]byte("secret"))
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*models.Jwt); ok {
		fmt.Println(claims.ID, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

	return token, err
}
