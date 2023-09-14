package user

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"request": req})
	}
	return param.ProfileResponse{Name: user.Name}, err

}
