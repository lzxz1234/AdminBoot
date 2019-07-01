package init

import (
	"git.lzxz1234.cn/lzxz1234/AdminBoot/models"
	"github.com/astaxie/beego/logs"
)

func init() {

	authID := initMenu("权限管理", "", 0, 0, "")
	initMenu("人员列表", "/auth/user/", authID, 0, "RBAC.USER.LIST")
	initMenu("角色列表", "/auth/role/", authID, 0, "RBAC.ROLE.LIST")

	initHref("查询人员列表", "/auth/user/list/", 0, "RBAC.USER.LIST", 1)
	initHref("添加修改人员信息", "/auth/user/*", 0, "RBAC.USER.MOD", 1)
	initHref("查询角色列表", "/auth/role/list/", 0, "RBAC.ROLE.LIST", 1)
	initHref("添加修改角色信息", "/auth/role/*", 0, "RBAC.ROLE.MOD", 1)
}

func initMenu(name string, href string, pid int, userID int, actionCode string) int {

	menu := models.AuthResource{
		Name:     name,
		ParentID: pid,
		UserID:   userID,
		Type:     0,
	}
	_, _, err := o.ReadOrCreate(&menu, "Name", "ParentID", "UserID", "Type")
	if err != nil {
		logs.Error("菜单初始化失败", name, err)
		return 0
	}
	menu.Href = href
	menu.ActionCode = actionCode
	o.Update(&menu)
	return menu.ID
}

func initHref(name string, href string, userID int, actionCode string, authType int) int {

	resource := models.AuthResource{
		Name:     name,
		ParentID: 0,
		UserID:   userID,
		Type:     1,
		AuthType: authType,
	}
	_, _, err := o.ReadOrCreate(&resource, "Name", "ParentID", "UserID", "Type")
	if err != nil {
		logs.Error("资源初始化失败", name, err)
		return 0
	}
	resource.Href = href
	resource.ActionCode = actionCode
	o.Update(&resource)
	return resource.ID
}
