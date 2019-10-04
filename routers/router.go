package routers

import (
	"BaseBeego/controllers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func TestFilter (ctx *context.Context) {
	_, ok := ctx.Input.Session("id").(int)
	if !ok && ctx.Request.RequestURI != "/login" {

		ctx.Redirect(302, "/")
	}
}



func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/testurl", &controllers.MainController{}, "get:TestUrl"),
		)
	beego.AddNamespace(ns)



	beego.InsertFilter("/testfilter", beego.BeforeRouter, TestFilter)
    beego.Router("/", &controllers.MainController{})
    beego.Router("/json", &controllers.MyJson{})
    beego.Router("/xsrf", &controllers.MyXSRF{})
    beego.Router("/param", &controllers.MyParam{})
    beego.Router("/upfile", &controllers.MyParam{}, "get:UpFile;post:UpFilePost")
    beego.Router("/session", &controllers.MySession{})
    beego.Router("/filter", &controllers.MyFilter{})
    beego.Router("/testfilter", &controllers.MyFilter{}, "get:TestFilter")
    beego.Router("/flash", &controllers.MyFlash{})
    beego.Router("/url", &controllers.MainController{}, "get:MyURL")
    beego.Router("/outType", &controllers.MainController{}, "get:OutType")
    beego.Router("/valid", &controllers.MyValid{})
    beego.Router("/error", &controllers.MyError{})
    beego.Router("/myerror", &controllers.MyError{}, "get:MyError")
    beego.Router("/httplib", &controllers.MainController{}, "get:MyHttplib")
    beego.Router("/context", &controllers.MainController{}, "get:MyContext")
    beego.Router("/log", &controllers.MyLog{})

}
