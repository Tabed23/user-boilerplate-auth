package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GenrateToken(role, email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
