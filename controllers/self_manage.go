package controllers

import (
	"github.com/lzxz1234/AdminBoot/utils/aes"
)

// SelfManageController 个人信息维护
type SelfManageController struct {
	BaseController
}

// Info 基础信息
// @router /info
func (c *SelfManageController) Info() {

	c.TplName = "self/info.tpl"
}

// Pass 密码信息
// @router /pass
func (c *SelfManageController) Pass() {

	c.TplName = "self/pass.tpl"
}

// PassMod 密码修改
// @router /pass [post]
func (c *SelfManageController) PassMod() {

	user := c.GetAuthUser()
	if user.Password != aes.EncryptString(c.GetString("OldPassword")) {
		c.Data["ErrMsg"] = "旧密码错误"
		c.Pass()
		return
	}
	user.Password = aes.EncryptString(c.GetString("NewPassword"))
	o.Update(user)

	c.DestroySession()
	c.Redirect("/auth/login", 302)
}

// InfoMod 基础信息
// @router /info [post]
func (c *SelfManageController) InfoMod() {

	user := c.GetAuthUser()

	if c.GetString("Portrait") != "" {
		user.Portrait = c.GetString("Portrait")
	}
	user.RealName = c.GetString("RealName")
	user.Signature = c.GetString("Signature")

	o.Update(user)

	c.Ctx.ResponseWriter.Write([]byte(`
	<html><script>window.top.location.reload();</script></html>
	`))
}
