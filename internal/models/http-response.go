package models

type HttpResponseModel struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func New(message string, data any) HttpResponseModel {
	return HttpResponseModel{
		Message: message,
		Data:    data,
	}
}
