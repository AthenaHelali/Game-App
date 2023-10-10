package user

import (
	"context"
	"fmt"
	entity2 "game-app/internal/entity"
	"game-app/internal/param"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(ctx context.Context, req param.RegisterRequest) (param.RegisterResponse, error) {
	//TODO - we should verify phone number by verification code

	pass := []byte(req.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	//create new user in storage
	createdUser, err := s.repo.RegisterUser(ctx, entity2.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(hashedPass),
		Role:        entity2.UserRole,
	})

	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	var resp param.RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber

	//return created user
	return resp, nil
}
