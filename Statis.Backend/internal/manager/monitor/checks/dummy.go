package checks

import (
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
)

type DummyCheck struct {
	ICheck[monitors.DummyMonitor]
}

func NewDummyCheck() *DummyCheck {
	return &DummyCheck{}
}

func (s *DummyCheck) Check(monitor model.Monitor[monitors.DummyMonitor]) {

}
