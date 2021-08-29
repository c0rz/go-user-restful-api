package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}
