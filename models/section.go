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

func FindSectionById(id int)(Section, error){
        o := orm.NewOrm()
        var section Section
        err := o.QueryTable("section").Filter("id", id).One(&section)
        return section, err
}


func CheckSectionExist(str string) bool{
        o := orm.NewOrm()
        exist := o.QueryTable("section").Filter("section", str).Exist()
        return exist
}

func AddSection(section string)(int64, error){
        o := orm.NewOrm()
        sql := "insert into section(section) values( ?)"
        res, err := o.Raw(sql, section).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }
}

func UpdateSection(id int, section string)(int64 , error){
        o := orm.NewOrm()
        sql := "update section set section=? where id=?"
        res, err := o.Raw(sql, section, id).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }

}

func DelSection(id int)(int64 , error){
        o := orm.NewOrm()
        if num, err := o.Delete(&Section{Id: id}); err == nil {
                return num, err
        }else{
                return 0, err
        }
}

