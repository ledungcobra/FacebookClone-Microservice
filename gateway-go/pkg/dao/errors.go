package dao

import (
	"ledungcobra/gateway-go/pkg/config"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func NewDBError(message string, err error) error {
	cfg := config.Cfg
	if cfg.Env == "production" {
		return DBError{Message: message}
	} else {
		return DBError{message, err}
	}
}

func handleError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return NewDBError("Record not found", err)
	case gorm.ErrInvalidTransaction:
		return NewDBError("Invalid transaction", err)
	case gorm.ErrNotImplemented:
		return NewDBError("Not implemented", err)
	case gorm.ErrMissingWhereClause:
		return NewDBError("Missing where clause", err)
	case gorm.ErrUnsupportedRelation:
		return NewDBError("Unsupported relations", err)
	case gorm.ErrPrimaryKeyRequired:
		return NewDBError("Primary key required", err)
	case gorm.ErrModelValueRequired:
		return NewDBError("Model value required", err)
	case gorm.ErrInvalidData:
		return NewDBError("Invalid data", err)
	case gorm.ErrUnsupportedDriver:
		return NewDBError("Unsupported driver", err)
	case gorm.ErrRegistered:
		return NewDBError("Registered", err)
	case gorm.ErrInvalidField:
		return NewDBError("Invalid field", err)
	case gorm.ErrEmptySlice:
		return NewDBError("Empty slice found", err)
	case gorm.ErrDryRunModeUnsupported:
		return NewDBError("Dry run mode unsupported", err)
	case gorm.ErrInvalidDB:
		return NewDBError("Invalid db", err)
	case gorm.ErrInvalidValue:
		return NewDBError("Invalid value", err)
	case gorm.ErrInvalidValueOfLength:
		return NewDBError("Invalid association values, length doesn't match", err)
	case gorm.ErrPreloadNotAllowed:
		return NewDBError("Preload is not allowed when count is used", err)
	}
	if error, ok := err.(*pgconn.PgError); ok {
		switch error.Code {
		case "23505":
			return NewDBError(error.Detail, err)
		case "23503":
			return NewDBError("Foreign key constraint violation", err)
		case "23P01":
			return NewDBError("Invalid text representation", err)
		case "23502":
			return NewDBError("Not null violation", err)
		case "23514":
			return NewDBError("Check violation", err)
		case "23P02":
			return NewDBError("Invalid numeric value", err)
		}
		return NewDBError("Unknow error", error)
	}
	return NewDBError("Another error", err)
}
