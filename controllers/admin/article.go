package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"strconv"
	"testblog/models"
	"testblog/utils/validations"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Get() {
	blog, err := models.FindAllBlog()
	c.CheckErr(err, "get all blog error")

	c.Data["Blog"] = blog
	c.Data["Pagetitle"] = "文章列表"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
	c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}

func (c *ArticleController) AddArticle() {
	section, err := models.FindAllSection()
	c.CheckErr(err, "query section from database error")
	label, err := models.FindAllLabel()
	c.CheckErr(err, "query label from database error")
	if c.isPost() {
		result := make(map[string]interface{})
		sid := c.GetString("section")
		cid := c.GetString("label")
		title := c.GetString("subject")
		content := c.GetString("content")
		sectionid, err := strconv.Atoi(sid)
		c.CheckErr(err, "sid atoi error")
		labelid, err := strconv.Atoi(cid)
		c.CheckErr(err, "labelid atoi error")

		valid := validation.Validation{}
		article := validations.ArticleForm{SectionId: sid, LabelId: cid, Title: title, Content: content}
		b, err := valid.Valid(&article)
		if err != nil {
			c.Resp(false, err.Error())
			return
		}
		if !b {
			c.Resp(false, "validation error")
			return
		}

		_, err = models.AddBlog(sectionid, labelid, title, content)
		if err == nil {
			c.Redirect("/admin/article", 302)
		} else {
			beego.Error("Insert into mysql error: ", err.Error())
			result["status"] = "false"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
			c.Data["json"] = result
			c.ServeJSON()
			c.StopRun()
		}
	}

	c.Data["Sections"] = section
	c.Data["Labels"] = label
	c.Data["Pagetitle"] = "新增文章"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/add.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "admin/editor.tpl"
}

func (c *ArticleController) EditArticle() {
	section, err := models.FindAllSection()
	c.CheckErr(err, "query section from database error")
	label, err := models.FindAllLabel()
	c.CheckErr(err, "query label from database error")

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.CheckErr(err, "get article id error")
	blog, err := models.FindBlogById(id)
	c.CheckErr(err, "get article content error")

	if c.isPost() {
		valid := validation.Validation{}
		result := make(map[string]interface{})
		sid := c.GetString("section")
		cid := c.GetString("label")
		title := c.GetString("subject")
		content := c.GetString("content")
		sectionid, err := strconv.Atoi(sid)
		c.CheckErr(err, "sid atoi error")
		labelid, err := strconv.Atoi(cid)
		c.CheckErr(err, "labelid atoi error")

		article := validations.ArticleForm{SectionId: sid, LabelId: cid, Title: title, Content: content}
		b, err := valid.Valid(&article)
		if err != nil {
			c.Resp(false, err.Error())
			return
		}
		if !b {

			for _, err := range valid.Errors {
				beego.Error("validation error: ", err.Key, err.Message)
				c.Data["json"] = &map[string]interface{}{"status": false, "key": err.Key, "message": err.Message}
			}
			c.ServeJSON()
			c.StopRun()
		}

		_, err = models.UpdateBlog(id, sectionid, labelid, title, content)
		if err == nil {
			c.Redirect("/admin/article", 302)
		} else {
			beego.Error("Insert into mysql error: ", err.Error())
			result["status"] = "false"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
			c.Data["json"] = result
			c.ServeJSON()
			c.StopRun()
		}
	}

	c.Data["Blog"] = blog
	c.Data["Sections"] = section
	c.Data["Labels"] = label
	c.Data["Pagetitle"] = "编辑文章"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/edit.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "admin/editor.tpl"
}

func (c *ArticleController) DelArticle() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		beego.Error("atoi error")
	}
	_, err = models.DelBlog(bid)
	if err == nil {
		beego.Informational("Delete blog success", id)
		c.Redirect("/admin/article", 302)
	}

}
