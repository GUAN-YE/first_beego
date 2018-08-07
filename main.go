package main

import (
	_ "ceshi_go/routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxiw/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:root@/orm_test?charset=utf8")
}

func main() {
	// beego.BConfig.WebConfig.Session.SessionOn = true
	o:=orm.NewOrm()
	o.Using("default")
	profile := new(Profile)
	profile.Age = 30
	user := new(User)
	user.Profile = profile
	user.Name = "slene"
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	beego.Run()
}

