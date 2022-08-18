package service

import (
	"context"
	"google.golang.org/grpc"
	"os"
	"time"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendMail(to, subject, body string) (*SendMailResponse, error) {
	conn, err := s.getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	timeOutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	mailClient := NewMailServiceClient(conn)
	return mailClient.SendMail(timeOutCtx, &Mail{
		From:    os.Getenv("GATEWAY_EMAIL"),
		To:      to,
		Subject: subject,
		Body:    body,
	})
}

func (s *NotificationService) getConn() (*grpc.ClientConn, error) {
	return grpc.Dial(os.Getenv("GATEWAY_NOTIFICATION_GRPC"), grpc.WithInsecure())
}
