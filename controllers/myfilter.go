package controllers

import "github.com/astaxie/beego"

type MyFilter struct {
	beego.Controller
}

func (this *MyFilter) Get(){
	this.TplName = "filter.html"
}

func (this *MyFilter) TestFilter(){

	this.TplName = "filter.html"
}