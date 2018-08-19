package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get-orders", GetOrders)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func GetOrders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	//orders := fetchEbayOrders(fetchEbayToken())
	orders := fetchEbayItemImage("273258129797", fetchEbayToken())
	w.Write(orders)
}

func fetchEbayOrders(ebayToken string) []byte {
	headers :=  map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	return ApiCallGet("https://api.ebay.com/sell/fulfillment/v1/order?filter=orderfulfillmentstatus:%7BNOT_STARTED%7CIN_PROGRESS%7D", headers)
}

func fetchEbayItemImage(itemId string, ebayToken string) []byte {
	headers :=  map[string]string{"Authorization": fmt.Sprintf("Bearer %s", ebayToken)}
	return ApiCallGet(fmt.Sprintf("https://api.ebay.com/buy/browse/v1/item/v1|%s|0", itemId), headers)
}

func fetchEbayToken() string {
	refreshToken := ""
	body := []byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&scope=https://api.ebay.com/oauth/api_scope https://api.ebay.com/oauth/api_scope/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sell.marketing https://api.ebay.com/oauth/api_scope/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sell.inventory https://api.ebay.com/oauth/api_scope/sell.account.readonly https://api.ebay.com/oauth/api_scope/sell.account https://api.ebay.com/oauth/api_scope/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sell.fulfillment https://api.ebay.com/oauth/api_scope/sell.analytics.readonly", refreshToken))
	encodedOAuthCreds := ""

	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Authorization": fmt.Sprintf("Basic %s", encodedOAuthCreds)}
	return string(ApiCallPost("https://api.ebay.com/identity/v1/oauth2/token", body, headers))
}
