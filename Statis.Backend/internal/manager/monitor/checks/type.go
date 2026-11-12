package checks

import "jespersen.zip.statis/internal/model"

type ICheck interface {
	Check(monitor *model.Monitor)
}
