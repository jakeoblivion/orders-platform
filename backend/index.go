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
	ebayGetOrdersUrl := "https://api.ebay.com/sell/fulfillment/v1/order?filter=orderfulfillmentstatus:%7BNOT_STARTED%7CIN_PROGRESS%7D"
	ebayToken := ""
	return ApiCallGet(ebayGetOrdersUrl, ebayToken)
}
