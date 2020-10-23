package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

//https://hackernoon.com/golang-sendmail-sending-mail-through-net-smtp-package-5cadbe2670e0
//https://msdn.microsoft.com/en-us/library/ms526560(v=exchg.10).aspx
func main() {
	var (
		serverAddr = "smtp.gmail.com"
		password   = "vnxjpkcflijpbiwa"
		emailAddr  = "needkopi@gmail.com"
		portNumber = 465
		tos        = []string{
			"chandra.rizky@students.amikom.ac.id",
		}
		cc = []string{
			"needkopi@gmail.com.com",
		}
		attachmentFilePath = "/home/needkopi/Development/go/src/github.com/learn/mail-smtp/Mekanisme WFH Signed.pdf"
		filename           = "Mekanisme WFH Signed.pdf"
		delimeter          = "**=myohmy689407924327"
	)

	log.Println("======= Test Gmail client (with attachment) =========")
	log.Println("NOTE: user need to turn on 'less secure apps' options")
	log.Println("URL:  https://myaccount.google.com/lesssecureapps\n\r")

	tlsConfig := tls.Config{
		ServerName:         serverAddr,
		InsecureSkipVerify: true,
	}

	log.Println("Establish TLS connection")
	conn, connErr := tls.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, portNumber), &tlsConfig)
	if connErr != nil {
		log.Panic(connErr)
	}
	defer conn.Close()

	log.Println("create new email client")
	client, clientErr := smtp.NewClient(conn, serverAddr)
	if clientErr != nil {
		log.Panic(clientErr)
	}
	defer client.Close()

	log.Println("setup authenticate credential")
	auth := smtp.PlainAuth("", emailAddr, password, serverAddr)

	if err := client.Auth(auth); err != nil {
		log.Panic(err)
	}

	log.Println("Start write mail content")
	log.Println("Set 'FROM'")
	if err := client.Mail(emailAddr); err != nil {
		log.Panic(err)
	}
	log.Println("Set 'TO(s)'")
	for _, to := range tos {
		if err := client.Rcpt(to); err != nil {
			log.Panic(err)
		}
	}

	writer, writerErr := client.Data()
	if writerErr != nil {
		log.Panic(writerErr)
	}

	//basic email headers
	sampleMsg := fmt.Sprintf("From: %s\r\n", emailAddr)
	sampleMsg += fmt.Sprintf("To: %s\r\n", strings.Join(tos, ";"))
	if len(cc) > 0 {
		sampleMsg += fmt.Sprintf("Cc: %s\r\n", strings.Join(cc, ";"))
	}
	sampleMsg += "Subject: Golang example send mail in HTML format with attachment\r\n"

	log.Println("Mark content to accept multiple contents")
	sampleMsg += "MIME-Version: 1.0\r\n"
	sampleMsg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	//place HTML message
	log.Println("Put HTML message")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: 7bit\r\n"

	// get template html
	t := template.New("action")

	var err error
	t, err = t.ParseFiles("/home/needkopi/Development/go/src/github.com/learn/mail-smtp/action.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	data := map[string]interface{}{
		"name": "halo ini testing",
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		log.Println(err.Error())
		return
	}

	sampleMsg += fmt.Sprintf("\r\n%s", tpl.String())

	//place file
	log.Println("Put file attachment")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: base64\r\n"
	sampleMsg += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"
	//read file
	rawFile, fileErr := ioutil.ReadFile(attachmentFilePath)
	if fileErr != nil {
		log.Panic(fileErr)
	}
	sampleMsg += "\r\n" + base64.StdEncoding.EncodeToString(rawFile)

	//write into email client stream writter
	log.Println("Write content into client writter I/O")
	if _, err := writer.Write([]byte(sampleMsg)); err != nil {
		log.Panic(err)
	}

	if closeErr := writer.Close(); closeErr != nil {
		log.Panic(closeErr)
	}

	client.Quit()

	log.Print("done.")
}
