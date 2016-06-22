package main

import "net/smtp"
import "log"
import "html/template"
import "bytes"

type email struct {
	Name string
	Body string
}

func send(body string, to string) {
	myEmail := email{"John", body}

	tmpl, err := template.New("email").ParseFiles("./template.html")
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, myEmail)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	from := "youremail@gmail.com"
	password := "yourpassword"

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: Your messages subject" + "\r\n\r\n" +
		buff.String() + "\r\n"

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("message sent")
}

func main() {
	send("Write your message here.", "recipient@gmail.com")
}
