package service

import (
	"authentication/domain"
	"authentication/internal/config"
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

func (e emailService) SendEmailVerification(to string, otp string) error {
	// Load file template
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	pathDir := curDir + "/internal/util/email_otp.html"
	tmpl, err := template.ParseFiles(pathDir)
	if err != nil {
		log.Fatalf("Error when load template %s", err.Error())
		return err
	}

	// Data untuk di-assign ke template
	data := struct {
		OTP string
	}{
		OTP: otp,
	}

	// Buat buffer untuk menampung hasil eksekusi template
	buffer := new(bytes.Buffer)

	// Eksekusi template dengan data yang telah di-assign
	err = tmpl.Execute(buffer, data)
	if err != nil {
		log.Fatalf("Error when execute template %s", err.Error())
		return err
	}

	// Menggunakan buffer.Bytes() untuk mendapatkan hasil template dalam bentuk []byte
	emailContent := buffer.Bytes()

	// Lakukan pengiriman email menggunakan emailContent
	// Contoh pengiriman email menggunakan library pihak ketiga, seperti "gomail" atau "email"

	auth := smtp.PlainAuth("", e.cnf.Email.Username, e.cnf.Email.Password, e.cnf.Email.Host)
	msg := []byte("" +
		"From: jajanin.sinii <" + e.cnf.Email.Username + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + "Email Verification" + "\r\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		string(emailContent) + "\r\n")

	return smtp.SendMail(e.cnf.Email.Host+":"+e.cnf.Email.Port, auth, e.cnf.Email.Username, []string{to}, msg)
}
