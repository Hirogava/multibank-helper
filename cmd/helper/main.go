package main

import (
	"os"

	"github.com/Hirogava/multibank-helper/internal/config/environment"
	"github.com/Hirogava/multibank-helper/internal/config/logger"
	postgres "github.com/Hirogava/multibank-helper/internal/repository/postgres"
	router "github.com/Hirogava/multibank-helper/internal/transport/http"
)

func main() {
	environment.LoadEnvFile(".env")

	logger.LogInit()
	logger.Logger.Info("Starting Swifty Gasprom backend server")

	dbConnStr := os.Getenv("DB_CONNECT_STRING")
	if dbConnStr == "" {
		logger.Logger.Fatal("DB_CONNECT_STRING environment variable is required")
	}
	logger.Logger.Info("Connecting to database", "connection_string", dbConnStr)

	manager := postgres.NewManager("postgres", dbConnStr)
	logger.Logger.Info("Database connection established successfully")

	logger.Logger.Info("Running database migrations")
	manager.Migrate()
	logger.Logger.Info("Database migrations completed successfully")

	logger.Logger.Info("Initializing HTTP router")
	r := router.CreateRouter(manager)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
		logger.Logger.Warn("SERVER_PORT not set, using default port 8080")
	}

	logger.Logger.Info("Starting HTTP server", "port", serverPort)
	if err := r.Run(serverPort); err != nil {
		logger.Logger.Fatal("Failed to start HTTP server", "error", err)
	}
}
