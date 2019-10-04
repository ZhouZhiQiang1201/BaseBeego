package controllers

import (
	"BaseBeego/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DBController struct {
	beego.Controller
}

func (this *DBController) Get() {
	this.Data["xsrf"] = this.XSRFToken()
	this.TplName = "db.html"
}

func (this *DBController) GetOne() {
	id, _:= this.GetInt("ID")
	fmt.Println(id)
	db := orm.NewOrm()
	user := models.User{Id:id}

	err := db.Read(&user)
	db.LoadRelated(&user, "Profile")
	data := make(map[string]interface{})
	if err == nil {
		data["ID"] = user.Id
		data["Name"] = user.Name
		data["Age"] = user.Profile.Age
	}
	data["error"] = err
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *DBController) Add(){
	name := this.GetString("name")
	age, _ := this.GetInt("age")
	fmt.Println(name)
	fmt.Println(age)
	user := models.User{}
	profile := models.Profile{}
	profile.Age = age
	user.Name = name
	user.Profile = &profile
	db := orm.NewOrm()
	_, err := db.Insert(&profile)
	_, err = db.Insert(&user)
	data := make(map[string]interface{})
	if err == nil {
		data["Name"] = user.Name
		data["ID"] = user.Id
		data["Age"] = profile.Age
	}
	data["msg"] = err
	this.Data["json"] = data
	this.ServeJSON()
}


func TestManyToMany(){
	db := orm.NewOrm()
	p := models.Post{}
	//t := models.Tag{}
	//	//p.Title = "GO"
	//	//db.Insert(&p)

	//	//t.Name = "beego"
	//	//db.Insert(&t)
	//	//fmt.Println(p.Id)
	//	//fmt.Println(t.Id)
	p.Id = 1
	db.Read(&p)
	//t.Name = "jdango"
	//db.Insert(&t)
	//
	m2m := db.QueryM2M(&p, "Tags")
	//m2m.Add(&t)

	if m2m.Exist(&models.Tag{Id:1}){
		fmt.Println("exist")
	}
}



func init()  {
	TestManyToMany()
	dbns := beego.NewNamespace("/db",
		beego.NSRouter("/user", &DBController{}),
		beego.NSRouter("/add", &DBController{},"post:Add"),
		beego.NSRouter("/one", &DBController{}, "get:GetOne"),
	)
	beego.AddNamespace(dbns)
}