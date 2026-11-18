package schedule

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
)

type WebhookScheduler struct {
	Scheduler gocron.Scheduler
	DB        *gorm.DB
	Jobs      map[string]WebhookJob
}

type WebhookJob struct {
	Job   gocron.Job
	State monitors.MonitorState
}

func NewWebhookScheduler(db *gorm.DB) *WebhookScheduler {
	s, _ := gocron.NewScheduler()
	return &WebhookScheduler{
		Scheduler: s,
		DB:        db,
		Jobs:      map[string]WebhookJob{},
	}
}

func (s *WebhookScheduler) Setup() *WebhookScheduler {
	var webhooks []model.WebhookModel
	s.DB.Find(&webhooks)
	for i := range webhooks {
		s.ScheduleWebhook(&webhooks[i])
	}

	println("[WEBHOOK CRON] Started!")
	s.Scheduler.Start()

	return s
}

func (s *WebhookScheduler) ScheduleWebhook(webhook *model.WebhookModel) {
	var monitorModels []model.MonitorModel
	s.DB.Find(&monitorModels)
	for i := range monitorModels {
		monitor := model.Monitor[any](monitorModels[i])
		s.newJob(webhook, &monitor)
	}
}

func (s *WebhookScheduler) ScheduleMonitor(monitor *model.Monitor[any]) {
	var webhooks []model.WebhookModel
	s.DB.Find(&webhooks)
	for i := range webhooks {
		s.newJob(&webhooks[i], monitor)
	}
}

func (s *WebhookScheduler) Key(webhook *model.WebhookModel, monitor *model.Monitor[any]) string {
	return string(rune(webhook.ID)) + webhook.Identifier + ":" + monitor.Identifier
}

func (s *WebhookScheduler) newJob(webhook *model.WebhookModel, monitor *model.Monitor[any]) {
	key := s.Key(webhook, monitor)

	if s.Jobs[key].Job != nil {
		s.Scheduler.RemoveJob(s.Jobs[key].Job.ID())
		s.Jobs[key] = WebhookJob{}
	}

	j, _ := s.Scheduler.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			func(webhookIdentifier string, monitorIdentifier string) {
				var monitorEntity model.MonitorModel
				var webhookEntity model.WebhookModel
				s.DB.Where("identifier = ?", monitorIdentifier).First(&monitorEntity)
				s.DB.Where("identifier = ?", webhookIdentifier).First(&webhookEntity)

				if util.CheckNull(monitorEntity.Identifier) ||
					util.CheckNull(webhookEntity.Identifier) ||
					util.CheckNull(webhookEntity.Url) {
					_ = s.Scheduler.RemoveJob(s.Jobs[key].Job.ID())
					return
				}

				if s.Jobs[key].State != monitors.ToState(monitorEntity.State) {
					s.Jobs[key] = WebhookJob{
						Job:   s.Jobs[key].Job,
						State: monitors.ToState(monitorEntity.State),
					}

					json := util.ReplacePlaceholders(
						webhookEntity.JSONTemplate,
						monitorEntity.Placeholders(),
					)

					reader := bytes.NewReader([]byte(json))
					post, err := http.Post(webhookEntity.Url, "application/json", reader)
					if err != nil {
						return
					}
					fmt.Println("Webhook ("+webhookEntity.Identifier+") has been send.. Code:", post.StatusCode)
				}
			},
			webhook.Identifier,
			monitor.Identifier,
		),
	)

	s.Jobs[key] = WebhookJob{
		Job:   j,
		State: monitors.ToState(monitor.State),
	}
}
