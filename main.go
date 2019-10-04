package main

import (
	_ "BaseBeego/routers"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
	_"BaseBeego/models"
)

func CustomError(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("myerror.html").ParseFiles("views/myerror.html")
	data := make(map[string]interface{})
	data["custom"] = "自定义错误"
	t.Execute(rw, data)

}

func main() {
	beego.ErrorHandler("myerror", CustomError)
	beego.Run()
}

