package auto_config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/services"
)

type AutoConfigureScheduler struct {
	WebhookService  *services.WebhookService
	MonitorService  *services.MonitorService
	IncidentService *services.IncidentService
	EventService    *services.EventService

	Scheduler gocron.Scheduler
	DB        *gorm.DB
	Jobs      map[string]gocron.Job
}

func NewAutoConfigureScheduler(
	db *gorm.DB,
	webhookService *services.WebhookService,
	monitorService *services.MonitorService,
	incidentService *services.IncidentService,
	eventService *services.EventService,
) *AutoConfigureScheduler {
	s, _ := gocron.NewScheduler()
	return &AutoConfigureScheduler{
		Scheduler:       s,
		DB:              db,
		WebhookService:  webhookService,
		MonitorService:  monitorService,
		IncidentService: incidentService,
		EventService:    eventService,
		Jobs:            map[string]gocron.Job{},
	}
}

func (s *AutoConfigureScheduler) Setup() *AutoConfigureScheduler {

	monitors, mOk := os.LookupEnv("MONITORS")
	if mOk {
		s.newJob(monitors, "MONITORS")
	}
	events, eOk := os.LookupEnv("EVENTS")
	if eOk {
		s.newJob(events, "EVENTS")
	}
	incidences, iOk := os.LookupEnv("INCIDENCES")
	if iOk {
		s.newJob(incidences, "INCIDENCES")
	}
	webhooks, wOk := os.LookupEnv("WEBHOOKS")
	if wOk {
		s.newJob(webhooks, "WEBHOOKS")
	}

	println("[AUTO_CONFIG CRON] Started!")
	s.Scheduler.Start()

	return s
}

func (s *AutoConfigureScheduler) newJob(jsonUrl string, jobType string) {
	var fun func()
	switch jobType {
	case "MONITORS":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, s.MonitorService.Validate, s.MonitorService.Add, s.MonitorService.Update)
			}
			break
		}
	case "EVENTS":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, s.EventService.Validate, s.EventService.Add, s.EventService.Update)
			}
			break
		}
	case "INCIDENCES":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, s.IncidentService.Validate, s.IncidentService.Add, s.IncidentService.Update)
			}
			break
		}
	case "WEBHOOKS":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, s.WebhookService.Validate, s.WebhookService.Add, s.WebhookService.Update)
			}
			break
		}
	default:
		{
			return
		}
	}

	seconds, ok := os.LookupEnv("AUTO_CONFIG_INTERVAL")
	if !ok {
		println("[AUTO CONFIG] Failed to start...")
		return
	}
	i, err := strconv.Atoi(seconds)
	if err != nil {
		return
	}

	j, _ := s.Scheduler.NewJob(
		gocron.DurationJob(
			time.Duration(i)*time.Second,
		),
		gocron.NewTask(
			func() {
				fun()
			},
		),
	)

	s.Jobs[jobType] = j
}

func (s *AutoConfigureScheduler) autoConfigure[DTO any](
	configUrl string,
	validateFunc func(dto *DTO) (int, string),
	addFunc func(dto *DTO) int,
	updateFunc func(identifier string, dto *DTO) int,
) {
	var dtos []DTO
	configData, _ := http.Get(configUrl)
	body, _ := io.ReadAll(configData.Body)
	_ = json.Unmarshal(body, &dtos)

	for i := range dtos {
		data := dtos[i]

		status, message := validateFunc(&data)
		switch status {
		case http.StatusConflict:
			refData := reflect.ValueOf(data)
			identifier := refData.FieldByName("Identifier").String()

			status = updateFunc(identifier, &data)
			if status == http.StatusNotFound {
				fmt.Println("[AUTO CONFIG] No data found for ")
				return
			}

			if status == http.StatusOK {
				fmt.Println("[AUTO CONFIG] Successfully Updated Auto Configure", configUrl)
			} else {
				fmt.Println("[AUTO CONFIG] Auto Configure Update failed with status", status, " | URL:", configUrl)
			}
			break
		case http.StatusOK:
			status = addFunc(&data)
			break
		case http.StatusBadRequest:
			fmt.Println(message)
			break
		}
	}
}
