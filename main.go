package main

import (
	"fmt"

	devto "github.com/ahmadnouh97/blog-scraper/scraper"
	"github.com/ahmadnouh97/blog-scraper/utils"
)

func main() {
	params := map[string]string{
		"per_page":       "60",
		"page":           "0",
		"sort_by":        "published_at",
		"sort_direction": "desc",
	}

	docs, err := devto.FetchBlogs(params)

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
