package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/defryheryanto/nebula/config"
)

func main() {
	var appServer *http.Server

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	go func() {
		db := setupDatabaseConnection(ctx)
		repo := setupRepositories(db)
		service := setupServices(repo)
		appServer = &http.Server{
			Addr:    fmt.Sprintf(":%s", config.Port),
			Handler: buildRoutes(setupHandler(service)),
		}

		slog.Info(fmt.Sprintf("starting server on %s", appServer.Addr))
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error starting server", "error", err)
		}
	}()

	<-done

	slog.Info("shutting down server")
	if err := appServer.Shutdown(ctx); err != nil {
		slog.Error("error shutting down server", "error", err)
	}
	cancel()

	slog.Info("server shutdown gracefully")
}
