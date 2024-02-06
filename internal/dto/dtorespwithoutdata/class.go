package dtorespwithoutdata

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-org/internal/dto"
)

type Class struct {
	Error dto.ServiceError `json:"error"`
}

func New(errCode string, errMessage string) *Class {
	return &Class{
		Error: dto.ServiceError{
			Code:    errCode,
			Message: errMessage,
		},
	}
}

func (c *Class) RenderTo(w http.ResponseWriter, httpStatusCode int) {
	responseBody, _ := json.Marshal(c)
	w.WriteHeader(httpStatusCode)
	w.Write(responseBody)
}
