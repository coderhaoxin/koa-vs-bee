package service

import (
	"time"
	"strconv"

	"github.com/astaxie/beego"

	"../model"
)

type View struct {
	beego.Controller
}

func (this *View) Get() {
	this.Data["title"] = "beego-app-lab"
	this.Data["description"] = "hello, hahaha"
	var items = []model.Item{}

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

	this.Data["items"] = items

	this.TplNames = "item.html"
}
