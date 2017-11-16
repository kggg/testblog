package  login

import (
        "github.com/astaxie/beego"
	"testblog/models"
	"github.com/gogather/com"
)



type LoginController struct {
        beego.Controller
}

type Login struct{
	name string
	password string
}

func (c *LoginController) Prepare() {
    c.EnableXSRF = false
}


func (c *LoginController) Get() {
        c.TplName = "login/index.tpl"
}

func (c *LoginController) Post(){
	result := make(map[string]interface{})
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
                result["status"] = "false"
                result["message"] = "用户和密码不能为空"
                c.Data["json"] = result
                c.ServeJSON()
		c.StopRun()
                return		
	}
	//need to verify with mysql user table
        user, err := models.FindUserByName(username)
        if err != nil {
                result["status"] = "false"
                result["message"] = "用户不存在"
                c.Data["json"] = result
                c.ServeJSON()		
              
	}else{
		pass  := com.Md5(password)
		if pass == user.Password {

		     c.SetSession("username", username)
		     c.SetSession("userid", user.Id)
	             c.Redirect("/admin", 302)
	        }else {
			c.Ctx.WriteString("passwd invalid")
		}
	}
}

func (this *LoginController) Logout() {
	this.SetSession("username", nil)
	this.DelSession("username")
	this.Ctx.Redirect(302, "/login")
}

