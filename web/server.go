package web

import "net/http"

type Server interface {
	Route(pattern string, handleFunc func(ctx *Context))
	Run(address string) error
}

type HttpServer struct {
	Name string
}

func NewHttpServer(name string) Server {
	server := new(HttpServer)
	server.Name = name
	return server
}

func (h *HttpServer) Route(pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		// 创建调用上下文
		ctx := NewContext(writer, request)
		// 触发外部传递进来的匿名函数
		handleFunc(ctx)
	})
}

func (h *HttpServer) Run(address string) error {
	return http.ListenAndServe(address, nil)
}

type Header map[string][]string
