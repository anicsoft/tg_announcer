package main

import (
	"anik/internal/app"
	"context"
	"log"
)

//	@title		Announcement bot API
//	@version	0.1
//	@description

//	@host		localhost:8080
//	@BasePath	/
func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init application:%s", err)
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run the application: %s", err)
	}
}
