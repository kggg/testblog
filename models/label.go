package models

import (
        "github.com/astaxie/beego/orm"
)

type Label struct {
        Id         int
	Section    *Section   `orm:"rel(fk)"`
	Label       string
        Created_at string
}

func init() {
        orm.RegisterModel(new(Label))
}

func FindAllLabel()([]Label, error){
        var label []Label
        o := orm.NewOrm()
        _, err := o.QueryTable("label").RelatedSel().All(&label)
        return label, err
}

