package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(status string, message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

func APIError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
