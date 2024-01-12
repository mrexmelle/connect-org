package localerror

import (
	"database/sql"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrAuthentication = errors.New("authentication_error")
	ErrParsingJson    = errors.New("error_parsing_json")
)

const (
	ErrSvcCodeDuplicatedKey      = "duplicated_key"
	ErrSvcCodeForeignKeyViolated = "foreign_key_violation"
	ErrSvcCodeRecordNotFound     = "record_not_found"

	ErrSvcCodeUnregistered = "unregistered"
	ErrSvcCodeNone         = "success"
)

var ErrorMap = map[error]CodePair{
	gorm.ErrDuplicatedKey:      NewCodePair(http.StatusBadRequest, ErrSvcCodeDuplicatedKey),
	gorm.ErrForeignKeyViolated: NewCodePair(http.StatusBadRequest, ErrSvcCodeForeignKeyViolated),
	gorm.ErrRecordNotFound:     NewCodePair(http.StatusNotFound, ErrSvcCodeRecordNotFound),
	sql.ErrNoRows:              NewCodePair(http.StatusNotFound, ErrSvcCodeRecordNotFound),
	ErrAuthentication:          NewCodePair(http.StatusUnauthorized, ErrAuthentication.Error()),
	ErrParsingJson:             NewCodePair(http.StatusBadRequest, ErrParsingJson.Error()),
}
