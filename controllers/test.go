package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	baseController
}

func (t *TestController) Get() {
	beego.Info("*** Inside test.Get() ****")
	token := t.Ctx.Input.Param("token")
	beego.Debug("token in test =", token)
	beego.Info("*** token in test =", token)

	t.TplName = "test.html"
}

func (t *TestController) ShowToTODO() {
	beego.Info("*** Inside test.ShowToTODO() ****")

	token := t.Ctx.Input.Param("token")
	beego.Debug("token in test =", token)
	beego.Info("*** token in test =", token)

	t.TplName = "tasklist.html"
}

