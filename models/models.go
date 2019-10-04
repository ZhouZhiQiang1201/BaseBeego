package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sqlite3"

type User struct {
	Id int
	Name string `orm:"size(20)"`
	Profile *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id int
	Age int
	User *User `orm:"reverse(one)"`

}

type Tag struct {
	Id int
	Name string
	Posts []*Post `orm:"reverse(many)"`

}

type Post struct {
	Id int
	Title string
	Tags []*Tag `orm:"rel(m2m)"`
}

func init(){
	orm.RegisterModel(new(User), new(Profile), new(Tag), new(Post))
	orm.RegisterDataBase("default", "sqlite3", "../basebeego.db")
	orm.RunSyncdb("default", false, true)
}