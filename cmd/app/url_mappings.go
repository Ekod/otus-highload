package main

import (
	"github.com/Ekod/otus-highload/internal/transport/controllers"
	"github.com/Ekod/otus-highload/internal/transport/controllers/ping"
	"github.com/Ekod/otus-highload/internal/transport/services"
	"github.com/Ekod/otus-highload/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	services *services.Services
}

func APIMux(cfg APIMuxConfig) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodDelete}
	router.Use(cors.New(config))
	router.Use(gin.Recovery())
	router.GET("/ping", ping.Ping)

	userHandlers := controllers.UserHandlers{
		Service: cfg.services,
		Logger:  cfg.Log,
	}

	middleware := middlewares.Middleware{
		Logger: cfg.Log,
	}

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/login", userHandlers.LoginUser)
		apiGroup.POST("/register", userHandlers.RegisterUser)
		apiGroup.GET("/users", middleware.CheckToken, userHandlers.GetUsers)
		apiGroup.GET("/info", middleware.CheckToken, userHandlers.GetCurrentUser)
		apiGroup.GET("/search-users", middleware.CheckToken, userHandlers.GetUsersByFullName)
	}

	return router
}
