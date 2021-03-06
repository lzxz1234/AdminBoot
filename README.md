# AdminBoot

基于 Beego + Rbac + QiNiu + LayUI 的项目模板

## 添加新权限

编辑 init/auth.go:

> initAction("RBAC.USER.LIST", "人员查看", "权限控制", "查看管理员列表，信息等")

四个参数分别是：

- 权限编码，用于后续权限校验操作
- 操作名称，角色编辑时展示
- 操作分组，角色编辑时展示
- 操作描述，角色编辑时展示

## 添加新资源

这个资源主要指和权限绑定的服务器资源，主要分两类，菜单 和 链接，菜单本身也是链接，不同的是菜单会展示在工作台左侧

编辑 init/menu.go 添加菜单 

> initMenu("人员列表", "/auth/user/", authID, 0, "RBAC.USER.LIST")

参数分别是：

- 菜单名称，展示在工作台左侧
- 菜单链接页面
- 父菜单 ID，父菜单 ID 来自 initMenu 的返回值，目前 UI 仅支持两层结构
- 用户 ID，可能某个菜单仅指定特定的人，就在此处设置用户 ID，一般默认 0 就可以了
- 权限编码，拥有对应权限的人可以看到对应菜单 

编辑 init/menu.go 添加链接 

> initHref("查询人员列表", "/auth/user/list/", 0, "RBAC.USER.LIST", 1)

参数分别是：

- 链接名称，方便管理
- 链接地址
- 用户 ID，可能某个菜单仅指定特定的人，就在此处设置用户 ID，一般默认 0 就可以了
- 权限编码，拥有对应权限的人可以看到对应菜单 
- 权限校验类型，0 必须， 1 充分，一般默认 1 即可