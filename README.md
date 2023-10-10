# Simple Web Scraper

This Go project is a simple web scraper that scrapes web content from web pages and finds the top 10 words in the content. It then saves this information to a file.

## Functionality

- Scrapes web content from specified web pages.
- Analyzes the content to find the top 10 most frequently occurring words.
- Saves the list of top 10 words to a file.


#### The maximum number of goroutines to use for scraping concurrently

In the code snippet above, boundChan is a channel used to control the maximum number of goroutines that can be used for scraping concurrently.

```go
// internal/controler/controller.go
boundChan := make(chan struct{}, opts.MaxGo)
c.boundChan <- struct{}{}
go c.processUrl(url)
```

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
  maxGo: 4        # The maximum number of goroutines to use for scraping concurrently.
  minLen: 4       # The minimum word length to consider when analyzing web content.
  maxLen: 10      # The maximum word length to consider when analyzing web content.
```

## Logs

```bash
time=2023-10-10T20:13:47.507+02:00 level=INFO msg="starting web scraper" env=local maxGo=4 "min word lenght"=4 "max word lenght"=10
time=2023-10-10T20:13:47.507+02:00 level=DEBUG msg="debug messages are enabled"
time=2023-10-10T20:13:47.507+02:00 level=INFO msg="processing url" url=https://en.wikipedia.org/wiki/C++23
time=2023-10-10T20:13:47.507+02:00 level=INFO msg="skipping url" url=https://en.wikipedia.org/wiki/C++23
time=2023-10-10T20:13:47.507+02:00 level=INFO msg="processing url" url=https://www.facebook.com/
time=2023-10-10T20:13:48.004+02:00 level=INFO msg="saving file"
time=2023-10-10T20:13:48.004+02:00 level=INFO msg="saved file" path=output.txt

```
