package controllers

import (
	"fmt"
	"strconv"

	"git.lzxz1234.cn/lzxz1234/AdminBoot/models"
	"git.lzxz1234.cn/lzxz1234/AdminBoot/utils"
	"git.lzxz1234.cn/lzxz1234/AdminBoot/utils/aes"
	"github.com/astaxie/beego/logs"
)

// AuthUserController user api
type AuthUserController struct {
	BaseController
}

// Index 用户管理首页
// @router /
func (c *AuthUserController) Index() {
	c.TplName = "auth/userList.tpl"
}

// Add 添加新用户
// @router /add [get]
func (c *AuthUserController) Add() {

	var roles []*models.AuthRole
	o.QueryTable(&models.AuthRole{}).All(&roles)
	c.Data["roles"] = &roles
	c.TplName = "auth/userAdd.tpl"
}

// DoAdd 添加新用户
// @router /add [post]
func (c *AuthUserController) DoAdd() {

	user := models.AuthUser{}
	if err := c.ParseForm(&user); err != nil {
		logs.Info("表单解析失败", err)
	}

	if o.QueryTable("t_auth_user").Filter("username", user.UserName).Exist() {
		c.Data["ErrMsg"] = "用户名不能重复"
		c.Add()
		return
	}
	user.Password = aes.EncryptString(user.Password)
	if c.GetString("initState") != "on" {
		user.State = 1
	}
	user.Portrait = "/static/img/portrait.jpg"
	o.Insert(&user)
	roleIDs := c.GetStrings("roleIDs")
	if roleIDs != nil {
		for _, roleID := range roleIDs {
			o.Raw("insert into t_auth_role_user(role_id, user_id) values(?, ?)", roleID, user.ID).Exec()
		}
	}
	c.Redirect("/auth/user/", 302)
}

// List users
// @router /list
func (c *AuthUserController) List() {

	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("limit", 10)

	count, _ := o.QueryTable(&models.AuthUser{}).Count()
	var users []*models.AuthUser
	o.QueryTable(&models.AuthUser{}).Limit(size, page*size-size).All(&users)
	for _, user := range users {
		o.QueryTable("t_auth_role").
			FilterRaw("id", fmt.Sprintf("in (select role_id from t_auth_role_user where user_id=%d)", user.ID)).
			All(&user.Roles, "Name")
	}

	c.Data["json"] = utils.NewPage(int(count), page, size, users)
	c.ServeJSON()
}

// ModState 修改用户状态
// @router /modState [post]
func (c *AuthUserController) ModState() {

	id, _ := c.GetInt("id")
	user := models.AuthUser{ID: id}
	o.Read(&user)

	state, _ := c.GetInt("state", -1)
	if state >= 0 {
		user.State = state
	}
	o.Update(&user)
	c.Data["json"] = utils.NewResult(0, "操作成功", nil)
	c.ServeJSON()
}

// Mod 修改用户信息
// @router /mod [get]
func (c *AuthUserController) Mod() {

	id, _ := c.GetInt("id")
	user := models.AuthUser{ID: id}
	o.Read(&user)
	o.QueryTable("t_auth_role").FilterRaw("id", "in (select role_id from t_auth_role_user where user_id="+strconv.Itoa(user.ID)+")").All(&user.Roles)

	var roles []*models.AuthRole
	o.QueryTable(&models.AuthRole{}).All(&roles)

	c.Data["user"] = &user
	c.Data["roles"] = &roles
	fmt.Println(&user)
	c.TplName = "auth/userMod.tpl"
}

// DoMod 修改用户信息
// @router /mod [post]
func (c *AuthUserController) DoMod() {

	id, _ := c.GetInt("id")
	user := models.AuthUser{ID: id}
	o.Read(&user)

	user.State = 0
	if c.GetString("state") != "on" {
		user.State = 1
	}
	if c.GetString("Password") != "" {
		user.Password = aes.EncryptString(c.GetString("Password"))
	}
	if c.GetString("Portrait") != "" {
		user.Portrait = c.GetString("Portrait")
	}

	user.RealName = c.GetString("RealName")
	user.Signature = c.GetString("Signature")

	o.Update(&user)

	o.Raw("delete from t_auth_role_user where user_id=?", user.ID).Exec()
	roleIDs := c.GetStrings("roleIDs")
	if roleIDs != nil {
		for _, roleID := range roleIDs {
			o.Raw("insert into t_auth_role_user(role_id, user_id) values(?, ?)", roleID, user.ID).Exec()
		}
	}
	c.Redirect("/auth/user/", 302)
}
