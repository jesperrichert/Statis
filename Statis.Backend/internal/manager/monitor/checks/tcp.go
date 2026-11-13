package checks

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
)

type TCPCheck struct {
	ICheck[monitors.TcpMonitor]
	DB *gorm.DB
}

func NewTCPCheck(db *gorm.DB) *TCPCheck {
	return &TCPCheck{
		DB: db,
	}
}

func (s *TCPCheck) Check(monitor model.Monitor[monitors.TcpMonitor]) {

}
