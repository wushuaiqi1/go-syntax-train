package base

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	res := 1
	if res == 1 {
		err := fmt.Errorf("res==1是真的,%d", res)
		fmt.Println(err)
	}
}

func TestPanic(t *testing.T) {
	flag := false
	defer func() {
		if has := recover(); has != nil {
			fmt.Printf("处理panic:%s\n", has)
		}
		fmt.Println("recover...")
	}()
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
		} else {
			flag = true
		}
		if flag {
			panic("触发panic")
		}
	}
}

func TestDelay(t *testing.T) {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		// 闭包拿到的是指向i的地址，延时绑定
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}
	for _, val := range fns {
		val()
	}
}
