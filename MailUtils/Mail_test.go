package MailUtils

import (
	"fmt"
	"testing"
)

func TestSendToMail(t *testing.T) {
	fmt.Println("开始测试邮件发送")
	t.Run("开始测试发送邮件:", test)
}

func test(t *testing.T) {
	mail := Mail{}
	mail.InitMailServer()
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
		</body>
		</html>`
	mail.InitMailBody("验证码", body, "1784929126@qq.com")
	mail.SendMail()
}

func testSendToMail(t *testing.T) {
	user := "liweijun0302@163.com"
	password := "ZZYWKVVUAFAMEXKA"
	host := "smtp.163.com:25"
	to := "1784929126@qq.com"

	subject := "使用Golang发送邮件"

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
		</body>
		</html>`

	sendUserName := "Micros0ft" //发送邮件的人名称
	fmt.Println("send email")
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
