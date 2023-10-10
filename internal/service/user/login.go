package user

import (
	"context"
	"fmt"
	param2 "game-app/internal/param"
	"game-app/internal/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(ctx context.Context, req param2.LoginRequest) (param2.LoginResponse, error) {
	// TODO - it's better to have separate methods for checking user existence and getting mysqluser by phone number
	// check the existence of phone number in repository
	//get the user by phone number
	const op = "userservice.login"
	user, err := s.repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return param2.LoginResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	//compare user.Password with req.Password

	if hErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); hErr != nil {
		return param2.LoginResponse{}, fmt.Errorf("password is not correct")

	}
	accessToken, tErr := s.auth.CreateAccessToken(user)
	if tErr != nil {
		return param2.LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}
	refreshToken, tErr := s.auth.CreateRefreshToken(user)
	if tErr != nil {
		return param2.LoginResponse{}, fmt.Errorf("unexpected error:%w", tErr)
	}

	response := param2.LoginResponse{
		User: param2.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name},
		Token: param2.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}
