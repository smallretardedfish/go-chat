package auth_handlers

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(id int64) (string, error) {
	expTime := time.Now().Add(30 * 24 * time.Hour).Unix() // 30 days just for local testing
	claims := Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
