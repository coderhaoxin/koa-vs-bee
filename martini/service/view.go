package service

import (
	"strconv"
	"time"

	"../model"

	"github.com/martini-contrib/render"
)

type ViewData struct {
	Title string
	Description string
	Items []model.Item
}

func View(r render.Render) {
	items := []model.Item{}
	for i := 0; i < 100; i++ {
		item := model.Item{
			Id:         int64(i),
			Name:       "tea-0" + strconv.Itoa(i),
			Quantity:   100 + int64(i),
			Price:      1000 + 10 * int64(i),
			CreateTime: time.Now().Unix(),
		}
		items = append(items, item)
	}

	viewData := ViewData{
		Title: "beego-app-lab",
		Description: "hello, hahaha",
		Items: items,
	}

	r.HTML(200, "item", viewData)
}
