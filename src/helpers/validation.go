package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func VariableRequiredValidation(param any) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validator.New().Var(param, "required")

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(param).Elem().FieldByName(err.Field())
			fieldName, _ := field.Tag.Lookup("json")

			errors = append(errors, &ErrorResponse{
				ErrorMessage: fmt.Sprintf("%s must contain %s", fieldName, err.ActualTag()),
			})
		}
	}

	return errors
}

func StructValidation(param any) []*ErrorResponse {
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

func PasswordValidation(password string, errors []*ErrorResponse) []*ErrorResponse {
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

func ImageValidation(file *multipart.FileHeader) []*ErrorResponse {
	var errors []*ErrorResponse

	if file.Size > int64(2<<20) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "image must contain less than 2 MB",
		})
	}

	fileHeader, err := file.Open()
	if err != nil {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "failed to open image",
		})
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512)
	if _, err := fileHeader.Read(buffer); err != nil {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "image failed to read image",
		})
	}

	fileType := http.DetectContentType(buffer)
	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg"}
	if !isValidFileType(validFileTypes, fileType) {
		errors = append(errors, &ErrorResponse{
			ErrorMessage: "image must contain .png|.jpg|.jpeg",
		})
	}

	return errors
}

func SizeUploadValidation(fileSize int64, maxFileSize int64) error {
	if fileSize > maxFileSize {
		return errors.New("file too large")
	}
	return nil
}

func TypeUploadValidation(buffer []byte, validFileTypes []string) error {
	fileType := http.DetectContentType(buffer)
	if !isValidFileType(validFileTypes, fileType) {
		return errors.New("type of file invalid; only png, jpg, jpeg, and pdf")
	}
	return nil
}

func isValidFileType(validFileTypes []string, fileType string) bool {
	for _, validType := range validFileTypes {
		if validType == fileType {
			return true
		}
	}
	return false
}
