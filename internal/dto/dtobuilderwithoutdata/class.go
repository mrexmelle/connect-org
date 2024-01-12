package dtobuilderwithoutdata

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-orgs/internal/dto"
	"github.com/mrexmelle/connect-orgs/internal/localerror"
)

type Class struct {
	Error             error
	LocalErrorService *localerror.Service
}

func New(err error) *Class {
	return &Class{
		Error: err,
	}
}

func (c *Class) RenderTo(w http.ResponseWriter) {
	errorInfo := c.LocalErrorService.Map(c.Error)

	responseBody, _ := json.Marshal(
		&dto.HttpResponseWithoutData{
			Error: dto.ServiceError{
				Code:    errorInfo.ServiceErrorCode,
				Message: errorInfo.ServiceErrorMessage,
			},
		},
	)
	if errorInfo.HttpStatusCode == http.StatusOK {
		w.Write(responseBody)
	} else {
		http.Error(w, string(responseBody), errorInfo.HttpStatusCode)
	}
}
