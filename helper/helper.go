package helper

import "github.com/go-playground/validator/v10"

type (
	Response struct {
		Meta Meta `json:"meta"`
		Data any  `json:"data"`
	}

	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}
)

func APIResponse(message string, code int, status string, data any) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	JsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return JsonResponse
}

func ErrorValidation(err error) []string {
	var errString []string

	for _, e := range err.(validator.ValidationErrors) {
		errString = append(errString, e.Error())
	}

	return errString
}
