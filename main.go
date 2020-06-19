package main

import (
	"github.com/astaxie/beego"
	"github.com/RyoLee/webcron/app/controllers"
	"github.com/RyoLee/webcron/app/jobs"
	_ "github.com/RyoLee/webcron/app/mail"
	"github.com/RyoLee/webcron/app/models"
	"html/template"
	"net/http"
)

const VERSION = "1.0.0"

func main() {
	models.Init()
	jobs.InitJobs()

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 生产环境不输出debug日志
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	beego.Router("/webcron/", &controllers.MainController{}, "*:Index")
	beego.Router("/webcron/login", &controllers.MainController{}, "*:Login")
	beego.Router("/webcron/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/webcron/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/webcron/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/webcron/help", &controllers.HelpController{}, "*:Index")
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
