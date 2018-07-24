package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	orders := fetchEbayOrders()
	w.Write(orders)
}

func fetchEbayOrders() []byte {
	return ApiCallGet("https://jsonplaceholder.typicode.com/posts/1", "")
}
