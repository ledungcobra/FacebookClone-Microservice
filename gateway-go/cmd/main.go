package main

import (
	"log"
	"os"

	"ledungcobra/gateway-go/pkg/app"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
	a := app.NewServer()
	err := a.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := a.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	err = a.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
