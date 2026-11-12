package checks

import (
	"jespersen.zip.statis/internal/model"
)

type DummyCheck struct {
	ICheck
}

func NewDummyCheck() *DummyCheck {
	return &DummyCheck{}
}

func (s *DummyCheck) Check(monitor *model.Monitor) {

}
