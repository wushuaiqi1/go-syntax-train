package base

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// LoginParams 登录参数
type LoginParams struct {
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
}

// TestJson 测试JSON转换
func TestJson(t *testing.T) {
	params := LoginParams{
		Account:  "123",
		Password: "wushuaiqi1123",
	}
	marshal, err := json.Marshal(params)
	if err != nil {
		return
	}
	log.Println(string(marshal))
	var res LoginParams
	err = json.Unmarshal(marshal, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(res.Account)
	log.Println(res.Password)
}
