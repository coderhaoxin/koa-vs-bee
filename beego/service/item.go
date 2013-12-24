package service

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"

	"../model"
)

type Item struct {
	beego.Controller
}

func (this *Item) Get() {
	item := model.Item{
		Id:         10086,
		Name:       "tea",
		Quantity:   100,
		Price:      1000,
		CreateTime: time.Now().Unix(),
	}

	this.Data["json"] = &item
	this.ServeJson()
}

func (this *Item) Post() {
	var item model.Item
	json.Unmarshal(this.Ctx.Input.RequestBody, &item)

	this.Data["json"] = &item
	this.ServeJson()
}
