package controllers

import (
	"github.com/astaxie/beego/logs"

	"github.com/lzxz1234/AdminBoot/comp/auth"
	"github.com/lzxz1234/AdminBoot/models"
	"github.com/lzxz1234/AdminBoot/utils"
	"github.com/lzxz1234/AdminBoot/utils/aes"
)

// AuthLoginController for login or logout
type AuthLoginController struct {
	BaseController
}

// Index 后台首页
// @router /
func (c *AuthLoginController) Index() {

	if c.GetAuthUser().ID == 0 {
		c.Redirect("/auth/login", 302)
		return
	}
	var menus []models.AuthResource
	o.QueryTable("t_auth_resource").Filter("type", 0).
		Filter("user_id__in", 0, c.GetAuthUser().ID).
		Filter("action_code__in", *c.GetAuthActionCodes()).All(&menus)
	c.Data["menus"] = &menus
	c.TplName = "manageIndex.tpl"
}

// GetLogin 打开登录页
// @router /login [get]
func (c *AuthLoginController) GetLogin() {

	c.TplName = "auth/login.tpl"
}

// Login login
// @router /login [post]
func (c *AuthLoginController) Login() {

	username := c.GetString("username")
	password := c.GetString("password")
	var user models.AuthUser
	err := o.QueryTable("t_auth_user").Filter("username", username).One(&user)
	if err != nil {
		c.Data["json"] = utils.NewResult(1, "用户不存在", nil)
		c.ServeJSON()
		logs.Info("用户名 %s 错误", username)
		return
	}

	if user.UserName != "admin" && user.State != 0 {
		c.Data["json"] = utils.NewResult(1, "用户已禁用", nil)
		c.ServeJSON()
		return
	}

	if user.Password == aes.EncryptString(password) {
		user.LastLoginTime = utils.Now()
		o.Update(&user)
		c.SetSession("authUser", &user)

		actions := auth.GetActions(&user)
		var authActions []string
		for _, action := range actions {
			authActions = append(authActions, action.Code)
		}
		authActions = append(authActions, "") // 空权限应该是所有人都有的
		c.SetSession("authActions", &authActions)
		c.Data["json"] = utils.NewResult(0, "登录成功", nil)
		logs.Info("用户 %s 登录系统成功", username)
	} else {
		c.Data["json"] = utils.NewResult(1, "密码错误", nil)
		logs.Info("用户 %s 登录密码错误", username)
	}
	c.ServeJSON()
}

// Logout logout
// @router /logout
func (c *AuthLoginController) Logout() {

	user := c.GetAuthUser()
	c.DestroySession()
	logs.Info("用户 %s 退出登录", user.UserName)
	c.Redirect("/auth/login", 302)
}
