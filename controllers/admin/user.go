package admin

import (
        //"html/template"
        "testblog/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"testblog/utils/validations"
)

type UserController struct {
        BaseController
}

func (c *UserController) Get(){
	users , err := models.FindAllUser()
	c.CheckErr(err, "get section info error")
        c.Data["Users"] = users
        c.Data["Pagetitle"] = "用户列表"
        c.TplName = "admin/user/index.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}

func (c *UserController) AddUser(){
	if c.isPost(){
		username := c.GetString("username")
		password := c.GetString("password")
		repassword := c.GetString("repassword")
                valid := validation.Validation{}
                u := validations.UserForm{Name: username, Password: password, Repassword: repassword}
                b, err := valid.Valid(&u)
                if err != nil {
                        c.Resp(false, err.Error())
                        return
                }
                if !b {
                        for _, err := range valid.Errors {
                                c.Data["json"] = &map[string]interface{}{"status": false, "key": err.Key, "message": err.Message}
                        }
                        c.ServeJSON()
                        c.StopRun()
                }
		exist := models.CheckUserExist(username)
		if exist {
			c.Resp(false, "the username has been existing")
		}
		_, err = models.AddUser(username, password)
                if err == nil {
                        c.Redirect("/admin/user", 302)
                } else {
			result := make(map[string]interface{})
                        beego.Error("Insert into mysql error: ", err.Error())
                        result["status"] = "false"
                        result["message"] = "数据创建失败"
                        result["debug"] = err.Error()
                        c.Data["json"] = result
                        c.ServeJSON()
                        c.StopRun()
                }
		

	}
        c.Data["Pagetitle"] = "新增用户"
        c.TplName = "admin/user/add.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = ""
        c.LayoutSections["Scripts"] = "admin/editor.tpl"
}

func (c *UserController) EditUser(){
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.CheckErr(err, "id atoi error")
	user, err := models.FindUserById(id)
	if c.isPost(){
                username := c.GetString("username")
                password := c.GetString("password")
                repassword := c.GetString("repassword")
                valid := validation.Validation{}
                u := validations.UserForm{Name: username, Password: password, Repassword: repassword}
                b, err := valid.Valid(&u)
                if err != nil {
                        c.Resp(false, err.Error())
                        return
                }
                if !b {
                        for _, err := range valid.Errors {
                                c.Data["json"] = &map[string]interface{}{"status": false, "key": err.Key, "message": err.Message}
                        }
                        c.ServeJSON()
                        c.StopRun()
                }
                _, err = models.UpdateUser(id, username, password)
                if err == nil {
                        c.Redirect("/admin/user", 302)
                } else {
                        result := make(map[string]interface{})
                        beego.Error("Insert into mysql error: ", err.Error())
                        result["status"] = "false"
                        result["message"] = "数据创建失败"
                        result["debug"] = err.Error()
                        c.Data["json"] = result
                        c.ServeJSON()
                        c.StopRun()
                }

	}
	c.Data["Users"] = user
        c.Data["Pagetitle"] = "编辑用户"
        c.TplName = "admin/user/edit.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}


func (c *UserController) DelUser(){
        id := c.Ctx.Input.Param(":id")
        bid, err := strconv.Atoi(id)
        if err != nil {
                beego.Error("atoi error")
        }
        _, err = models.DelUser(bid)
        if err == nil {
                beego.Informational("Delete blog success", id)
                c.Redirect("/admin/user", 302)
        }
}


