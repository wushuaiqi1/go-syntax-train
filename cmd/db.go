package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect(name string) *gorm.DB {
	// dsn 格式 username:password@tcp(localhost:1)/db_name %2F代表/正斜杠
	open, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:1)/" + name + "?parseTime=True&loc=Asia%2FShanghai"))
	if err != nil {
		panic(err)
	}
	return open
}

func ConnectTest() *gorm.DB {
	return connect("test_docker")
}
