package uservalidator

import (
	"context"
	"fmt"
	"game-app/internal/param"
	"game-app/internal/pkg/errormessage"
	"game-app/internal/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) error {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),

		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]{10,}$"))),

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(IRPhoneNumberRegex)).
				Error(errormessage.ErrorMsgPhoneNumberIsNotUnique),
			validation.By(v.checkPhoneNumberUniqueness)),
	); err != nil {
		return richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req}).WithMessage(errormessage.ErrorMsgInvalidInput).WithKind(richerror.KindInvalid)
	}
	return nil
}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	uniqueness, err := v.repo.IsPhoneNumberUnique(context.Background(), phoneNumber)
	if err != nil {
		return fmt.Errorf(errormessage.ErrorMsgNotFound)
	}
	if !uniqueness {
		return fmt.Errorf(errormessage.ErrorMsgPhoneNumberIsNotUnique)
	}
	return nil
}
