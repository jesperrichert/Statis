package checks

import (
	"net"
	"strconv"

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

func (s *TCPCheck) Check(monitor *model.Monitor[monitors.TcpMonitor]) {
	m := monitor
	tcp, err := net.ResolveTCPAddr("tcp", m.TypeData.Host+":"+strconv.Itoa(m.TypeData.Port))
	conn, err := net.DialTCP("tcp", nil, tcp)
	if err != nil {
		m.State = monitors.Offline.ToString()
	} else {
		defer conn.Close()
		_, err = conn.Write([]byte("Ping"))

		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			m.State = monitors.Degraded.ToString()
		} else {
			m.State = monitors.Online.ToString()
		}
	}

	s.DB.Where("identifier = ?", monitor.Identifier).Updates(&monitor)
}
