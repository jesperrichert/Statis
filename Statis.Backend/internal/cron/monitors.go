package cron

import (
	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

type MonitorCronConfig struct {
	Scheduler gocron.Scheduler
	DB        *gorm.DB
	Jobs      map[string]*gocron.Job
}

func NewMonitorScheduler(db *gorm.DB) *MonitorCronConfig {
	s, _ := gocron.NewScheduler()
	return &MonitorCronConfig{
		Scheduler: s,
		DB:        db,
		Jobs:      map[string]*gocron.Job{},
	}
}

func (s *MonitorCronConfig) Schedule() {
	/*
		j, err := s.Scheduler.NewJob(
				gocron.DurationJob(
					10*time.Second,
				),
				gocron.NewTask(
					func() {
						// TODO: Add support for scheduling for monitors...
					},
				),
			)
	*/
}

func (s *MonitorCronConfig) Start() {
	println("[MONITOR CRON] Started!")
	s.Scheduler.Start()
}
