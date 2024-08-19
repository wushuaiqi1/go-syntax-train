package base

import (
	"fmt"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	Go()
}

func Go() {
	defer func() {
		fmt.Println("defer")
	}()
	// 不影响主进程
	go func() {
		defer fmt.Println("111")
		time.Sleep(10 * time.Second)
	}()
}
