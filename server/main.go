package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/piaohua/snowflake-golang/server/controllers"
)

func main() {
	beego.BConfig.EnableGzip = true
	beego.BConfig.RunMode = "dev"
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Listen.HTTPPort = 8085

	//初始化 namespace
	ns :=
		beego.NewNamespace("/v1",
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.Domain() == "localhost" {
					return true
				}
				return false
			}),
			beego.NSBefore(func(ctx *context.Context) {
				if ctx.Input.Header("x-snowflake-access-token") != "snowflake" {
					//ctx.Output.Body([]byte("notAllowed"))
				}
			}),
			beego.NSNamespace("/id",
				beego.NSRouter("/", &controllers.MainController{}),
				beego.NSNamespace("/batch",
					beego.NSRouter("/:count([0-9]+)", &controllers.MainController{}, "post:Count"),
				),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns)
	beego.Run()
}
