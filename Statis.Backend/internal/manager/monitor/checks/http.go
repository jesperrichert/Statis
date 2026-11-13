package checks

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
)

type HttpCheck struct {
	ICheck
	DB *gorm.DB
}

func NewHttpCheck(db *gorm.DB) *HttpCheck {
	return &HttpCheck{
		DB: db,
	}
}

func (s *HttpCheck) Check(monitor *model.Monitor[any]) {

}
