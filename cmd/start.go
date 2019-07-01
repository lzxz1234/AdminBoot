package cmd

import (
	"net/http"
	"strings"

	// init
	_ "git.lzxz1234.cn/lzxz1234/AdminBoot/init"
	_ "git.lzxz1234.cn/lzxz1234/AdminBoot/routers"

	"git.lzxz1234.cn/lzxz1234/AdminBoot/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
)

var resourceMap = make(map[string]*models.AuthResource)

// Start Server
func Start() {
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)

	beego.ErrorHandler("401", needLogin)
	beego.InsertFilter("/*", beego.BeforeRouter, filter)
	beego.Run()
}

func needLogin(rw http.ResponseWriter, r *http.Request) {

	http.Redirect(rw, r, "/auth/login", 302)
}

func filter(ctx *context.Context) {
	reqPath := ctx.Request.RequestURI // 可能是 /auth/self/modify
	if strings.HasSuffix(reqPath, "/") {
		reqPath = reqPath + "/"
	} // 一定是 /auth/self/modify/
	if resourceMap[reqPath] != nil {
		res := resourceMap[reqPath]
		if res != nil {
			if hasAction(ctx.Input.CruSession, res.ActionCode) {
				if res.AuthType == 1 {
					return
				}
			} else {

				needLogin(ctx.ResponseWriter, ctx.Request)
				panic(beego.ErrAbort)
			}
		}
	}
	// 开始模糊匹配 /auth/self/* /auth/*
	for ; strings.Count(reqPath, "/") >= 2; reqPath = reqPath[0:strings.LastIndex(reqPath[0:len(reqPath)-1], "/")] {
		res := resourceMap[reqPath+"/*"]
		if res != nil {
			if hasAction(ctx.Input.CruSession, res.ActionCode) {
				if res.AuthType == 1 {
					return
				}
			} else {
				needLogin(ctx.ResponseWriter, ctx.Request)
				panic(beego.ErrAbort)
			}
		}
	}
}

func hasAction(sess session.Store, code string) bool {

	if code == "" {
		return true
	}

	authActions := sess.Get("authActions")
	if authActions == nil {
		return false
	}
	for _, authAction := range *authActions.(*[]string) {
		if code == authAction {
			return true
		}
	}
	return false
}

func init() {

	// Href 要求，全以 / 开头，以 / 或者 /* 结尾
	var o = orm.NewOrm()
	var resources []*models.AuthResource
	o.QueryTable(&models.AuthResource{}).All(&resources)
	for _, resource := range resources {

		if resourceMap[resource.Href] != nil {
			logs.Error("资源链接重复", resource.Href)
		}
		resourceMap[resource.Href] = resource
	}
}
