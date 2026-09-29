package model

import (
	"errors"
	"fmt"
	"net/url"
)

// StringComment represents a Crowdin string comment.
type StringComment struct {
	ID          int        `json:"id"`
	Text        string     `json:"text"`
	UserID      int        `json:"userId"`
	StringID    int        `json:"stringId"`
	User        *ShortUser `json:"user"`
	String      *String    `json:"string"`
	ProjectID   int        `json:"projectId"`
	LanguageID  string     `json:"languageId"`
	Type        string     `json:"type"`
	IssueType   string     `json:"issueType"`
	IssueStatus string     `json:"issueStatus"`
	ResolverID  int        `json:"resolverId"`
	Resolver    *ShortUser `json:"resolver"`
	ResolvedAt  string     `json:"resolvedAt"`
	CreatedAt   string     `json:"createdAt"`

	IsShared             *bool         `json:"isShared,omitempty"`
	SenderOrganization   *Organization `json:"senderOrganization,omitempty"`
	ResolverOrganization *Organization `json:"resolverOrganization,omitempty"`

	// List of attachments added to the comment.
	Attachments []*StringCommentAttachment `json:"attachments,omitempty"`
	// File identifier. It is set for asset comments only.
	FileID *int `json:"fileId,omitempty"`
	// File object. It is set for asset comments only.
	File *StringCommentFile `json:"file,omitempty"`
}

// StringCommentAttachment represents a file attached to a string comment.
type StringCommentAttachment struct {
	// Attachment ID.
	ID int `json:"id"`
	// Original file name.
	Name string `json:"name"`
	// MIME type.
	Mime string `json:"mime"`
	// File size in bytes.
	Size int `json:"size"`
	// Attachment category. Enum: image, video, audio, document, other.
	Category string `json:"category"`
	// Thumbnail URL.
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// Preview URL.
	URL string `json:"url"`
	// Download URL.
	DownloadURL string `json:"downloadUrl"`
}

// StringCommentFile represents the asset file of an asset comment.
type StringCommentFile struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Context string `json:"context"`
}

type String struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Type    string `json:"type"`
	Context string `json:"context"`
	FileID  int    `json:"fileId"`
}

type Organization struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
}

// StringCommentsResponse defines the structure of of the response when
// getting a single string comment.
type StringCommentsResponse struct {
	Data *StringComment `json:"data"`
}

// StringCommentsListResponse defines the structure of the response when
// getting a list of string comments.
type StringCommentsListResponse struct {
	Data []*StringCommentsResponse `json:"data"`
}

// StringCommentsListOptions specifies the optional parameters to the
// StringCommentsService.List method.
type StringCommentsListOptions struct {
	// Sort results by specified field.
	// Enum: id, text, type, createdAt, resolvedAt, issueStatus, issueType.
	// Example: orderBy=createdAt desc,text
	OrderBy string `json:"orderBy,omitempty"`
	// String Identifier.
	StringID int `json:"stringId,omitempty"`
	// Defines string comment type.
	// Enum: comment, issue.
	// Note: `type=comment` can't be used with `issueType` or `issueStatus`
	// in same request.
	Type string `json:"type,omitempty"`
	// Defines issue type. It can be one issue type or multiple issue types.
	// Enum: general_question, translation_mistake, context_request, source_mistake.
	// Example: issueType=general_question,translation_mistake
	IssueType []string `json:"issueType,omitempty"`
	// Defines issue resolution status.
	// Enum: resolved, unresolved.
	IssueStatus string `json:"issueStatus,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the StringCommentsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *StringCommentsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Set("orderBy", o.OrderBy)
	}
	if o.StringID != 0 {
		v.Set("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.Type != "" {
		v.Set("type", o.Type)
	}
	if o.IssueType != nil {
		v.Set("issueType", JoinSlice(o.IssueType))
	}
	if o.IssueStatus != "" {
		v.Set("issueStatus", o.IssueStatus)
	}

	return v, len(v) > 0
}

// StringCommentsAddRequest defines the structure of the request to add a
// new string comment.
type StringCommentsAddRequest struct {
	// Text of the comment.
	Text string `json:"text"`
	// String Identifier.
	// Note: Can't be used with `fileId` in the same request.
	StringID int `json:"stringId,omitempty"`
	// File Identifier. Use it to add a comment to an asset file.
	// Note: Can't be used with `stringId` in the same request.
	FileID int `json:"fileId,omitempty"`
	// Target Language Identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// Defines comment or issue.
	// Enum: comment, issue.
	Type string `json:"type"`
	// Defines issue type.
	// Enum: general_question, translation_mistake, context_request, source_mistake.
	// Default: general_question.
	IssueType string `json:"issueType,omitempty"`
	// Defines shared comment or issue.
	IsShared *bool `json:"isShared,omitempty"`
	// List of attachments to be added to the comment.
	Attachments []*StringCommentAttachmentRequest `json:"attachments,omitempty"`
}

// StringCommentAttachmentRequest defines an attachment to be added
// to a string comment.
type StringCommentAttachmentRequest struct {
	// Storage Identifier.
	ID int `json:"id"`
}

// Validate checks if the StringCommentsAddRequest is valid.
// It implements the crowdin.RequestValidator interface.
func (r *StringCommentsAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Text == "" {
		return errors.New("text is required")
	}
	if r.StringID == 0 && r.FileID == 0 {
		return errors.New("stringId or fileId is required")
	}
	if r.StringID != 0 && r.FileID != 0 {
		return errors.New("stringId and fileId cannot be used in the same request")
	}
	if r.TargetLanguageID == "" {
		return errors.New("targetLanguageId is required")
	}
	if r.Type == "" {
		return errors.New("type is required")
	}
	for _, a := range r.Attachments {
		if a == nil || a.ID == 0 {
			return errors.New("attachment id is required")
		}
	}

	return nil
}
