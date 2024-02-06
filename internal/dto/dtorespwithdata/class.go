package dtorespwithdata

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-org/internal/dto"
)

type Class[T any] struct {
	Data  *T               `json:"data"`
	Error dto.ServiceError `json:"error"`

	PreWriteHook func(*T) `json:"-"`
}

func New[T any](data *T, errCode string, errMessage string) *Class[T] {
	return &Class[T]{
		Data: data,
		Error: dto.ServiceError{
			Code:    errCode,
			Message: errMessage,
		},
		PreWriteHook: nil,
	}
}

func NewError(errCode string, errMessage string) *Class[any] {
	return New[any](nil, errCode, errMessage)
}

func (c *Class[T]) WithPrewriteHook(hook func(*T)) *Class[T] {
	c.PreWriteHook = hook
	return c
}

func (c *Class[T]) RenderTo(w http.ResponseWriter, httpStatusCode int) {
	if httpStatusCode == http.StatusOK && c.PreWriteHook != nil {
		c.PreWriteHook(c.Data)
	}
	responseBody, _ := json.Marshal(c)
	w.WriteHeader(httpStatusCode)
	w.Write(responseBody)
}
