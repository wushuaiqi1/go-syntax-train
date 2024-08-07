package web

import "net/http"

type Server interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
	Run(address string) error
}

type HttpServer struct {
	name         string
	routeHandler *RouteHandler
}

func NewHttpServer(name string) Server {
	server := new(HttpServer)
	server.name = name
	server.routeHandler = NewRouterHandler()
	return server
}

func (h *HttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	h.routeHandler.AddRoute(method, pattern, handleFunc)
}

// BuildHandler 构建请求上下文处理
func buildHandler(method string, pattern string, handleFunc func(ctx *Context)) {
	// TODO 为什么外部变量method数据状态可以维护，闭包特性可以维护上下文访问的数据，直到闭包销毁
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		// 创建调用上下文
		ctx := NewContext(writer, request)
		if request.Method != method {
			err := ctx.OfFail("Method not Allowed")
			if err != nil {
				return
			}
			return
		}
		// 触发外部传递进来的匿名函数，即业务方法
		handleFunc(ctx)
	})
}

func (h *HttpServer) Run(address string) error {
	http.Handle("/", h.routeHandler)
	return http.ListenAndServe(address, nil)
}

type Header map[string][]string
