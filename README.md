# Blog Scraper

A simple Golang tool for scraping and storing blog posts. This tool extracts titles, authors, publication dates, and content, saving them to a local SQLite3 database.

## Features

- Scrapes blog content with ease
- Stores blog data in SQLite3 for quick access

## Tech Stack

- **Golang**: Core programming language
- **goose**: For DB Migrations
- **SQLite3**: For lightweight data storage
- **html-to-markdown**: For Parsing html blog content to markdown

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
  goose -dir ./db/migrations up
  ```

3. Roll Back the last Migration:
  ```bash
  goose -dir ./db/migrations down
  ```

4. Check Migration Status:
  ```bash
  goose -dir ./db/migrations status
  ```