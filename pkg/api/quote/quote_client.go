package quote

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	zQuotesRandomAPIURL = "https://zenquotes.io/api/random"
	userAgent           = "psmagicmain/1.0"
)

// Gets a random quote from zenquotes.io
func (qs *QuoteService) GetRandomQuote() (*ZenQuote, error) {
	req, err := http.NewRequest("GET", zQuotesRandomAPIURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "application/json")

	resp, err := qs.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "requesting quote")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}
	defer resp.Body.Close()

	var quotes []ZenQuote
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling quote response body")
	}

	if len(quotes) == 0 {
		return nil, errors.New("no quotes returned")
	}

	return &quotes[0], nil
}
