package Otp

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
)

type Otp struct {
	Id    string
	Url   string
	Intro string
	Uname string
	Pwd   string
	Email string
	Db    *sql.DB
}

func (o *Otp) InitMysql() {
	o.Db, err = sql.Open("mysql", "root:660967@tcp(192.168.88.129:3306)/users") // @ TODO 记得修改数据库地址，最好是固定一个ip
	if err != nil {
		fmt.Println("初始化数据库失败:", err)
	}
}

func (o *Otp) AddOtp(url string, intro string, uname string, pwd string, email string) {
	o.InitMysql()
	prepare, err := o.Db.Prepare("insert into PwdTable(Url,Intro,Uname,Pwd,Email) values (?,?,?,?,?)")
	if err != nil {
		fmt.Println("预编译sql失败:", err)
		return
	}
	_, err = prepare.Exec(url, intro, uname, pwd, email)
	if err != nil {
		fmt.Println("插入数据库失败:", err)
		return
	}
	fmt.Println("插入数据库成功")
}
