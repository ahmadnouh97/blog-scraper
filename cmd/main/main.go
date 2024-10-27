package main

import (
	"fmt"
	"log"

	"github.com/ahmadnouh97/blog-scraper/internal"
	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func main() {
	// Initialize the database
	db, err := internal.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer db.Close()

	// Initialize blog repository
	blogRepo := blog.NewRepository(db)

	// Scrape Dev.to
	params := map[string]string{
		"per_page":       "60",
		"page":           "0",
		"sort_by":        "published_at",
		"sort_direction": "desc",
	}

	devToBlogs, err := scraper.FetchBlogs(params)

	if err != nil {
		log.Fatal("Failed to fetch Dev.to blogs: ", err)
	}

	// Save blogs to database
	for _, devToBlog := range devToBlogs {
		newBlog := &blog.Blog{
			ID:        devToBlog.ID,
			Title:     devToBlog.Title,
			Content:   devToBlog.Content,
			CreatedAt: devToBlog.CreatedAt,
		}

		if _, err := blogRepo.AddBlog(newBlog); err != nil {
			log.Fatal("Failed to save blog to database: ", err)
		}
	}

	// Load blogs from database
	blogs, err := blogRepo.GetBlogs()

	if err != nil {
		log.Fatal("Failed to load blogs from database: ", err)
	}

	// Save blogs to JSON file
	err = utils.SaveJSON(blogs, "blogs.json")

	if err != nil {
		log.Fatal("Failed to save blogs to JSON file: ", err)
	}

	fmt.Println("Blogs saved to JSON file: blogs.json")

	// Load blogs from JSON file
	blogs, err = utils.LoadJSON[[]*blog.Blog]("blogs.json")

	if err != nil {
		log.Fatal("Failed to load blogs from JSON file: ", err)
	}

	// Print the length of the loaded blogs
	fmt.Printf("Loaded %d blogs from JSON file\n", len(blogs))
}
