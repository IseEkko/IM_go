package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//定义返回的结构体
type Hjson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//将json抽象出来
func Respons(writer http.ResponseWriter, code int, msg string, data interface{}) {
	h := Hjson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	rsp, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}

	//设置头部
	writer.Header().Set("Content-Type", "application/json")
	//设置状态
	writer.WriteHeader(http.StatusOK)
	//设置返回
	writer.Write(rsp)
}
