package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// import db driver
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	driverName := beego.AppConfig.String("database.driver")
	dataSource := beego.AppConfig.String("database.source")

	orm.RegisterDataBase("default", driverName, dataSource, 30)

	orm.RegisterModelWithPrefix("t_", new(AuthUser), new(AuthRole),
		new(AuthAction), new(AuthRoleAction), new(AuthRoleUser), new(AuthResource))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
