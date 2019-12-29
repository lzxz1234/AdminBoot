package init

import (
	"github.com/astaxie/beego/logs"
	"github.com/lzxz1234/AdminBoot/models"
)

func init() {

	authID := NewMenu("权限管理", "", 0, 0, "")
	NewMenu("人员列表", "/auth/user/", authID, 0, "RBAC.USER.LIST")
	NewMenu("角色列表", "/auth/role/", authID, 0, "RBAC.ROLE.LIST")

	NewHref("查询人员列表", "/auth/user/list/", 0, "RBAC.USER.LIST", 1)
	NewHref("添加修改人员信息", "/auth/user/*", 0, "RBAC.USER.MOD", 1)
	NewHref("查询角色列表", "/auth/role/list/", 0, "RBAC.ROLE.LIST", 1)
	NewHref("添加修改角色信息", "/auth/role/*", 0, "RBAC.ROLE.MOD", 1)
}

func NewMenu(name string, href string, pid int, userID int, actionCode string) int {

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

func NewHref(name string, href string, userID int, actionCode string, authType int) int {

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
