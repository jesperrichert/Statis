package schedule

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/manager"
	"jespersen.zip.statis/internal/model"
)

type MonitorScheduler struct {
	Scheduler      gocron.Scheduler
	DB             *gorm.DB
	MonitorManager *manager.MonitorManager
	Jobs           map[string]gocron.Job
}

func NewMonitorScheduler(db *gorm.DB, monitorManager *manager.MonitorManager) *MonitorScheduler {
	s, _ := gocron.NewScheduler()
	return &MonitorScheduler{
		Scheduler:      s,
		DB:             db,
		MonitorManager: monitorManager,
		Jobs:           map[string]gocron.Job{},
	}
}

func (s *MonitorScheduler) Setup() {
	var m []model.Monitor[any]
	s.DB.Find(&m)
	for i := range m {
		s.ScheduleMonitor(m[i])
	}

	println("[MONITOR CRON] Started!")
	s.Scheduler.Start()
}

func (s *MonitorScheduler) ScheduleMonitor(monitor model.Monitor[any]) {
	s.newJob(monitor)
}

func (s *MonitorScheduler) newJob(monitor model.Monitor[any]) {
	key := string(rune(monitor.ID)) + monitor.Identifier
	j, _ := s.Scheduler.NewJob(
		gocron.DurationJob(
			time.Duration(monitor.ScheduleTime)*time.Second,
		),
		gocron.NewTask(
			func(monitorIdentifier string) {
				var monitorEntity model.Monitor[any]
				s.DB.Where("identifier = ?", monitorIdentifier).First(&monitorEntity)
				s.MonitorManager.Check(&monitorEntity)
			},
			monitor.Identifier,
		),
	)

	s.Jobs[key] = j
}
