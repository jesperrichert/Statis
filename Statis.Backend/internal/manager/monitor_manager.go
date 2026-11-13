package manager

import (
	"jespersen.zip.statis/internal/manager/monitor/checks"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
)

type MonitorManager struct {
	Checker map[monitors.MonitorType]checks.ICheck
}

func NewMonitorManager(checker map[monitors.MonitorType]checks.ICheck) *MonitorManager {
	return &MonitorManager{
		Checker: checker,
	}
}

func (m *MonitorManager) Check(monitor *model.Monitor[any]) {
	m.Checker[monitors.ToType(monitor.Type)].Check(monitor)
}
