package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"github.com/beego/i18n"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}


// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for resetting values
func (bc *baseController) Prepare() {
	//bc.Layout = "layout.tpl"
	bc.EnableRender = true
	fmt.Println("*** Inside default.Prepare() ***")
	// Reset language option.
	bc.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := bc.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			bc.Lang = al
		}
	}

	// 2. Default language is English.
	if len(bc.Lang) == 0 {
		bc.Lang = "en-US"
	}

	// Set template level language option.
	bc.Data["Lang"] = bc.Lang
}


type MainController struct {
	baseController
}

func (m *MainController) Get() {
	fmt.Println("*** Inside default.Get() ***")
	m.TplName = "login.html"
}
