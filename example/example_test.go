package example

import (
	"github.com/shikingram/ci-taskmodel/model"
	"testing"
)

func TestNewDeployment(t *testing.T) {
	m := model.NewBaseModel()
	h := NewDeployment(m)
	err := h.Handle("create")
	if err != nil {
		t.Fail()
	}

	h.SetAction("update")
	t.Log(h.Action())

	if h.Action() != "update" {
		t.Fail()
	}
}

func TestNewTaskDeployment(t *testing.T) {
	m := model.NewTaskBaseModel()
	h := NewTaskDeployment(m)
	h.SetMetadata(map[string]interface{}{
		"action": "create",
	})
	h.Handle()
}
