package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func Home(repo *blog.Repository, logger *utils.CustomLogger) func(http.ResponseWriter, *http.Request) {
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

		logger.Info("Blogs fetched successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(blogs); err != nil {
			logger.Error("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		logger.Info("Response encoded successfully")

	}
}
