package configs

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var path string = ""

func init() {
	// 先定义要解析的东西
	flag.StringVar(&path, "config", "", "invalid or missing config file")
}

func loadConfigByDefaultPath() {
	viper.SetConfigName("configs")
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
		currentWorkDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		path = currentWorkDir + "/configs/config.toml"
	}
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	port := viper.GetString("server.name")
	log.Println("load viper config success,the server name is", port)
}

func LoadTomlConfig() {
	// 加载配置
	loadConfigByCommandLineParams()
}
