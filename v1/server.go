package v1

import (
	"net/http"
)

type Handler interface {
	Register(method string, pattern string, handle func(ctx *Context))
	Start(address string)
	http.Handler
}

type HttpServer struct {
	name string
	RequestHandler
	RouteHandler
	RunHandler
}

func NewHttpServer(name string) Handler {
	return &HttpServer{
		name:           name,
		RequestHandler: NewRequestHandler(),
		RouteHandler:   NewRouteHandlerMap(),
		RunHandler:     NewRunHandler(),
	}
}

// Register 注册动作
func (server *HttpServer) Register(method string, pattern string, handle func(ctx *Context)) {
	server.RegisterRoute(method, pattern, handle)
}

// ServeHTTP 实现原生http.Handler接口
func (server *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	funcByRoute := server.GetFuncByRoute(r.Method, r.URL.Path)
	if funcByRoute == nil {
		ctx.OfFail("404")
		return
	}
	// 执行业务方法
	funcByRoute(ctx)
}

// Start 启动动作
func (server *HttpServer) Start(address string) {
	server.Bind(server)
	server.Run(address)
}
