package initlation

import (
	"html/template"
	"log"
	"net/http"
)

func Tpl() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
	//这里是静态资源的目录支持，第一个参数代表的是支持以"/"开头的目录，然后后面的参数代表的是该目录下的文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
}
