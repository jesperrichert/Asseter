package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/dto"
	"go.Asseter/internal/services"
	"go.Asseter/internal/util"
	"gorm.io/gorm"
)

type AuthController struct {
	DB      *gorm.DB
	Service *services.AuthService
}

func NewAuthController(db *gorm.DB, service *services.AuthService) *AuthController {
	return &AuthController{
		DB:      db,
		Service: service,
	}
}

// POST
func (e *AuthController) Register(ctx *gin.Context) {
	if e.Service.Config.AllowRegister == "false" {
		util.GenerateResponse(
			ctx,
			http.StatusMethodNotAllowed,
			"Registration is Disabled",
			true,
			nil,
		)
		return
	}
	var data dto.LocalUserDto
	json.NewDecoder(ctx.Request.Body).Decode(&data)
	defer ctx.Request.Body.Close()
	e.Service.Register(ctx, data.Username, data.Password)
}

// GET
func (e *AuthController) Login(ctx *gin.Context) {
	var data dto.LocalUserDto
	json.NewDecoder(ctx.Request.Body).Decode(&data)
	defer ctx.Request.Body.Close()
	e.Service.Login(ctx, data.Username, data.Password)
}

// GET
func (e *AuthController) Oidc(ctx *gin.Context) {
	e.Service.Oidc(ctx, ctx.Query("code"))
}

// GET
func (e *AuthController) OidcConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"authenticationUrl": e.Service.GenerateOidcAuthorizationUrl(),
	})
}
