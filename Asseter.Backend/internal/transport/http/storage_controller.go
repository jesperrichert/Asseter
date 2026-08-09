package http

import (
	"github.com/gin-gonic/gin"
	"go.Asseter/internal/services"
	"gorm.io/gorm"
)

type StorageController struct {
	db      *gorm.DB
	service *services.StorageService
}

func NewStorageController(db *gorm.DB, service *services.StorageService) *StorageController {
	return &StorageController{
		db:      db,
		service: service,
	}
}

// POST
func (e *StorageController) Post(ctx *gin.Context) {

}

// GET
func (e *StorageController) Get(ctx *gin.Context) {

}

// PUT
func (e *StorageController) Put(ctx *gin.Context) {

}

// DELETE
func (e *StorageController) Delete(ctx *gin.Context) {

}
