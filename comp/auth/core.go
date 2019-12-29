package auth

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/lzxz1234/AdminBoot/models"
	"github.com/lzxz1234/AdminBoot/utils/aes"
)

var o = orm.NewOrm()

const (
	// Wrong 参数错误
	Wrong = -1
	// NotExists 找不到用户
	NotExists = -2
	// Disabled 账户被禁用
	Disabled = -3
	// Logout 单点登录，表示已登出
	Logout = -4
)

// GetUser decrypt userId from token string and fetch object
func GetUser(token string) (*models.AuthUser, int) {

	if token == "" {
		return &models.AuthUser{}, Wrong
	}
	raw, err := aes.DecryptString(token)
	if err != nil {
		return &models.AuthUser{}, Wrong
	}
	userID, _ := strconv.Atoi(strings.Split(raw, "_")[0])
	unix, _ := strconv.ParseInt(strings.Split(raw, "_")[1], 10, 64)
	user := models.AuthUser{ID: userID}
	if o.Read(&user) != nil {
		return &user, NotExists
	}
	if user.State != 0 {
		return &user, Disabled
	}
	if user.LastLoginTime.Unix() != unix {
		return &user, Logout
	}
	return &user, 0
}

// GetActions 查询用户拥有的全部权限
func GetActions(user *models.AuthUser) []*models.AuthAction {

	var actions []*models.AuthAction
	o.Raw(`select * from t_auth_action a where a.id in(
		select action_id from t_auth_role_action ra where ra.role_id in (
			select role_id from t_auth_role_user where user_id = ?
		)
	)`, user.ID).QueryRows(&actions)
	return actions
}

// UserHasAction has permision or not
func UserHasAction(user *models.AuthUser, action string) bool {

	if user.UserName == "admin" {
		return true
	}
	var maps []orm.Params
	r := o.Raw(`select count(1) as cnt
	              from t_auth_role_action 
			     where action_id = (
					   select id from t_auth_action where code=?)
			       and role_id in(
					   select id from t_auth_role_user where user_id=?)`, action, user.ID)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		return maps[0]["cnt"].(int) > 0
	}
	return false
}
