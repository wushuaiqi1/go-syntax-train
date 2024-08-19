package model

import (
	"go-syntax-train/cmd"
	"go-syntax-train/model"
	"strconv"
	"testing"
)

// 模型映射表
func TestAutoMigrate(t *testing.T) {
	db := cmd.ConnectTest()
	err := db.AutoMigrate(&model.Message{})
	if err != nil {
		return
	}
}

func TestBatchCreate(t *testing.T) {
	db := cmd.ConnectTest()
	for i := 0; i < 100; i++ {
		message := model.NewMessage("我是消息"+strconv.Itoa(i), 1)
		db.Create(&message)
	}
}
