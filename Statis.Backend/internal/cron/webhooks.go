package cron

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

type WebhookCronConfig struct {
	Scheduler gocron.Scheduler
	DB        *gorm.DB
	Jobs      map[string]WebhookJob
}

type WebhookJob struct {
	Job   gocron.Job
	State monitors.MonitorState
}

func NewWebhookScheduler(db *gorm.DB) *WebhookCronConfig {
	s, _ := gocron.NewScheduler()
	return &WebhookCronConfig{
		Scheduler: s,
		DB:        db,
		Jobs:      map[string]WebhookJob{},
	}
}

func (s *WebhookCronConfig) Schedule() {
	var webhooks []model.Webhook
	s.DB.Find(&webhooks)
	for i := range webhooks {
		webhook := webhooks[i]

		var m []model.Monitor
		s.DB.Find(&m)

		for i := range m {
			monitor := m[i]
			s.AddJob(webhook, monitor)
		}
	}
}

func (s *WebhookCronConfig) AddJob(webhook model.Webhook, monitor model.Monitor) {
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

func (s *WebhookCronConfig) Start() {
	println("[WEBHOOK CRON] Started!")
	s.Scheduler.Start()
}
