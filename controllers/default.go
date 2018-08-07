package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}
type FormContrller struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
// @router / [get]
func (c *FormContrller) Post(){
	f,h, err := c.GetFile ("uploadname")
	if err !=nil {
		fmt.Println("00000000000000")
	}
	defer f.Close()
	c.SaveToFile("uploadname","static/photos/" + h.Filename)
}

func (this *UserController) Get() {
	// this.Ctx.WriteString("厉害了我的哥")
	// this.Data["username"]=this.GetString("username")
	username :=this.GetString("username")
	this.Data["age"]= this.GetString("age")
	this.Data["username"] = username
	fmt.Println(this.Data["username"])
	this.TplName = "user.tpl"
}

func (this *ReadController) Get() {
	username := this.GetString("username")
	o := orm.NewOrm()
	user :=User{Id:username}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}else {
		fmt.Println(user.Id, user.Name)
	}
}

func (this *ReadOrCreate) Get() {
	username := this.GetString("username")
	o :=orm.NewOrm()
	user := User{Name:username}
	if created,id,err := o.ReadOrCreate(&user,"Name"); err ==nil {
		if created {
			fmt.Println("new insert an object id",id)
		} else {
			fmt.Print("get an object,Id:",id)
		}
	}
}

func (this *InsertCreate) Get() {
	o := orm.NewOrm()
	username := this.GetString("username")
	user := User{Name:username}
	id , err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}

func (this *InserMulti) Get() {
	username := this.GetString("username")
	user := []User{
		{Name:username},
		{Name:username},
	}
	successNums, err := o.InsertMulti(100,users)
}

func (this *Update) Get() {
	username := this.GetSrting("username")
	o := orm.NewOrm()
	user := User{Id:username}
	if o.Read(&user) == nil {
		user.Name = "MyName"
		if num, err :=o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}
}

func (this *Delete) Get() {
	o :=orm.NewOrm
	username := this.GetSrting("username")
	if num, err := o.Delete(&User{Id:username}); err == nil{
		fmt.Println(num)
	}
}

