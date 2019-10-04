package controllers

import (
	"github.com/astaxie/beego"
)

type MyValid struct {
	beego.Controller
}

type User struct {
	Id int `form:"-"`
	Name interface{} `form:"username,,名称:"`
	Age int `form:"age,text,年龄:"`
	Sex string `sex,,性别:`
	Email string `form:"enail,,邮箱:"`
	Intro string `form:",textarea"`
}

func (this *MyValid) Get(){
	this.Data["Form"] = &User{}
	this.TplName = "validation.html"
}

type user struct {
	userName string
	Age int
	email string
	mobile string
	ip string
}

func (this *MyValid) Post() {
	//userName := this.GetString("userName")
	//age := this.GetString("Age")
	//email := this.GetString("email")
	//IP := this.GetString("IP")
	//mobile := this.GetString("mobile")
	//u := user{}
	//valid := validation.Validation{}
	//if v := valid.Min(u.Age, 18 "age")


	this.TplName = "validation.html"
}