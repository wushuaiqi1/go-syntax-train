package v1

import (
	"fmt"
	"go-syntax-train/utils"
)

// RouteHandler 路由解析和处理
type RouteHandler interface {
	RegisterRoute(method string, pattern string, handleFunc func(ctx *Context))
	GetFuncByRoute(method string, pattern string) func(ctx *Context)
}

type RouteHandlerMap struct {
	routes map[string]func(ctx *Context)
}

func NewRouteHandlerMap() RouteHandler {
	return &RouteHandlerMap{
		routes: make(map[string]func(ctx *Context)),
	}
}

// RegisterRoute 注册路由
func (h *RouteHandlerMap) RegisterRoute(method string, pattern string, handle func(ctx *Context)) {
	uri := utils.GetUri(method, pattern)
	_, ok := h.routes[uri]
	if ok {
		panic(fmt.Sprintf("Route repeat add:%s", uri))
	}
	h.routes[uri] = handle
}

// GetFuncByRoute 获取业务方法
func (h *RouteHandlerMap) GetFuncByRoute(method string, pattern string) func(ctx *Context) {
	uri := utils.GetUri(method, pattern)
	if routeFunc, ok := h.routes[uri]; ok {
		return routeFunc
	}
	return nil
}
