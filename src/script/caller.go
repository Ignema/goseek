package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Payload struct {
    Issue string
    Link string
    Query string
}

func main() {
	// Check if arguments number is valid
	if len(os.Args) != 5 {
		log.Fatal("Invalid number of arguments...")
	}

	// Create json string
	payload := Payload{os.Args[2], os.Args[3], os.Args[4]}

	// Create the POST request body
	body, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", os.Args[1], bytes.NewBuffer(body))
	if err != nil {
	log.Fatal(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print the response status and body
	log.Println(resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	log.Fatal(err)
	}
	log.Println(string(respBody))
}