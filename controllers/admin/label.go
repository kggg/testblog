package admin

import (
        "github.com/astaxie/beego/validation"
        //"html/template"
        "testblog/models"
        "testblog/utils/validations"
	"strconv"
	"github.com/astaxie/beego"
)

type LabelController struct {
        BaseController
}

func (c *LabelController) Get(){
	label, err := models.FindAllLabel()
	c.CheckErr(err, "get all label info error")
	section , err := models.FindAllSection()
	c.CheckErr(err, "get section info error")
        c.Data["Label"] = label
        c.Data["Sections"] = section
        c.Data["Pagetitle"] = "标签列表"
        c.TplName = "admin/label/index.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}

func (c *LabelController) AddLabel(){
        section , err := models.FindAllSection()
        c.CheckErr(err, "get section info error")
	if c.isPost(){
		label := c.GetString("label")
		sid := c.GetString("section")
		section, err := strconv.Atoi(sid)
		c.CheckErr(err, "sid atoi error")
		valid := validation.Validation{}
		ll := validations.LabelForm{Label: label,  Section: sid}
                b, err := valid.Valid(&ll)
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
		exist := models.CheckLabelExist(label)
		if exist {
			c.Resp(false, "the label name has been existing")
		}
                _, err = models.AddLabel(section, label)
                if err == nil {
                        c.Redirect("/admin/label", 302)
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
        c.Data["Sections"] = section
        c.Data["Pagetitle"] = "新增标签"
        c.TplName = "admin/label/add.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = ""
        c.LayoutSections["Scripts"] = "admin/editor.tpl"
}

func (c *LabelController) EditLabel(){
        section , err := models.FindAllSection()
        c.CheckErr(err, "get section info error")
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	labels , err := models.FindLabelById(id)
	c.CheckErr(err, "id atoi error")
        if c.isPost(){
                label := c.GetString("label")
                sid := c.GetString("section")
                section, err := strconv.Atoi(sid)
                c.CheckErr(err, "sid atoi error")
                valid := validation.Validation{}
                ll := validations.LabelForm{Label: label,  Section: sid}
                b, err := valid.Valid(&ll)
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
                //exist := models.CheckLabelExist(label)
                //if exist {
                //        c.Resp(false, "the label name has been existing")
                //}
                _, err = models.UpdateLabel(id, section, label)
                if err == nil {
                        c.Redirect("/admin/label", 302)
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
	c.Data["Label"] = labels
        c.Data["Sections"] = section
        c.Data["Pagetitle"] = "编辑标签"
        c.TplName = "admin/label/edit.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/datatable-css.tpl"
        c.LayoutSections["Scripts"] = "admin/datatable-js.tpl"
}


func (c *LabelController) DelLabel(){
        id := c.Ctx.Input.Param(":id")
        bid, err := strconv.Atoi(id)
        if err != nil {
                beego.Error("atoi error")
        }
        _, err = models.DelLabel(bid)
        if err == nil {
                beego.Informational("Delete blog success", id)
                c.Redirect("/admin/label", 302)
        }
}


