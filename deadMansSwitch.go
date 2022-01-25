package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"sync"
	"time"
)

var timeIncrement time.Duration = 5
var password string = "password"
var foo int = 0
var wg sync.WaitGroup

func sendPotentiallyTreasonousEmail() {
	from := "whistleblower@cia.gov"
	password := "emailpassword123"
	to := []string{
		"journalist@mail.com",
	}
	smtpHost := "email.com"
	smtpPort := "993"
	message := []byte("powerful impact boom from da cannon")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully. Rest in peace : ( . . .")
}

func waitSomeTime() {
	wg.Add(1)
	time.Sleep(timeIncrement * time.Second)
	if foo == 1 {
		foo = 0
		main()
	} else {
		fmt.Println("\nRest in peace : (\nCOVER ME!!!! dududududududu ch ch pew pew dududu pew pew dudu")
		sendPotentiallyTreasonousEmail()
		os.Exit(1)
	}
	wg.Done()
}

func main() {
	go waitSomeTime()
	var stra string
	fmt.Printf("Enter Password:")
	fmt.Scanln(&stra)
	if stra == password {
		fmt.Println("Password Correct!")
		foo = 1
		stra = "meow"
	} else {
		sendPotentiallyTreasonousEmail()
		foo = 0
		os.Exit(1)
	}
	wg.Wait()
}
