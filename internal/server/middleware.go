package server

import (
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
