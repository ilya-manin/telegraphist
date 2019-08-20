package telegram

/*
WebhookInfo contains information about the current status of a webhook.
*/
type WebhookInfo struct {
	URL                  string   `json:"url"`                          // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate bool     `json:"has_custom_certificate"`       // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount   int64    `json:"pending_update_count"`         // Number of updates awaiting delivery
	LastErrorDate        int64    `json:"last_error_date,omitempty"`    // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage     string   `json:"last_error_message,omitempty"` // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	MaxConnections       int64    `json:"max_connections,omitempty"`    // Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`    // Optional. A list of update types the bot is subscribed to. Defaults to all update types
}
