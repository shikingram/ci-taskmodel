package main

import (
	"ci-taskmodel/model"
	"fmt"
)

type Deployment struct {
	model.IBaseModel
	name string
}

func NewDeployment(handler model.IBaseModel) *Deployment {
	return &Deployment{
		IBaseModel: handler,
	}
}

func (d *Deployment) Handle(action string) error {
	fmt.Println(action)
	return nil
}

func (d *Deployment) Callback(requestId, fromId, fromType, targetType string) error {
	return nil
}

func (d *Deployment) SetModelPending() error {
	return nil
}

func (d *Deployment) DirRun() bool {
	return false
}

func main() {
	handler := model.NewIHandler()
	h := NewDeployment(handler)
	err := h.Handle("create")
	if err != nil {
		panic(err)
	}
	h.SetAction("update")
	fmt.Println(h.Action())
}
