package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
)

const webhookSecret = "zquak_secret_key_matatuyangu"


func generateSignature(payload []byte) []byte {
	h := hmac.New(sha256.New, []byte(webhookSecret))
	h.Write(payload)

	return h.Sum(nil)
}


func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Received webhook:")
	fmt.Println(string(body))


	signature := generateSignature(body)
	_ = signature

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))

}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
