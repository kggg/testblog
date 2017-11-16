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

    beego.Router("/admin/section",&admin.SectionController{})
    beego.Router("/admin/section/add",&admin.SectionController{},"*:AddSection")
    beego.Router("/admin/section/edit/:id([0-9]+)",&admin.SectionController{},"*:EditSection")
    beego.Router("/admin/section/delete/:id([0-9]+)",&admin.SectionController{},"*:DelSection")

    beego.Router("/admin/user",&admin.UserController{})
    beego.Router("/admin/user/add",&admin.UserController{},"*:AddUser")
    beego.Router("/admin/user/edit/:id([0-9])+",&admin.UserController{},"*:EditUser")
    beego.Router("/admin/user/delete/:id([0-9])+",&admin.UserController{},"*:DelUser")
    /*
    beego.Router("/admin/system",&admin.SystemController{})
    beego.Router("/admin/backup",&admin.BackupController{})
    */

}
