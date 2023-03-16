package main

import (
	"net/http"
	"os"

	"github.com/Ekod/otus-highload/internal/transport/controllers"
	"github.com/Ekod/otus-highload/internal/transport/controllers/ping"
	"github.com/Ekod/otus-highload/internal/transport/services"
	"github.com/Ekod/otus-highload/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown    chan os.Signal
	Log         *zap.SugaredLogger
	RedisClient *redis.Client
	Services    *services.Services
}

func APIMux(cfg APIMuxConfig) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut}
	router.Use(cors.New(config))
	router.Use(gin.Recovery())
	router.GET("/ping", ping.Ping)

	userHandlers := controllers.UserHandlers{
		Service: cfg.Services,
		Logger:  cfg.Log,
	}

	friendHandlers := controllers.FriendHandlers{
		Service: cfg.Services,
		Logger:  cfg.Log,
	}

	postHandlers := controllers.PostHandlers{
		Service:     cfg.Services,
		Logger:      cfg.Log,
		RedisClient: cfg.RedisClient,
	}

	middleware := middlewares.Middleware{
		Logger: cfg.Log,
	}

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/users", middleware.CheckToken, userHandlers.GetUsers)
		apiGroup.GET("/search-users", middleware.CheckToken, userHandlers.GetUsersByFullName)

		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/login", userHandlers.LoginUser)
			userGroup.POST("/register", userHandlers.RegisterUser)
			userGroup.GET("/info", middleware.CheckToken, userHandlers.GetCurrentUser)
		}

		friendGroup := apiGroup.Group("/friend")
		{
			friendGroup.GET("/list", middleware.CheckToken, friendHandlers.GetFriends)
			friendGroup.POST("/add", middleware.CheckToken, friendHandlers.MakeFriends)
			friendGroup.DELETE("/remove/:id", middleware.CheckToken, friendHandlers.RemoveFriend)
		}

		postGroup := apiGroup.Group("/post")
		{
			postGroup.GET("/get/:id", postHandlers.GetPost)
			postGroup.POST("/create", postHandlers.CreatePost)
			postGroup.PUT("/update/:id", postHandlers.UpdatePost)
			postGroup.DELETE("/delete/:id", postHandlers.DeletePost)
			postGroup.POST("/feed", postHandlers.FeedPost)
		}
	}

	return router
}
