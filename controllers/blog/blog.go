package blog

import (
	"testblog/models"
	"strconv"
	"strings"
)

type MainController struct {
	BaseController
}


func (c *MainController) Get() {
	page := c.Ctx.Input.Param(":page")
	intpage, err := strconv.Atoi(page)
        if err != nil {
		intpage = 1
	}
        start := (intpage -1) * 5
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


func (c *MainController) FindSection() {
        var page int
	c.Ctx.Input.Bind(&page, "page")
        if page == 0  {
                page = 1
        }
        start := (page -1) * 20
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.CheckErr(err, "get section id error")
        count, err := models.CountBlogBySectionId(id)
        c.CheckErr(err, "get blog count error")
        pagetotal := int(count/20)
        if b := count % 20; b != 0 {
                pagetotal += 1;
        }
        blog, err := models.FindBlogBySectionId(id, start)
        c.CheckErr(err, "get blog error")
	c.Data["Sid"] = id
        c.Data["Blog"] = blog
        c.Data["Page"] = page
        c.Data["Pagetotal"] = pagetotal
        c.Layout = "layout/main.tpl"
        c.TplName = "blog/section/section.tpl"
}

func (c *MainController) FindLabel() {
        var page int
        c.Ctx.Input.Bind(&page, "page")
        if page == 0  {
                page = 1
        }
        start := (page -1) * 20
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
        c.CheckErr(err, "get Label id error")
        count, err := models.CountBlogByLabelId(id)
        c.CheckErr(err, "get blog count error")
        pagetotal := int(count/20)
        if b := count % 20; b != 0 {
                pagetotal += 1;
        }
        blog, err := models.FindBlogByLabelId(id, start)
        c.CheckErr(err, "get blog error")
        c.Data["Lid"] = id
        c.Data["Blog"] = blog
        c.Data["Page"] = page
        c.Data["Pagetotal"] = pagetotal
        c.Layout = "layout/main.tpl"
        c.TplName = "blog/label/label.tpl"
}


func (c *MainController) SearchBlog() {
	search := c.GetString("search")
	search = strings.TrimSpace(search)
	//blog, num, err := models.SearchBlog(search)
	blog, err := models.SearchBlog(search)
	c.CheckErr(err, "get search key from data error")
        /*
	*/
	c.Data["Blog"] = blog
        c.Layout = "layout/main.tpl"
        c.TplName = "blog/search/search.tpl"
}


type ShowController struct {
	BaseController
}

func (c *ShowController) Get(){
	pid, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.CheckErr(err, "blog id atoi error")
	blog, err := models.FindBlogById(pid)
	c.CheckErr(err, "get blog by id error")
	c.Data["Blog"] = blog
	c.Layout = "layout/main.tpl"
	c.TplName = "blog/show.tpl"

}
