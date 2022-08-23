package users

import (
	"errors"
	"gorm.io/gorm"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"strings"
)

type UserDaoStub struct {
	interfaces.IUserDAO
	users     map[uint]*models.User
	CurrentID uint
}

const (
	TestPassword = "test123"
	TestEmail    = "test@gmail.com"
)

func NewUserDaoStub() *UserDaoStub {
	hashedPassword, _ := common.HashPassword(TestPassword)
	u := &UserDaoStub{
		users: map[uint]*models.User{
			1: {
				Email:    TestEmail,
				Password: hashedPassword,
				Detail:   models.Detail{},
				Post:     nil,
				Model: gorm.Model{
					ID: 0,
				},
				Verified: false,
			},
		},
	}
	u.CurrentID = 2
	return u
}

func (u *UserDaoStub) Save(user *models.User) error {
	if user.ID == 0 {
		for _, v := range u.users {
			if v.Email == user.Email {
				return errors.New("email is already exist")
			}
		}
	}

	user.ID = u.CurrentID
	u.users[user.ID] = user
	u.CurrentID++
	return nil
}

func (u *UserDaoStub) Find(query any, args ...any) (*models.User, error) {
	queryStr, ok := query.(string)
	if !ok {
		return nil, errors.New("query is not string")
	}
	switch {
	case strings.Contains(queryStr, "id"):
		return u.users[args[0].(uint)], nil
	case strings.Contains(queryStr, "email"):
		for _, v := range u.users {
			if v.Email == args[0].(string) {
				return v, nil
			}
		}
	case strings.Contains(queryStr, "user_name"):
		for _, v := range u.users {
			if v.UserName == args[0].(string) {
				return v, nil
			}
		}
	}
	return nil, dao.ErrRecordNotFound
}

type NotificationServiceStub struct {
	interfaces.INotificationService
}

type CommonDaoStub[T any] struct {
	interfaces.ICommonDao[T]
}

func NewCommonDaoStub[T any]() *CommonDaoStub[T] {
	return &CommonDaoStub[T]{}
}

func (c *CommonDaoStub[T]) Save(object *T) error {
	return nil
}

func (c *CommonDaoStub[T]) Delete(object any) error {
	return nil
}

func (c *CommonDaoStub[T]) Find(query any, args ...any) (*T, error) {
	return nil, nil
}

func (c *CommonDaoStub[T]) RawQuery(query string, args ...any) error {
	return nil
}
