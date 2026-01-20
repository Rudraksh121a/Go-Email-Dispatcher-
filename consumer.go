package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "LocalHost"
		smtpPort := "1025"

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Woker :%d error  parsing  template for  %s  ", id, recipient.Email)
			continue
		}
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "rudraksh@rud.com", []string{recipient.Email}, []byte(msg))
		fmt.Printf("Worker %d: sending email to  %s \n", id, recipient.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, recipient)
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: sent email to  %s \n", id, recipient.Email)

	}

}
