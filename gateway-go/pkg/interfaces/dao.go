package interfaces

import "ledungcobra/gateway-go/pkg/models"

type IUserDAO interface {
	SaveUser(user *models.User) error
	Find(user *models.User, query any, args any) error
}
