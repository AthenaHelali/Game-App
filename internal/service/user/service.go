package user

import (
	"context"
	"game-app/internal/entity"
)

type repository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error)
	GetUserByID(ctx context.Context, UserID uint) (entity.User, error)
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
