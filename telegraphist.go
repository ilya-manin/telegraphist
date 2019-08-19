package telegraphist

import (
	"net/http"
	"time"

	"github.com/xamut/telegraphist/telegram"
)

const (
	// DefaultBaseURL is endpoint URL to Telegram Bot API
	DefaultBaseURL string = "https://api.telegram.org/bot"
	// DefaultTimeout is period of time that will be allowed to wait a response
	DefaultTimeout time.Duration = 10 * time.Second
)

// Config represents settings of API client
type Config struct {
	BaseURL  string
	BotToken string
	Timeout  time.Duration
}

// New creates a new instance of API client with Config
func New(config *Config) (*telegram.Client, error) {
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	httpClient := &http.Client{Timeout: config.Timeout}

	return telegram.NewClient(config.BaseURL, config.BotToken, httpClient)
}
