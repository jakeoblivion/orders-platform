package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ApiCall(method string, url string, body []byte, headers map[string]string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	//setup headers
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	//make API request
	response, err := client.Do(request)

	if response.StatusCode == 401 {
		fmt.Printf("Failed to parse token: %s", response)
	}

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return body
}

func ApiCallGet(url string, headers map[string]string) []byte {
	return ApiCall("GET", url, nil, headers)
}

func ApiCallPost(url string, body []byte, headers map[string]string) []byte {
	return ApiCall("POST", url, body, headers)
}
