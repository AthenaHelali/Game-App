package uservalidator

import (
	"context"
	"fmt"
	"game-app/param"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) error {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(IRPhoneNumberRegex)).Error(errormessage.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.doesPhoneNumberExist)),
		validation.Field(&req.Password, validation.Required),
	); err != nil {
		return richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req}).WithMessage(errormessage.ErrorMsgInvalidInput).WithKind(richerror.KindInvalid)
	}
	return nil
}

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)

	_, err := v.repo.GetUserByPhoneNumber(context.Background(), phoneNumber)
	if err != nil {
		return fmt.Errorf(errormessage.ErrorMsgNotFound)
	}
	return nil
}
