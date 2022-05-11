package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"httpWeb_IM/global"
	"httpWeb_IM/model"
	"httpWeb_IM/util"
	"math/rand"
)

type UserService struct {
}

//用户注册
func (users *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	//查询账号是不是存在
	var tmp = model.User{Mobile: mobile}
	result := global.DB.Find(&user)
	//判断是不是没有数据
	returnok := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if returnok {
		return tmp, errors.New("用户已存在")
	}
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	insert := global.DB.Create(&tmp)
	err = insert.Error
	if err != nil {
		return tmp, errors.New("插入失败")
	}
	return tmp, nil
}

//用户登录
func (users *UserService) Login(mobile, password string) (ok bool, err error, user model.User) {
	var tmp = model.User{Mobile: mobile}
	result := global.DB.First(&tmp)
	if result.Error != nil {
		return false, errors.New("用户登录查询失败"), user
	}
	if result.RowsAffected == 0 {
		return false, errors.New("用户不存在"), user
	}
	passwords := util.MakePasswd(password, tmp.Salt)
	if tmp.Passwd == passwords {
		return true, err, tmp
	}
	return false, errors.New("用户密码错误"), user
}
