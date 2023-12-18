package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type AeroflotApi struct {
	client http.Client
}

func (a *AeroflotApi) Get(url string) (body []byte, err error) {
	resp, err := a.client.Get(url)
	if err != nil {
		log.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	// log.Printf("Body : %s", body)
	return
}

func (a *AeroflotApi) Post(url string, req_type string, data []byte) (body []byte, err error) {
	myData := bytes.NewBuffer(data)
	resp, err := a.client.Post(url, req_type, myData)
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	// log.Printf("Body : %s", body)
	return
}

func (a *AeroflotApi) PostJson(url string, data []byte) (body []byte, err error) {
	body, err = a.Post(url, "application/json", data)
	return
}

func New() (api AeroflotApi) {
	api = AeroflotApi{
		client: http.Client{
			Timeout:   time.Duration(30) * time.Second,
			Transport: NewTransport(nil),
		},
	}
	// Cookies
	api.Get("https://www.aeroflot.ru/sb/exchg/app/ru-ru")
	return
}
