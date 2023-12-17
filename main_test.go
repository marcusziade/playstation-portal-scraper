package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var webhookURL string

func TestSendWebhookNotification(t *testing.T) {
	os.Setenv("IFTTT_WEBHOOK_KEY", "testkey")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/trigger/") {
			t.Errorf("Incorrect endpoint: got %v want %v", r.URL.Path, "/trigger/")
		}

		if r.Method != http.MethodPost {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}

		if r.Header.Get("IFTTT_WEBHOOK_KEY") != "testkey" {
			t.Errorf("Incorrect IFTTT key: got %v want %v", r.Header.Get("IFTTT_WEBHOOK_KEY"), "testkey")
		}
	}))
	defer server.Close()

	webhookURL = server.URL

	testCases := []struct {
		event   string
		message string
	}{
		{"test_event", "Test message"},
		{"", "Empty event"},
		{"test_event", ""},
	}

	for _, tc := range testCases {
		sendWebhookNotification(tc.event, tc.message)
	}

	os.Unsetenv("IFTTT_WEBHOOK_KEY")
}
