package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ebayToken = fetchEbayToken()

type EbayOrders struct {
	Total  int `json:"total"`
	Orders []struct {
		OrderID                      string    `json:"orderId"`
		LegacyOrderID                string    `json:"legacyOrderId"`
		CreationDate                 time.Time `json:"creationDate"`
		FulfillmentStartInstructions []struct {
			ShippingStep struct {
				ShipTo struct {
					FullName       string `json:"fullName"`
					ContactAddress struct {
						AddressLine1 string `json:"addressLine1"`
						AddressLine2 string `json:"addressLine2"`
						City         string `json:"city"`
						PostalCode   string `json:"postalCode"`
						CountryCode  string `json:"countryCode"`
					} `json:"contactAddress"`
					PrimaryPhone struct {
						PhoneNumber string `json:"phoneNumber"`
					} `json:"primaryPhone"`
					Email string `json:"email"`
				} `json:"shipTo"`
			} `json:"shippingStep"`
		} `json:"fulfillmentStartInstructions"`
		OrderItems []struct {
			OrderItemID  string `json:"lineItemId"`
			LegacyItemID string `json:"legacyItemId"`
			Title        string `json:"title"`
			LineItemCost struct {
				Value string `json:"value"`
			} `json:"lineItemCost"`
			Quantity int `json:"quantity"`
			Total    struct {
				Value string `json:"value"`
			} `json:"total"`
			LineItemFulfillmentInstructions struct {
				MaxEstimatedDeliveryDate time.Time `json:"maxEstimatedDeliveryDate"`
				ShipByDate               time.Time `json:"shipByDate"`
			} `json:"lineItemFulfillmentInstructions"`
		} `json:"lineItems"`
	} `json:"orders"`
}

type Orders struct {
	Orders []Order
}

type Order struct {
	OrderItems      []OrderItem
	OrderDate       string
	OrderTotal      string
	ShipByDate      string
	ShippingAddress ShippingAddress
	Platform        string
}

type OrderItem struct {
	ItemName  string
	ImageUrl  string
	ItemPrice string
}

type ShippingAddress struct {
	Name        string
	AddressLine string
	City        string
	PostCode    string
	Country     string
}

var mockedEbayOrders = Orders{Orders: []Order{
	{OrderItems: []OrderItem{
		{
			ItemName:  "Butterfly Frame",
			ImageUrl:  "https://i.ebayimg.com/00/s/MTYwMFgxNjAw/z/w58AAOSw6CJbFG32/$_1.JPG",
			ItemPrice: "99.99"},
	},
		OrderDate:  "2018-08-22T01:34:06Z",
		OrderTotal: "99.99",
		ShipByDate: "2018-08-24T01:34:06Z",
		ShippingAddress: ShippingAddress{
			Name:        "John Smith",
			AddressLine: "3 Ellen Street",
			City:        "Runcorn",
			PostCode:    "SW17TQ",
			Country:     "UK",
		},
		Platform: "ebay",
	},
	{OrderItems: []OrderItem{
		{
			ItemName:  "Moth Frame",
			ImageUrl:  "https://www.taxidermyart.co.uk/wp-content/uploads/2015/08/Atlas-Moth-300x285.jpg",
			ItemPrice: "89.99"},
	},
		OrderDate:  "2018-08-22T01:34:06Z",
		OrderTotal: "89.99",
		ShipByDate: "2018-08-24T01:34:06Z",
		ShippingAddress: ShippingAddress{
			Name:        "Jane Doe",
			AddressLine: "8 Ellen Street",
			City:        "Acherfield",
			PostCode:    "CR4 1GB",
			Country:     "UK",
		},
		Platform: "etsy",
	}},
}

func main() {
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("content-type", "application/json")
	//ebayOrders := fetchEbayOrders()
	//ebayItemImage := fetchEbayItemImage("273258129797")xa

	fmt.Println("MOCKED ORDERS: ", mockedEbayOrders)
	orders, _ := json.Marshal(mockedEbayOrders)
	w.Write(orders)
}

func fetchEbayOrders() []byte {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	response := ApiCallGet("https://api.ebay.com/sell/fulfillment/v1/order?filter=orderfulfillmentstatus:%7BNOT_STARTED%7CIN_PROGRESS%7D", headers)

	var ebayOrders EbayOrders

	err := json.Unmarshal(response, &ebayOrders)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}
	fmt.Printf("\n\n Ebay Orders Struct: %+v", ebayOrders)

	ebayOrdersByte, err2 := json.Marshal(ebayOrders)
	if err2 != nil {
		fmt.Println("There was an error MARSHALLING:", err)
	}
	fmt.Println("\n\n Ebay Orders stringfied:", string(ebayOrdersByte))

	return ebayOrdersByte
}

func fetchEbayItemImage(itemId string) []byte {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	return ApiCallGet(fmt.Sprintf("https://api.ebay.com/buy/browse/v1/item/v1|%s|0", itemId), headers)
}

func fetchEbayToken() string {
	refreshToken := ""
	body := []byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&scope=https://api.ebay.com/oauth/api_scope https://api.ebay.com/oauth/api_scope/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sell.marketing https://api.ebay.com/oauth/api_scope/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sell.inventory https://api.ebay.com/oauth/api_scope/sell.account.readonly https://api.ebay.com/oauth/api_scope/sell.account https://api.ebay.com/oauth/api_scope/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sell.fulfillment https://api.ebay.com/oauth/api_scope/sell.analytics.readonly", refreshToken))
	encodedOAuthCreds := ""

	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Authorization": fmt.Sprintf("Basic %s", encodedOAuthCreds)}
	return string(ApiCallPost("https://api.ebay.com/identity/v1/oauth2/token", body, headers))
}
