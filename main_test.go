package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebhookValidSignature(t *testing.T) {
	payload := []byte(`{"product_id":"PHONE-001","quantity":19}`)
	signature := generateSignature(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/webhook",
		bytes.NewReader(payload),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Webhook-Signature", signature)

	recorder := httptest.NewRecorder()

	webhookHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
}
