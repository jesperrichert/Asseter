package services

import "log"

type StorageService struct {
}

func NewStorageService() *StorageService {
	return &StorageService{}
}

func (e *StorageService) Log(message string) {
	log.Println(message)
}
