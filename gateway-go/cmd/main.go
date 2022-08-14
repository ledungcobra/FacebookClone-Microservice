package main

import (
	"ledungcobra/gateway-go/pkg/app"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := app.NewServer()
	err := app.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
