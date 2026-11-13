package checks

import (
	"encoding/json"

	"jespersen.zip.statis/internal/model"
)

type ICheckBase interface {
	CheckAny(monitor model.Monitor[any])
}

type ICheck[T any] interface {
	Check(monitor model.Monitor[T])
}

type CheckAdapter[T any] struct {
	ICheckBase
	inner ICheck[T]
}

func NewCheckAdapter[T any](inner ICheck[T]) ICheckBase {
	return &CheckAdapter[T]{inner: inner}
}

func (ca *CheckAdapter[T]) CheckAny(monitor model.Monitor[any]) {
	var bytes []byte
	bytes, _ = json.Marshal(monitor)
	var m model.Monitor[T]
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return
	}
	ca.inner.Check(m)
}
