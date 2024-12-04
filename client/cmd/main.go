package main

import (
	"log"

	"github.com/AlexBlackNn/authloyalty/client/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = application.Start()
	if err != nil {
		log.Fatal(err)
	}

}
