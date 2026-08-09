package router

import (
	"github.com/gin-gonic/gin"
	"go.Asseter/internal/transport/http"
	"go.Asseter/internal/transport/http/middleware"
	"gorm.io/gorm"
)

type RouterConfig struct {
	App               *gin.Engine
	DB                *gorm.DB
	AuthController    *http.AuthController
	StorageController *http.StorageController
	//Register Controller Here
}

func (c *RouterConfig) Setup() {
	if c.App == nil {
		c.App = gin.Default()
	}

	authMiddleware := middleware.NewAuthMiddleware(c.DB)

	auth := c.App.Group("/auth")
	{
		auth.GET("/", c.AuthController.Get)
		auth.GET("/oidc", c.AuthController.Oidc)
		auth.POST("/", c.AuthController.Post)
	}

	api := c.App.Group("/api")
	{
		api.POST("/storage", authMiddleware.Handle, c.StorageController.Post)
		api.GET("/storage/:fileName", authMiddleware.Handle, c.StorageController.Get)
		api.PUT("/storage/:fileName", authMiddleware.Handle, c.StorageController.Put)
		api.DELETE("/storage/:fileName", authMiddleware.Handle, c.StorageController.Delete)
	}

}
