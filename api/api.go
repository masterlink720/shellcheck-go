package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"io/ioutil"
)

const (
	Url    = "http://www.shellcheck.net/shellcheck.php"                            // Base URL used to access shellcheck's web site
	Method = "POST"                                                                // The http method used when sending requests
	Agent  = "shellcheckgo/0.0.1 (https://github.com/masterlink720/shellcheck-go)" // Might as well be honest for their logs
)

func Check(script string) (string) {

	// Setup a "form" to send
	/*
	form := url.Values{
		"script":	script,
	}
	*/
	form := url.Values{}
	 form.Add("script", script)

	// URL encode the form
	body := form.Encode()

	fmt.Println("Submitting the script...")
	req, err := http.NewRequest(Method, Url, strings.NewReader(body))

	// Fail?
	onError(err)

	// Setup some request headers
	req.Header.Set("Content-Length", string(len(body)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json, text/javascript, */*")

	// Pass the request to a client handler
	client := &http.Client{}
	resp,err := client.Do(req)

	// Fail?
	onError(err)

	// Close the request when this thread completes
	defer resp.Body.Close()

	// Read the response
	buffer, err := ioutil.ReadAll(resp.Body)

	// Fail?
	onError(err)

	// Parse the response
	body = string(buffer[:])

	return body

	//fmt.Printf("Response!\n\tStatus: %v\n\tHeaders: %v\n\tResponse:...\n%s", resp.Status, resp.Header, body)
}

func onError(err error) {
	if err != nil {
		log.Fatal("Fatal API error", err)
		os.Exit(1)
	}
}