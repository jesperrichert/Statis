package config

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/cron"
)

type SchedulerConfig struct {
	WebHookScheduler *cron.WebhookCronConfig
	MonitorScheduler *cron.MonitorCronConfig
}

func Cron(db *gorm.DB) *SchedulerConfig {
	webhookScheduler := cron.NewWebhookScheduler(db)
	monitorScheduler := cron.NewMonitorScheduler(db)

	config := &SchedulerConfig{
		WebHookScheduler: webhookScheduler,
		MonitorScheduler: monitorScheduler,
	}

	config.MonitorScheduler.Schedule()
	config.WebHookScheduler.Schedule()

	config.MonitorScheduler.Start()
	config.WebHookScheduler.Start()

	return config
}
