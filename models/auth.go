package models

import (
	"git.lzxz1234.cn/lzxz1234/AdminBoot/utils"
)

// AuthUser system administrators
type AuthUser struct {
	ID            int `orm:"pk;column(id);auto"`
	State         int // 0 Enabled 1 Disabled
	RealName      string
	UserName      string `orm:"unique"`
	Password      string
	Portrait      string
	Signature     string
	CreateTime    utils.Time `orm:"auto_now_add"`
	LastLoginTime utils.Time `orm:"auto_now"`

	Roles []*AuthRole `orm:"rel(m2m);rel_through(git.lzxz1234.cn/lzxz1234/AdminBoot/models.AuthRoleUser)"`
}

// AuthRole system roles
type AuthRole struct {
	ID          int `orm:"pk;column(id);auto"`
	Name        string
	Description string

	Actions []*AuthAction `orm:"rel(m2m);rel_through(git.lzxz1234.cn/lzxz1234/AdminBoot/models.AuthRoleAction)"`
	Users   []*AuthUser   `orm:"reverse(many)"`
}

// AuthAction system actions
type AuthAction struct {
	ID          int `orm:"pk;column(id);auto"`
	Name        string
	Code        string
	Group       string
	Description string

	Roles []*AuthRole `orm:"reverse(many)"`
}

// AuthRoleAction role action relations
type AuthRoleAction struct {
	ID     int         `orm:"pk;column(id);auto"`
	Role   *AuthRole   `orm:"rel(fk)"`
	Action *AuthAction `orm:"rel(fk)"`
}

// AuthRoleUser role user relations
type AuthRoleUser struct {
	ID   int       `orm:"pk;column(id);auto"`
	Role *AuthRole `orm:"rel(fk)"`
	User *AuthUser `orm:"rel(fk)"`
}

// AuthResource 需要权限校验的资源
type AuthResource struct {
	ID         int `orm:"pk;column(id);auto"`
	UserID     int `orm:"column(user_id)"` // 0 全员，其它指定人员
	ParentID   int `orm:"column(parent_id)"`
	Type       int // 0 菜单， 1 仅路径
	AuthType   int // 0 必须， 1 充分，区别是是否还进行其它资源权限校验
	Name       string
	Href       string
	ActionCode string
}
