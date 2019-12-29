package routers

import (
	"github.com/astaxie/beego"
	"github.com/lzxz1234/AdminBoot/controllers"
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
	beego.Router("/", &controllers.BootController{})
}
