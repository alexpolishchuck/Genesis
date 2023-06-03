package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"io"
	"log"
	"net/http"
	"net/url"
)

type http_request_handler struct {
	Port         string
	File_handler file_handler
}

// currency rate
//subscribe
//send new currency

func (h http_request_handler) get_subscribe(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		load_template(writer, "templates/subscribe.html", nil)

	case "POST":
		request.ParseForm()
		email := request.Form["email"][0]
		h.File_handler.save_email(email)

	case "DELETE":
		params, _ := url.ParseQuery(request.URL.RawQuery)
		h.File_handler.delete_email(params.Get("email"))
	}
}

func (h http_request_handler) get_welcome_page(writer http.ResponseWriter, request *http.Request) {
	load_template(writer, "templates/welcome_page.html", nil)
}

func (h http_request_handler) send_to_all(writer http.ResponseWriter, request *http.Request) {

	currency_rate := h.request_currency_rate()

	email_body := "BTS to UAH: " + fmt.Sprintf("%f", currency_rate.Rates.BTC)

	emails := h.File_handler.read_all()

	for _, i := range emails {
		for _, email := range i {
			m := gomail.NewMessage()

			m.SetHeader("From", "currencyrateservice@gmail.com")
			m.SetHeader("To", email)
			m.SetHeader("Subject", "BTS to UAH rate")
			m.SetBody("text/plain", email_body)
			d := gomail.NewDialer("smtp.gmail.com", 587, "currencyrateservice@gmail.com", "rbbloqipornibpoj")
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

			if err := d.DialAndSend(m); err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
	}

}
func (h http_request_handler) request_currency_rate() currency_rate_struct {
	resp, err := http.Get("http://api.coinlayer.com/live?access_key=2496bb3fc39ac7fec2ce96bed66e2654&target=UAH&symbols=BTC")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	json_body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var res currency_rate_struct

	err = json.Unmarshal(json_body, &res)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
func (h http_request_handler) get_currency(writer http.ResponseWriter, request *http.Request) {

	currency_rate := h.request_currency_rate()

	load_template(writer, "templates/currency_rates_page.html", currency_rate)

}

func (handler http_request_handler) start_server() error {
	fmt.Println("Starting a server")

	http.HandleFunc("/subscribe", handler.get_subscribe)
	http.HandleFunc("/", handler.get_welcome_page)
	http.HandleFunc("/get_currency", handler.get_currency)
	http.HandleFunc("/send_to_all", handler.send_to_all)

	if err := http.ListenAndServe(":"+handler.Port, nil); err != nil {
		return fmt.Errorf("Could not start client API server on port %s: %w", handler.Port, err)
	}

	return nil
}
