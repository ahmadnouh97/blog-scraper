package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tmc/langchaingo/chains"

	"github.com/ahmadnouh97/blog-scraper-qa/internal/blog"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/llm"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/utils"
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
		// Read query parameters
		defaults := map[string]int{
			"page":      1,
			"page_size": 10,
		}
		queryParams := r.URL.Query()

		getOrDefault := func(key string) int {
			if value := queryParams.Get(key); value != "" {
				number, err := strconv.Atoi(value)
				if err != nil {
					return defaults[key]
				}
				return number
			}
			return defaults[key]
		}

		page := getOrDefault("page")
		pageSize := getOrDefault("page_size")

		// Get blogs
		blogs, totalCount, err := repo.GetBlogs(page, pageSize)

		if err != nil {
			logger.Error("Failed to get blogs: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Blogs fetched from database successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		response := &blog.BlogsPaginationResponse{
			Blogs:      blogs,
			Page:       page,
			PageSize:   pageSize,
			TotalItems: totalCount,
			TotalPages: (totalCount + pageSize - 1) / pageSize,
			HasMore:    totalCount > page*pageSize,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
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

func CountBlogs(repo *blog.Repository, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := repo.CountBlogs()

		if err != nil {
			logger.Error("Failed to count blogs: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(count); err != nil {
			logger.Error("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func AskQuestion(ctx context.Context, chain *chains.SQLDatabaseChain, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		question := queryParams.Get("question")

		if question == "" {
			http.Error(w, "Please provide a question", http.StatusBadRequest)
			return
		}

		answer, err := llm.Run(ctx, chain, question)

		if err != nil {
			logger.Error("Failed to answer the question: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		llmAnswer := &llm.LLMResponse{
			Answer: answer,
		}

		if err := json.NewEncoder(w).Encode(llmAnswer); err != nil {
			logger.Error("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
