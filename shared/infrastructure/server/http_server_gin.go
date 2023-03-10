package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
)

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(log logger.Logger, address string, appData gogen.ApplicationData) GinHTTPHandler {

	router := gin.Default()
	// PING API
	router.Use(cors.New(cors.Config{
		ExposeHeaders:    []string{"Data-Length"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000", "https://goarisan-staging.up.railway.app"},
	}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, appData)
	})
	// contentStatic, _ := fs.Sub(web.StaticFiles, "dist")
	// router.StaticFS("/web", http.FS(contentStatic))

	// CORS

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(log, router, address),
		Router:             router,
	}
}
