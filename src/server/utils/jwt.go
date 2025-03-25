package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func getSecretKey() []byte {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		secretKey = []byte("tmp-secret-key")
	}
	return secretKey
}

func GenerateJWT(userID uint, email string) (string, error) {
	expirationTimestamp := time.Now().Add(24 * time.Hour)

	claims := &JWTClaims{
		UserID:         fmt.Sprintf("%d", userID),
		Email:          email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTimestamp.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := getSecretKey()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	secretKey := getSecretKey()

	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (any, error) {
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
