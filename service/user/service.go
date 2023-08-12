package user

import (
	"fmt"
	"game-app/dto"
	"game-app/entity"
	"game-app/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
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

func New(authGenerator AuthGenerator, repo repository) *Service {
	return &Service{auth: authGenerator, repo: repo}
}

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

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User  dto.UserInfo `json:"user"`
	Token Tokens       `json:"token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it's better to have separate methods for checking user existence and getting user by phone number
	// check the existence of phone number in repository
	//get the user by phone number
	const op = "userservice.login"
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")
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
		User: dto.UserInfo{
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
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"request": req})
	}
	return ProfileResponse{Name: user.Name}, err

}
