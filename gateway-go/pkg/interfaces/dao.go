package interfaces

import "ledungcobra/gateway-go/pkg/models"

type IUserDAO interface {
	SaveUser(user *models.User) error
}
