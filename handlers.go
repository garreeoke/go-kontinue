package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Status codes
const (
	timeFormat = "01-02-2006"
	success    = 200
	failure    = 400
)

// Healthz return ok without auth for testing connectivity
func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Health request received")
	w.WriteHeader(200)
	data := []byte("OK")
	w.Write(data)

}

// PostExample create a new record nd store it in elastic
func PostExample(w http.ResponseWriter, r *http.Request) {

	log.Println("PostExample received")
	payload := Payload{
		Code:       success,
		W:          w,
		Error:      nil,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		payload.processError(err, "reading body")
	}
	payload.Data = body
	payload.completeRequest()
}

// PutExample create a new record nd store it in elastic
func PutExample(w http.ResponseWriter, r *http.Request) {

	log.Println("PutExample received")
	payload := Payload{
		Code:       success,
		W:          w,
		Error:      nil,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		payload.processError(err, "reading body")
	}
	payload.Data = body
	payload.completeRequest()
}

// DeleteExample create a new record nd store it in elastic
func DeleteExample(w http.ResponseWriter, r *http.Request) {

	log.Println("DeleteExample received")
	data := "{\"deleted\": \"yes\"}"
	payload := Payload{
		Code:       success,
		W:          w,
		Error:      nil,
		Data: []byte(data),
	}
	payload.completeRequest()
}

func (p *Payload) completeRequest() {
	// Check for error message and end accordingly
	if p.Error != nil {
		p.Code = failure
		log.Println("Error: ", p.Error)
	}
	log.Println("Request completed")
	p.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p.W.WriteHeader(p.Code)
	p.W.Write(p.Data)
}

func (p *Payload) processError(err error, msg string) {
	p.Error = errors.New(msg + ":" + err.Error())
	p.completeRequest()
}
