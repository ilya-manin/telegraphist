package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Client keeps data to perform Bot API methods
type Client struct {
	baseURL    string
	botToken   string
	httpClient *http.Client
}

func (c *Client) urlFor(apiMethod string) (string, error) {
	var err error

	if apiMethod == "" {
		return "", fmt.Errorf("apiMethod should not be empty")
	}

	url := fmt.Sprintf("%s%s/%s", c.baseURL, c.botToken, apiMethod)
	return url, err
}

type files map[string]*os.File

func (f files) add(name string, file *os.File) {
	if f == nil {
		f = make(files)
	}
	f[name] = file
}

func (c *Client) performRequest(apiMethod string, params interface{}, filesToUpload *files, result interface{}) error {
	var err error

	url, err := c.urlFor(apiMethod)
	if err != nil {
		return err
	}

	body, contentType, err := buildRequestBody(params, filesToUpload)
	if err != nil {
		return err
	}

	httpResp, err := c.httpClient.Post(url, contentType, body)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	var resp response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return fmt.Errorf("Error code: %v, %s", resp.ErrorCode, resp.Description)
	}

	return json.Unmarshal(resp.Result, result)
}

// NewClient creates a new object with all Bot API methods
func NewClient(baseURL string, botToken string, httpClient *http.Client) (*Client, error) {
	var err error

	if baseURL == "" {
		return nil, fmt.Errorf("baseURL should not be empty")
	}

	if botToken == "" {
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
