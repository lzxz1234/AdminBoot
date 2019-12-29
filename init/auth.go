package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/lzxz1234/AdminBoot/models"
	"github.com/lzxz1234/AdminBoot/utils"
	"github.com/lzxz1234/AdminBoot/utils/aes"
)

var o = orm.NewOrm()

func init() {

	admin := models.AuthUser{UserName: "admin"}
	if created, _, err := o.ReadOrCreate(&admin, "UserName"); err == nil {
		if created {
			admin.State = 0
			admin.RealName = "Administrator"
			admin.Password = aes.EncryptString("admin123!")
			admin.Portrait = "/static/img/portrait.jpg"
			admin.Signature = ""
			admin.CreateTime = utils.Now()
			admin.LastLoginTime = utils.Now()
			o.Update(&admin)
		} else if beego.AppConfig.String("boot.password") != "" {
			admin.Password = aes.EncryptString(beego.AppConfig.String("boot.password"))
		}
		o.Update(&admin)
	}

	role := models.AuthRole{Name: "超级管理员"}
	if created, _, err := o.ReadOrCreate(&role, "Name"); err == nil {
		if created {
			role.Description = "拥有系统所有数据权限和操作权限"
			o.Update(&role)
		}
	}

	if !o.QueryTable("t_auth_role_user").Filter("role_id", role.ID).Filter("user_id", admin.ID).Exist() {
		o.Raw("insert into t_auth_role_user(role_id, user_id) values(?,?)", role.ID, admin.ID).Exec()
	}

	NewAction("RBAC.USER.LIST", "人员查看", "权限控制", "查看管理员列表，信息等")
	NewAction("RBAC.USER.MOD", "人员修改", "权限控制", "修改管理员信息，角色列表，执行禁用等操作")
	NewAction("RBAC.ROLE.LIST", "角色查看", "权限控制", "查看系统角色列表，信息等")
	NewAction("RBAC.ROLE.MOD", "角色修改", "权限控制", "修改角色信息、对应权限等操作")
}

func NewAction(code string, name string, group string, description string) {

	adminRole := models.AuthRole{Name: "超级管理员"}
	err := o.Read(&adminRole, "Name")
	if err != nil {
		logs.Error("Load Administrator Role Failed", err)
	}
	action := models.AuthAction{Code: code}
	if created, _, err := o.ReadOrCreate(&action, "Code"); err == nil {
		if created {
			o.Raw("insert into t_auth_role_action(role_id, action_id) values(?,?)", adminRole.ID, action.ID).Exec()
		}
		action.Name = name
		action.Group = group
		action.Description = description
		o.Update(&action)
	} else {
		logs.Error("Init Action Error : ", err)
	}
}
