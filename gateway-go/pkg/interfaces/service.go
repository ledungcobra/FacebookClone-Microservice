package interfaces

import "ledungcobra/gateway-go/pkg/service"

type INotificationService interface {
	SendMail(to, subject, body string) (*service.SendMailResponse, error)
}
