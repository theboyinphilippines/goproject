package admin

import (
	"beegoxiaomi/models"
	"strings"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdHeight = 100
	cpt.StdHeight = 40

}

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {

	c.TplName = "admin/login/login.html"
}

func (c *LoginController) DoLogin() {
	//1. 验证用户输入验证码是否正确
	flag := cpt.VerifyReq(c.Ctx.Request)
	if flag {
		//2. 获取表单传过来的用户名和密码, 去掉首尾空格

		username := strings.Trim(c.GetString("username"), "")
		password := models.Md5(strings.Trim(c.GetString("password"), ""))

		//3. 去数据库匹配
		manager := []models.Manager{}
		models.DB.Where("username= ? and password=? ", username, password).Find(&manager)
		if len(manager) > 0 {
			// 登录成功 1.保存用户信息  2. 跳转到后台系统
			c.SetSession("userinfo", manager[0])

			c.Success("登录成功", "/")

		} else {
			c.Error("用户名或密码错误", "/login")
		}

	} else {
		c.Error("验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登录成功", "/login")

}
