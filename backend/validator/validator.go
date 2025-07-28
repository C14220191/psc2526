package validator

// TODO : Add error parser for validation failures so that the return format is prettier

import (
	"context"
	"reflect"
	"strings"
	"github.com/go-playground/validator/v10"
)

// Initiate only one instance of validator
var validate *validator.Validate

func InitValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterTagNameFunc(RegisterJSONTagNameFunc)
	}

	return validate
}

func GetValidator() *validator.Validate {
	if validate == nil {
		return InitValidator()
	} else {
		return validate
	}
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.Struct(s)
}

// For cases where we need to use custom validation rules, make sure to register them

// Custom validation tag name to be used in struct tag
type CustomValidatorName string

func RegisterJSONTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

type EchoValidator struct {
	Val *validator.Validate
}

func (cv *EchoValidator) Validate(i interface{}) error {
	if err := cv.Val.Struct(i); err != nil {
		return err
	}
	return nil
}

func GetEchoAdapter() *EchoValidator {
	return &EchoValidator{
		Val: GetValidator(),
	}
}
