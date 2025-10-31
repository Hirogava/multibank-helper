package http

import (
	"time"

	"github.com/Hirogava/multibank-helper/internal/config/logger"
	"github.com/Hirogava/multibank-helper/internal/repository/postgres"
	"github.com/Hirogava/multibank-helper/internal/handler/helper"
	"github.com/Hirogava/multibank-helper/internal/handler/analys"
	"github.com/Hirogava/multibank-helper/internal/handler/best"
	"github.com/Hirogava/multibank-helper/internal/handler/prognozer"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CreateRouter(manager *postgres.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	logger.Logger.Debug("Registering Helper handlers")
	helper.InitHelperHandlers(r, manager)

	logger.Logger.Debug("Registering Best handlers")
	best.InitBestHandlers(r, manager)

	logger.Logger.Debug("Registering Analys handlers")
	analys.InitAnalysHandlers(r, manager)

	logger.Logger.Debug("Registering Prognozer handlers")
	prognozer.InitPrognozerHandlers(r, manager)

	logger.Logger.Info("HTTP router created successfully")
	return r
}
