package quote

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/psmagicman/terminal-dashboard-app/pkg/config"
	"github.com/psmagicman/terminal-dashboard-app/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetRandomQuote(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   string
		responseStatus int
		expectedQuote  *ZenQuote
		expectError    bool
		errorMessages  []string
	}{
		{
			name:           "Successful quote response",
			responseBody:   `[{"q":"Test quote","a":"Test author"}]`,
			responseStatus: http.StatusOK,
			expectedQuote:  &ZenQuote{"Test quote", "Test author"},
			expectError:    false,
		},
		{
			name:           "Empty response",
			responseBody:   `[]`,
			responseStatus: http.StatusOK,
			expectedQuote:  nil,
			expectError:    true,
			errorMessages:  []string{"no quotes returned"},
		},
		{
			name:           "HTTP error",
			responseBody:   ``,
			responseStatus: http.StatusInternalServerError,
			expectedQuote:  nil,
			expectError:    true,
			errorMessages:  []string{"requesting quote"},
		}, {
			name:           "Invalid JSON",
			responseBody:   `invalid json`,
			responseStatus: http.StatusOK,
			expectedQuote:  nil,
			expectError:    true,
			errorMessages:  []string{"unmarshalling quote response body"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			testConfig := &config.Config{}
			testConfig.Set("TEST_ZENQUOTES_API_URL", "https://www.example.com")
			testConfig.Set("TEST_USER_AGENT", "TestAgent/1.0")
			mockClient := new(MockHTTPClient)
			mockResponse := &http.Response{
				StatusCode: tt.responseStatus,
				Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
			}
			if tt.name == "HTTP error" {
				mockError := errors.New("requesting quote")
				mockClient.On("Do", mock.Anything).Return(mockResponse, mockError)
			} else {
				mockClient.On("Do", mock.Anything).Return(mockResponse, nil)
			}

			quoteService := NewQuoteService(mockClient, testConfig)

			// When
			quote, err := quoteService.GetRandomQuote()

			// Then
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, quote)
				testutils.TestErrorMessageContains(t, err, tt.errorMessages...)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedQuote, quote)
			}

			mockClient.AssertExpectations(t)
		})
	}
}
