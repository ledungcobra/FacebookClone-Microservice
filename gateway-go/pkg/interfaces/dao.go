package interfaces

import "ledungcobra/gateway-go/pkg/models"

type IUserDAO interface {
	Save(user *models.User) error
	Find(query any, args ...any) (*models.User, error)
}
