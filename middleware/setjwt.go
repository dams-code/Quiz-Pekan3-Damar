package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var setKodeJWT = []byte("damar123!")

type ClaimsDataBuku struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, username string) (string, error) {
	claims := ClaimsDataBuku{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(setKodeJWT)
}

func ValidasiJWT(tokenString string) (*ClaimsDataBuku, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsDataBuku{}, func(t *jwt.Token) (interface{}, error) {
		return setKodeJWT, nil
	})

	if err != nil {
		return nil, err
	}

	cekClaims, ok := token.Claims.(*ClaimsDataBuku)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid atau expired")
	}

	return cekClaims, nil
}
