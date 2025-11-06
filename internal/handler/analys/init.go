package analys

import (
	"github.com/Hirogava/multibank-helper/internal/handler/middleware"
	"github.com/Hirogava/multibank-helper/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func InitAnalysHandlers(r *gin.Engine, manager *postgres.Manager) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		
	}
}
	