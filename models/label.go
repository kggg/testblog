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

func FindLabelByName(str string)(Label, error){
	o := orm.NewOrm()
	var label Label
	err := o.QueryTable("label").Filter("label", str).One(&label)
	return label, err
}

func FindLabelById(id int)(Label, error){
        o := orm.NewOrm()
        var label Label
        err := o.QueryTable("label").Filter("id", id).RelatedSel().One(&label)
        return label, err
}


func CheckLabelExist(str string) bool{
	o := orm.NewOrm()
	exist := o.QueryTable("label").Filter("label", str).Exist()
	return exist
}

func AddLabel(sid int, label string)(int64, error){
	o := orm.NewOrm()
        sql := "insert into label(section_id, label) values( ?, ?)"
        res, err := o.Raw(sql, sid,label).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }
}

func UpdateLabel(id int, sid int, label string)(int64 , error){
        o := orm.NewOrm()
        sql := "update label set section_id=?,label=? where id=?"
        res, err := o.Raw(sql, sid,label, id).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }

}

func DelLabel(id int)(int64 , error){
        o := orm.NewOrm()
        if num, err := o.Delete(&Label{Id: id}); err == nil {
                return num, err
        }else{
                return 0, err
        }
}
