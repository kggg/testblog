package admin

import (
        "github.com/astaxie/beego"
        //"testblog/models"
)

type BaseController struct {
        beego.Controller
}

func (c *BaseController) Prepare() {
        username := c.GetSession("username")
        if username == false || username == nil {
                c.Redirect("/login", 302)
        }
        c.Data["Username"] = username
	
}


func (this *BaseController) Resp(status bool, str string) {
        this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
        this.ServeJSON()
        this.StopRun()
}

func (this *BaseController) CheckErr(err error, str string){
        if err != nil {
                this.Resp(false, str)
        }
}

