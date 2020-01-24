package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/auditoria"
	dbConnManager "github.com/udistrital/cuentas_contables_crud/db"
	_ "github.com/udistrital/cuentas_contables_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
	"github.com/udistrital/utils_oas/responseformat"
)

//var mainDb = beego.AppConfig.String("mongo_db")

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
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// migrations
	if _, err := dbConnManager.RunMigrations(); err != nil {
		logs.Error("Migrations Error: ", err.Error())
	} else {
		logs.Info("Migration process success !")
	}
	// Custom libs
	beego.BConfig.RecoverFunc = responseformat.GlobalResponseHandler
	auditoria.InitMiddleware()
	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()

	beego.Run()
}
