### Obtaining ETSY OAuth 1.0 token via Postman:
1. Create etsy app
2. Using the consumer key (KEYSTRING) and consumer secret (SHARED SECRET) in the Authorization headers, hit the request token endpoint via Postman: `https://openapi.etsy.com/v2/oauth/request_token?scope=transactions_r email_r listings_r` 
3. Decode the response using utf8_decode and past in the browser `https://www.urldecoder.org/`
4. You will be navigated to an etsy page, click "Allow"
5. Copy the verification code from the success page and copy the url (contains OAuth token and secret)
6. Hit the request access token endpoint: `https://openapi.etsy.com/v2/oauth/access_token?oauth_verifier=<verification code>` with the Authorisation from the OAuth token and secret 

