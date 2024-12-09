package main

import (
	"fmt"
	"sync"

	"github.com/mohanrajnambe/web-crawler-go/models"
	"github.com/mohanrajnambe/web-crawler-go/utils"
)

type CrawlerService struct {
	Crawler *models.Crawler
}

func NewCrawlService() *CrawlerService {
	return &CrawlerService{
		Crawler: &models.Crawler{
			Visited: make(map[string]bool),
		},
	}
}

func (cs *CrawlerService) Crawl(url string, depth int, results chan<- models.PageMetadata) {
	if depth == 0 {
		return
	}

	//check if URL has a already been visited
	cs.Crawler.Mutex.Lock()
	if cs.Crawler.Visited[url] {
		cs.Crawler.Mutex.Unlock()
		return
	}

	cs.Crawler.Visited[url] = true
	cs.Crawler.Mutex.Unlock()

	//Fetch page and parse metadata
	page, links, err := utils.FetchPage(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	results <- *page

	var wg sync.WaitGroup
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			cs.Crawl(link, depth-1, results)
		}(link)
	}

	wg.Wait()

}
