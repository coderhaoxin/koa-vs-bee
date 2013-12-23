package service

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
)

type ItemModel struct {
	Id         int64  `json:id`
	Name       string `json:name`
	Quantity   int64  `json:quantity`
	Price      int64  `json:price`
	CreateTime int64  `json:createTime`
}

type View struct {
	beego.Controller
}

func (this *View) Get() {
	this.Data["title"] = "beego-app-lab"
	this.Data["description"] = "hello, hahaha"

	items = []ItemModel{}
	this.Data["items"] = items
	item := ItemModel{
		Id:         10086,
		Name:       "tea",
		Quantity:   100,
		Price:      1000,
		CreateTime: time.Now().Unix(),
	}

	this.TplNames = "item.html"
}
