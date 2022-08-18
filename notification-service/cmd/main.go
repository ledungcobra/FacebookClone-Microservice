package main

import (
	"github.com/joho/godotenv"
	"github.com/ledungcobra/notification-service/pkg/app"
	"github.com/ledungcobra/notification-service/pkg/service"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	a := app.NewApp(service.NewMailService())
	err := a.Listen()
	if err != nil {
		log.Fatalln("Error listening: ", err)
	}
}
