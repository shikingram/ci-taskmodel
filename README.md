# ci-taskmodel

worker公用模型包提取

## 安装

```
go get github.com/shikingram/ci-taskmodel
```

## 使用说明

### 定义结构体

嵌入IBaseModel

```golang
type Deployment struct {
	model.IBaseModel
	name string
}

func NewDeployment(m model.IBaseModel) *Deployment {
    return &Deployment{
    IBaseModel: m,
  }
}
```

### 注入实例

```golang
m := model.NewBaseModel()
h := NewDeployment(m)
err := h.Handle("create")

```
