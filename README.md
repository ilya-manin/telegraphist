# Telegraphist
[![GoDoc Telegraphist](https://godoc.org/github.com/xamut/telegraphist?status.svg)](https://godoc.org/github.com/xamut/telegraphist)
[![CI status](https://github.com/xamut/telegraphist/workflows/CI/badge.svg)](https://github.com/xamut/telegraphist/actions)

Bindings for Telegram Bot API on Go. [API v4.4 (29 July 2019)](https://core.telegram.org/bots/api#july-29-2019) is supported.
All methods and types represent their alter ego from Telegram Bot API so you would expect what you see in the official documentation.
And you don't need to study the new API to use this library.

## Documentation

* Official Telegram Bot API [documentation](https://core.telegram.org/bots/api)
* [telegraphist](https://godoc.org/github.com/xamut/telegraphist)
* [telegraphist/telegram](https://godoc.org/github.com/xamut/telegraphist/telegram)

## Installation

```bash
go get -u github.com/xamut/telegraphist
```

## Usage

```golang
package main

import (
	"encoding/json"
	"fmt"
	"os"

	Telegraphist "github.com/xamut/telegraphist"
	Telegram "github.com/xamut/telegraphist/telegram"
)

func printResp(resp interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	respAsJSON, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(respAsJSON))
}

func main() {
	telegraphist, err := Telegraphist.NewClient(&Telegraphist.ClientConfig{
		BotToken: "YOUR_BOT_TOKEN",
	})

	if err != nil {
		fmt.Println(err)
	}

	printResp(telegraphist.GetMe()) // Simple method to obtain basic bot info
					// https://core.telegram.org/bots/api#getme
					// Returns an User struct

	chatID := int64(<YOUR_CHAT_ID>)
	printResp(telegraphist.SendMessage(&Telegram.SendMessageParams{ // Send a text message to a chat
		ChatID: chatID,  				        // https://core.telegram.org/bots/api#sendmessage
		Text:   "Hello there!", 				// Returns a Message struct
	}))

	photo, _ := os.Open("<PATH_TO_PHOTO>")
	printResp(telegraphist.SendPhoto(&Telegram.SendPhotoParams{ // Send a photo to a chat
		ChatID:  chatID,              			    // https://core.telegram.org/bots/api#sendphoto
		Photo:   photo,					    // Returns a Message struct
		Caption: "Just a nice photo",
	}))
}
```

## Getting updates

### Webhook
* Find your IP

```bash
curl ifconfig.co
```

* Use [nip.io](https://nip.io) to pretend you have a domain, for example, "app.X.X.X.X.nip.io"

* Generate a certificate, replace X.X.X.X with your IP:

```bash
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/O=Organization/CN=app.X.X.X.X.nip.io"
```

* Run server

```golang
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	Telegraphist "github.com/xamut/telegraphist"
	Telegram "github.com/xamut/telegraphist/telegram"
)

func combineURL(host string, parts ...string) string {
	baseURL, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}
	parts = append([]string{baseURL.Path}, parts...)
	baseURL.Path = path.Join(parts...)

	return baseURL.String()
}

func main() {
	telegraphist, err := Telegraphist.NewClient(&Telegraphist.ClientConfig{
		BotToken: "<YOUR_BOT_TOKEN>",
	})

	if err != nil {
		log.Fatal(err)
	}

	host := "https://app.X.X.X.X.nip.io" // For example: "https://example.com"
	webhookPath := "<YOUR_WEBHOOK_PATH>" // For example: "/telegram_webhook"

	cert, _ := os.Open("./cert.pem")
	ok, err := telegraphist.SetWebhook(&Telegram.SetWebhookParams{
		URL:         combineURL(host, webhookPath),
		Certificate: cert,
	})

	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		log.Println("Something went wrong and the webhook wasn't set up")
	}

	handlers := map[string]func(update Telegram.Update) error{
		// "message":              func(update Telegram.Update) error {},
		// "edited_message":       func(update Telegram.Update) error {},
		// "channel_post":         func(update Telegram.Update) error {},
		// "edited_channel_post":  func(update Telegram.Update) error {},
		// "inline_query":         func(update Telegram.Update) error {},
		// "chosen_inline_result": func(update Telegram.Update) error {},
		// "callback_query":       func(update Telegram.Update) error {},
		// "shipping_query":       func(update Telegram.Update) error {},
		// "pre_checkout_query":   func(update Telegram.Update) error {},
		// "poll":                 func(update Telegram.Update) error {},
		"always": func(update Telegram.Update) error {
			js, _ := json.MarshalIndent(update, "", "\t")
			fmt.Println(string(js))
			return nil
		},
	}

	err = Telegraphist.NewServer(&Telegraphist.ServerConfig{
		EnableHTTPS: true,
		Webhook:     webhookPath,
	}, &handlers)

	if err != nil {
		log.Fatal(err)
	}
}
```

NOTE: Use this example only for testing purposes, for production use something like nginx or traefik to serve HTTPS (and/or to (re)generate Let's Encrypt certificate).

## License

Telegraphist is released under the [MIT License](https://opensource.org/licenses/MIT).
