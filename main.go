package main

import (
	_ "github.com/udistrital/cuentas_contables_crud/routers"

	"github.com/astaxie/beego"
	dbConnManager "github.com/udistrital/cuentas_contables_crud/db"
)

func main() {

	dbConnManager.InitDB(beego.AppConfig.String("mongo_host"),
		"27017",
		beego.AppConfig.String("mongo_user"),
		beego.AppConfig.String("mongo_pass"),
		beego.AppConfig.String("mongo_db_auth"),
	)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
