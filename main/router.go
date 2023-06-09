package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"one_time_pad_service/MailUtils"
	"one_time_pad_service/Otp"
	"one_time_pad_service/User"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.Handle("/css/img/", http.StripPrefix("/css/img/", http.FileServer(http.Dir("../css/img/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))
	http.Handle("/mod/", http.StripPrefix("/mod/", http.FileServer(http.Dir("../mod"))))

	http.HandleFunc("/loginIndex", LoginIndex)

	http.HandleFunc("/login", Login)

	http.HandleFunc("/registerIndex", RegisterIndex)

	http.HandleFunc("/register", Register)

	http.HandleFunc("/", LoginIndex)

	http.HandleFunc("/sendVerifyCode", SendVerifyCode)

	http.HandleFunc("/add", Add)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
}

func RegisterIndex(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/register.html")
	if err != nil {
		fmt.Println("解析模版失败：", err)
	}
	files.Execute(w, "")
}

func SendVerifyCode(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	fmt.Println("监测到发送验证码按钮：")
	data := r.URL.RawQuery
	//fmt.Println(data)
	rawMail := strings.Split(data, "=")
	//fmt.Println(rawMail)
	mail := rawMail[1]
	fmt.Println("mail", mail)
	rand.Seed(time.Now().UnixNano())
	VerifyNum := rand.Intn(900000) + 99999

	send := MailUtils.Mail{}
	send.InitMailServer()
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			验证码: ` + strconv.Itoa(VerifyNum) + "\n" + `
		</body>
		</html>`
	send.InitMailBody("Micros0ft验证码", body, mail)
	send.SendMail()
	fmt.Println("发送成功")

	// 接入数据库
	user := User.User{}
	user.InitMysql()
	prepare, err := user.Db.Prepare("insert into user(username,password,email,verify_code,is_verify) values (?,?,?,?,?)")
	if err != nil {
		fmt.Println("sql预编译错误:", err)
		return
	}
	_, err = prepare.Exec("", "", mail, strconv.Itoa(VerifyNum), 0)
	if err != nil {
		fmt.Println("插入数据库失败:", err)
		return
	}
	fmt.Println("成功插入数据库:")
}

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("../mod/login.html")
	files.Execute(w, "")
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析Form表单失败:", err)
		return
	}
	form := r.PostForm
	var email string
	var password string
	for k, v := range form {
		fmt.Printf("[%v : %v]\n", k, v)
		if k == "email" {
			email = v[0]
		} else if k == "password" {
			password = v[0]
		}
	}

	user := User.User{}
	user.InitMysql()
	prepare, err := user.Db.Prepare("select password,is_verify from user where email=?")
	if err != nil {
		fmt.Println("解析sql语句错误:", err)
		return
	}
	row := prepare.QueryRow(email)
	var databasePassword string
	var databaseIsVeryfy string
	err = row.Scan(&databasePassword, &databaseIsVeryfy)
	if err != nil {
		fmt.Println("读取数据库失败:", err)
		return
	}
	fmt.Printf("数据库中的数据[%T : %v][%T : %v]\n", databasePassword, databasePassword, databaseIsVeryfy, databaseIsVeryfy)
	file1, err := template.ParseFiles("../mod/login.html")
	file2, err := template.ParseFiles("../mod/index.html")

	if databaseIsVeryfy == "1" {
		if databasePassword == password {
			fmt.Println("身份验证成功:")
			err := file2.Execute(w, email)
			if err != nil {
				fmt.Println("登录失败:", err)
				return
			}
		} else {
			fmt.Println("信息不匹配:")
			file1.Execute(w, "信息不匹配")
		}
	} else {
		fmt.Println("账号还未完成注册:")
		file1.Execute(w, "账号还未完成注册")
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.PostForm
	var nickname string
	var email string
	var password string
	var verifyCode string

	for k, v := range form {
		fmt.Printf("[%v : %v]\n", k, v)
		if k == "nickname" {
			nickname = v[0]
		} else if k == "email" {
			email = v[0]
		} else if k == "verifyCode" {
			verifyCode = v[0]
		} else if k == "password" {
			password = v[0]
		}
	}
	user := User.User{}
	user.InitMysql()
	prepare, err := user.Db.Prepare("select verify_code from user where email=?")
	if err != nil {
		fmt.Println("解析sql语句失败:", err)
		return
	}
	row := prepare.QueryRow(email)
	var databaseVerifyCode string
	err = row.Scan(&databaseVerifyCode)
	if err != nil {
		fmt.Println("获取数据库数据失败:", err)
		return
	}

	files, _ := template.ParseFiles("../mod/register.html")

	if databaseVerifyCode == verifyCode {
		fmt.Println("验证成功：准备存入数据库")
		stmt, err := user.Db.Prepare("update user set username=?,password=?,is_verify=? where email=?")
		if err != nil {
			fmt.Println("解析sql语句失败:", err)
			return
		}
		fmt.Println("Password 到底去哪了:", password)
		_, err = stmt.Exec(nickname, password, strconv.Itoa(1), email)
		if err != nil {
			fmt.Println("修改数据库失败:", err)
			return
		}
		files.Execute(w, "注册成功")
	} else {
		fmt.Println("验证码错误")
		files.Execute(w, "验证码错误,请再次输入或重新获取验证码")
	}

}

func Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析Form表单失败:", err)
		return
	}
	form := r.PostForm
	var url string
	var intro string
	var uname string
	var pwd string
	var email string
	for i, v := range form {
		//fmt.Printf("[%v : %v]\n", i, v)
		if i == "url" {
			url = v[0]
		} else if i == "intro" {
			intro = v[0]
		} else if i == "uname" {
			uname = v[0]
		} else if i == "pwd" {
			pwd = v[0]
		} else if i == "email" {
			email = v[0]
		}
	}
	fmt.Println("url:", url)
	fmt.Println("intro:", intro)
	fmt.Println("uname:", uname)
	fmt.Println("pwd:", pwd)
	fmt.Println("email:", email)
	otp := Otp.Otp{}
	otp.InitMysql()
	otp.AddOtp(url, intro, uname, pwd, email)

	var tmpOtpObjs []Otp.Otp
	prepare, err := otp.Db.Prepare("select * from PwdTable where email=?")
	if err != nil {
		fmt.Println("预编译sql失败:", err)
		return
	}
	rows, err := prepare.Query(email)
	defer rows.Close()
	if err != nil {
		fmt.Println("执行sql失败", err)
		return
	}
	for rows.Next() {
		var tmp Otp.Otp
		err = rows.Scan(&tmp.Id, &tmp.Url, &tmp.Intro, &tmp.Uname, &tmp.Pwd, &tmp.Email)
		tmpOtpObjs = append(tmpOtpObjs, tmp)
	}

	files, err := template.ParseFiles("../mod/otp.html")
	if err != nil {
		fmt.Println("解析模板失败:", err)
		return
	}
	err = files.Execute(w, tmpOtpObjs)
	if err != nil {
		fmt.Println("执行模板失败:", err)
		return
	}
}
