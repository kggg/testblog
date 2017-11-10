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
	_, err := o.QueryTable("blog").RelatedSel().OrderBy("-id").Limit(5, start).All(&blog)
	return blog, err
}


func FindBlogById(id int) (Blog, error) {
	o := orm.NewOrm()
	blog := Blog{Id: id}
	err := o.Read(&blog, "id")
	return blog, err
}

func FindBlogBySectionId(id, start int)([]Blog, error){
        var blog []Blog
        o := orm.NewOrm()
        _, err := o.QueryTable("blog").Filter("section_id", id).RelatedSel().OrderBy("-id").Limit(20, start).All(&blog)
        return blog, err
}

func FindBlogByLabelId(id, start int)([]Blog, error){
        var blog []Blog
        o := orm.NewOrm()
        _, err := o.QueryTable("blog").Filter("label_id", id).RelatedSel().OrderBy("-id").Limit(20, start).All(&blog)
        return blog, err
}

func SearchBlog(str string)([]Blog,   error){
        var blog []Blog
        o := orm.NewOrm()
	cond := orm.NewCondition()
        cond1 := cond.And("title__icontains", str).Or("content__icontains", str)
        qs := o.QueryTable("blog")
        qs = qs.SetCond(cond1)
	_, err := qs.All(&blog)
	//sql := "select * from blog where title like '%?%' or content like '%?%'"
        //num, err := o.Raw(sql, str, str).QueryRows(&blog)
        //return blog, num, err
        //_, err := o.Raw(sql, str, str).QueryRows(&blog)
        return blog, err
}


func CountBlog()(int64, error){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("blog").Count()
	return cnt, err
}

func CountBlogBySectionId(id int)(int64 , error){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("blog").Filter("section_id", id).Count()
	return cnt, err
}

func CountBlogByLabelId(id int)(int64 , error){
        o := orm.NewOrm()
        cnt, err := o.QueryTable("blog").Filter("label_id", id).Count()
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
