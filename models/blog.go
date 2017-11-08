package models

import (
	"github.com/astaxie/beego/orm"
)

type Blog struct {
	Id         int
	Section    *Section  `orm:"rel(fk)"`
	Label   *Label `orm:"rel(fk)"`
	Title      string
	Content    string
	Created_at string
	Updated_at string
}

func init() {
	orm.RegisterModel(new(Blog))
}



func FindAllBlog()([]Blog, error) {
        var blog []Blog
        o := orm.NewOrm()
	_, err := o.QueryTable("blog").RelatedSel().All(&blog)
	return blog, err

}

func FindBlogByPageId( start int)([]Blog, error){
	var blog []Blog
	o := orm.NewOrm()
	_, err := o.QueryTable("blog").RelatedSel().OrderBy("-id").Limit(10, start).All(&blog)
	return blog, err
}


func FindBlogById(id int) (Blog, error) {
	o := orm.NewOrm()
	o.Using("default")
	blog := Blog{Id: id}
	err := o.Read(&blog, "id")
	return blog, err
}

func CountBlog()(int64, error){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("blog").Count()
	return cnt, err
}

/*
func AddBlog(userid int, subid int, cateid int, title string, content string)(int64 , error){
	o := orm.NewOrm()
	o.Using("default")
	sql := "insert into blog(user_id, subject_id, category_id, title, content) values( ?, ?, ?, ?, ?)"
	res, err := o.Raw(sql, userid, subid,cateid,title,content).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}


func DelBlog(id int)(int64, error){
        o := orm.NewOrm()
        o.Using("default")
        if num, err := o.Delete(&Blog{Id: id}); err == nil {
                return num, err
	}else{
		return 0, err
	}
}
*/
