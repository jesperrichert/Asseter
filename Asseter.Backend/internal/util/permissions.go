package util

import (
	"slices"

	"go.Asseter/internal/model"
	"gorm.io/gorm"
)

func Permission(db gorm.DB, session, permission string) (bool, *model.User) {
	var apiAccess model.APIAccess
	db.Where("token = ?", session).Preload("User").First(&apiAccess)
	if slices.Contains(apiAccess.Permissions, permission) {
		return true, apiAccess.User
	} else {
		return false, nil
	}
}
