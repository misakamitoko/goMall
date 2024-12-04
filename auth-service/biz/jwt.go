package biz

import (
	"fmt"
	"log"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GetJwtToken(secretKey string, iat, seconds int64, payload any) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseJwt(tokenString string, secretKey string) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	log.Fatalf("invalid token")
	return false
}