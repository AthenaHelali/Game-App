package user

import (
	"game-app/dto"
	"game-app/pkg/richerror"
)

func (s Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"request": req})
	}
	return dto.ProfileResponse{Name: user.Name}, err

}
