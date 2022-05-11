package controller

import (
	"fmt"
	"httpWeb_IM/model"
	"httpWeb_IM/service"
	"httpWeb_IM/util"
	"log"
	"math/rand"
	"net/http"
)

var userServices service.UserService

//用户注册
func Register(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	user, err := userServices.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.Respons(writer, 100, "注册失败", err.Error())
	} else {
		util.Respons(writer, 200, "注册成功", user)
	}
}

//用户登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	ok, err, user := userServices.Login(mobile, passwd)
	if err != nil {
		util.Respons(writer, 100, "密码错误", nil)
		return
	}
	if ok {
		data := make(map[string]interface{})
		data["id"] = user.Model.ID
		data["token"] = user.Token
		util.Respons(writer, 200, "密码正确", data)
	} else {
		util.Respons(writer, 100, "密码错误", nil)
		return
	}

}
