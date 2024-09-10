package model

import (
	"fmt"
	"go-syntax-train/config"
	"testing"
)

type User struct {
	Id      int `gorm:"user"`
	Name    string
	Age     int
	Address string
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}

func TestCreate(t *testing.T) {
	db := config.ConnectTest()
	db.Create(&User{
		Name:    "武帅祺",
		Age:     20,
		Address: "河北省邯郸市广平县",
	})
}

func TestAscFirst(t *testing.T) {
	db := config.ConnectTest()
	u1 := User{}
	db.First(&u1)
	fmt.Println(u1)
	u1 = User{}
	db.Take(&u1, "id", 2)
	fmt.Println(u1)
}

func TestFind(t *testing.T) {
	db := config.ConnectTest()
	users := make([]User, 0)
	db.Limit(1).Find(&users)
	fmt.Println(users)
}
