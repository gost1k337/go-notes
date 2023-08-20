package app

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
	"tg_bot/configs"
	"tg_bot/internal/handlers"
	"tg_bot/internal/repository"
	"tg_bot/internal/service"
	"tg_bot/pkg/httpserver"
	"tg_bot/pkg/lib"
	"tg_bot/pkg/postgres"
)

// Run
//
// @title           			Notes Service
// @version         			1.0
// @description     			This is a service for creating notes with reminders and notifications.
// @host      					localhost:3000
// @BasePath  					/
func Run(cfg *configs.Config) {
	logger := lib.NewLogger(cfg)

	db, err := postgres.New(cfg.PostgresDSN)
	if err != nil {
		logger.Fatal(err)
	}

	// Repos
	logger.Info("Initializing repositories...")
	repositories := repository.NewRepositories(db, logger)

	// Services
	logger.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps, logger)

	// Handlers
	logger.Info("Initializing handlers and routes...")
	handler := handlers.New(services, logger)

	// HTTP server
	logger.Info("Starting http server...")
	logger.Debugf("Server port: %s", cfg.ServerPort)
	httpServer := httpserver.New(handler.HTTP(), httpserver.Port(cfg.ServerPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	logger.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
