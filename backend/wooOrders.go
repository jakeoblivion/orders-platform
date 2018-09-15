package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type WooOrders []struct {
	Total    string `json:"total"`
	DatePaid string `json:"date_paid"`
	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"shipping"`
	LineItems []struct {
		Name      string `json:"name"`
		ProductID int    `json:"product_id"`
		Total     string `json:"total"`
	} `json:"line_items"`
}

type WooImageUrl struct {
	Product struct {
		PictureURL string `json:"featured_src"`
	} `json:"product"`
}

func fetchWooOrders() []Order {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://taxidermyart.co.uk//wp-json/wc/v2/orders", nil)
	req.SetBasicAuth(wooConsumerKey, wooSecretKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	response, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	var wooOrders WooOrders
	err3 := json.Unmarshal(response, &wooOrders)
	if err3 != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}

	return Adapter.OrdersAdapter(wooOrders)
}

func (wo WooOrders) OrdersAdapter() []Order {
	var orders []Order
	for _, wooOrder := range wo {
		var orderItems []OrderItem
		for _, wooOrderItem := range wooOrder.LineItems {
			orderItems = append(orderItems, OrderItem{
				ItemName:  wooOrderItem.Name,
				ImageUrl:  fetchWooItemImageUrl(wooOrderItem.ProductID),
				ItemPrice: wooOrderItem.Total,
			})
		}

		shipByDate, _ := time.Parse("2006-01-02T15:04:05", wooOrder.DatePaid)
		var order = Order{
			OrderItems: orderItems,
			OrderTotal: wooOrder.Total,
			OrderDate:  wooOrder.DatePaid,
			ShipByDate: shipByDate.AddDate(0, 0, 3).String(),
			ShippingAddress: ShippingAddress{
				Name:        fmt.Sprintf("%s %s", wooOrder.Shipping.FirstName, wooOrder.Shipping.LastName),
				AddressLine: wooOrder.Shipping.Address1,
				City:        wooOrder.Shipping.City,
				PostCode:    wooOrder.Shipping.Postcode,
				Country:     countryCodeConverter(wooOrder.Shipping.Country),
			},
			Platform: "woo",
		}
		orders = append(orders, order)
	}
	return orders
}

func fetchWooItemImageUrl(itemId int) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://taxidermyart.co.uk/wc-api/v3/products/%v?fields=featured_src", itemId), nil)
	req.SetBasicAuth(wooConsumerKey, wooSecretKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	response, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	var wooImageUrl WooImageUrl
	err3 := json.Unmarshal(response, &wooImageUrl)
	if err3 != nil {
		fmt.Println("There was an error UNMARSHALLING:", err)
	}
	imageUrl := wooImageUrl.Product.PictureURL
	index := len(imageUrl) - 4                                            //.jpg
	imageUrlThumbnail := imageUrl[:index] + "-300x300" + imageUrl[index:] //TODO: Find a better way to get the thumbnails
	fmt.Println(imageUrlThumbnail)
	return imageUrlThumbnail
}
