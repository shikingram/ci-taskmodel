package example

import (
	"fmt"
	"github.com/shikingram/ci-taskmodel/model"
)

type TaskDeployment struct {
	model.ITaskBaseModel
	name string
}

func NewTaskDeployment(m model.ITaskBaseModel) *TaskDeployment {
	return &TaskDeployment{
		ITaskBaseModel: m,
	}
}

func (d *TaskDeployment) Handle() {
	// do something
	fmt.Println(d.GetAction())
}

func (d *TaskDeployment) DirRun() bool {
	return false
}
