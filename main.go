package main

import (
	"fmt"

	"github.com/mohanrajnambe/web-crawler-go/models"
)

func main() {
	startURL := "https://www.wikipedia.org/"
	maxDepth := 3

	service := NewCrawlService()

	results := make(chan models.PageMetadata)
	go func() {
		service.Crawl(startURL, maxDepth, results)
		close(results)
	}()

	fmt.Println("Crawling started")
	for result := range results {
		fmt.Printf("\nURL: %s\nTitle: %s\nSnippet: %s\n", result.URL, result.Title, result.Snippet)
	}
	fmt.Println("Crawling finished")

}
