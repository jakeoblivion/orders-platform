package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Adapter interface {
	OrdersAdapter() []Order
}

type Order struct {
	OrderItems      []OrderItem     `json:"orderItems"`
	OrderDate       string          `json:"orderDate"`
	OrderTotal      string          `json:"orderTotal"`
	ShipByDate      string          `json:"shipByDate"`
	ShippingAddress ShippingAddress `json:"shippingAddress"`
	Platform        string          `json:"platform"`
}

type OrderItem struct {
	ItemName  string `json:"itemName"`
	ImageUrl  string `json:"imageUrl"`
	ItemPrice string `json:"itemPrice"`
}

type ShippingAddress struct {
	Name        string `json:"name"`
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	PostCode    string `json:"postCode"`
	Country     string `json:"country"`
}

func main() {
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("content-type", "application/json")
	var orders []Order

	orders = append(orders, fetchWooOrders()...)
	orders = append(orders, fetchEbayOrders()...)

	adaptedOrders, err := json.Marshal(orders)
	if err != nil {
		fmt.Println("There was an error MARSHALLING:", err)
	}
	w.Write(adaptedOrders)
}
