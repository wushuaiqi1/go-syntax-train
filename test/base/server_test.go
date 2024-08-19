package base

import (
	v2 "go-syntax-train/v2"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	server := v2.NewHttpServer("test-server")
	server.Register(http.MethodGet, "/home/user/", Home)
	server.Register(http.MethodPost, "/hi", Hi)
	server.Register(http.MethodGet, "/home/hi", Hi)
	server.Start(":8080")
}

func Hi(ctx *v2.Context) {
	ctx.OfSuccess("hello")
}

func Home(ctx *v2.Context) {
	json, err := ctx.ReadJson()
	if err != nil {
		return
	}
	err = ctx.WriteJson(200, json)
	if err != nil {
		return
	}
}
