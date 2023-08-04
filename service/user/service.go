package user

import (
	"fmt"
	"gam-app/entity"
	"gam-app/pkg/phonenumber"
)

type repository interface {
	isPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (entity.User, error)
}

type Service struct {
	repo repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//TODO - we should verify phone number by verification code
	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	//check uniqueness of phone number
	if isUnique, err := s.repo.isPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}
	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be greater than 3")
	}

	//create new user in storage
	createdUser, err := s.repo.RegisterUser(entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.PhoneNumber,
	})

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	//return created user
	return RegisterResponse{createdUser}, nil

}
