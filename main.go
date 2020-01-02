package main

import (
	_ "github.com/udistrital/cuentas_contables_crud/routers"

	"github.com/astaxie/beego"
	dbConnManager "github.com/udistrital/cuentas_contables_crud/db"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

var mainDb = beego.AppConfig.String("mongo_db")

func main() {

	dbConnManager.InitDB(beego.AppConfig.String("mongo_host"),
		"27017",
		beego.AppConfig.String("mongo_user"),
		beego.AppConfig.String("mongo_pass"),
		beego.AppConfig.String("mongo_db_auth"),
		beego.AppConfig.String("mongo_db"),
	)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Custom libs
	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()

	beego.Run()
}
