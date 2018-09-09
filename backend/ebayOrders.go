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

type EbayImageUrl struct {
	Item struct {
		PictureURL []string `json:"pictureURL"`
	} `json:"item"`
}

func fetchEbayOrders() []Order {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	response := ApiCallGet("https://api.ebay.com/sell/fulfillment/v1/order?limit=2", headers)
	//response := ApiCallGet("https://api.ebay.com/sell/fulfillment/v1/order?filter=orderfulfillmentstatus:%7BNOT_STARTED%7CIN_PROGRESS%7D", headers)

	var ebayOrders EbayOrders

	err := json.Unmarshal(response, &ebayOrders)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}

	return Adapter.OrdersAdapter(ebayOrders)
}

func (eo EbayOrders) OrdersAdapter() []Order {
	var orders []Order
	for _, ebayOrder := range eo.Orders {
		var orderItems []OrderItem
		for _, ebayOrderItem := range ebayOrder.OrderItems {
			orderItems = append(orderItems, OrderItem{
				ItemName:  ebayOrderItem.Title,
				ImageUrl:  fetchEbayItemImageUrl(ebayOrderItem.LegacyItemID),
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
				Country:     countryCodeConverter(ebayOrder.FulfillmentStartInstructions[0].ShippingStep.ShipTo.ContactAddress.CountryCode),
			},
			Platform: "ebay",
		}
		orders = append(orders, order)
	}
	return orders
}

func fetchEbayItemImageUrl(itemId string) string {
	var ebayListing = ApiCallGet(fmt.Sprintf("http://open.api.ebay.com/shopping?callname=GetSingleItem&responseencoding=JSON&appid=JacobNew-Taxiderm-PRD-0e0516b4c-e789e949&siteid=0&version=939&ItemID=%s", itemId), map[string]string{})
	var ebayImageUrl EbayImageUrl

	err := json.Unmarshal(ebayListing, &ebayImageUrl)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}

	return ebayImageUrl.Item.PictureURL[0]
}

func fetchEbayToken() string {
	body := []byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&scope=https://api.ebay.com/oauth/api_scope https://api.ebay.com/oauth/api_scope/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sell.marketing https://api.ebay.com/oauth/api_scope/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sell.inventory https://api.ebay.com/oauth/api_scope/sell.account.readonly https://api.ebay.com/oauth/api_scope/sell.account https://api.ebay.com/oauth/api_scope/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sell.fulfillment https://api.ebay.com/oauth/api_scope/sell.analytics.readonly", refreshToken))
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Authorization": fmt.Sprintf("Basic %s", encodedOAuthCreds)}

	return string(ApiCallPost("https://api.ebay.com/identity/v1/oauth2/token", body, headers))
}
