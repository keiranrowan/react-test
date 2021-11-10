package server

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}
