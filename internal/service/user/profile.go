package user

import (
	"context"
	"game-app/internal/param"
	"game-app/internal/pkg/richerror"
)

func (s Service) Profile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"
	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"request": req})
	}
	return param.ProfileResponse{Name: user.Name}, err

}
