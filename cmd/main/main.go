package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmadnouh97/blog-scraper/cmd/handlers"
	"github.com/ahmadnouh97/blog-scraper/cmd/middlewares"
	"github.com/ahmadnouh97/blog-scraper/internal"
	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	logger := utils.NewCustomLogger()
	err := godotenv.Load()

	if err != nil {
		logger.Error("Failed to load .env file !")
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	db, err := internal.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	defer db.Close()

	blogRepo := blog.NewRepository(db, logger)

	mux := http.NewServeMux()

	homeHandler := http.HandlerFunc(handlers.Home(blogRepo, logger))
	getBlogsRoute := http.HandlerFunc(handlers.GetBlogs(blogRepo, logger))

	getBlogsHandler := middlewares.ApiKeyMiddleware(getBlogsRoute, SECRET_KEY)

	mux.Handle("GET /", homeHandler)
	mux.Handle("GET /blogs", getBlogsHandler)

	logger.Info("Server is running on http://localhost:8000")

	err = http.ListenAndServe(":8000", mux)

	if err != nil {
		logger.Error("Failed to start server: ", err)
		return
	}
}
