package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"os"
)

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
	} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
			}
			w.Write(contents)
	}
}

func main() {
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3000", nil))}
