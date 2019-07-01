package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"],
        beego.ControllerComments{
            Method: "GetLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthLoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "Mod",
            Router: `/mod`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthRoleController"],
        beego.ControllerComments{
            Method: "DoMod",
            Router: `/mod`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Mod",
            Router: `/mod`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "DoMod",
            Router: `/mod`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "ModState",
            Router: `/modState`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:CmnFileController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:CmnFileController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"],
        beego.ControllerComments{
            Method: "InfoMod",
            Router: `/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"],
        beego.ControllerComments{
            Method: "Pass",
            Router: `/pass`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"] = append(beego.GlobalControllerRouter["git.lzxz1234.cn/lzxz1234/AdminBoot/controllers:SelfManageController"],
        beego.ControllerComments{
            Method: "PassMod",
            Router: `/pass`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
