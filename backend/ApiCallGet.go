package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

func ApiCallGet(url string, token string ) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := client.Do(request)

	if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
	} else {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
			}
			return body
	}
	return nil
}
