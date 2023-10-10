package backofficeuserservice

import (
	entity2 "game-app/internal/entity"
)

type Service struct {
}

func New() Service {
	return Service{}
}
func (s Service) ListAllUsers() ([]entity2.User, error) {
	//TODO - implement me
	list := make([]entity2.User, 0)
	list = append(list, entity2.User{
		ID:          0,
		PhoneNumber: "fake",
		Name:        "fake",
		Password:    "123",
		Role:        entity2.AdminRole,
	})
	return list, nil
}
