package telegraphist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/xamut/telegraphist/telegram"
)

var allowedEvents = [...]string{
	"always",
	"callback_query",
	"channel_post",
	"chosen_inline_result",
	"edited_channel_post",
	"edited_message",
	"inline_query",
	"message",
	"poll",
	"pre_checkout_query",
	"shipping_query",
}

// ServerConfig represents settings of the webhook server
type ServerConfig struct {
	Debug       bool
	EnableHTTPS bool
	Logger      Logger
	Port        int
	Timeout     time.Duration
	Webhook     string
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type eventsRouter struct {
	handlers *map[string]func(update telegram.Update) error
	logger   *serverLogger
	webhook  string
	debug    bool
}

func (router *eventsRouter) handle(event string, update telegram.Update) error {
	if (*(router.handlers))[event] != nil {
		router.logger.info("Process \"%s\" event with update_id = %v", event, update.UpdateID)
		return (*(router.handlers))[event](update)
	}

	router.logger.warn("Unknown handler for \"%s\" event with update_id = %v", event, update.UpdateID)
	return nil
}

func (router *eventsRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if router.debug {
		rawBody, err := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
		if err != nil {
			router.logger.err(err.Error())
		}
		router.logger.debug("Received data: %s", string(rawBody))
	}

	defer req.Body.Close()

	if req.Method != "POST" || path.Clean(req.URL.Path) != router.webhook {
		router.logger.err("Received unknown %s request on %s", req.Method, req.URL)
		res.WriteHeader(http.StatusNotFound)
		return
	}

	var update telegram.Update
	err := json.NewDecoder(req.Body).Decode(&update)
	if err != nil {
		router.logger.err("Error occurred while parsing an update message: %s", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	events := []string{}

	if (*(router.handlers))["always"] != nil {
		events = append(events, "always")
	}

	switch {
	case update.Message != nil:
		events = append(events, "message")
	case update.EditedMessage != nil:
		events = append(events, "edited_message")
	case update.ChannelPost != nil:
		events = append(events, "channel_post")
	case update.EditedChannelPost != nil:
		events = append(events, "edited_channel_post")
	case update.InlineQuery != nil:
		events = append(events, "inline_query")
	case update.ChosenInlineResult != nil:
		events = append(events, "chosen_inline_result")
	case update.CallbackQuery != nil:
		events = append(events, "callback_query")
	case update.ShippingQuery != nil:
		events = append(events, "shipping_query")
	case update.PreCheckoutQuery != nil:
		events = append(events, "pre_checkout_query")
	case update.Poll != nil:
		events = append(events, "poll")
	}

	for _, event := range events {
		err = router.handle(event, update)
		if err != nil {
			router.logger.err(err.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.WriteHeader(http.StatusOK)
}

type serverLogger struct {
	Logger
}

func (l *serverLogger) info(format string, v ...interface{}) {
	l.log("info", format, v...)
}

func (l *serverLogger) err(format string, v ...interface{}) {
	l.log("error", format, v...)
}

func (l *serverLogger) warn(format string, v ...interface{}) {
	l.log("warn", format, v...)
}

func (l *serverLogger) debug(format string, v ...interface{}) {
	l.log("debug", format, v...)
}

func (l *serverLogger) log(logLevel string, format string, v ...interface{}) {
	l.Printf("%s: %s", strings.ToUpper(logLevel), fmt.Sprintf(format, v...))
}

func contains(list [11]string, item string) bool {
	for _, currentItem := range list {
		if currentItem == item {
			return true
		}
	}
	return false
}

func validateHandlers(handlers *map[string]func(update telegram.Update) error) error {
	handlersCounter := 0
	for event, fn := range *handlers {
		if contains(allowedEvents, event) && fn != nil {
			handlersCounter++
		}
	}

	if handlersCounter == 0 {
		return fmt.Errorf("Please configure at least one event handler")
	}

	return nil
}

// NewServer create a new instance of the web server to process Telegram API webhook
func NewServer(config *ServerConfig, handlers *map[string]func(update telegram.Update) error) error {
	if config.Logger == nil {
		config.Logger = log.New(os.Stdout, "[telegraphist] ", log.Ldate|log.Ltime|log.LUTC)
	}

	if config.Port == 0 {
		config.Port = DefaultPort
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	if config.Webhook == "" {
		return fmt.Errorf("Webhook should not be empty")
	}

	err := validateHandlers(handlers)
	if err != nil {
		return err
	}

	sLogger := &serverLogger{config.Logger}

	handler := &eventsRouter{
		debug:    config.Debug,
		handlers: handlers,
		logger:   sLogger,
		webhook:  config.Webhook,
	}

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", config.Port),
		Handler:      handler,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}

	sLogger.Printf("Server is staring on %s", server.Addr)

	if config.EnableHTTPS {
		return server.ListenAndServeTLS("cert.pem", "key.pem")
	}

	return server.ListenAndServe()
}
