package routers

import (
	"github.com/astaxie/beego"
	//"blog/controllers/admin"
        "testblog/controllers/blog"
        //"blog/controllers/login"
)

func init() {
        //beego.Router("/login", &login.LoginController{})
    //beego.Router("/logout", &login.LoginController{}, "*:Logout")

    beego.Router("/", &blog.MainController{})
    beego.Router("/page/?:page", &blog.MainController{})
    //beego.Router("/show/:id", &blog.ShowController{})

    /*
    ns :=
    beego.NewNamespace("/admin",
            beego.NSRouter("/", &admin.IndexController{}),
            beego.NSNamespace("/article",
                   beego.NSRouter("/", &admin.ArticleController{}),
                   beego.NSRouter("/add", &admin.ArticleController{},"*:AddArticle"),
                   beego.NSRouter("/edit/:id", &admin.ArticleController{},"*:EditArticle"),
                   beego.NSRouter("/delete/:id", &admin.ArticleController{},"*:DelArticle")
	    ),
	    beego.NSNamespace("/user",
	           beego.NSRouter("/", &admin.UserController{}),
		   beego.NSRouter("/setting", &admin.UserController{},"*:EditUser"),
		   beego.NSRouter("/changepassword", &admin.UserController{},"*:ChangePassword"),

	    ),
     )
     beego.AddNamespace(ns)
     */
}
