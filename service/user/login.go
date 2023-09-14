package user

import (
	"fmt"
	"game-app/dto"
	"game-app/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// TODO - it's better to have separate methods for checking user existence and getting user by phone number
	// check the existence of phone number in repository
	//get the user by phone number
	const op = "userservice.login"
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	//compare user.Password with req.Password

	if hErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); hErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("password is not correct")

	}
	accessToken, tErr := s.auth.CreateAccessToken(user)
	if tErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}
	refreshToken, tErr := s.auth.CreateRefreshToken(user)
	if tErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}

	response := dto.LoginResponse{
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name},
		Token: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}
