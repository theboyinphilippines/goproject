package routers

import (
	"beegoxiaomi/controllers/index"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{})
	beego.Router("/login", &index.LoginController{})
}
