package interfaces

import (
	"io"
	"ledungcobra/gateway-go/pkg/app_service"
	"ledungcobra/gateway-go/pkg/cloudinary"
)

type INotificationService interface {
	SendMail(to, subject, body string) (*app_service.SendMailResponse, error)
}

type IUploadImageService interface {
	UploadFromBytes(reader io.Reader) (*cloudinary.UploadResult, error)
}
