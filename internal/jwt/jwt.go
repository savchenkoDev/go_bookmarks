package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

const JWT_EXPIRATION_HOURS = 1

func GenerateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": id,
		"exp": time.Now().Add(time.Hour * time.Duration(JWT_EXPIRATION_HOURS)).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println("Failed to parse token:", err)
		return 0, err
	}
	if !token.Valid {
		log.Println("Invalid token")
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Invalid token claims")
		return 0, errors.New("invalid token claims")
	}
	return int64(claims["uid"].(float64)), nil
}
