package config

import (
	"github.com/gin-gonic/gin"
	"go.Asseter/internal/services"
	"go.Asseter/internal/transport/http"
	"go.Asseter/internal/transport/http/router"
	"gorm.io/gorm"
)

type Appconfig struct {
	App *gin.Engine
	DB  *gorm.DB
}

func Build(config *Appconfig) {

	//Register Services
	authService := services.NewAuthService()
	storageService := services.NewStorageService()

	//Register Controller
	authController := http.NewAuthController(config.DB, authService)
	storageController := http.NewStorageController(config.DB, storageService)

	routeConfig := router.RouterConfig{
		App:               config.App,
		DB:                config.DB,
		AuthController:    authController,
		StorageController: storageController,
	}

	routeConfig.Setup()
}
