package routers

import (
	"git.lzxz1234.cn/lzxz1234/AdminBoot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/auth",
		beego.NSNamespace("/role",
			beego.NSInclude(&controllers.AuthRoleController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.AuthUserController{}),
		),
		beego.NSInclude(&controllers.AuthLoginController{}),
	)
	beego.AddNamespace(ns)
	ns = beego.NewNamespace("/cmn",
		beego.NSInclude(&controllers.CmnFileController{}),
	)
	beego.AddNamespace(ns)
	ns = beego.NewNamespace("/self",
		beego.NSInclude(&controllers.SelfManageController{}),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
}
