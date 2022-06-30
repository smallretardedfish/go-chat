package auth_handlers

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(id int64, jwtKey string) (string, error) {
	expTime := time.Now().Add(30 * 24 * time.Hour).Unix() // 30 days for local testing
	claims := Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
