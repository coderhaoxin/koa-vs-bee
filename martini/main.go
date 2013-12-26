package main

import (
	"github.com/codegangsta/martini"
	// "github.com/codegangsta/martini-contrib/render"

	"./service"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("static"))
	// m.Use(render.Renderer(render.Options{
	// 	Directory: "view",
	// }))

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello, " + params["name"]
	})
	m.Get("/view", service.View)
	m.Get("/item", service.GetItem)
	m.Post("/item", service.PostItem)
	m.Run()
	// http.ListenAndServe(":8080", m)
}
