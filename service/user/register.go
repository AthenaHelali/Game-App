package user

import (
	"fmt"
	"game-app/dto"
	"game-app/entity"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	//TODO - we should verify phone number by verification code

	pass := []byte(req.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	//create new user in storage
	createdUser, err := s.repo.RegisterUser(entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(hashedPass),
	})

	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	var resp dto.RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber

	//return created user
	return resp, nil
}
