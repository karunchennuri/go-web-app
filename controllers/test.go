package controllers

import (
	"fmt"
)

type TestController struct {
	baseController
}

func (t *TestController) Get() {
	fmt.Println("*** Inside test.Get() ****")
	t.TplName = "test.html"
}

func (t *TestController) ShowToTODO() {
	fmt.Println("*** Inside test.ShowToTODO() ****")
	t.TplName = "tasklist.html"
}