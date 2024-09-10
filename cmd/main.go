package main

import (
	"flag"
	"go-syntax-train/config"
)

func main() {
	// 需要注意main函数的执行顺序
	flag.Parse()
	config.LoadConfig()
}
