package blog

import (
	"github.com/astaxie/beego"
	"testblog/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	section, err := models.FindAllSection()
	c.CheckErr(err, "get section error")
	label, err := models.FindAllLabel()
	c.CheckErr(err, "get label error")
	c.Data["section"] = section
	c.Data["label"] = label
}

func (c *MainController) Get() {
	page := c.Ctx.Input.Param(":page")
	intpage, err := strconv.Atoi(page)
        if err != nil {
		intpage = 1
	}
        start := (intpage -1) * 10
	count, err := models.CountBlog()
        c.CheckErr(err, "get blog count error")
	pagetotal := int(count/10)
	if b := count % 10; b != 0 {
		pagetotal += 1;
	}
	blog, err := models.FindBlogByPageId(start)
        c.CheckErr(err, "get blog error")
	c.Data["Blog"] = blog
	c.Data["Page"] = intpage
	c.Data["Pagetotal"] = pagetotal
	c.Layout = "layout/main.tpl"
	c.TplName = "blog/index.tpl"
}

func (this *MainController) Resp(status bool, str string) {
        this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
        this.ServeJSON()
        this.StopRun()
}

func (this *MainController) CheckErr(err error, str string){
	if err != nil {
		this.Resp(false, str)
	}
}
