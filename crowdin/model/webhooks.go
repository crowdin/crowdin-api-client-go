package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Event is a type that represents an event that happens
// in a Crowdin project.
type Event string

const (
	// Project file is added.
	EventFileAdded Event = "file.added"
	// Project file is updated.
	EventFileUpdated Event = "file.updated"
	// Project file is reverted.
	EventFileReverted Event = "file.reverted"
	// Project file is deleted.
	EventFileDeleted Event = "file.deleted"
	// Project file is fully translated.
	EventFileTranslated Event = "file.translated"
	// Project file is fully reviewed.
	EventFileApproved Event = "file.approved"
	// All strings in project are translated.
	EventProjectTranslated Event = "project.translated"
	// All strings in project are approved.
	EventProjectApproved Event = "project.approved"
	// Project are successfully built.
	EventProjectBuilt Event = "project.built"
	// Final translation of string is updated (using Replace in suggestions feature).
	EventTranslationUpdated Event = "translation.updated"
	// Source string is added.
	EventStringAdded Event = "string.added"
	// Source string is updated.
	EventStringUpdated Event = "string.updated"
	// Source string is deleted.
	EventStringDeleted Event = "string.deleted"
	// String comment/issue is added.
	EventStringCommentCreated Event = "stringComment.created"
	// String comment/issue is updated.
	EventStringCommentUpdated Event = "stringComment.updated"
	// String comment/issue is deleted.
	EventStringCommentDeleted Event = "stringComment.deleted"
	// String comment/issue is restored.
	EventStringCommentRestored Event = "stringComment.restored"
	// One of source strings is translated.
	EventSuggestionAdded Event = "suggestion.added"
	// Translation for source string is updated (using Replace in suggestions feature).
	EventSuggestionUpdated Event = "suggestion.updated"
	// One of translations is deleted.
	EventSuggestionDeleted Event = "suggestion.deleted"
	// Translation for string is approved.
	EventSuggestionApproved Event = "suggestion.approved"
	// Approval for previously added translation is removed.
	EventSuggestionDisapproved Event = "suggestion.disapproved"
	// Task is added.
	EventTaskAdded Event = "task.added"
	// Task status was changed.
	EventTaskStatusChanged Event = "task.statusChanged"
	// Task is deleted.
	EventTaskDeleted Event = "task.deleted"

	// Organization webhook events.
	// Project is created.
	EventProjectCreated Event = "project.created"
	// Project is deleted.
	EventProjectDeleted Event = "project.deleted"
)

// ContentType is a type that represents the content type of a webhook.
type WebhookContentType string

const (
	ContentTypeJSON      WebhookContentType = "application/json"
	ContentTypeForm      WebhookContentType = "application/x-www-form-urlencoded"
	ContentTypeMultipart WebhookContentType = "multipart/form-data"
)

// Webhook represents a webhook in Crowdin projects or account.
type Webhook struct {
	ID              int               `json:"id"`
	ProjectID       int               `json:"projectId"`
	Name            string            `json:"name"`
	URL             string            `json:"url"`
	Events          []string          `json:"events"`
	Headers         map[string]string `json:"headers"`
	Payload         map[string]any    `json:"payload"`
	IsActive        bool              `json:"isActive"`
	BatchingEnabled bool              `json:"batchingEnabled"`
	RequestType     string            `json:"requestType"`
	ContentType     string            `json:"contentType"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
}

// UnmarshalJSON unmarshals the JSON data into the Webhook structure.
// Headers property can be either an object (map) or an empty array.
func (w *Webhook) UnmarshalJSON(data []byte) error {
	type alias Webhook
	aux := struct {
		Headers json.RawMessage `json:"headers"`
		*alias
	}{
		alias: (*alias)(w),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("webhook unmarshal error: %w", err)
	}

	if len(aux.Headers) > 0 {
		headers := make(map[string]string)

		// Check if "headers" property is an array.
		if aux.Headers[0] == '[' {
			w.Headers = headers
		} else {
			if err := json.Unmarshal(aux.Headers, &headers); err != nil {
				return fmt.Errorf("webhook headers unmarshal error: %w", err)
			}
			w.Headers = headers
		}
	}

	return nil
}

// WebhookResponse defines the structure of the response when
// getting a single webhook.
type WebhookResponse struct {
	Data *Webhook `json:"data"`
}

// WebhooksListResponse defines the structure of the response
// when getting a list of webhooks.
type WebhooksListResponse struct {
	Data []*WebhookResponse `json:"data"`
}

// WebhookAddRequest defines the structure of the request
// when adding a new webhook.
type WebhookAddRequest struct {
	// Webhook name.
	Name string `json:"name"`
	// Webhook URL.
	URL string `json:"url"`
	// List of events.
	Events []Event `json:"events"`
	// Webhook request type.
	// Enum: GET, POST.
	RequestType string `json:"requestType"`
	// Indicates whether webhook is active.
	// Default: true.
	IsActive *bool `json:"isActive,omitempty"`
	// Indicates whether webhook batching is enabled.
	// Default: false.
	BatchingEnabled *bool `json:"batchingEnabled,omitempty"`
	// Webhook content type.
	// Enum: application/json, application/x-www-form-urlencoded, multipart/form-data.
	// Default: application/json.
	ContentType WebhookContentType `json:"contentType,omitempty"`
	// Webhook headers.
	Headers map[string]string `json:"headers,omitempty"`
	// Custom webhook payload.
	// For more details, see https://developer.crowdin.com/webhooks.
	Payload any `json:"payload,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *WebhookAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.URL == "" {
		return errors.New("url is required")
	}
	if len(r.Events) == 0 {
		return errors.New("events is required")
	}
	if r.RequestType == "" {
		return errors.New("requestType is required")
	} else if r.RequestType != "GET" && r.RequestType != "POST" {
		return errors.New("requestType must be GET or POST")
	}

	return nil
}
