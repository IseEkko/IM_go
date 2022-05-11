package initlation

import (
	"httpWeb_IM/global"
	"httpWeb_IM/model"
)

//数据表加载
func Auto() {
	_ = global.DB.AutoMigrate(&model.User{})
	_ = global.DB.AutoMigrate(&model.Contact{})
}
