package app

import (
	"github.com/ledungcobra/notification-service/pkg/service"
	"google.golang.org/grpc"
	"net"
	"os"
)

type App struct {
	mailService service.MailServiceServer
}

func NewApp(mailService service.MailServiceServer) *App {
	a := &App{
		mailService: mailService,
	}
	return a
}

func (a *App) Listen() error {
	gRpcPort := os.Getenv("NOTIFICATION_PORT")
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+gRpcPort)
	if err != nil {
		return err
	}
	service.RegisterMailServiceServer(s, a.mailService)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
