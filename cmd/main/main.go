package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ahmadnouh97/blog-scraper/cmd/handlers"
	"github.com/ahmadnouh97/blog-scraper/cmd/middlewares"
	"github.com/ahmadnouh97/blog-scraper/internal"
	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/scraper"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
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

	getBlogsHandler := middlewares.ApiKeyMiddleware(getBlogsRoute, SECRET_KEY)
	scrapeBlogsHandler := middlewares.ApiKeyMiddleware(scrapeBlogsRoute, SECRET_KEY)

	mux.Handle("GET /status", statusHandler)
	mux.Handle("GET /blogs", getBlogsHandler)
	mux.Handle("GET /scrape", scrapeBlogsHandler)

	logger.Info("Server is running ✅, check http://localhost:8000/status for status")

	err = http.ListenAndServe(":8000", mux)

	if err != nil {
		logger.Error("Failed to start server: ", err)
		return
	}
}
