package authservice

import (
	"fmt"
	"game-app/entity"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Config struct {
	SignKey               string
	AccessSubject         string
	RefreshSubject        string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
}
type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}

}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(s.config.AccessSubject, user.ID, s.config.AccessExpirationTime)
}
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(s.config.RefreshSubject, user.ID, s.config.RefreshExpirationTime)

}
func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("userID: %v,ExpiresAt: %v\n ", claims.UserID, claims.RegisteredClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}
}
func (s Service) createToken(subject string, userID uint, expiresAtDuration time.Duration) (string, error) {
	// TODO -   replace with rsa 256 RS256

	claims := &Claims{
		Subject: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAtDuration)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	return tokenString, err
}
