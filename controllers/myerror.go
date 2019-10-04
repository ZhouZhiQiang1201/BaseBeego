package controllers

import (
	"github.com/astaxie/beego"
)

type MyError struct {
	beego.Controller
}

func (this *MyError) Get(){
	this.TplName = "error.html"
}

func (this *MyError) MyError() {
	this.Abort("myerror")
}

