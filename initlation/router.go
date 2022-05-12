package initlation

import (
	"httpWeb_IM/controller"
	"net/http"
)

func RouterInit() {
	//用户登录注册
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.Register)
	//好友添加
	http.HandleFunc("/contact/addfriend", controller.User_Contact_Controller)
	//查询用户好友
	http.HandleFunc("/contact/loadfriend", controller.User_Contact_Search)
	http.HandleFunc("/chat", controller.Chat)
}
