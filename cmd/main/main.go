package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ahmadnouh97/blog-scraper-qa/cmd/handlers"
	"github.com/ahmadnouh97/blog-scraper-qa/cmd/middlewares"
	"github.com/ahmadnouh97/blog-scraper-qa/internal"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/blog"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/llm"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper-qa/internal/utils"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func initScheduler(blogRepo *blog.Repository, logger *utils.CustomLogger) (gocron.Scheduler, error) {
	location, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		return nil, err
	}

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(location))

	if err != nil {
		return nil, err
	}

	logger.Info("Scheduler initialized successfully")

	job, err := scheduler.NewJob(
		gocron.DurationJob(
			25*time.Second,
		),
		gocron.NewTask(
			scraper.ScrapeBlogs,
			blogRepo,
			logger,
		),
	)

	if err != nil {
		return nil, err
	}

	logger.Info("Job %s initialized successfully", job.ID())

	return scheduler, nil
}

func runScheduler(scheduler gocron.Scheduler, logger *utils.CustomLogger) {
	scheduler.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(1 * time.Minute):
	}

	// when you're done, shut it down
	err := scheduler.Shutdown()

	if err != nil {
		logger.Error("Failed to shutdown scheduler: ", err)
		os.Exit(1)
	}

	logger.Info("Scheduler shutdown successfully")
}

func main() {
	ctx := context.Background()
	logger := utils.NewCustomLogger()
	err := godotenv.Load()

	if err != nil {
		logger.Error("Failed to load .env file !")
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")
	GEMINI_API_KEY := os.Getenv("GEMINI_API_KEY")

	filepath := "./db/blogs.db"
	db, err := internal.InitDB(filepath)

	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	chain, err := llm.InitSQLDatabaseChain(ctx, GEMINI_API_KEY, filepath, 100)

	defer db.Close()

	if err != nil {
		log.Fatal("Failed to initialize chain: ", err)
	}

	blogRepo := blog.NewRepository(db, logger)

	// scheduler, err := initScheduler(blogRepo, logger)

	if err != nil {
		logger.Error("Failed to initialize scheduler: ", err)
		return
	}

	// go runScheduler(scheduler, logger)

	mux := http.NewServeMux()

	statusHandler := http.HandlerFunc(handlers.CheckStatus(blogRepo, logger))
	getBlogsRoute := http.HandlerFunc(handlers.GetBlogs(blogRepo, logger))
	scrapeBlogsRoute := http.HandlerFunc(handlers.ScrapeBlogs(blogRepo, logger))
	countBlogsRoute := http.HandlerFunc(handlers.CountBlogs(blogRepo, logger))
	askQuestionRoute := http.HandlerFunc(handlers.AskQuestion(ctx, chain, logger))

	getBlogsHandler := middlewares.ApiKeyMiddleware(getBlogsRoute, SECRET_KEY)
	scrapeBlogsHandler := middlewares.ApiKeyMiddleware(scrapeBlogsRoute, SECRET_KEY)
	countBlogsHandler := middlewares.ApiKeyMiddleware(countBlogsRoute, SECRET_KEY)
	askQuestionHandler := middlewares.ApiKeyMiddleware(askQuestionRoute, SECRET_KEY)

	mux.Handle("GET /status", statusHandler)
	mux.Handle("GET /blogs", getBlogsHandler)
	mux.Handle("GET /scrape", scrapeBlogsHandler)
	mux.Handle("GET /count", countBlogsHandler)
	mux.Handle("GET /ask", askQuestionHandler)

	logger.Info("Server is running ✅, check http://localhost:8000/status for status")

	err = http.ListenAndServe(":8000", mux)

	if err != nil {
		logger.Error("Failed to start server: ", err)
		return
	}
}
