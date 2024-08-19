package base

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string, 1)
	go func() {
		ch <- "ssks"
		time.Sleep(time.Second)
		ch <- "hello,world"
	}()
	time.Sleep(time.Second)
	msg := <-ch
	fmt.Println(msg)
	fmt.Println(time.Now().UnixMilli())
	msg = <-ch
	fmt.Println(time.Now().UnixMilli())
}

func TestSelect(t *testing.T) {
	ch1 := make(chan string, 1)
	go func() {
		ch1 <- "hello,ch1"
	}()
	ch2 := make(chan string, 2)
	go func() {
		ch2 <- "hello,ch2"
	}()
	select {
	case msg := <-ch2:
		{
			fmt.Println(msg)
			return
		}
	case msg := <-ch1:
		{
			//msg := <-ch1
			fmt.Println(msg)
			return
		}
	}
}
