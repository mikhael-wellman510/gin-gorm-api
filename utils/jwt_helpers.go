package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userId string) (string, error) {
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	log.Println("Apa isi claims : ", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println("isi token : ", token)
	log.Println("isi jwtSecret : ", jwtSecret)
	return token.SignedString(jwtSecret)
}
