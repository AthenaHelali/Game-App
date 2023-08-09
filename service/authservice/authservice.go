package authservice

import (
	"fmt"
	"game-app/entity"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Service struct {
	signKey               string
	accessSubject         string
	refreshSubject        string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
}

func New(signKey, accessSubject, refreshSubject string, accessExpirationTime, refreshExpirationTime time.Duration) Service {
	return Service{
		signKey:               signKey,
		accessSubject:         accessSubject,
		refreshSubject:        refreshSubject,
		accessExpirationTime:  accessExpirationTime,
		refreshExpirationTime: refreshExpirationTime,
	}

}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(s.accessSubject, user.ID, s.accessExpirationTime)
}
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(s.refreshSubject, user.ID, s.refreshExpirationTime)

}
func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
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
	tokenString, err := accessToken.SignedString([]byte(s.signKey))
	return tokenString, err
}
