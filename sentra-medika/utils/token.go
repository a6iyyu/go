// Package utils provides utility functions for various common tasks.
package utils

import (
	"errors"
	"os"
	t "time"

	"github.com/golang-jwt/jwt/v5"
	u "github.com/google/uuid"
)

// Claims struct untuk JWT yang bertujuan untuk menyimpan informasi pengguna
type Claims struct {
	UserID u.UUID `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenWithExpiry(user u.UUID, role string, duration t.Duration) (string, error) {
	// Menentukan waktu kedaluwarsa token berdasarkan durasi yang diberikan
	exp := t.Now().Add(duration)

	// Membuat klaim token
	claims := &Claims{
		UserID: user,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(t.Now()),
			Issuer:    "sentra-medika",
		},
	}

	// Membuat token JWT dengan algoritma penandatanganan HMAC SHA256 dan klaim yang telah dibuat
	// di atas, lalu menandatanganinya dengan kunci rahasia dari variabel lingkungan.
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*Claims, error) {
	// Parsing dan validasi token JWT
	claims := &Claims{}

	// Memparsing token dengan klaim yang diberikan
	t, err := jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (any, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Memeriksa kesalahan selama parsing
	if err != nil {
		return nil, err
	}

	// Memeriksa validitas token
	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	// Mengembalikan klaim token jika valid
	return claims, nil
}