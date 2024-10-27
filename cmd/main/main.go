package main

import (
	"fmt"

	"github.com/ahmadnouh97/blog-scraper/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func main() {
	params := map[string]string{
		"per_page":       "60",
		"page":           "0",
		"sort_by":        "published_at",
		"sort_direction": "desc",
	}

	docs, err := scraper.FetchBlogs(params)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	utils.SaveJSON(docs, "devto.json")

	// devtoData, err := utils.LoadJSON[[]devto.DevToData]("devto.json")

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }

	// fmt.Println(len(devtoData))
}
