package quote

import (
	"net/http"

	"github.com/psmagicman/terminal-dashboard-app/pkg/config"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type QuoteService struct {
	client HTTPClient
	config *config.Config
}

func NewQuoteService(client HTTPClient, config *config.Config) *QuoteService {
	return &QuoteService{
		client: client,
		config: config,
	}
}
