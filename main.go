package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/jseiser/justin-go/application"
)

func main() {
	app := application.New()
	fmt.Println("Created new app")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
