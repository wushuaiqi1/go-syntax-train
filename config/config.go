package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var path string = ""

func init() {
	// 先定义要解析的东西
	flag.StringVar(&path, "config", "", "invalid or missing config file")
}

func loadConfigByDefaultPath() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/Users/tal/GolandProjects/go-syntax-train/")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in)
		fmt.Println(viper.GetInt("server.port"))
	})
	for {
		time.Sleep(10 * time.Second)
	}
}

func loadConfigByCommandLineParams() {
	if len(path) == 0 {
		panic("missing path")
	}
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	port := viper.GetString("server.port")
	fmt.Println("viper读取配置文件成功，端口号是", port)
}

func LoadConfig() {
	// 加载配置
	loadConfigByCommandLineParams()
}
