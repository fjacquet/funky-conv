package main

import "net/smtp"

func sendMain(text string) {
	msg := "Subject: Encoding Failure Alert\n\nThe encoding job failed."
	err := smtp.SendMail("smtp.example.com:587", auth, "alert@example.com", []string{"admin@example.com"}, []byte(msg))

}
