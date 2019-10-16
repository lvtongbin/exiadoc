package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	host := beego.AppConfig.String("db.host")
	port := beego.AppConfig.String("db.port")
	password := beego.AppConfig.String("db.password")

	user := beego.AppConfig.String("db.user")
	name := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	if port == "" {
		port = "3306"
	}

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8"

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
