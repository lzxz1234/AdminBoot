package controllers

import (
	"strconv"

	"github.com/lzxz1234/AdminBoot/models"
	"github.com/lzxz1234/AdminBoot/utils"
)

// AuthRoleController role api
type AuthRoleController struct {
	BaseController
}

// Index 角色管理首页
// @router /
func (c *AuthRoleController) Index() {
	c.TplName = "auth/roleList.tpl"
}

// List roles
// @router /list
func (c *AuthRoleController) List() {

	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 10)

	count, _ := o.QueryTable(&models.AuthRole{}).Count()
	var roles []*models.AuthRole
	o.QueryTable(&models.AuthRole{}).Limit(size, page*size-size).All(&roles)
	for _, role := range roles {
		o.QueryTable("t_auth_action").
			FilterRaw("id", "in (select action_id from t_auth_role_action where role_id="+strconv.Itoa(role.ID)+")").
			GroupBy("Group").
			All(&role.Actions, "Group")
		o.QueryTable("t_auth_user").
			FilterRaw("id", "in (select user_id from t_auth_role_user where role_id="+strconv.Itoa(role.ID)+")").
			All(&role.Users, "RealName", "UserName")
	}

	c.Data["json"] = utils.NewPage(int(count), page, size, roles)
	c.ServeJSON()
}

// Add 添加新角色
// @router /add [get]
func (c *AuthRoleController) Add() {

	var actions []*models.AuthAction
	o.QueryTable("t_auth_action").All(&actions)
	actionMap := make(map[string][]*models.AuthAction)
	for _, action := range actions {
		actionMap[action.Group] = append(actionMap[action.Group], action)
	}
	c.Data["actions"] = &actionMap
	c.TplName = "auth/roleAdd.tpl"
}

// DoAdd 添加新角色
// @router /add [post]
func (c *AuthRoleController) DoAdd() {

	role := models.AuthRole{
		Name:        c.GetString("Name"),
		Description: c.GetString("Description"),
	}
	o.Insert(&role)
	actionIDs := c.GetStrings("actionIds")
	if actionIDs != nil {
		for _, actionID := range actionIDs {
			o.Raw("insert into t_auth_role_action(role_id, action_id) values(?, ?)", role.ID, actionID).Exec()
		}
	}
	c.Redirect("/auth/role/", 302)
}

// Mod 修改角色信息
// @router /mod [get]
func (c *AuthRoleController) Mod() {

	id, _ := c.GetInt("id")
	role := models.AuthRole{ID: id}
	o.Read(&role)

	o.QueryTable("t_auth_action").
		FilterRaw("id", "in (select action_id from t_auth_role_action where role_id="+strconv.Itoa(role.ID)+")").
		All(&role.Actions, "ID")

	var actions []*models.AuthAction
	o.QueryTable("t_auth_action").All(&actions)
	actionMap := make(map[string][]*models.AuthAction)
	for _, action := range actions {
		actionMap[action.Group] = append(actionMap[action.Group], action)
	}

	c.Data["actions"] = &actionMap
	c.Data["role"] = &role
	c.TplName = "auth/roleMod.tpl"
}

// DoMod 修改角色信息
// @router /mod [post]
func (c *AuthRoleController) DoMod() {

	id, _ := c.GetInt("id")
	role := models.AuthRole{ID: id}
	o.Read(&role)
	role.Name = c.GetString("Name")
	role.Description = c.GetString("Description")
	o.Update(&role)
	o.Raw("delete from t_auth_role_action where role_id=?", role.ID).Exec()
	actionIDs := c.GetStrings("actionIDs")
	if actionIDs != nil {
		for _, actionID := range actionIDs {
			o.Raw("insert into t_auth_role_action(role_id, action_id) values(?, ?)", role.ID, actionID).Exec()
		}
	}
	c.Redirect("/auth/role/", 302)
}

// Info detail for role
// @router /detail
func (c *AuthRoleController) Info() {

	id, _ := c.GetInt("id", 0)
	role := models.AuthRole{ID: id}
	o.Read(&role)
	c.Data["json"] = &role
	c.ServeJSON()
}
