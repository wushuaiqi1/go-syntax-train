package web

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
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

func (h *HttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}

func (h *HttpServer) Run(address string) error {
	return http.ListenAndServe(address, nil)
}

type Header map[string][]string
