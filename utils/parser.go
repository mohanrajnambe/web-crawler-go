package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mohanrajnambe/web-crawler-go/models"
	"golang.org/x/net/html"
)

func FetchPage(url string) (*models.PageMetadata, []string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to fetch %s: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("non-200 response for %s: %d", url, resp.StatusCode)
	}

	links, title, snippet := parseHTML(resp.Body)
	return &models.PageMetadata{
		Title:   title,
		URL:     url,
		Snippet: snippet,
	}, links, err
}

func parseHTML(body io.Reader) ([]string, string, string) {
	var links []string
	var title, snippet string
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return links, title, snippet
			}
			return nil, "", ""
		case html.StartTagToken:
			token := z.Token()
			switch token.Data {
			case "a":
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						href := attr.Val
						if strings.HasPrefix(href, "http") {
							links = append(links, href)
						}
					}
				}
			case "title":
				z.Next()
				title = string(z.Text())
			}
		case html.TextToken:
			if len(snippet) < 150 {
				snippet += string(z.Text())
			}
		}
	}
}
