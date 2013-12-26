package service

import (
	"github.com/codegangsta/martini-contrib/render"
)

func View(r render.Render) {
	r.HTML(200, "view", "gogogo!!!hahaha")
}
