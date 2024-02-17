package main

import (
	"context"
	"fmt"
	"microservices/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
