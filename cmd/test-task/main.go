package main

import (
	"fmt"
	"github.com/egorbokov/middleware/internal/pkg/app"
	"log"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		fmt.Println("Error during starting web-server!")
	}
}
