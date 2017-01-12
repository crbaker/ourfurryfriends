package controllers

import (
	"crypto/tls"

	"github.com/revel/revel"
	gomail "gopkg.in/gomail.v2"
)

type App struct {
	*revel.Controller
}

type serviceBooking struct {
	name    string
	phone   string
	email   string
	address string
	details string
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) BookService(name string, phone string, email string, address string, details string) revel.Result {

	var newBooking = serviceBooking{name: name, phone: phone, email: email, address: address, details: details}

	sendMail(newBooking)

	return c.Render()
}

func sendMail(newBooking serviceBooking) {
	m := gomail.NewMessage()

	m.SetHeader("From", newBooking.email)
	// m.SetHeader("To", "emma@ourfurryfriends.co.za")
	m.SetHeader("To", "crbaker@gmail.com")
	m.SetHeader("Subject", "New Booking Request - ")
	m.SetBody("text/html", buildEmailBody(newBooking))

	d := gomail.NewDialer("127.0.0.1", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func buildEmailBody(newBooking serviceBooking) string {
	d := "<h1>Service Request</h1>"

	d += "From: " + newBooking.name + "</br>"
	d += "Email: " + newBooking.email + "</br>"
	d += "Address: " + newBooking.address + "</br>"
	d += "details: " + newBooking.details + "</br>"

	return d
}
