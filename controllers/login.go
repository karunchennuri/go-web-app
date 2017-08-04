package controllers

import (
	"encoding/json"
	"fmt"
	"go-web-app/util"
	"github.com/astaxie/beego"
)

type LoginController struct {
	baseController
}

type User struct {
	name string
	role string
}

func (l *LoginController) SignIn() {
	beego.Info("*** Inside GET operation SignIn() ****")
	//l.Ctx.Redirect(302, "/task")
	l.TplName = "test.html"
	//return
	//beego.ReadFromRequest(&l.Controller)
	//l.Redirect("/test", 302)
	//l.TplName = "tasklist.html"
	//l.TplName = "test.html"
}

func (l *LoginController) ValidateToken() {
	beego.Info("*** Inside -POST- ValidateToken() ****")
	// Get form value.
	req := struct{ Token string }{}
	if err := json.Unmarshal(l.Ctx.Input.RequestBody, &req); err != nil {
		l.Ctx.Output.SetStatus(400)
		l.Ctx.Output.Body([]byte("empty title"))
		return
	}
	token := req.Token

	beego.Debug("token =", token)
	// Check valid.
	/*if len(token) == 0 {
		l.Redirect("/", 302)
		return
	}*/

	flash := beego.NewFlash()
	beego.Debug("token validation start")
	beego.Info("token validation start")
	i,err := util.NewIAMClient("https://iam.cn-north-1.myhwclouds.com/")
	if err != nil {
		beego.Error(fmt.Sprintf("IAM Client error : %v", err))
		beego.Info(fmt.Sprintf("IAM Client error : %v", err))
		flash.Error("Error occured")
		flash.Store(&l.Controller)
		//l.Ctx.Redirect( 302, "/")
		l.Abort("401")
		return
	}

	t, err := i.ValidateIAMToken(token)

	if err != nil {
		beego.Error(fmt.Sprintf("validation of token error : %v", err))
		beego.Info(fmt.Sprintf("validation of token error : %v", err))
		flash.Error("Login failed")
		flash.Store(&l.Controller)
		//l.Ctx.Redirect( 302, "/")
		l.Abort("401")
		return
	}

	beego.Debug("validation of token complete....")
	beego.Debug("t=", string(t))

	flash.Success("Login Successful")
	flash.Store(&l.Controller)
	u := &User{}
	u.name = "Karun Chennuri"
	u.role = "Admin"
	l.SetSession("User", u)
	l.Ctx.Redirect( 302, "/login")
	//l.Redirect("/test", 302)
	//l.TplName = "test.html"
	//l.Redirect("/test", 302)
	//l.Redirect("/task?token="+token, 302)

	// Usually put return after redirect.
	return

	//this.Ctx.Redirect(200, "/task/")
	//l.Redirect("/task", 302)
	//this.TplName = "tasklist.html"
	//this.Render()
	//return
}
