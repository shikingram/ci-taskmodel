package model

import "github.com/shikingram/ci-taskmodel/types"

/*
IBaseHandler is a base interface of handler，all handler should implement those functions
*/
type IBaseHandler interface {
	Handle(action string) error
	Callback(requestId, fromId, fromType, targetType string) error
}

/*
IBaseModel is a base interface of model, all handler should contain those fields
*/
type IBaseModel interface {
	SetCilDeployChain(traceId, action, parentUniqId, uniqId string)
	SetRetryNum(RetryNum int64)
	SetAction(action string)
	CilDeployChain() *types.CilDeployChain
	Action() string
	RetryNum() int64
}

type baseModelImpl struct {
	retryNum       int64
	action         string
	cilDeployChain *types.CilDeployChain
	//todo not export
	Details types.Details
}

var _ IBaseModel = &baseModelImpl{}

func (d *baseModelImpl) RetryNum() int64 {
	return d.retryNum
}

func (d *baseModelImpl) Action() string {
	return d.action
}

func (d *baseModelImpl) CilDeployChain() *types.CilDeployChain {
	return d.cilDeployChain
}

func NewBaseModel() IBaseModel {
	return &baseModelImpl{
		cilDeployChain: &types.CilDeployChain{},
		Details:        types.Details{},
	}
}

func (d *baseModelImpl) SetCilDeployChain(traceId, action, parentUniqId, uniqId string) {
	d.cilDeployChain = &types.CilDeployChain{
		TraceId:      traceId,
		Action:       action,
		ParentUniqid: parentUniqId,
		Uniqid:       uniqId,
	}
}

func (d *baseModelImpl) SetAction(action string) {
	d.action = action
}

func (d *baseModelImpl) SetRetryNum(retryNum int64) {
	d.retryNum = retryNum
}
