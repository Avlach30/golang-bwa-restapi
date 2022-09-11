package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"` //* Formatting property of json response
	Data interface{} `json:"data"` //* Any data type
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ApiResponse(isSuccess bool, message string, data any) Response {
	meta := Meta{
		Success: isSuccess,
		Message: message,
	}

	res := Response{
		Meta: meta,
		Data: data,
	}

	return res
}

func ErrorValidationResponse(err error) []string {
	//* Menghandle error yg disebabkan validasi
	var errors []string
	for _, error := range err.(validator.ValidationErrors) {
		errors = append(errors, error.Error())
	}

	return errors
}