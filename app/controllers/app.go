package controllers

import (
	"crypto/tls"

	"fmt"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
	"github.com/revel/revel"
	gomail "gopkg.in/gomail.v2"
)

type App struct {
	*revel.Controller
}

type serviceBooking struct {
	name     string
	phone    string
	email    string
	fromDate string
	toDate   string
	address  string
	details  string
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) BookService(name string, phone string, email string, fromDate string, toDate string, address string, details string, recaptchaResponse string) revel.Result {

	recaptcha.Init("6LeSnxEUAAAAACMUFijc4z0iA1J3tzfRYW7DfLqv")

	captchaValid := false

	if len(recaptchaResponse) > 0 {
		captchaValid = recaptcha.Confirm("", recaptchaResponse)
	}

	if captchaValid {
		var newBooking = serviceBooking{name: name, phone: phone, email: email, fromDate: fromDate, toDate: toDate, address: address, details: details}

		go sendMail(newBooking)

		c.Response.Status = 201
		return c.Render()
	}

	c.Response.Status = 400
	return c.Render()
}

func sendMail(newBooking serviceBooking) {
	m := gomail.NewMessage()

	m.SetHeader("From", "info@ourfurryfriends.co.za")
	m.SetHeader("To", "emma@ourfurryfriends.co.za")
	m.SetHeader("Subject", "New Booking Request - "+newBooking.name)
	m.SetBody("text/html", buildEmailBody(newBooking))

	d := gomail.NewDialer("127.0.0.1", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Could not send email")
	}
}

func buildEmailBody(newBooking serviceBooking) string {
	d := "<h1>Service Request</h1>"

	d += "<p><b>From</b>: " + newBooking.name + "</p>"
	d += "<p><b>Phone</b>: " + newBooking.phone + "</p>"
	d += "<p><b>Email</b>: " + newBooking.email + "</p>"
	d += "<p><b>From Date</b>: " + newBooking.fromDate + "</p>"
	d += "<p><b>To Date</b>: " + newBooking.toDate + "</p>"
	d += "<p><b>Address</b>: " + newBooking.address + "</p>"
	d += "<p><b>details</b>: " + newBooking.details + "</p>"

	return d
}
