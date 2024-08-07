package web

import (
	"fmt"
	"net/http"
)

type Handler interface {
	http.Handler
	RouteHandler
}

type RouteHandler interface {
	RegisterRoute(method string, pattern string, handleFunc func(ctx *Context))
}

type RouteHandlerMap struct {
	routes map[string]func(ctx *Context)
}

func NewRouterHandlerMap() *RouteHandlerMap {
	return &RouteHandlerMap{
		routes: make(map[string]func(ctx *Context), 4),
	}
}

// HandleHttp 解析http请求
func (h *RouteHandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *RouteHandlerMap) RegisterRoute(method string, pattern string, handle func(ctx *Context)) {
	uri := h.getUri(method, pattern)
	_, ok := h.routes[uri]
	if ok {
		panic(fmt.Sprintf("Route repeat add:%s", uri))
	}
	h.routes[uri] = handle
}

func (h *RouteHandlerMap) getUri(method string, pattern string) (res string) {
	return method + "#" + pattern
}
