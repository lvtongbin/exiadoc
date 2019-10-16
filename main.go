package main

import (
	_ "trism/longpoll"
	_ "trism/models"
	_ "trism/myredis"
	_ "trism/routers"
	_ "trism/statistics"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
