package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/psmagicman/terminal-dashboard-app/pkg/api/quote"
	"github.com/psmagicman/terminal-dashboard-app/pkg/config"
)

func main() {
	setupEnvVariables()
	cfg, err := config.LoadConfig("THIS_APP_")
	if err != nil {
		log.Fatal(err, "failed to load config")
		panic(err)
	}

	timeout, err := strconv.Atoi(cfg.Get("request_timeout"))
	if err != nil {
		log.Fatal(err, "failed to convert timeout string to integer")
		panic(err)
	}

	client := getDefaultHttpClient(timeout)

	quoteService := quote.NewQuoteService(client, cfg)
	quote, err := quoteService.GetRandomQuote()
	if err != nil {
		log.Fatal(err, "failed to get quote")
		panic(err)
	}
	fmt.Printf("%s - %s", quote.Quote, quote.Author)
}

func setupEnvVariables() {
	// TODO: add to env variables before running
	os.Setenv("THIS_APP_ZENQUOTES_API_URL", "https://zenquotes.io/api")
	os.Setenv("THIS_APP_USER_AGENT", "ExampleAgent/1.0")
	os.Setenv("THIS_APP_REQUEST_TIMEOUT", "90")
}

func getDefaultHttpClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}
