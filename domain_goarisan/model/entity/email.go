package entity

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *User, toEmail string, data *EmailData, file, templatee string) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	from := os.Getenv("FromEmailAddr")
	password := os.Getenv("SMTPpwd")
	to := toEmail

	host := "smtp.gmail.com"
	port := 587

	var body bytes.Buffer

	template, err := ParseTemplateDir(templatee)
	if err != nil {
		log.Fatal("Could not parse template", err)
	}
	template.ExecuteTemplate(&body, file, &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	// m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(host, port, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
