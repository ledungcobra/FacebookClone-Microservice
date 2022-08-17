package dao

import (
	i "ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = DBError{Message: "Record not found"}
)

type DBError struct {
	Message string
	err     error
}

func (db DBError) Error() string {
	return db.Message + ": " + db.err.Error()
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) i.IUserDAO {
	return &UserDAO{
		db: db,
	}
}

func (u *UserDAO) Find(user *models.User, query any, args any) error {
	if result := u.db.Where(query, args).First(user); result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return ErrRecordNotFound
	}
	return nil
}

func (u *UserDAO) SaveUser(user *models.User) error {
	if result := u.db.Save(&user); result.Error != nil {
		return handleError(result.Error)
	}
	return nil
}
