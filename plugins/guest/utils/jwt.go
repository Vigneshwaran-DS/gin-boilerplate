package utils

import (
	"errors"
	"time"

	"gin-boilerplate/config"
	"github.com/golang-jwt/jwt/v5"
)

// GuestClaims represents JWT claims for guest users
type GuestClaims struct {
	GuestID  uint   `json:"guest_id"`
	GuestUID string `json:"guest_uid"`
	UserType string `json:"user_type"` // Always "guest"
	jwt.RegisteredClaims
}

// GenerateGuestToken generates JWT Token for guest users
func GenerateGuestToken(guestID uint, guestUID string) (string, error) {
	expireTime := time.Duration(config.AppConfig.JWT.ExpireTime) * time.Hour
	claims := GuestClaims{
		GuestID:  guestID,
		GuestUID: guestUID,
		UserType: "guest",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWT.Secret))
}

// ParseGuestToken parses guest JWT Token
func ParseGuestToken(tokenString string) (*GuestClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &GuestClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*GuestClaims); ok && token.Valid {
		// Verify this is actually a guest token
		if claims.UserType != "guest" {
			return nil, errors.New("not a guest token")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
