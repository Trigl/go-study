package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

type MainController1 struct {
	beego.Controller
}

func (this *MainController1) Get() {
	this.Ctx.WriteString("hello world1")
}

func main() {
	beego.Router("/*", &MainController{})
	beego.Router("/nihao", &MainController1{})
	beego.Run()
}
