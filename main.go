package main

import (
	"os"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Check if arguments number is valid
	if len(os.Args) != 4 {
		log.Fatal("Invalid number of arguments...")
	}

	// Create the POST request body
	body := []byte(os.Args[3])

	// Create a new HTTP request
	req, err := http.NewRequest(os.Args[1], os.Args[2], bytes.NewBuffer(body))
	if err != nil {
	log.Fatal(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "text/plain")

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