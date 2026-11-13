package checks

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
)

type PingCheck struct {
	ICheck[monitors.PingMonitor]
	DB *gorm.DB
}

func NewPingCheck(db *gorm.DB) *PingCheck {
	return &PingCheck{
		DB: db,
	}
}

func (s *PingCheck) Check(monitor model.Monitor[monitors.PingMonitor]) {

}
