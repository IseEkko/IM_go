package model

import "gorm.io/gorm"

const (
	CONCAT_CATE_USER     = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)

//好友和群都存在这个表里面
//可根据具体业务做拆分
type Contact struct {
	gorm.Model
	//谁的10000
	Ownerid int64 `form:"ownerid" json:"ownerid"` // 记录是谁的
	//对端,10001
	Dstobj int64 `form:"dstobj" json:"dstobj"` // 对端信息
	//
	Cate int    `form:"cate" json:"cate"` // 什么类型
	Memo string `form:"memo" json:"memo"` // 备注
}
