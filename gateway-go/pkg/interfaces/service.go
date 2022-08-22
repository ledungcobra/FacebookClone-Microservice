package interfaces

import "ledungcobra/gateway-go/pkg/app_service"

type INotificationService interface {
	SendMail(to, subject, body string) (*app_service.SendMailResponse, error)
}
