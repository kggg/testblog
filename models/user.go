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



func FindUserByName(username string) (User, error) {
	o := orm.NewOrm()
	user := User{Name: username}
	err := o.Read(&user, "name")

	return user, err
}

func FindUserById(id int) (User, error) {
        o := orm.NewOrm()
        user := User{Id: id}
        err := o.Read(&user, "id")

        return user, err
}

func FindAllUser()([]User, error){
	o := orm.NewOrm()
	var users []User
	_, err := o.QueryTable("user").All(&users);
	return users, err
}


func CheckUserExist(str string) bool{
        o := orm.NewOrm()
        exist := o.QueryTable("user").Filter("name", str).Exist()
        return exist
}

func AddUser(name, password string)(int64, error){
        o := orm.NewOrm()
        sql := "insert into user(name, password) values( ?, ?)"
        res, err := o.Raw(sql, name, password).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }
}

func UpdateUser(id int,  name, password string)(int64 , error){
        o := orm.NewOrm()
        sql := "update user set name=?,password=? where id=?"
        res, err := o.Raw(sql, name,password, id).Exec()
        if nil != err {
                return 0, err
        } else {
                return res.LastInsertId()
        }

}


func DelUser(id int)(int64, error){
        o := orm.NewOrm()
        if num, err := o.Delete(&User{Id: id}); err == nil {
                return num, err
        }else{
                return 0, err
        }
}

