package matchingvalidator

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) error {
	const op = "matchingvalidator.AddTowaitingListRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.Category,
			validation.Required,
			validation.By(v.IsCategoryValid)),
	); err != nil {
		return richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req}).WithMessage(errormessage.ErrorMsgInvalidInput).WithKind(richerror.KindInvalid)
	}
	return nil
}

func (v Validator) IsCategoryValid(value interface{}) error {
	category := value.(entity.Category)

	if !category.IsValid() {
		return fmt.Errorf(errormessage.ErrorMsgCategoryIsNotValid)
	}

	return nil
}
