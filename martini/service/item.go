package service

import (
	// "encoding/json"
	"time"
	"fmt"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"../model"
)

func GetItem(r render.Render) {
	item := model.Item{
		Id:         10086,
		Name:       "tea",
		Quantity:   100,
		Price:      1000,
		CreateTime: time.Now().Unix(),
	}

	r.JSON(200, item)
}

func PostItem(r render.Render, context martini.Context) {
	fmt.Println(context)
	// var item model.Item
	// json.Unmarshal(this.Ctx.Input.RequestBody, &item)

	r.JSON(200, "item")
}
