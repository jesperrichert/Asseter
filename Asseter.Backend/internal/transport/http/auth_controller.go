package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/services"
	"gorm.io/gorm"
)

type AuthController struct {
	db      *gorm.DB
	service *services.AuthService
}

func NewAuthController(db *gorm.DB, service *services.AuthService) *AuthController {
	return &AuthController{
		db:      db,
		service: service,
	}
}

// POST
func (e *AuthController) Post(ctx *gin.Context) {

}

// GET
func (e *AuthController) Get(ctx *gin.Context) {

}

// GET
func (e *AuthController) Oidc(ctx *gin.Context) {
	e.service.Oidc(ctx, ctx.Query("code"))
}

// GET
func (e *AuthController) OidcConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"authenticationUrl": e.service.GenerateOidcAuthorizationUrl(),
	})
}
