package config

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

func Cron(db *gorm.DB) {
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("[CRON] init scheduler error", err)
	}
	j, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				// TODO: Add support for scheduling for monitors...
				// TODO: Add webhook support for event publish
				// TODO: Add Config polling from JSON Endpoints
				// (Add Support for external data provider...)
			},
		),
	)
	if err != nil {
		fmt.Println("[CRON] init job error", err)
	}
	fmt.Println("[CRON] running job id", j.ID(), "for uptime checks.")
	s.Start()
}
