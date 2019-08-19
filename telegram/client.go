package telegram

import (
	"fmt"
	"net/http"
)

// Client keeps data to perform Bot API methods
type Client struct {
	baseURL    string
	botToken   string
	httpClient *http.Client
}

func (c *Client) urlFor(apiMethod string) (string, error) {
	var err error

	if len(apiMethod) == 0 {
		return "", fmt.Errorf("apiMethod should not be empty")
	}

	url := fmt.Sprintf("%s%s/%s", c.baseURL, c.botToken, apiMethod)
	return url, err
}

func (c *Client) performRequest() error {
	var err error

	url, err := c.urlFor("getMe")
	fmt.Printf("Performing apiMethod by calling %s", url)

	return err
}

// NewClient creates a new object with all Bot API methods
func NewClient(baseURL string, botToken string, httpClient *http.Client) (*Client, error) {
	var err error

	if len(baseURL) == 0 {
		return nil, fmt.Errorf("baseURL should not be empty")
	}

	if len(botToken) == 0 {
		return nil, fmt.Errorf("botToken should not be empty")
	}

	if httpClient == nil {
		return nil, fmt.Errorf("httpClient should not be nil")
	}

	return &Client{
		baseURL:    baseURL,
		botToken:   botToken,
		httpClient: httpClient,
	}, err
}
