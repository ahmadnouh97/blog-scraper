package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func CheckStatus(repo *blog.Repository, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Checking database connection..")
		w.WriteHeader(http.StatusOK)

		if err := repo.CheckConnection(); err != nil {
			logger.Error("Failed to connect to database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Database connection checked successfully")

		w.Write([]byte("Service is up and running :D"))
	}
}

func GetBlogs(repo *blog.Repository, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Fetching blogs..")
		blogs, err := repo.GetBlogs()

		if err != nil {
			logger.Error("Failed to get blogs: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Blogs fetched from database successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(blogs); err != nil {
			logger.Error("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		logger.Info("Blogs fetched successfully")
	}
}

func ScrapeBlogs(repo *blog.Repository, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		// Default values
		defaults := map[string]string{
			"per_page":       "60",
			"page":           "0",
			"sort_by":        "published_at",
			"sort_direction": "desc",
		}

		// Helper function to get a query parameter or default value
		getOrDefault := func(key string) string {
			if value := queryParams.Get(key); value != "" {
				return value
			}
			return defaults[key]
		}

		// Extract query parameters with defaults
		perPage := getOrDefault("per_page")
		page := getOrDefault("page")
		sortBy := getOrDefault("sort_by")
		sortDirection := getOrDefault("sort_direction")

		logger.Info("Scraping...")
		devToBlogs, err := scraper.ScrapeDevToBlogs(perPage, page, sortBy, sortDirection)

		if err != nil {
			logger.Error("Failed to scrape Dev.to: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Blogs scraped successfully")

		// Save blogs to database
		logger.Info("Saving blogs to database..")
		scraper.SaveDevToBlogs(devToBlogs, repo, logger)
		logger.Info("Blogs saved to database successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(devToBlogs); err != nil {
			logger.Error("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		logger.Info("Blogs scrapped and saved successfully")
	}
}
