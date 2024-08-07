package v1

import "net/http"

// RequestHandler 请求处理
type RequestHandler interface {
	ConcreteEmit(ctx *Context, routeFunc func(ctx *Context))
	Bind(handler http.Handler)
}

func NewRequestHandler() RequestHandler {
	return &RequestHandlerImpl{}
}

type RequestHandlerImpl struct {
}

// ConcreteEmit 具体业务方法处理
func (h *RequestHandlerImpl) ConcreteEmit(ctx *Context, routeFunc func(ctx *Context)) {
	routeFunc(ctx)
}

// Bind 绑定全局请求解析
func (h *RequestHandlerImpl) Bind(handler http.Handler) {
	http.Handle("/", handler)
}
