package dtobuilderwithdata

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-org/internal/dto"
	"github.com/mrexmelle/connect-org/internal/localerror"
)

type Class[T any] struct {
	Data              *T
	Error             error
	PreWriteHook      func(*T)
	LocalErrorService *localerror.Service
}

func New[T any](data *T, err error) *Class[T] {
	return &Class[T]{
		Data:         data,
		Error:        err,
		PreWriteHook: nil,
	}
}

func (c *Class[T]) WithPrewriteHook(hook func(*T)) *Class[T] {
	c.PreWriteHook = hook
	return c
}

func (c *Class[T]) RenderTo(w http.ResponseWriter) {
	errorInfo := c.LocalErrorService.Map(c.Error)

	responseBody, _ := json.Marshal(
		&dto.HttpResponseWithData[T]{
			Data: c.Data,
			Error: dto.ServiceError{
				Code:    errorInfo.ServiceErrorCode,
				Message: errorInfo.ServiceErrorMessage,
			},
		},
	)
	if errorInfo.HttpStatusCode == http.StatusOK {
		if c.PreWriteHook != nil {
			c.PreWriteHook(c.Data)
		}
		w.Write(responseBody)
	} else {
		http.Error(w, string(responseBody), errorInfo.HttpStatusCode)
	}
}
