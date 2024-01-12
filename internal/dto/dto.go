package dto

type HttpResponseWithData[T any] struct {
	Data  *T           `json:"data"`
	Error ServiceError `json:"error"`
}

type HttpResponseWithoutData struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
