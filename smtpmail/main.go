package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "needkopi@gmail.com"
const CONFIG_PASSWORD = "vnxjpkcflijpbiwa"

func main() {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_EMAIL)
	mailer.SetHeader("To", "muhammadrudyy@gmail.com")
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Ngajak gelot")
	mailer.SetBody("text/html", "<h1>HALLO CUK!! PIE KABARE? ISEH PENAK JAMANKU TO?</h1>")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_EMAIL,
		CONFIG_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Mail Sent!!")
}
