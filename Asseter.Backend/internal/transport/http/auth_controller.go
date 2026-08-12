package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/dto"
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
func (e *AuthController) Register(ctx *gin.Context) {
	var data dto.LocalUserDto
	json.NewDecoder(ctx.Request.Body).Decode(&data)
	defer ctx.Request.Body.Close()
	e.service.Register(ctx, data.Username, data.Password)
}

// GET
func (e *AuthController) Login(ctx *gin.Context) {
	var data dto.LocalUserDto
	json.NewDecoder(ctx.Request.Body).Decode(&data)
	defer ctx.Request.Body.Close()
	e.service.Login(ctx, data.Username, data.Password)
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
