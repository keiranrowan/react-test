package client

import (
	"github.com/urfave/cli/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func Init(ctx *cli.Context) error {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	
	r.Run(":8000")

	return nil
}
