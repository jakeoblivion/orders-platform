package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var ebayToken = fetchEbayToken()

type EbayOrders struct {
	Total  int `json:"total"`
	Orders []struct {
		OrderID        string    `json:"orderId"`
		LegacyOrderID  string    `json:"legacyOrderId"`
		CreationDate   time.Time `json:"creationDate"`
		PricingSummary struct {
			Total struct {
				ConvertedFromValue string `json:"convertedFromValue"`
			} `json:"total"`
		} `json:"pricingSummary"`
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
				ConvertedFromValue string `json:"convertedFromValue"`
			} `json:"lineItemCost"`
			Quantity int `json:"quantity"`
			Total    struct {
				ConvertedFromValue string `json:"convertedFromValue"`
			} `json:"total"`
			LineItemFulfillmentInstructions struct {
				MaxEstimatedDeliveryDate time.Time `json:"maxEstimatedDeliveryDate"`
				ShipByDate               time.Time `json:"shipByDate"`
			} `json:"lineItemFulfillmentInstructions"`
		} `json:"lineItems"`
	} `json:"orders"`
}

type EbayListing struct {
	Image struct {
		ImageURL string `json:"imageUrl"`
	} `json:"image"`
}

//var mockedEbayOrders = Orders{Orders: []Order{
//	{OrderItems: []OrderItem{
//		{
//			ItemName:  "REAL King Swallowtail in Frame UK - Yellow, Butterfly, Entomology, Taxidermy",
//			ImageUrl:  "https://i.ebayimg.com/00/s/MTYwMFgxNjAw/z/w58AAOSw6CJbFG32/$_1.JPG",
//			ItemPrice: "99.99"},
//	},
//		OrderDate:  "2000-08-22T01:34:06Z",
//		OrderTotal: "99.99",
//		ShipByDate: "2000-08-24T01:34:06Z",
//		ShippingAddress: ShippingAddress{
//			Name:        "John Smith",
//			AddressLine: "3 Ellen Street",
//			City:        "Runcorn",
//			PostCode:    "SW17TQ",
//			Country:     "UK",
//		},
//		Platform: "ebay",
//	},
//	{OrderItems: []OrderItem{
//		{
//			ItemName:  "Atlas Moth in Frame (Attacus atlas)",
//			ImageUrl:  "https://www.taxidermyart.co.uk/wp-content/uploads/2015/08/Atlas-Moth-300x285.jpg",
//			ItemPrice: "89.99"},
//	},
//		OrderDate:  "2018-08-22T01:34:06Z",
//		OrderTotal: "89.99",
//		ShipByDate: "2018-08-27T01:34:06Z",
//		ShippingAddress: ShippingAddress{
//			Name:        "Jane Doe",
//			AddressLine: "8 Ellen Street",
//			City:        "Acherfield",
//			PostCode:    "CR4 1GB",
//			Country:     "UK",
//		},
//		Platform: "etsy",
//	}}}

func fetchEbayOrders() []byte {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	response := ApiCallGet("https://api.ebay.com/sell/fulfillment/v1/order?filter=orderfulfillmentstatus:%7BNOT_STARTED%7CIN_PROGRESS%7D", headers)

	var ebayOrders EbayOrders

	err := json.Unmarshal(response, &ebayOrders)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}

	return ebayOrdersAdapter(ebayOrders)
}

func ebayOrdersAdapter(ebayOrders EbayOrders) []byte {
	var orders Orders
	for _, ebayOrder := range ebayOrders.Orders {
		var orderItems []OrderItem
		for _, ebayOrderItem := range ebayOrder.OrderItems {
			orderItems = append(orderItems, OrderItem{
				ItemName:  ebayOrderItem.Title,
				ImageUrl:  fetchEbayItemImage(ebayOrderItem.LegacyItemID),
				ItemPrice: ebayOrderItem.Total.ConvertedFromValue,
			})
		}

		var order = Order{
			OrderItems: orderItems,
			OrderDate:  ebayOrder.CreationDate.String(),
			OrderTotal: ebayOrder.PricingSummary.Total.ConvertedFromValue,
			ShipByDate: ebayOrder.OrderItems[0].LineItemFulfillmentInstructions.ShipByDate.String(),
			ShippingAddress: ShippingAddress{
				Name:        ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.FullName,
				AddressLine: ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.ContactAddress.AddressLine1,
				City:        ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.ContactAddress.City,
				PostCode:    ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.ContactAddress.PostalCode,
				Country:     ebayCountryCodeConverter(ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.ContactAddress.CountryCode),
			},
			Platform: "ebay",
		}
		orders.Orders = append(orders.Orders, order)
	}

	adaptedOrders, err := json.Marshal(orders)
	if err != nil {
		fmt.Println("There was an error MARSHALLING:", err)
	}

	//fmt.Println("\n\n Adapted orders:", string(adaptedOrders))

	return adaptedOrders
}

func fetchEbayItemImage(itemId string) string {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	var ebayListing = ApiCallGet(fmt.Sprintf("https://api.ebay.com/buy/browse/v1/item/v1|%s|0", itemId), headers)

	var ebayImageUrl EbayListing

	err := json.Unmarshal(ebayListing, &ebayImageUrl)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}

	return ebayImageUrl.Image.ImageURL
}

func fetchEbayToken() string {
	refreshToken := ""
	body := []byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&scope=https://api.ebay.com/oauth/api_scope https://api.ebay.com/oauth/api_scope/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sell.marketing https://api.ebay.com/oauth/api_scope/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sell.inventory https://api.ebay.com/oauth/api_scope/sell.account.readonly https://api.ebay.com/oauth/api_scope/sell.account https://api.ebay.com/oauth/api_scope/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sell.fulfillment https://api.ebay.com/oauth/api_scope/sell.analytics.readonly", refreshToken))
	encodedOAuthCreds := ""

	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Authorization": fmt.Sprintf("Basic %s", encodedOAuthCreds)}
	return string(ApiCallPost("https://api.ebay.com/identity/v1/oauth2/token", body, headers))
}
