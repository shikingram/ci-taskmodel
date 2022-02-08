package model

import "github.com/shikingram/ci-taskmodel/types"

type ITaskHandler interface {
	Handle()
	GetSelf() map[string]interface{}
	GetState() string
	SetCilDeployChain(traceId, action, parentUniqId, uniqId string)
	GetDeployType() string
	GetUniqueId() string
	GetAction() string
}

/*
IBaseHandler 独立函数实现
*/
type IBaseHandler interface {
	Handle(action string) error
	Callback(requestId, fromId, fromType, targetType string) error
}

/*
IBaseModel 公用的model
*/
type IBaseModel interface {
	SetCilDeployChain(traceId, action, parentUniqId, uniqId string)
	SetRetryNum(RetryNum int64)
	SetAction(action string)
	CilDeployChain() *types.CilDeployChain
	Action() string
	RetryNum() int64
}

type handlerImpl struct {
	retryNum       int64
	action         string
	cilDeployChain *types.CilDeployChain
	//todo 不导出
	Details types.Details
}

func (d *handlerImpl) RetryNum() int64 {
	return d.retryNum
}

func (d *handlerImpl) Action() string {
	return d.action
}

func (d *handlerImpl) CilDeployChain() *types.CilDeployChain {
	return d.cilDeployChain
}

func NewIHandler() IBaseModel {
	return &handlerImpl{
		cilDeployChain: &types.CilDeployChain{},
		Details:        types.Details{},
	}
}

func (d *handlerImpl) SetCilDeployChain(traceId, action, parentUniqId, uniqId string) {
	d.cilDeployChain = &types.CilDeployChain{
		TraceId:      traceId,
		Action:       action,
		ParentUniqid: parentUniqId,
		Uniqid:       uniqId,
	}
}

func (d *handlerImpl) SetAction(action string) {
	d.action = action
}

func (d *handlerImpl) SetRetryNum(retryNum int64) {
	d.retryNum = retryNum
}
