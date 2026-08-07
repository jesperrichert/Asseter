package repository

import (
	"go.Asseter/internal/model"
)

type ExampleRepository struct {
	Repository[model.Example]
}

func NewExample() *ExampleRepository {
	return &ExampleRepository{}
} 

//func(e* ExampleRepository)Extend(db *gorm.DB) {}
