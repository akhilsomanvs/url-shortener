package models

type ApiResponseModel struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewApiResponseModel(message string, data any) ApiResponseModel {
	return ApiResponseModel{
		Message: message,
		Data:    data,
	}
}
