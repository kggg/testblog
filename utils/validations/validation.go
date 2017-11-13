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


