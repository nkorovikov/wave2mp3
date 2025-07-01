package main

import (
	"fmt"
	"log"

	"github.com/nkorovikov/wave2mp3/cmd/app"
)

func main() {
	application := app.New()

	err := application.Init()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot init app. Err: %w", err))
	}

	err = application.Run()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot run app. Err: %w", err))
	}
}
