package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

func ApiCallGet(url string, authorization string ) []byte {
	response, err := http.Get(url)

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
