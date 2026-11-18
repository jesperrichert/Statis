package schedule

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/manager"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
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

func (s *MonitorScheduler) Setup() *MonitorScheduler {
	var m []model.MonitorModel
	s.DB.Find(&m)
	for i := range m {
		monitor := model.Monitor[any](m[i])
		s.ScheduleMonitor(&monitor)
	}

	println("[MONITOR CRON] Started!")
	s.Scheduler.Start()

	return s
}

func (s *MonitorScheduler) ScheduleMonitor(monitor *model.Monitor[any]) {

	s.newJob(monitor)
}

func (s *MonitorScheduler) Key(monitor *model.Monitor[any]) string {
	return string(rune(monitor.ID)) + monitor.Identifier
}

func (s *MonitorScheduler) newJob(m *model.Monitor[any]) {
	monitor := m

	if s.Jobs[s.Key(m)] != nil {
		s.Scheduler.RemoveJob(s.Jobs[s.Key(m)].ID())
		s.Jobs[s.Key(m)] = nil
	}

	if monitors.ToType(monitor.Type) == monitors.Group ||
		monitors.ToType(monitor.Type) == monitors.Unkown {
		return

	}

	key := s.Key(monitor)
	j, _ := s.Scheduler.NewJob(
		gocron.DurationJob(
			time.Duration(monitor.ScheduleTime)*time.Second,
		),
		gocron.NewTask(
			func(monitorIdentifier string) {
				var monitorEntity model.MonitorModel
				s.DB.Where("identifier = ?", monitorIdentifier).First(&monitorEntity)
				me := model.Monitor[any](monitorEntity)

				if util.CheckNull(monitorEntity.Identifier) {
					_ = s.Scheduler.RemoveJob(s.Jobs[key].ID())
					return
				}
				s.MonitorManager.Check(&me)
			},
			monitor.Identifier,
		),
	)

	s.Jobs[key] = j
}
