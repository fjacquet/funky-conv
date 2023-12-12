package main

import "net/smtp"

func sendMain(text string) error {
	// msg := "Subject: Encoding Failure Alert\n\nThe encoding job failed."
	return smtp.SendMail(cfg.SMTP.Server, nil, cfg.SMTP.From, []string{cfg.SMTP.To}, []byte(text))

}
