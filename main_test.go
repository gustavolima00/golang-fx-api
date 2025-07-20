package main

import (
	"context"
	"testing"
	"time"

	"go.uber.org/fx"
)

func TestStartApp(t *testing.T) {
	var app *fx.App

	app = fx.New(
		fx.Invoke(func() {
			// This function will be called when the application starts.
		}),
	)

	go func() {
		startApp(app)
	}()

	// Wait for the application to start.
	time.Sleep(1 * time.Second)

	// Stop the application.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}
