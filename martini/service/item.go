package service

import (
	"fmt"
	"time"

	"../model"

	"github.com/martini-contrib/render"
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

func PostItem(r render.Render, item model.Item) {
	fmt.Println(item)
	r.JSON(200, item)
}
