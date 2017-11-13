package routers

import (
	"github.com/astaxie/beego"
	"testblog/controllers/admin"
        "testblog/controllers/blog"
        "testblog/controllers/login"
)

func init() {
    beego.Router("/login", &login.LoginController{})
    beego.Router("/logout", &login.LoginController{}, "*:Logout")

    beego.Router("/", &blog.MainController{})
    beego.Router("/page/?:page([0-9]+)", &blog.MainController{})
    beego.Router("/show/:id([0-9]+)", &blog.ShowController{})
    beego.Router("/section/:id([0-9]+)", &blog.MainController{}, "*:FindSection")
    beego.Router("/label/:id([0-9]+)", &blog.MainController{}, "*:FindLabel")
    beego.Router("/search", &blog.MainController{}, "*:SearchBlog")

    beego.Router("/admin",&admin.IndexController{})
    beego.Router("/admin/article",&admin.ArticleController{})
    beego.Router("/admin/article/add",&admin.ArticleController{},"*:AddArticle")
    beego.Router("/admin/article/edit/:id([0-9]+)",&admin.ArticleController{},"*:EditArticle")
    beego.Router("/admin/article/delete/:id([0-9]+)",&admin.ArticleController{},"*:DelArticle")

    beego.Router("/admin/label",&admin.LabelController{})
    beego.Router("/admin/label/add",&admin.LabelController{},"*:AddLabel")
    beego.Router("/admin/label/edit/:id([0-9]+)",&admin.LabelController{},"*:EditLabel")
    beego.Router("/admin/label/delete/:id([0-9]+)",&admin.LabelController{},"*:DelLabel")
    /*
    */

}
