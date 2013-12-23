package main

import (
	"./service"

	"github.com/astaxie/beego"
)

func main() {
	beego.RESTRouter("/item", &service.Item{})
	beego.SetStaticPath("/", "static")
	beego.ViewsPath = "view"
	beego.Run()
}
