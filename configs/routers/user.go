package routers

import (
	"fmt"
	"go-syntax-train/configs"
	"go-syntax-train/internel/controller"
)

func init() {
	fmt.Println("user route init...")
	configs.RegisterRouteBindSentry("/user", controller.GetUser)
}
