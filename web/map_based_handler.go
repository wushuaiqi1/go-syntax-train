package web

import (
	"fmt"
	"net/http"
)

// RouteHandler 路由解析器
type RouteHandler struct {
	routes map[string]func(ctx *Context)
}

func NewRouterHandler() *RouteHandler {
	return &RouteHandler{
		make(map[string]func(ctx *Context)),
	}
}

// HandleHttp 解析http请求
func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := h.getUri(r.Method, r.URL.Path)
	handle, ok := h.routes[uri]
	if !ok {
		_, err := w.Write([]byte("404 route not found"))
		if err != nil {
			return
		}
		return
	}
	ctx := NewContext(w, r)
	handle(ctx)
}

// AddRoute 添加路由
func (h *RouteHandler) AddRoute(method string, pattern string, handle func(ctx *Context)) {
	uri := h.getUri(method, pattern)
	_, ok := h.routes[uri]
	if ok {
		panic(fmt.Sprintf("Route repeat add:%s", uri))
	}
	h.routes[uri] = handle
}

func (h *RouteHandler) getUri(method string, pattern string) (res string) {
	return method + "#" + pattern
}
