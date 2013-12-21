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

type Item struct {
	beego.Controller
}

func (this *Item) Get() {
	item := ItemModel{
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
	var item ItemModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &item)

	this.Data["json"] = &item
	this.ServeJson()
}
