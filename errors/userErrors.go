package errors

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func CustomError(field string, msg string) *ApiError {
	return &ApiError{
		Field: field,
		Msg:   msg,
	}
}

func (e *ApiError) Code() string {
	if e == nil {
		return ""
	}
	return e.Field
}

func (e *ApiError) Error() string {
	if e == nil {
		return ""
	}
	return e.Msg
}

func NewUserError(code, text string, args ...interface{}) *ApiError {

	return &ApiError{code, fmt.Sprintf(text, args...)}
}

func GetErrors(err error) []ApiError {
	var validationArray validator.ValidationErrors

	if errors.As(err, &validationArray) {
		errorObjArray := make([]ApiError, len(validationArray))

		fmt.Println(validationArray, "validationArray")

		for index, value := range validationArray {

			errorObjArray[index] = ApiError{
				Field: value.Field(),
				Msg:   value.Tag(),
			}
		}

		return errorObjArray
	}

	return nil

}
