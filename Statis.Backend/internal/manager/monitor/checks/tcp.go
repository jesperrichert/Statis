package checks

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
)

type TCPCheck struct {
	ICheck
	DB *gorm.DB
}

func NewTCPCheck(db *gorm.DB) *TCPCheck {
	return &TCPCheck{
		DB: db,
	}
}

func (s *TCPCheck) Check(monitor *model.Monitor[any]) {

}
