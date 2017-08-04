package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"go-web-app/models"
)

type TaskController struct {
	baseController
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (tc *TaskController) ListTasks() {
	//tc.TplName = "tasklist.html"
	u := tc.GetSession("User")
	beego.Info("*** session obj=", u)
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	tc.Data["json"] = res
	tc.ServeJSON()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (tc *TaskController) NewTask() {
	beego.Info("**** Inside NewTask ***")
	tc.TplName = "tasklist.html"
	req := struct{ Title string }{}
	if err := json.Unmarshal(tc.Ctx.Input.RequestBody, &req); err != nil {
		tc.Ctx.Output.SetStatus(400)
		tc.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t, err := models.NewTask(req.Title)
	if err != nil {
		tc.Ctx.Output.SetStatus(400)
		tc.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beego.Info("before save =", t)
	models.DefaultTaskList.Save(t)
}

// Examples:
//
//   req: GET /task/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /task/42
//   res: 404 task not found
func (tc *TaskController) GetTask() {
	tc.TplName = "tasklist.html"
	id := tc.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		tc.Ctx.Output.SetStatus(404)
		tc.Ctx.Output.Body([]byte("task not found"))
		return
	}
	tc.Data["json"] = t
	tc.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (tc *TaskController) UpdateTask() {
	tc.TplName = "tasklist.html"
	id := tc.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(tc.Ctx.Input.RequestBody, &t); err != nil {
		tc.Ctx.Output.SetStatus(400)
		tc.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		tc.Ctx.Output.SetStatus(400)
		tc.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		tc.Ctx.Output.SetStatus(400)
		tc.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}
