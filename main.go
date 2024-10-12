package main

import (
	"fmt"

	"github.com/ahmadnouh97/blog-scraper/scraper"
	"github.com/ahmadnouh97/blog-scraper/utils"
)

func main() {
	url := "https://dev.to/search/feed_content?per_page=60&page=0&sort_by=published_at&sort_direction=desc"
	docs, err := devto.Scrape(url)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printDocs(docs)
}

func printDocs(docs []*devto.DevToData) {
	for _, doc := range docs {
		pretty_doc, err := utils.ToJSON(doc, true)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println(pretty_doc)
	}
}
