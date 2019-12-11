## Orders platform API

This API fetches orders from etsy, ebay and woocommerce.
```
GET /get-orders
```

### Obtaining ETSY OAuth 1.0 token via Postman:
1. Create etsy app
2. Using the consumer key (KEYSTRING) and consumer secret (SHARED SECRET) in the Authorization headers, hit the request token endpoint via Postman: `https://openapi.etsy.com/v2/oauth/request_token?scope=transactions_r email_r listings_r` 
3. Decode the response using utf8_decode and past in the browser `https://www.urldecoder.org/`
4. You will be navigated to an etsy page, click "Allow"
5. Copy the verification code from the success page and copy the url (contains OAuth token and secret)
6. Hit the request access token endpoint: `https://openapi.etsy.com/v2/oauth/access_token?oauth_verifier=<verification code>` with the Authorisation from the OAuth token and secret 

### Obtaining EBAY token:
1. Log into `https://developer.ebay.com/my/auth/?env=production&index=0&auth_type=oauth` and obtain a user token.
2. Click `Get a Token from eBay via Your Application` and Test Sign-In
3. Grant Application Access by clicking "I agree"
4. Copy the `code` from the URL once access is granted. This code is url-encoded and will need to be decoded
5. POST to `https://api.ebay.com/identity/v1/oauth2/token` with:

 ```
Headers:
    Content-Type: application/x-www-form-urlencoded
    Authorization: Basic base64Encoded({clientId}:{secretId})

Body (application/x-www-form-urlencoded):
    grant_type: authorization_code
    redirect_uri: {runName}
    code: {the decoded `code` from above}
```


 
