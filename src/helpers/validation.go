package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func ValidateStruct(param any) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validator.New().Struct(param)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(param).Elem().FieldByName(err.Field())
			fieldName, _ := field.Tag.Lookup("json")
			var message string
			if err.Param() == "" {
				message = fmt.Sprintf("%s must contain %s", fieldName, err.ActualTag())
			} else {
				message = fmt.Sprintf("%s must contain %s=%s", fieldName, err.ActualTag(), err.Param())
			}

			errors = append(errors, &ErrorResponse{
				ErrorMessage: message,
			})
		}
	}
	return errors
}

func ValidatePassword(password string, errors []*ErrorResponse) []*ErrorResponse {
	for _, err := range errors {
		if strings.Contains(fmt.Sprintf("%s", err), "password") {
			return errors
		}
	}

	uppercasePassword := regexp.MustCompile(`[A-Z]`)
	spaceProhibitedPassword := regexp.MustCompile(`[\s]`)
	numberPassword := regexp.MustCompile(`[0-9]`)
	specialPassword := regexp.MustCompile(`[\W_]`)

	if !uppercasePassword.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "password must contain at least one uppercase letter",
		})
	} else if spaceProhibitedPassword.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "password must contain no space",
		})
	} else if !numberPassword.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "password must contain at least one digit number",
		})
	} else if !specialPassword.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "password must contain at least one special letter",
		})
	}

	return errors
}
