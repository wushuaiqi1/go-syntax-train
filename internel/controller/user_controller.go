package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-syntax-train/internel/service"
	"net/http"
	"time"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// traceId接入
	service.TraceMark("获取用户信息", map[any]any{
		"traceId": uuid.New().String(),
	})

	time.Sleep(2 * time.Second)
	//if true {
	//	panic("hello")
	//}

	resp.WriteHeader(http.StatusOK)
	data := map[string]any{
		"name": "张三",
		"age":  19,
		"addr": "昌平区",
	}
	res, _ := json.Marshal(data)
	_, err := resp.Write(res)
	if err != nil {
		return
	}
}
