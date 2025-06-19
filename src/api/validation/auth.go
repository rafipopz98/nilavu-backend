package validation

import (
	request_auth "github.com/nilavu-backend/src/api/requests/auth"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateSignUpRequest(body request_auth.SignUP_Request) error{
	return  validation.ValidateStruct(&body,
	validation.Field(&body.First_Name, validation.Required, validation.Length(2, 50)),
	validation.Field(&body.Last_Name, validation.Required, validation.Length(2, 50)),
	validation.Field(&body.Email, validation.Required, is.Email),
	validation.Field(&body.Password, validation.Required, validation.Length(6, 50)),
	validation.Field(&body.Role, validation.Required),
	)
}