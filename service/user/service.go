package user

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}

type Service struct {
	repo repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//TODO - we should verify phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	//check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
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

	//TODO - check the password with regex pattern
	//validate password
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password length should be greater than 8")
	}

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
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	//return created user
	return RegisterResponse{createdUser}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO it's better to have separate methods for checking user existence and getting user by phone number
	// check the existence of phone number in repository
	//get the user by phone number

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	//compare user.Password with req.Password

	if hErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); hErr != nil {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")

	}
	return LoginResponse{}, err
}
