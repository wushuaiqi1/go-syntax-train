package main

import (
	"flag"
	"fmt"
	"go-syntax-train/configs"
	"go-syntax-train/configs/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 需要注意main函数的执行顺序
	flag.Parse()
	configs.LoadTomlConfig()
	configs.LoadJaegerConfig()
	configs.LoadSentryConfig()
	routers.Init()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
	ProgramExitAction()
}

func ProgramExitAction() {
	fmt.Println("阻塞式处理程序退出动作～")
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case tmp := <-exitSignal:
		switch tmp {
		case syscall.SIGHUP:
			log.Println("挂起信号，一般继续处理")
		case syscall.SIGQUIT:
			log.Println("退出信号，Ctrl+\\触发")
		case syscall.SIGTERM:
			log.Println("终止信号")
		case syscall.SIGINT:
			log.Println("中断信号,Ctrl+C触发")
		default:
			log.Println("默认不处理")
			return
		}
	}
}
