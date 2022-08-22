package interfaces

import "ledungcobra/gateway-go/pkg/models"

type ICommonDao[T any] interface {
	Save(object *T) error
	Delete(object any) error
	Find(query any, args ...any) (*T, error)
	RawQuery(query string, args ...any) error
}

type IUserDAO interface {
	ICommonDao[models.User]
}
