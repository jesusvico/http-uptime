package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/jesusvico/http-uptime/internal/config"
	"github.com/jesusvico/http-uptime/internal/metrics"
	"github.com/jesusvico/http-uptime/internal/server"
)

func main() {
	// Initialize the logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	// Parse the flags
	port := flag.String("p", "8080", "Port to listen on")
	configFile := flag.String("c", "config.yaml", "Configuration file")
	flag.Parse()

	// Read the configuration file
	conf, err := config.New(*configFile)
	if err != nil {
		slog.Error("Error reading configuration file", "error", err)
		os.Exit(1)
	}

	// Initialize metrics
	metrics.Init()

	// Start the server
	if err := server.Start(conf, *port); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
