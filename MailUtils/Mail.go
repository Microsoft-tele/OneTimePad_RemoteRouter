package MailUtils

import (
	"fmt"
	"net/smtp"
	"strings"
)

// Mail user := "liweijun0302@163.com"
//password := "ZZYWKVVUAFAMEXKA"
//host := "smtp.163.com:25"
//to := "1784929126@qq.com"

// Mail 邮件类
type Mail struct {
	ServerSmtpUsername string
	ServerSmtpPassword string
	ServerSmtpHost     string
	SendUserName       string // 发送邮件的人的名称
	ServerSmtpTo       string
	Subject            string
	Body               string
}

// InitMailServer 初始化邮箱服务器
func (m *Mail) InitMailServer() {
	m.ServerSmtpUsername = "liweijun0302@163.com"
	m.ServerSmtpPassword = "ZZYWKVVUAFAMEXKA"
	m.ServerSmtpHost = "smtp.163.com:25"
}

// InitMailBody subject:主题 body:邮件主体 to:目的邮箱
func (m *Mail) InitMailBody(subject string, body string, to string) {
	m.Subject = subject
	m.SendUserName = "Micros0ft"
	m.Body = body
	m.ServerSmtpTo = to
}

func (m *Mail) SendMail() {
	err := SendToMail(m.ServerSmtpUsername, m.SendUserName, m.ServerSmtpPassword, m.ServerSmtpHost, m.ServerSmtpTo, m.Subject, m.Body, "html")
	if err != nil {
		fmt.Println("发送邮件失败，请检查原因：", err)
		return
	}
}

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
