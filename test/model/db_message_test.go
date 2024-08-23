package model

import (
	"go-syntax-train/cmd"
	"go-syntax-train/model"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
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
	start := time.Now().UnixMilli()
	db := cmd.ConnectTest()
	res := sync.WaitGroup{}
	for i := 1; i <= 151; i++ {
		res.Add(1)
		go func(i int) {
			message := model.NewMessage("我是消息"+strconv.Itoa(i), 1)
			db.Create(&message)
			res.Add(-1)
		}(i)
	}
	// 阻塞式等待
	res.Wait()
	end := time.Now().UnixMilli()
	log.Println("批量插入数据耗时：", end-start, " ms")
}
