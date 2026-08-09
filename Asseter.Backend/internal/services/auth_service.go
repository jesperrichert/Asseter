package services

import "log"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (e *AuthService) Log(message string) {
	log.Println(message)
}
