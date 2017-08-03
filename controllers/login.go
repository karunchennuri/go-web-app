package controllers

import (
	"fmt"
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
	fmt.Println("*** Inside GET operation SignIn() ****")
	//l.Ctx.Redirect(302, "/task")
	l.TplName = "test.html"
	//return
	//beego.ReadFromRequest(&l.Controller)
	//l.Redirect("/test", 302)
	//l.TplName = "tasklist.html"
	//l.TplName = "test.html"
}

func (l *LoginController) ValidateToken() {
	fmt.Println("*** Inside POST ValidateToken() ****")
	// Get form value.
	token := l.GetString("token")
	fmt.Println("*** token 1 =", token)
	// Check valid.
/*	if len(token) == 0 {
		l.Redirect("/", 302)
		return
	}
*/
	flash := beego.NewFlash()
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
