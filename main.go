package main

import (
	"github.com/astaxie/beego"
	"ligang/functions/dc"
	"ligang/routers"

	_ "ligang/models"
	_ "ligang/routers"
	_ "ligang/tasks"

	"ligang/sysinit"
	"ligang/utils"
)

func init() {
	utils.InitCache()
	utils.InitLogs()
	sysinit.InitDB()
	routers.SplitMethod()
	dc.InitDc()
}

func main() {
	set := sysinit.Settings
	go dc.Loop()
	beego.BConfig.RunMode = set.Runmode

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "ligangmex"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/"] = "static"

	// beego.BConfig.Listen.Graceful = true
	beego.Run(":" + set.Httpport)
}
