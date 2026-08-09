package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/model"
	"go.Asseter/internal/util"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	DB *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		DB: db,
	}
}

func (middleware *AuthMiddleware) Handle(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 {
		util.GenerateResponse(
			ctx,
			http.StatusUnauthorized,
			"Unauthorized",
			true,
			nil,
		)
		return
	}
	var access model.APIAccess
	middleware.DB.First(&access, "token = ?", token)

	if len(access.Token) != 0 {
		ctx.Next()
	}

	util.GenerateResponse(
		ctx,
		http.StatusUnauthorized,
		"Unauthorized",
		true,
		nil,
	)
}
