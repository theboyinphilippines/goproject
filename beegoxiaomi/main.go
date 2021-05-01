package main

import (
	"beegoxiaomi/models"
	_ "beegoxiaomi/routers"

	"github.com/astaxie/beego"
)

func main() {
	//注册模版函数，在html中使用
	beego.AddFuncMap("unixToDate", models.UnixToDate)

	beego.Run()
}
