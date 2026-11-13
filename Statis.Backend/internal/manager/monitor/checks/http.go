package checks

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
)

type HttpCheck struct {
	ICheck[monitors.HttpMonitor]
	DB *gorm.DB
}

func NewHttpCheck(db *gorm.DB) *HttpCheck {
	return &HttpCheck{
		DB: db,
	}
}

func (s *HttpCheck) Check(monitor model.Monitor[monitors.HttpMonitor]) {
	m := monitor
	client := http.Client{}
	req, err := http.NewRequest(m.TypeData.Method, m.TypeData.Url, strings.NewReader(m.TypeData.Body))
	res, err := client.Do(req)
	if err != nil {
		m.State = monitors.Offline.ToString()
	} else {
		defer res.Body.Close()
		validCode := 200
		if m.TypeData.Status != 0 {
			validCode = m.TypeData.Status
		}
		if res.StatusCode == validCode {
			m.State = monitors.Online.ToString()
		} else {
			m.State = monitors.Degraded.ToString()
		}
	}

	s.DB.Where("identifier = ?", m.Identifier).Updates(&m)
}
