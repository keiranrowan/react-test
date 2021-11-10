package server

import (
	"github.com/urfave/cli/v2"
	"github.com/gin-gonic/gin"
)

func Init(ctx *cli.Context) error {
	r := gin.New()

	RegisterMiddleware(r)
	RegisterRoutes(r)
	
	r.Run(":8080")
	
	return nil
}
