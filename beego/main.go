package main

import (
	"./service"

	"github.com/astaxie/beego"
)

func main() {
	beego.RESTRouter("/item", &service.Item{})
	beego.RESTRouter("/view", &service.View{})
	beego.DirectoryIndex = true
	beego.SetStaticPath("/static", "static/")
	beego.ViewsPath = "view"
	beego.Run()
}
