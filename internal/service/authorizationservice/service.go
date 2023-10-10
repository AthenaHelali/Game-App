package authorizationservice

import (
	"context"
	entity2 "game-app/internal/entity"
	"game-app/internal/pkg/richerror"
)

type Repository interface {
	GetUserPermissionTitles(userID uint, role entity2.Role) ([]entity2.PermissionTitle, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckAccess(ctx context.Context, userID uint, role entity2.Role, permissions ...entity2.PermissionTitle) (bool, error) {
	const op = "authorizationservice.CheckAccess"

	permissionTitles, err := s.repo.GetUserPermissionTitles(userID, role)

	if err != nil {
		return false, richerror.New(op).WithError(err)
	}

	for _, pt := range permissionTitles {
		for _, p := range permissions {
			if p == pt {
				return true, nil
			}
		}
	}
	return false, nil
}
