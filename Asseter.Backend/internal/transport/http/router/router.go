package router

import (
	"github.com/gin-gonic/gin"
	"go.Asseter/internal/transport/http"
)

type RouterConfig struct {
	App               *gin.Engine
	ExampleController *http.ExampleController 
	//Register Controller Here
}

func (c *RouterConfig) Setup() {
	if c.App == nil {
		c.App = gin.Default()
	}

	api := c.App.Group("/api")
	{
		api.GET("/:id", c.ExampleController.Get)
		api.GET("/", c.ExampleController.List)
		api.POST("/", c.ExampleController.Post)
	}

}
