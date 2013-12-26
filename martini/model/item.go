package model

type Item struct {
	Id         int64  `json:id`
	Name       string `json:name`
	Quantity   int64  `json:quantity`
	Price      int64  `json:price`
	CreateTime int64  `json:createTime`
}
