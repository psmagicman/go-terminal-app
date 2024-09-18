package quote

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type QuoteService struct {
	client HTTPClient
}

func NewQuoteService(client HTTPClient) *QuoteService {
	return &QuoteService{
		client: client,
	}
}
