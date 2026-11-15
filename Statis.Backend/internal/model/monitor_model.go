package model

import (
	"strconv"
	"strings"
)

type MonitorModel Monitor[any]

func (MonitorModel) TableName() string { return "monitors" }

type Monitor[T any] struct {
	ID                 int `gorm:"primaryKey"`
	Identifier         string
	IsParent           bool
	Group              string
	Name               string
	Description        string
	Type               string // MonitorType
	ScheduleTime       int
	TypeData           T                 `gorm:"serializer:json"`
	State              string            // MonitorState
	Events             []*EventModel     `gorm:"many2many:monitor_events;"`
	AffectedByExternal map[string]string `gorm:"serializer:json"`
	ExternalUrl        string
}

func (m *Monitor[T]) ToModel() MonitorModel {
	return MonitorModel{
		ID:                 m.ID,
		Identifier:         m.Identifier,
		IsParent:           m.IsParent,
		Group:              m.Group,
		Name:               m.Name,
		Description:        m.Description,
		Type:               m.Type,
		ScheduleTime:       m.ScheduleTime,
		TypeData:           m.TypeData,
		State:              m.State,
		Events:             m.Events,
		AffectedByExternal: m.AffectedByExternal,
		ExternalUrl:        m.ExternalUrl,
	}
}

func (m MonitorModel) Placeholders() map[string]string {

	return map[string]string{
		"monitor.identifier":    m.Identifier,
		"monitor.name":          m.Name,
		"monitor.description":   m.Description,
		"monitor.group":         m.Group,
		"monitor.type":          m.Type,
		"monitor.state":         m.State,
		"monitor.schedule_time": strconv.Itoa(m.ScheduleTime),
		"monitor.is_parent":     strconv.FormatBool(m.IsParent),
		"monitor.external_url":  m.ExternalUrl,
		"monitor.affected_by":   m.formatAffectedBy(),
	}
}

func (m MonitorModel) formatAffectedBy() string {
	var s []string
	for k, _ := range m.AffectedByExternal {
		s = append(s, k)
	}
	return strings.Join(s, ", ")
}
