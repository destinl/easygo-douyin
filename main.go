package main

import (
	"easygo-douyin/config"
	"easygo-douyin/router"
	"fmt"
)

func main() {
	r := router.Init()
	err := r.Run(fmt.Sprintf(":%d", config.Global.Port))
	if err != nil {
		return
	}
}
