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

func (s *WebhookScheduler) Setup() {
	var webhooks []model.Webhook
	s.DB.Find(&webhooks)
	for i := range webhooks {
		s.ScheduleWebhook(webhooks[i])
	}

	println("[WEBHOOK CRON] Started!")
	s.Scheduler.Start()
}

func (s *WebhookScheduler) ScheduleWebhook(webhook model.Webhook) {
	var m []model.Monitor
	s.DB.Find(&m)
	for i := range m {
		s.newJob(webhook, m[i])
	}
}

func (s *WebhookScheduler) ScheduleMonitor(monitor model.Monitor) {
	var webhooks []model.Webhook
	s.DB.Find(&webhooks)
	for i := range webhooks {
		s.newJob(webhooks[i], monitor)
	}
}

func (s *WebhookScheduler) newJob(webhook model.Webhook, monitor model.Monitor) {
	key := string(rune(webhook.ID)) + webhook.Identifier + monitor.Identifier
	j, _ := s.Scheduler.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			func(webhookIdentifier string, monitorIdentifier string) {
				var monitorEntity model.Monitor
				var webhookEntity model.Webhook
				s.DB.Where("identifier = ?", monitorIdentifier).First(&monitorEntity)
				s.DB.Where("identifier = ?", webhookIdentifier).First(&webhookEntity)

				if s.Jobs[key].State != monitors.ToState(monitorEntity.State) {
					s.Jobs[key] = WebhookJob{
						Job:   s.Jobs[key].Job,
						State: monitors.ToState(monitorEntity.State),
					}

					reader := bytes.NewReader([]byte(webhookEntity.JSONTemplate))
					post, _ := http.Post(webhookEntity.Url, "application/json", reader)
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
