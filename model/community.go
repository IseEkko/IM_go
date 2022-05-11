package model

import "time"

type Community struct {
	Id int64 `form:"id" json:"id"`
	//名称
	Name string `form:"name" json:"name"`
	//群主ID
	Ownerid int64 `form:"ownerid" json:"ownerid"` // 什么角色
	//群logo
	Icon string `form:"icon" json:"icon"`
	//como
	Cate int `form:"cate" json:"cate"` // 什么角色
	//描述
	Memo string `form:"memo" json:"memo"` // 什么角色
	//
	Createat time.Time `form:"createat" json:"createat"` // 什么角色
}

const (
	COMMUNITY_CATE_COM = 0x01
)
