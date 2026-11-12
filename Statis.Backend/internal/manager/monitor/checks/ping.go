package checks

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
)

type PingCheck struct {
	ICheck
	DB *gorm.DB
}

func NewPingCheck(db *gorm.DB) *PingCheck {
	return &PingCheck{
		DB: db,
	}
}

func (s *PingCheck) Check(monitor *model.Monitor) {

}
