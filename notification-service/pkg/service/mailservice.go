package service

import "context"

type MailService struct {
	UnimplementedMailServiceServer
}

func (s *MailService) SendMail(ctx context.Context, mail *Mail) (*SendMailResponse, error) {
	err := SendMail(mail.From, mail.To, mail.Subject, mail.Body, "text/html")
	return &SendMailResponse{Success: err == nil}, err
}

func NewMailService() MailServiceServer {
	return &MailService{}
}
