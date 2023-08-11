package user

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
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

type UserInfo struct {
	ID          uint      `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
}
type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

func New(authGenerator AuthGenerator, repo repository) *Service {
	return &Service{auth: authGenerator, repo: repo}
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
	var resp RegisterResponse
	resp.User.ID = createdUser.ID
	resp.User.Name = createdUser.Name
	resp.User.PhoneNumber = createdUser.PhoneNumber

	//return created user
	return resp, nil
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token Tokens   `json:"token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it's better to have separate methods for checking user existence and getting user by phone number
	// check the existence of phone number in repository
	//get the user by phone number

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username is not correct")
	}

	//compare user.Password with req.Password

	if hErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); hErr != nil {
		return LoginResponse{}, fmt.Errorf("password is not correct")

	}
	accessToken, tErr := s.auth.CreateAccessToken(user)
	if tErr != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}
	refreshToken, tErr := s.auth.CreateRefreshToken(user)
	if tErr != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}

	response := LoginResponse{
		User: UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name},
		Token: Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

// Profile all request inputs for service should be sanitized.
func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return ProfileResponse{Name: user.Name}, err

}
