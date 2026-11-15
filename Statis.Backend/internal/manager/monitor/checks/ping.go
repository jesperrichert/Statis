package checks

import (
	probing "github.com/prometheus-community/pro-bing"
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

func (s *PingCheck) Check(monitor *model.Monitor[monitors.PingMonitor]) {
	m := monitor
	pinger, err := probing.NewPinger(m.TypeData.Host)
	if err != nil {
		m.State = monitors.Offline.ToString()
	} else {
		pinger.Count = 4
		err = pinger.Run() // Blocks until finished.
		if err != nil {
			m.State = monitors.Offline.ToString()
		} else {
			stats := pinger.Statistics()
			if stats.PacketsRecv > 1 {
				m.State = monitors.Online.ToString()
			} else {
				m.State = monitors.Degraded.ToString()
			}
		}
	}
	s.DB.Where("identifier = ?", m.Identifier).Updates(&m)
}
