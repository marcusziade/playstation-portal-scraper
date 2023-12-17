package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
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

	if unavailable {
		fmt.Println("The page has the 'Currently Unavailable' label.")
	} else {
		fmt.Println("The 'Currently Unavailable' label is not present on the page.")
	}
}
