package server

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BlackHole1/sesmate/internal/server/controller"
	"github.com/BlackHole1/sesmate/internal/server/middleware"
)

type Context struct {
	dir  string
	addr string
}

func New(dir, host string, port int) *Context {
	return &Context{
		dir:  dir,
		addr: host + ":" + strconv.Itoa(port),
	}
}

func (c *Context) Execute() error {
	engine := gin.New()

	engine.Use(gin.Logger(), middleware.RawBody(), middleware.Template(c.dir))
	c.router(engine)

	return engine.Run(c.addr)
}

func (c *Context) router(engine *gin.Engine) {
	{
		aws := engine.Group("/aws", middleware.AWSRequestId())
		v2 := aws.Group("/v2")
		// SendEmail
		v2.POST("/email/outbound-emails", controller.AWSV2SendEmail)
	}

	{
		engine.GET("/", controller.ViewEmail)
	}
}
