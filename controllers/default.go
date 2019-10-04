package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}



func (c *MainController) Get() {

	c.TplName = "index.html"
}

func (this *MainController) MyURL(){
	this.TplName = "url.html"
}


func (this *MainController) OutType() {
	this.TplName = "out.html"

}

func (this *MainController) MyHttplib() {
	this.TplName = "http.html"
}
func (this *MainController) MyContext() {
	this.TplName = "context.html"
}

func (this *MainController) TestUrl() {
	data := make(map[string]interface{})
	data["name"] = "test url"
	data["message"] = "前缀URL设置或版本"
	data["url"] = "/v1/testUrl"
	this.Data["json"]= data
	this.ServeJSON()
}
