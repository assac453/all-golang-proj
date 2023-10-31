package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	from := "asdasd@mail.ru"
	pass := "asdsad"

	to := []string{
		"asdasd@mial.ru",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message1 := []byte("Text for message. some text to send")

	auth1 := smtp.PlainAuth("", from, pass, smtpHost)

	err := smtp.SendMail(fmt.Sprintf("%v:%v", smtpHost, smtpPort), auth1, from, to, message1)
	if err != nil {
		panic(err)
	}
	fmt.Println("email sent success")
}
