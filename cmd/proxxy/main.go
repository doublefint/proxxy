//Package proxxy - see readme.md
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"proxxy/internal/app"
)

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {

	// TODO: "go.uber.org/zap"
	logger := log.New(os.Stdout, "proxxy> ", log.LstdFlags)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	port := getEnv("PROXXY_PORT", "8080")
	proxxy := app.New(port, logger)

	logger.Println("listen at", proxxy.Addr())

	if err := proxxy.Start(ctx); err != nil {
		logger.Fatalf("an error occured while running the app: %v", err)
		os.Exit(1)
	}

}
