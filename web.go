package main

import (
	"httpWeb_IM/initlation"
	"net/http"
)

func main() {
	//模版加载
	initlation.Tpl()
	//数据库链接
	initlation.InitDB()
	//路由方法
	initlation.RouterInit()
	http.ListenAndServe(":8080", nil)
}
