package model

import "github.com/shikingram/ci-taskmodel/types"

/*
ITaskBaseHandler is a base interface of handler，all task handler should implement those functions
*/
type ITaskBaseHandler interface {
	Handle()
}

/*
ITaskBaseModel is a base interface of model, all task handler should contain those fields
*/
type ITaskBaseModel interface {
	Metadata() map[string]interface{}
	SetMetadata(metadata map[string]interface{})
	CilDeployChain() *types.CilDeployChain
	GetSelf() map[string]interface{}
	GetState() string
	SetCilDeployChain(traceId, action, parentUniqId, uniqId string)
	GetDeployType() string
	GetUniqueId() string
	GetAction() string
}

type taskBaseModelImpl struct {
	metadata       map[string]interface{}
	cilDeployChain *types.CilDeployChain
}

var _ ITaskBaseModel = &taskBaseModelImpl{}

func NewTaskBaseModel() ITaskBaseModel {
	return &taskBaseModelImpl{
		cilDeployChain: &types.CilDeployChain{},
	}
}

func (t *taskBaseModelImpl) Metadata() map[string]interface{} {
	return t.metadata
}

func (t *taskBaseModelImpl) SetMetadata(metadata map[string]interface{}) {
	t.metadata = metadata
}

func (t *taskBaseModelImpl) CilDeployChain() *types.CilDeployChain {
	return t.cilDeployChain
}

func (t *taskBaseModelImpl) SetCilDeployChain(traceId, action, parentUniqId, uniqId string) {
	t.cilDeployChain = &types.CilDeployChain{
		TraceId:      traceId,
		Action:       action,
		ParentUniqid: parentUniqId,
		Uniqid:       uniqId,
	}
}

func (t *taskBaseModelImpl) GetSelf() map[string]interface{} {
	return t.metadata
}

func (t *taskBaseModelImpl) GetState() string {
	return getMustStringFromMap(t.metadata, "state")
}

func (t *taskBaseModelImpl) GetDeployType() string {
	return getMustStringFromMap(t.metadata, "deploy_type")
}

func (t *taskBaseModelImpl) GetUniqueId() string {
	return getMustStringFromMap(t.metadata, "uniqid")
}

func (t *taskBaseModelImpl) GetAction() string {
	return getMustStringFromMap(t.metadata, "action")
}

func getMustStringFromMap(m map[string]interface{}, key string) string {
	if m[key] == nil {
		return "unknown"
	} else {
		if v, ok := m[key].(string); ok {
			return v
		} else {
			return "unknown"
		}
	}
}
