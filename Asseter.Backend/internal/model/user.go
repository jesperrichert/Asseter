package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string
	AccessToken string
	IsOidc      bool
	APIAccessId int
	APIAccess   *APIAccess
}
