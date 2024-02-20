package main

import (
	"context"
	"fmt"

	"github.com/jseiser/justin-go/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
