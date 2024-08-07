package main

import (
	"context"
	"log"
	"tg_announcer/internal/app"
)

//	@title		GetAnnouncement bot API
//	@version	0.1
//	@description

// @host		localhost:8888
// @BasePath	/backend
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
