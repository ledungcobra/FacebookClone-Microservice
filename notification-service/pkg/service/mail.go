package service

type IMailService interface {
	SendMail(from string, to string, subject string, body string, templateType string) error
}
