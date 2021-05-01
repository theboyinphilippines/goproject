package admin

import (
	"beegoxiaomi/models"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.TplName = "admin/role/index.html"
}

func (c *RoleController) Add() {
	c.TplName = "admin/role/add.html"
}

func (c *RoleController) DoAdd() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		// 此时用return，后面代码将不再执行
		return
	}
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())

	err := models.DB.Create(&role).Error
	if err != nil {
		c.Error("增加角色失败", "/role/add")
	} else {
		c.Success("增加角色成功", "/role")
	}

}

func (c *RoleController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	c.Data["role"] = role

	c.TplName = "admin/role/edit.html"
}

func (c *RoleController) DoEdit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}

	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	if title == "" {
		c.Error("标题不能为空", "/role/add")
		// 此时用return，后面代码将不再执行
		return
	}

	//修改
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	err2 := models.DB.Save(&role).Error
	if err2 != nil {
		c.Error("修改角色失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改角色成功", "/role")
	}

}

func (c *RoleController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}

	//删除
	role := models.Role{Id: id}
	models.DB.Delete(&role)
	c.Success("删除角色成功", "/role")

}

func (c *RoleController) Auth() {
	//1.获取角色id 在模版中提交时，需要传入角色id
	roleId, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}

	//2.获取全部权限
	access := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	c.Data["accessList"] = access
	c.Data["roleId"] = roleId

	//3.获取当前角色拥有的权限，并把权限id房放在一个map里面
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Find(&roleAccess)

	roleAccessMap := make(map[int]int)

	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	//4.循环遍历所有的权限数据，判断当前权限是否在角色权限map对象中，如果是的话给当前数据加入checked属性

	//先遍历模块，再遍历模块下的菜单，操作
	for i := 0; i < len(access); i++ {

		if _, ok := roleAccessMap[access[i].Id]; ok {
			access[i].Checked = true
		}
		for j := 0; j < len(access[i].AccessItem); j++ {
			if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
				access[i].AccessItem[j].Checked = true
			}
		}
	}
	//5、渲染权限数据以及角色 Id
	c.Data["accessList"] = access
	c.Data["roleId"] = roleId

	c.TplName = "admin/role/auth.html"

}

func (c *RoleController) DoAuth() {
	//1.获取参数post传过来的角色id 和 权限切片
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	//接收权限id
	accessNode := c.GetStrings("access_node")

	//2.修改角色权限---删除当前角色下面的所有权限(因为是多对多的中间表)
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id = ?", roleId).Delete(&roleAccess)

	//3.执行增加数据
	//权限id是字符串切片，需要遍历
	for _, v := range accessNode {
		accessId, _ := strconv.Atoi(v)
		roleAccess.RoleId = roleId
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)

	}
	c.Success("授权成功", "/role/auth?id="+strconv.Itoa(roleId))

}
