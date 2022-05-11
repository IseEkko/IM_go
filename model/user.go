package model

import (
	"gorm.io/gorm"
)

const (
	SEX_WOMEN = "W"
	SEX_MEN   = "M"
	//
	SEX_UNKNOW = "U"
)

type User struct {
	//用户ID
	gorm.Model
	Mobile   string `form:"mobile" json:"mobile"`
	Passwd   string `form:"passwd" json:"passwd"` // 什么角色
	Avatar   string `form:"avatar" json:"avatar"`
	Sex      string `form:"sex" json:"sex"`           // 什么角色
	Nickname string `form:"nickname" json:"nickname"` // 什么角色
	//加盐随机字符串6
	Salt   string `form:"salt" json:"-"`        // 什么角色
	Online int    `form:"online" json:"online"` //是否在线
	//前端鉴权因子,
	Token string `form:"token" json:"token"` // 什么角色
	Memo  string `form:"memo" json:"memo"`   // 什么角色
}
