package main

import (
	"github.com/astaxie/beego"
	_ "go-web-app/routers"
	"github.com/beego/i18n"
)

const (
	APP_VER = "1.0"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
