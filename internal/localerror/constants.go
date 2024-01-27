package localerror

import (
	"database/sql"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrAuthentication = errors.New("authentication_error")
	ErrBadJson        = errors.New("bad_json")
	ErrBadHierarchy   = errors.New("bad_hierarchy")
	ErrAlreadyMax     = errors.New("already_max")
	ErrIdNotInteger   = errors.New("id_not_integer")
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

	ErrAuthentication: NewCodePair(http.StatusUnauthorized, ErrAuthentication.Error()),
	ErrBadJson:        NewCodePair(http.StatusBadRequest, ErrBadJson.Error()),
	ErrBadHierarchy:   NewCodePair(http.StatusBadRequest, ErrBadHierarchy.Error()),
	ErrAlreadyMax:     NewCodePair(http.StatusForbidden, ErrAlreadyMax.Error()),
	ErrIdNotInteger:   NewCodePair(http.StatusBadRequest, ErrIdNotInteger.Error()),
}
