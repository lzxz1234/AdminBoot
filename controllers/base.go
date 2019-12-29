package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lzxz1234/AdminBoot/models"
	"github.com/lzxz1234/AdminBoot/utils"
)

var o = orm.NewOrm()

// BootController 首页
type BootController struct {
	beego.Controller
}

// BaseController 封装通用方法
type BaseController struct {
	beego.Controller
}

// Render 覆盖 Render 方法
func (c *BaseController) Render() error {

	if !c.EnableRender {
		return nil
	}
	c.Data["siteName"] = beego.AppConfig.String("boot.siteName")
	c.Data["siteSubName"] = beego.AppConfig.String("boot.siteSubName")

	c.Data["me"] = c.GetAuthUser()
	authActions := c.GetSession("authActions")
	if authActions != nil {
		for _, authAction := range *authActions.(*[]string) {
			c.Data[strings.Replace(authAction, ".", "_", -1)] = true
		}
	}

	rb, err := c.RenderBytes()
	if err != nil {
		return err
	}

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	return c.Ctx.Output.Body(rb)
}

// HasAction 判断有没有权限
func (c *BaseController) HasAction(action string) bool {

	for _, authAction := range *c.GetAuthActionCodes() {
		if action == authAction {
			return true
		}
	}
	c.Data["json"] = utils.NewResult(-1, "没有权限", nil)
	c.ServeJSON()
	return false
}

// GetAuthActionCodes 获取全部权限代码
func (c *BaseController) GetAuthActionCodes() *[]string {

	authActions := c.GetSession("authActions")
	if authActions == nil {
		c.Abort("401")
	}
	return authActions.(*[]string)
}

// GetAuthUser 获取当前登录管理员用户
func (c *BaseController) GetAuthUser() (user *models.AuthUser) {

	userInSession := c.GetSession("authUser")
	if userInSession == nil {
		user = &models.AuthUser{}
		return
	}
	user = userInSession.(*models.AuthUser)
	return
}
