package models

import (
    "github.com/astaxie/beego/orm"
    //"github.com/gogather/com"

)

type User struct {
    Id          int
    Name        string
    Password      string
    Created_at  string
}


func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(User))
}



func FindUser(username string) (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Name: username}
	err := o.Read(&user, "name")

	return user, err
}

