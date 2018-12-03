package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
	http.Handle("/", http.FileServer(http.Dir("/root/dist")))
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	var orders []Order

	start := time.Now()

	ordersChannel := make(chan []Order)

	go func() {
		ordersChannel <- fetchEbayOrders()
	}()
	go func() {
		ordersChannel <- fetchEtsyOrders()
	}()
	go func() {
		ordersChannel <- fetchWooOrders()
	}()

	for i := 0; i < 3; i++ {
		ordersFromChannel := <-ordersChannel

		orders = append(orders, ordersFromChannel...)
	}

	adaptedOrders, err := json.Marshal(orders)
	if err != nil {
		fmt.Println("There was an error MARSHALLING all the orders:", err)
	}

	fmt.Println("Time taken: ", time.Since(start).Seconds())
	_, err = w.Write(adaptedOrders)
	if err != nil {
		fmt.Println("There was an error returning orders:", err)
	}
}
