package telegraphist

import (
	"net/http"
	"time"

	"github.com/xamut/telegraphist/telegram"
)

const (
	// DefaultBaseURL is endpoint URL to Telegram Bot API
	DefaultBaseURL string = "https://api.telegram.org/bot"
	// DefaultPort is port for serving the webhook
	DefaultPort int = 443
	// DefaultTimeout is period of time that will be allowed to wait a response
	DefaultTimeout time.Duration = 5 * time.Second
)

// ClientConfig represents settings of API client
type ClientConfig struct {
	BaseURL  string
	BotToken string
	Timeout  time.Duration
}

// NewClient creates a new instance of API client with ClientConfig
func NewClient(config *ClientConfig) (*telegram.Client, error) {
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	httpClient := &http.Client{Timeout: config.Timeout}

	return telegram.NewClient(config.BaseURL, config.BotToken, httpClient)
}
