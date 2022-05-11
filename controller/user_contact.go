package controller

import (
	"httpWeb_IM/request_model"
	"httpWeb_IM/service"
	"httpWeb_IM/util"
	"net/http"
)

var UsercontactController service.ContactService

//添加好友功能
func User_Contact_Controller(writer http.ResponseWriter, request *http.Request) {
	//定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg request_model.ContactArg
	util.Bind(request, &arg)
	//调用service
	err := UsercontactController.UserContact(arg.Userid, arg.Dstid)
	//
	if err != nil {
		util.Respons(writer, 100, "好友添加失败", err.Error())
	} else {
		util.Respons(writer, 200, "好友添加成功", nil)
	}
}

//查询好友功能
func User_Contact_Search(writer http.ResponseWriter, request *http.Request) {
	var arg request_model.User_seach_Contact
	util.Bind(request, &arg)
	userRes := UsercontactController.SeachContactUser(arg.Userid)
	util.RespOkList(writer, userRes, len(userRes))
}
