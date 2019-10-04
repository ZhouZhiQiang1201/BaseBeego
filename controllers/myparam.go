package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"time"
)

type MyParam struct {
	beego.Controller
}

func (this *MyParam) Get() {
	this.TplName = "param.html"
}

func (this *MyParam) UpFile() {
	//this.Data["xsrfdata"] = this.XSRFFormHTML()
	this.Data["xsrf"] = this.XSRFToken()
	this.TplName = "upfile.html"
}

func (this *MyParam) UpFilePost() {
	//this.EnableXSRF = false
	file, header, err := this.GetFile("files")

	if err != nil {
		fmt.Println(err)
		this.Data["error"] = "读取失败"
		this.Data["xsrf"] = this.XSRFToken()
		this.TplName = "upfile.html"
		return
	}
	//关闭文件
	defer file.Close()
	//获取后缀
	ext := path.Ext(header.Filename)

	fileName := time.Now().Format("20060102150405") + ext
	fmt.Println(fileName)
	err = this.SaveToFile("files", "static/images/" + fileName)
	if err != nil {
		fmt.Println(err)
		this.Data["error"] = "保存失败"
	}
	this.Data["xsrf"] = this.XSRFToken()
	this.TplName = "upfile.html"


}