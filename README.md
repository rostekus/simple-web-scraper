# Simple Web Scraper

This Go project is a simple web scraper that scrapes web content from web pages and finds the top 10 words in the content. It then saves this information to a file.

## Functionality

- Scrapes web content from specified web pages.
- Analyzes the content to find the top 10 most frequently occurring words.
- Saves the list of top 10 words to a file.

## Getting Started

To use this project, follow these steps:

1. Clone the repository to your local machine:

```shell
git clone https://github.com/rostekus/simple-web-scraper.git
```
2. Run the application

```shell
go run cmd/web-sraper/main.go
```

## Configuration

You can configure the behavior of the web scraper by editing the `config.yml` file. Here's an explanation of the available configuration options:

```yaml
env: "local"

scraper:
  timeout: 4s     # The timeout for making HTTP requests to web pages (e.g., "4s" for 4 seconds).
  maxGo: 4        # The maximum number of goroutines to use for scraping concurrently.
  minLen: 4       # The minimum word length to consider when analyzing web content.
  maxLen: 10      # The maximum word length to consider when analyzing web content.
