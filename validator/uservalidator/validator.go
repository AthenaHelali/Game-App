package uservalidator

import (
	"fmt"
	"game-app/dto"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return Validator{repo: repository}
}

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) error {
	const op = "uservalidator.ValidateRegisterRequest"

	//TODO - add 3 to config
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]{10,}$"))),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09[0-9]{9}$")), validation.By(v.checkPhoneNumberUniqueness)),
	); err != nil {
		return richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req}).WithMessage(errormessage.ErrorMsgInvalidInput).WithKind(richerror.KindInvalid)
	}
	return nil
}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)
	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(errormessage.ErrorMsgPhoneNumberIsNotUnique)

		}
	}
	return nil
}
