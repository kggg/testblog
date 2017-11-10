package blog

import (
        "github.com/astaxie/beego"
        "testblog/models"
)

type BaseController struct {
        beego.Controller
}

func (c *BaseController) Prepare() {
        section, err := models.FindAllSection()
        c.CheckErr(err, "get section error")
        type SectionCount struct {
                Id    int
                Section  string
                Count   int
        }
        var SC []SectionCount
        for _, v := range section{
                cont, err := models.CountBlogBySectionId(v.Id)
                c.CheckErr(err, "count section id error")
                cnt := int(cont)
                SC = append(SC, SectionCount{Id:v.Id,Section:v.Section,Count: cnt})
        }

        label, err := models.FindAllLabel()
        c.CheckErr(err, "get label error")
        c.Data["section"] = SC
        c.Data["label"] = label
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

