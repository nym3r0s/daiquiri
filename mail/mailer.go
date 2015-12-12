package mail

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
)

type MailConfig struct {
	SmtpUsername string
	SmtpPassword string
	SmtpHostname string
	SmtpId       string
}

// Global Var Helper
var mail_cred MailConfig

func Config_Init(path string) {

	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	config := MailConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	// set global var
	mail_cred = config
	// Debug
	fmt.Println("Mail Config", config)
}

func SendMail(addr string, mailbody string) {
	auth := smtp.PlainAuth(mail_cred.SmtpId, mail_cred.SmtpUsername, mail_cred.SmtpPassword, mail_cred.SmtpHostname)
	msg := []byte(mailbody)

	to := []string{addr}
	go func() {
		err := smtp.SendMail(mail_cred.SmtpHostname+":1025", auth, "noreply@api.daiquiri.org", to, msg)
		if err != nil {
			fmt.Println("Mail Error", err)
		} else {
			fmt.Println("Mail Success", err, auth)
		}
	}()
}
