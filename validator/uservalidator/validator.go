package uservalidator

import "game-app/entity"

const (
	IRPhoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}
type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return Validator{repo: repository}
}
