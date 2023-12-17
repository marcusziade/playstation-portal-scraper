package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type WebhookRequest struct {
	Value1 string `json:"value1"`
}

func main() {
	for {
		url := "https://direct.playstation.com/en-us/buy-accessories/playstation-portal-remote-player"

		logger := log.New(os.Stdout, "", log.LstdFlags)

		c := colly.NewCollector(
			colly.AllowedDomains("direct.playstation.com"),
		)

		var unavailable bool
		foundTexts := make(map[string]bool)

		c.OnRequest(func(r *colly.Request) {
			logger.Printf("Requesting URL: %s\n", r.URL.String())
		})

		c.OnResponse(func(r *colly.Response) {
			logger.Printf("Received response: %d from %s\n", r.StatusCode, r.Request.URL)
		})

		c.OnHTML("p.sony-text-body-1", func(e *colly.HTMLElement) {
			text := strings.TrimSpace(e.Text)
			if text != "" && !foundTexts[text] {
				logger.Printf("Found element with text: '%s'\n", text)
				foundTexts[text] = true
				if strings.Contains(text, "Currently Unavailable") {
					unavailable = true
				}
			}
		})

		c.OnError(func(r *colly.Response, err error) {
			logger.Printf("Error: %v when requesting URL: %s\n", err, r.Request.URL)
		})

		err := c.Visit(url)
		if err != nil {
			logger.Printf("Error visiting URL: %v\n", err)
			return
		}

		if !unavailable {
			message := "The 'Currently Unavailable' label is not present on the page. The product may be available!"
			fmt.Println(message)
			sendWebhookNotification("portal_available", message)
		} else {
			fmt.Println("The page has the 'Currently Unavailable' label. The product is not available.")
			prettyDate := time.Now().Format("Monday, January 2, 2006 3:04:05 PM MST")
			fmt.Printf("Last checked: %s\n", prettyDate)
		}

		time.Sleep(1 * time.Hour)
	}
}

func sendWebhookNotification(event string, message string) {
	keyInEnv := os.Getenv("IFTTT_WEBHOOK_KEY")
	if keyInEnv == "" {
		fmt.Println("IFTTT_WEBHOOK_KEY environment variable is not set. Skipping webhook notification.")
		return
	}

	webhookURL := "https://maker.ifttt.com/trigger/" + event + "/with/key/" + keyInEnv

	data := WebhookRequest{
		Value1: message,
	}
	requestBody, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Error sending webhook: %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Webhook notification sent.")
}
