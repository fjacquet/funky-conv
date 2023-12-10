package main

import "net/smtp"

func sendMain(text string) error {
	// msg := "Subject: Encoding Failure Alert\n\nThe encoding job failed."
	return smtp.SendMail(Cfg.SMTP.Server, nil, Cfg.SMTP.From, []string{Cfg.SMTP.To}, []byte(text))

}
