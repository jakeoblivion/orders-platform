package main

import "os"

var ebayRefreshToken string
var encodedOAuthCreds string

var etsyAccessToken string
var etsyTokenSecret string

var etsyConsumerKey string
var etsyConsumerSecret string

var wooConsumerKey string
var wooSecretKey string

func setupCredentials() {
	ebayRefreshToken = os.Getenv("EBAY_REFRESH_TOKEN")
	encodedOAuthCreds = os.Getenv("EBAY_ENCODED_AUTH_CREDS")
	etsyAccessToken = os.Getenv("ETSY_ACCESS_TOKEN")
	etsyTokenSecret = os.Getenv("ETSY_TOKEN_SECRET")
	etsyConsumerKey = os.Getenv("ETSY_CONSUMER_KEY")
	etsyConsumerSecret = os.Getenv("ETSY_CONSUMER_SECRET")
	wooConsumerKey = os.Getenv("WOO_CONSUMER_KEY")
	wooSecretKey = os.Getenv("WOO_SECRET_KEY")
}
