# Blog Scraper

A simple Golang tool for scraping and storing blog posts. This tool extracts titles, authors, publication dates, and content, saving them to a local SQLite3 database.

## Features

- Scrapes blog content with ease
- Stores blog data in SQLite3 for quick access

## Tech Stack

- **Golang**: Core programming language
- **goose**: For DB Migrations
- **SQLite3**: For saving scraped data to sqlite database.
- **gocron**: For Scheduling Scraping Jobs.
- **langchaingo**: To Perform QA with LLMs on the Scraped Data.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/blog-scraper.git
   ```

2. Install dependencies:
  ```bash
  go mod tidy
  ```

3. Run the scraper:
  ```bash
  go run cmd/main/main.go
  ```

## Migrations

1. Create a Migration File:
  ```bash
  goose -dir ./db/migrations create <migration_name> sql
  ```

2. Run the Migration:
  ```bash
  goose -dir ./db/migrations sqlite3 ./db/blogs.db up
  ```

3. Roll Back the last Migration:
  ```bash
  goose -dir ./db/migrations sqlite3 ./db/blogs.db down
  ```

4. Check Migration Status:
  ```bash
  goose -dir ./db/migrations sqlite3 ./db/blogs.db status
  ```


## Build & Run with Docker

- Build:
  ```bash
  docker build --tag IMAGE_NAME .
  ```

- Run:
  ```bash
  docker run -e SECRET_KEY="SECRET_KEY" -p 8000:8000 IMAGE_NAME
  ```

  ***SECRET_KEY***: A secret key for authorization.
