package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jesusvico/http-uptime/internal/config"
	"github.com/jesusvico/http-uptime/internal/metrics"
)

func Start(cfg *config.Config, port string) error {
	// Start a Goroutine for each endpoint
	for _, endpoint := range cfg.Endpoints {
		go func() {
			for {
				metrics.Collect(endpoint)
				time.Sleep(10 * time.Second)
			}
		}()
	}

	// Expose metrics endpoint
	http.Handle("/metrics", metrics.Handler())

	// Start the server
	server := &http.Server{Addr: fmt.Sprintf(":%s", port)}
	go func() {
		slog.Info("Starting server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %v", err)
	}

	slog.Info("Server stopped")
	return nil
}
