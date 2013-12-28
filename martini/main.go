package main

import (
	"./model"
	"./service"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("static"))
	m.Use(render.Renderer(render.Options{
		Directory:  "view",
		Extensions: []string{".html"},
	}))

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello, " + params["name"]
	})
	m.Get("/view", service.View)
	m.Get("/item", service.GetItem)
	m.Post("/item", binding.Json(model.Item{}), service.PostItem)
	m.Run()
	// http.ListenAndServe(":8080", m)
}
