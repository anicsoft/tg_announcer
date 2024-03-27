package main

import (
	"anik/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app:%s", err)
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run the app: %s", err)
	}
}
