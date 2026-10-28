package repository

import (
	"jespersen.zip.statis/internal/model"
)

type ExampleRepository struct {
	Repository[model.Example]
}

func NewExample() *ExampleRepository {
	return &ExampleRepository{}
} 

//func(e* ExampleRepository)Extend(db *gorm.DB) {}
