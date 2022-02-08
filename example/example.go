package example

import (
	"fmt"
	"github.com/shikingram/ci-taskmodel/model"
)

type Deployment struct {
	model.IBaseModel
	name string
}

func NewDeployment(m model.IBaseModel) *Deployment {
	return &Deployment{
		IBaseModel: m,
	}
}

func (d *Deployment) Handle(action string) error {
	// do something
	fmt.Println(action)
	return nil
}

func (d *Deployment) Callback(requestId, fromId, fromType, targetType string) error {
	// do something
	return nil
}

func (d *Deployment) DirRun() bool {
	return false
}
