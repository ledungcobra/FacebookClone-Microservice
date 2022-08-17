package main

import (
	"log"
	"os"

	"ledungcobra/gateway-go/pkg/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
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
