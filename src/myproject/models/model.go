package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	driverName := beego.AppConfig.String("driverName")

	orm.RegisterDriver(driverName, orm.DRMySQL)

	user := beego.AppConfig.String("root")
	password := beego.AppConfig.String("")
	host := beego.AppConfig.String("localhost")
	port := beego.AppConfig.String("3306")
	dbName := beego.AppConfig.String("demo")

	dbConn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		fmt.Println("failed")
	}
	fmt.Println("succeeded")
}
