package model

import (
	"errors"
	"fmt"
	"net/url"
)

// StringCorrection represents a correction of a source string.
type StringCorrection struct {
	ID                 int        `json:"id"`
	Text               string     `json:"text"`
	PluralCategoryName string     `json:"pluralCategoryName"`
	User               *ShortUser `json:"user"`
	CreatedAt          string     `json:"createdAt"`
}

// StringCorrectionResponse defines the structure of a response when
// getting a string correction.
type StringCorrectionResponse struct {
	Data *StringCorrection `json:"data"`
}

// StringCorrectionsListResponse defines the structure of a response when
// getting a list of string corrections.
type StringCorrectionsListResponse struct {
	Data []*StringCorrectionResponse `json:"data"`
}

// StringCorrectionsListOptions specifies the parameters to the
// StringCorrectionsService.List method.
type StringCorrectionsListOptions struct {
	// String Identifier. Required.
	StringID int `json:"stringId"`
	// Sort corrections by the specified field.
	// Example: orderBy=createdAt desc,text.
	OrderBy string `json:"orderBy,omitempty"`
	// Enable denormalize placeholders.
	// Enum: 0, 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of StringCorrectionsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *StringCorrectionsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.StringID > 0 {
		v.Add("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		v.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}

	return v, len(v) > 0
}

// StringCorrectionAddRequest defines the structure of a request to
// add a string correction.
type StringCorrectionAddRequest struct {
	// String Identifier.
	StringID int `json:"stringId"`
	// Correction text.
	Text string `json:"text"`
	// Plural form. Enum: zero, one, two, few, many, other.
	// Note: This field becomes required for strings with plural forms.
	PluralCategoryName string `json:"pluralCategoryName,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *StringCorrectionAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StringID == 0 {
		return errors.New("stringId is required")
	}
	if r.Text == "" {
		return errors.New("text is required")
	}

	return nil
}
