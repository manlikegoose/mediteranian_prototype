#Meridian pivot - syncing services

##Overview
In developing this mini-prototype i have used Golang as my programing language for learning and also demonstrating a syncing service using webhook signature verification.

This prototype demonstrates a warehouse that sends inventory updates to a Go HTTP server.

##Structure

warehouse
    |
    |POST/webhook
    |Json Payload
    |HMAC-SHA256 signature
    |
Go Webhook server
    |
    |
    signature Verification    
    /               \    
    valid            invalid
    |                  |
    200                401


###features
-HTTP webhook receiver
-JSON request handling
- HMAC-SHA256 signature generation
- Webhook signature verification
- Mock inventory webhook data
- Valid and invalid signature tests

 ## Running the Prototype

Start the server:

```bash
go run .


The server runs on http://localhost:8080