package validations

import (
    "github.com/astaxie/beego/validation"
    "strings"
)

type  ArticleForm struct{
        SectionId   string   `valid:"Required;Numeric"`
        LabelId  string      `valid:"Required;Numeric"`
        Title  string     `valid:"Required;MaxSize(100)"`
        Content  string   `valid:"Required"`
}


//  Article
func (art *ArticleForm) Valid(v *validation.Validation){
    art.Title = strings.TrimSpace(art.Title)
}

type LabelForm struct{
	Label  string  `valid:"Required;MaxSize(100)`
	Section  string  `valid:"Required;Numeric"`
}

func (l *LabelForm) Valid(v *validation.Validation){
    l.Label = strings.TrimSpace(l.Label)
}


type UserForm struct{
	Name     string   `valid:"Required;MaxSize(100)`
	Password  string   `valid:"Required;MaxSize(150)`
	Repassword  string  `valid:"Required;MaxSize(150)`
}

func (u *UserForm) Valid(v *validation.Validation){
    if u.Password != u.Repassword {
        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
        v.SetError("Password","the password not same")
    }
    u.Name = strings.TrimSpace(u.Name)
    u.Password = strings.TrimSpace(u.Password)
    u.Repassword = strings.TrimSpace(u.Repassword)
}
