package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func createSignature(payload []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payload)

	return hex.EncodeToString(h.Sum(nil))
}

func sendWebhook() {
	payload, err := os.ReadFile("test_data/webhook.json")
	if err != nil {
		fmt.Println("Error reading mock webhook:", err)
		return
	}
	signature := createSignature(payload, webhookSecret)

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/webhook",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Webhook-Signature", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending webhook:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
}
