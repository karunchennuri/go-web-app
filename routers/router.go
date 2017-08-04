package routers

import (
	"go-web-app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Register routers.
	beego.Router("/", &controllers.MainController{})
	// Indicate LoginController.ValidateToken method to handle POST requests
	beego.Router("/login", &controllers.LoginController{}, "get:SignIn")
	beego.Router("/login", &controllers.LoginController{}, "post:ValidateToken")
	// tasks
	beego.Router("/task", &controllers.TaskController{}, "get:ListTasks")
	beego.Router("/task", &controllers.TaskController{}, "post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")


	beego.Router("/test", &controllers.TestController{})
	beego.Router("/test/todo", &controllers.TestController{}, "get:ShowToTODO")

}
