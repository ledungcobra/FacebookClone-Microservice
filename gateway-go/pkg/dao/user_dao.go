package dao

import (
	i "ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"

	"gorm.io/gorm"
)

type UserError struct {
	Message string
	err     error
}

// Error implements error
func (u UserError) Error() string {
	return u.Message + ": " + u.err.Error()
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) i.IUserDAO {
	return &UserDAO{
		db: db,
	}
}

func (u *UserDAO) SaveUser(user *models.User) error {
	if result := u.db.Save(&user); result.Error != nil {
		return UserError{"Cannot save user", result.Error}
	}
	return nil
}
