package main

import (
	"log"
	"net/http"

	"github.com/ahmadnouh97/blog-scraper/cmd/handlers"
	"github.com/ahmadnouh97/blog-scraper/internal"
	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func main() {
	logger := utils.NewCustomLogger()

	db, err := internal.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	defer db.Close()

	blogRepo := blog.NewRepository(db, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Home(blogRepo, logger))
	mux.HandleFunc("GET /blogs", handlers.GetBlogs(blogRepo, logger))

	logger.Info("Server is running on http://localhost:8000")

	err = http.ListenAndServe(":8000", mux)

	if err != nil {
		logger.Error("Failed to start server: ", err)
		return
	}
}
