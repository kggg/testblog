package admin

import (
        //"html/template"
        "testblog/models"
	"strconv"
	"github.com/astaxie/beego"
)

type SectionController struct {
        BaseController
}

func (c *SectionController) Get(){
	section , err := models.FindAllSection()
	c.CheckErr(err, "get section info error")
        c.Data["Sections"] = section
        c.Data["Pagetitle"] = "板块列表"
        c.TplName = "admin/section/index.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}

func (c *SectionController) AddSection(){
	if c.isPost(){
		section := c.GetString("section")
		exist := models.CheckSectionExist(section)
		if exist {
			c.Resp(false, "the section name has been existing")
		}
		_, err := models.AddSection(section)
                if err == nil {
                        c.Redirect("/admin/section", 302)
                } else {
			result := make(map[string]interface{})
                        result["status"] = "false"
                        result["message"] = "数据创建失败"
                        result["debug"] = err.Error()
                        c.Data["json"] = result
                        c.ServeJSON()
                        c.StopRun()
                }

	}
        c.Data["Pagetitle"] = "新增板块"
        c.TplName = "admin/section/add.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = ""
        c.LayoutSections["Scripts"] = "admin/editor.tpl"
}

func (c *SectionController) EditSection(){
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	sections , err := models.FindSectionById(id)
	c.CheckErr(err, "id atoi error")
        if c.isPost(){
                section := c.GetString("section")
                _, err = models.UpdateSection(id, section)
                if err == nil {
                        c.Redirect("/admin/section", 302)
                } else {
                        result := make(map[string]interface{})
                        result["status"] = "false"
                        result["message"] = "数据创建失败"
                        result["debug"] = err.Error()
                        c.Data["json"] = result
                        c.ServeJSON()
                        c.StopRun()
                }
	}
        c.Data["Sections"] = sections
        c.Data["Pagetitle"] = "编辑板块"
        c.TplName = "admin/section/edit.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}


func (c *SectionController) DelSection(){
        id := c.Ctx.Input.Param(":id")
        bid, err := strconv.Atoi(id)
        if err != nil {
                beego.Error("atoi error")
        }
        _, err = models.DelSection(bid)
        if err == nil {
                beego.Informational("Delete blog success", id)
                c.Redirect("/admin/section", 302)
        }
}


