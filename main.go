package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	sendWithNoAuth()
	sendWithAuth()
}

func sendWithAuth() {
	// 配置 MailPit 的 SMTP 服务器和凭据
	smtpServer := "localhost"
	smtpPort := 1026
	username := "user1"
	password := "password1"

	// 创建认证
	auth := smtp.PlainAuth("", username, password, smtpServer)

	// 创建电子邮件消息
	from := "sender@example.com"
	to := []string{"recipient@example.com"}
	subject := "Hello from Go!"
	body := "This is the email content."

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to[0], subject, body)

	// 发送邮件
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, to, []byte(message))
	if err != nil {
		log.Fatal("发送邮件时出现错误:", err)
	}

	fmt.Println("邮件发送成功")
}

func sendWithNoAuth() {
	// 配置 MailPit 的 SMTP 服务器和凭据
	smtpServer := "localhost"
	smtpPort := 1025

	// 创建电子邮件消息
	from := "sender@example.com"
	to := []string{"recipient@example.com"}
	subject := "Hello from Go!"
	body := "This is the email content."

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to[0], subject, body)

	// 发送邮件
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), nil, from, to, []byte(message))
	if err != nil {
		log.Fatal("发送邮件时出现错误:", err)
	}
	fmt.Println("邮件发送成功")
}
