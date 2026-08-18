package router

import (
	"github.com/gin-gonic/gin"
	"go.Asseter/internal/env"
	"go.Asseter/internal/transport/http"
	"go.Asseter/internal/transport/http/middleware"
	"gorm.io/gorm"
)

type RouterConfig struct {
	Config            *env.Config
	App               *gin.Engine
	DB                *gorm.DB
	AuthController    *http.AuthController
	StorageController *http.StorageController
}

func (c *RouterConfig) Setup() {
	if c.App == nil {
		c.App = gin.Default()
	}

	authMiddleware := middleware.NewAuthMiddleware(c.DB)

	api := c.App.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", c.AuthController.Login)
			auth.POST("/register", c.AuthController.Register)
			auth.GET("/oidc/callback", c.AuthController.Oidc)
			auth.GET("/oidc", c.AuthController.OidcConfig)
		}

		api.POST("/storage", authMiddleware.Handle, c.StorageController.Post)
		api.GET("/storage/:fileName", authMiddleware.Handle, c.StorageController.Get)
		api.DELETE("/storage/:fileName", authMiddleware.Handle, c.StorageController.Delete)
	}

}
