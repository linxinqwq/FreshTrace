package models

import "gorm.io/gorm"

type AddGoodsRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Store  int64  `json:"store"`
	Price  int64  `json:"price"`
	Images string `json:"images"`
	Uint   string `json:"uint"`   // 单位
	Origin string `json:"origin"` // 产地
	Kind   string `json:"kind"`
}

type UpdateGoodsInfoRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Store  int64  `json:"store"`
	Price  int64  `json:"price"`
	Images string `json:"images"`
	Kind   string `json:"kind"`
	Uint   string `json:"uint"`   // 单位
	Origin string `json:"origin"` // 产地
}

type RemoveGoodsRequest struct {
	Id     int64 `json:"id"`
	Status int64 `json:"status"`
}

type BuyGoodsRequest struct {
	Id      int64  `json:"id"`
	Store   int64  `json:"store"`
	Address string `json:"address"`
}

type GetUserOrdersResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Senders     string `json:"senders"`
	BlockNumber int64  `json:"blockNumber"`
	Number      int64  `json:"number"`
	Image       string `json:"image"`
	Price       int64  `json:"price"`
	TotalPrice  int64  `json:"totalPrice"`
	CreateTime  string `json:"createTime"`
	Address     string `json:"address"`
	Status      int64  `json:"status"`
}

type UpdateCheckStatusRequest struct {
	Id     int64 `json:"id"`
	Status int64 `json:"status"`
}

type UpdateGoodsProcessRequest struct {
	Id       int64  `json:"id"`
	ImageId  string `json:"imageId"`
	Operator string `json:"operator"`
	Status   int64  `json:"status"`
}

type GoodsCard struct {
	gorm.Model
	GoodsId  int64  `json:"goods_id"`
	UserName string `json:"user_name"`
	Number   int64  `json:"number"` // 购买数量
}

type GoodsCardDo struct {
	Id      int64  `json:"id"`
	GoodsId int64  `json:"goods_id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Store   int64  `json:"store"`
	Price   int64  `json:"price"`
	Images  string `json:"images"`
	Uint    string `json:"uint"`   // 单位
	Origin  string `json:"origin"` // 产地
	Time    string `json:"time"`   // 购买时间
	Check   bool   `json:"check"`  // 审核状态
}

type BuyFromCartRequest struct {
	Ids     []uint `json:"ids"`
	Address string `json:"address"`
}

type UpdateUserStatusRequest struct {
	UserName string `json:"username"`
	Status   bool   `json:"status"`
}

type CheckOrderRequest struct {
	Id int64 `json:"id"`
}
