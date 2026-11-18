package schedule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

type AutoConfigureScheduler struct {
	Scheduler gocron.Scheduler
	DB        *gorm.DB
	Jobs      map[string]gocron.Job
}

func NewAutoConfigureScheduler(db *gorm.DB) *AutoConfigureScheduler {
	s, _ := gocron.NewScheduler()
	return &AutoConfigureScheduler{
		Scheduler: s,
		DB:        db,
		Jobs:      map[string]gocron.Job{},
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
				s.autoConfigure(jsonUrl, "http://localhost:3000/api/status/monitors", "http://localhost:3000/api/status/monitors/{identifier}")
			}
			break
		}
	case "EVENTS":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, "http://localhost:3000/api/status/events", "http://localhost:3000/api/status/events/{identifier}")
			}
			break
		}
	case "INCIDENCES":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, "http://localhost:3000/api/status/incidences", "http://localhost:3000/api/status/incidences/{identifier}")
			}
			break
		}
	case "WEBHOOKS":
		{
			fun = func() {
				s.autoConfigure(jsonUrl, "http://localhost:3000/api/notifications/webhooks", "http://localhost:3000/api/status/webhooks/{identifier}")
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

func (s *AutoConfigureScheduler) autoConfigure(configUrl, postUrl, updateUrl string) {
	authToken, ok := os.LookupEnv("AUTH_TOKEN")
	if !ok {
		println("[AUTO CONFIG] Failed to start no auth token")
		return
	}
	client := &http.Client{}

	var dtos []map[string]any

	configData, _ := http.Get(configUrl)
	body, _ := io.ReadAll(configData.Body)

	_ = json.Unmarshal(body, &dtos)
	for i := range dtos {
		b, _ := json.Marshal(dtos[i])

		req, _ := http.NewRequest("POST", postUrl, bytes.NewReader(b))
		req.Header.Add("Authorization", authToken)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("[AUTO_CONFIG CRON] Failed auto configure the webhook ", err.Error())
			return
		}
		if res.StatusCode == http.StatusConflict {
			var identifier string
			for k, v := range dtos[i] {
				if k == "identifier" {
					identifier = v.(string)
				}
			}
			req, _ = http.NewRequest("POST", strings.ReplaceAll(updateUrl, "{identifier}", identifier), bytes.NewReader(b))
			req.Header.Add("Authorization", authToken)
			_, _ = client.Do(req)
		}
	}
}
