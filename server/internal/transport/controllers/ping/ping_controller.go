package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping для хост платформы, чтобы проверять состояние сервиса
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
