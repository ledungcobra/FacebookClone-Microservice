package dao

import (
	i "ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = DBError{Message: "Record not found", err: gorm.ErrRecordNotFound}
)

type DBError struct {
	Message string `json:"message"`
	err     error
}

func (db DBError) Error() string {
	return db.Message + ": " + db.err.Error()
}

type UserDAO struct {
	CommonDao[models.User]
}

func NewUserDao(db *gorm.DB) i.IUserDAO {
	return &UserDAO{
		*NewCommonDao[models.User](db),
	}
}
