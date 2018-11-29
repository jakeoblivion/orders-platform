package main

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

var etsyToken = oauth.AccessToken{etsyAccessToken, etsyTokenSecret, nil}
var consumer = oauth.NewConsumer(etsyConsumerKey, etsyConsumerSecret, oauth.ServiceProvider{})
var latestEtsyOrders = LatestEtsyItemOrders{}

type LatestEtsyItemOrders struct {
	Results []struct {
		TransactionID  int    `json:"transaction_id"`
		Title          string `json:"title"`
		Price          string `json:"price"`
		ImageListingID int    `json:"image_listing_id"`
		ListingID      int    `json:"listing_id"`
		ReceiptID      int    `json:"receipt_id"`
	} `json:"results"`
}

type EtsyOrders struct {
	Results []struct {
		ReceiptID   int    `json:"receipt_id"`
		CreationTsz int    `json:"creation_tsz"`
		Name        string `json:"name"`
		FirstLine   string `json:"first_line"`
		City        string `json:"city"`
		Zip         string `json:"zip"`
		CountryID   int    `json:"country_id"`
		Grandtotal  string `json:"grandtotal"`
		ShippedDate int    `json:"shipped_date"`
	} `json:"results"`
}

type EtsyImage struct {
	Results []struct {
		Thumbnail string `json:"url_570xN"`
	}
}

func fetchEtsyOrders() []Order {
	etsyOrders := EtsyOrders{}

	etsyOrdesChannel := make(chan EtsyOrders)
	latestEtsyOrdersChannel := make(chan LatestEtsyItemOrders)

	go func() {
		latestEtsyOrdersChannel <- getLatestEtsyOrdersItem()
	}()

	go func() {
		etsyOrdesChannel <- getEtsyOrders()
	}()

	latestEtsyOrders = <-latestEtsyOrdersChannel
	etsyOrders = <-etsyOrdesChannel

	return Adapter.OrdersAdapter(etsyOrders)
}

func etsyApiCall(url string) []byte {
	response, err := consumer.Get(url, nil, &etsyToken)

	if response.StatusCode == 401 {
		fmt.Printf("Failed to parse Etsy token: %s", response)
	}

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return body
}

func (eo EtsyOrders) OrdersAdapter() []Order {
	var orders []Order

	numberOfOrders := len(eo.Results)
	var wg sync.WaitGroup
	wg.Add(numberOfOrders)
	for index := 0; index < numberOfOrders; index++ {
		etsyOrder := eo.Results[index]

		go func(index int) {
			defer wg.Done()

			var orderItems []OrderItem
			for _, etsyOrderItem := range latestEtsyOrders.Results {

				if etsyOrderItem.ReceiptID == etsyOrder.ReceiptID {
					orderItems = append(orderItems, OrderItem{
						ItemName:  etsyOrderItem.Title,
						ImageUrl:  getImageUrl(etsyOrderItem.ListingID),
						ItemPrice: etsyOrderItem.Price,
					})
				}
			}

			var order = Order{
				OrderItems: orderItems,
				OrderDate:  time.Unix(int64(etsyOrder.CreationTsz), 0).Format(time.RFC3339),
				OrderTotal: etsyOrder.Grandtotal,
				ShipByDate: time.Unix(int64(etsyOrder.ShippedDate), 0).Format(time.RFC3339),
				ShippingAddress: ShippingAddress{
					Name:        etsyOrder.Name,
					AddressLine: etsyOrder.FirstLine,
					City:        etsyOrder.City,
					PostCode:    etsyOrder.Zip,
					Country:     countryCodeConverter(strconv.Itoa(etsyOrder.CountryID)),
				},
				Platform: "etsy",
			}
			orders = append(orders, order)

		}(index)
	}
	wg.Wait()
	fmt.Println("Finished fetching Etsy Orders")

	return orders

}

func getEtsyOrders() EtsyOrders {
	response := etsyApiCall("https://openapi.etsy.com/v2/shops/TaxidermyArt/receipts/open")

	var etsyOrders EtsyOrders

	err := json.Unmarshal(response, &etsyOrders)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING etsyOrders:", err)
	}
	return etsyOrders
}

func getLatestEtsyOrdersItem() LatestEtsyItemOrders {
	fmt.Println("Called getLatestEtsyOrdersItem")
	response := etsyApiCall("https://openapi.etsy.com/v2/shops/TaxidermyArt/transactions/?limit=30")

	var latestEtsyItemOrders LatestEtsyItemOrders

	err := json.Unmarshal(response, &latestEtsyItemOrders)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING latestEtsyItemOrders:", err)
	}

	return latestEtsyItemOrders
}

func getImageUrl(listingId int) string {
	response := etsyApiCall("https://openapi.etsy.com/v2/listings/" + strconv.Itoa(listingId) + "/images")
	var etsyImage EtsyImage

	err := json.Unmarshal(response, &etsyImage)
	if err != nil {
		fmt.Println("There was an error UNMARSHALLING etsyImage:", err)
	}

	return etsyImage.Results[0].Thumbnail
}
