package models

import (
        "github.com/astaxie/beego/orm"
)

type Section struct {
        Id         int
	Section    string
        Created_at string
}

func init() {
        orm.RegisterModel(new(Section))
}

func FindAllSection()([]Section, error){
        var section []Section
        o := orm.NewOrm()
        _, err := o.QueryTable("section").All(&section)
        return section, err
}
