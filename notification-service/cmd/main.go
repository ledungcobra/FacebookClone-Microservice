package main

import (
	"log"

	"github.com/ledungcobra/notification-service/pkg/service"
)

type args struct {
	from         string
	to           string
	subject      string
	body         string
	templateType string
}
type input struct {
	name    string
	args    args
	wantErr bool
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	tt :=
		input{name: "test sending mail should success", args: args{from: "Ledung", to: "ledungcobra@gmail.com", subject: "Test",
			body: `http://google.com`, templateType: "text/html"}, wantErr: false}

	if err := service.SendMail(tt.args.from, tt.args.to, tt.args.subject, tt.args.body, tt.args.templateType); (err != nil) != tt.wantErr {
		log.Fatalf("%s: unexpected error: %v", tt.name, err)
	}
}
