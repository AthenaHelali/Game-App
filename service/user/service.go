package user

import (
	"game-app/entity"
)

type repository interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(UserID uint) (entity.User, error)
}
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo repository
}

func New(authGenerator AuthGenerator, repo repository) *Service {
	return &Service{auth: authGenerator, repo: repo}
}
