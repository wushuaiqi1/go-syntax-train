package v2

import "net/http"

// RunHandler 启动处理
type RunHandler interface {
	Run(address string)
}

type RunHandlerImpl struct {
}

func NewRunHandler() RunHandler {
	return &RunHandlerImpl{}
}

// Run 启动Web项目
func (h *RunHandlerImpl) Run(address string) {
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err.Error())
	}
}
