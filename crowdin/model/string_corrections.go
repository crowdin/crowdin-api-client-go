package model

import (
        "errors"
        "fmt"
        "net/url"
)

type StringCorrection struct {
	ID                  int       `json:"id"`
	Text                string    `json:"text"`
	PluralCategoryName  *string   `json:"pluralCategoryName,omitempty"`
	User                *User     `json:"user,omitempty"`
	CreatedAt           string    `json:"createdAt"`
}

type StringCorrectionGetResponse struct {
	Data *StringCorrection `json:"data"`
}

type StringCorrectionsListResponse struct {
	Data       []*StringCorrectionResponse `json:"data"`
	Pagination *Pagination                 `json:"pagination,omitempty"`
}

type StringCorrectionResponse struct {
	Data *StringCorrection `json:"data"`
}

type StringCorrectionsListOptions struct {
	StringID                int    `json:"stringId,omitempty"`
	Limit                   int    `json:"limit,omitempty"`
	Offset                  int    `json:"offset,omitempty"`
	OrderBy                 string `json:"orderBy,omitempty"`
	DenormalizePlaceholders *int   `json:"denormalizePlaceholders,omitempty"`
}

type StringCorrectionGetOptions struct {
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`
}

type StringCorrectionAddRequest struct {
	StringID            int     `json:"stringId"`
	Text                string  `json:"text"`
	PluralCategoryName  *string `json:"pluralCategoryName,omitempty"`
}

type StringCorrectionsDeleteOptions struct {
	StringID int `json:"stringId"`
}

func (o *StringCorrectionsListOptions) Values() (url.Values, bool) {
	v := url.Values{}
	if o == nil {
		return v, false
	}
	if o.StringID > 0 {
		v.Set("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", o.Limit))
	}
	if o.Offset > 0 {
		v.Set("offset", fmt.Sprintf("%d", o.Offset))
	}
	if o.OrderBy != "" {
		v.Set("orderBy", o.OrderBy)
	}
	if o.DenormalizePlaceholders != nil {
		v.Set("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	return v, len(v) > 0
}

func (o *StringCorrectionGetOptions) Values() (url.Values, bool) {
	v := url.Values{}
	if o == nil {
		return v, false
	}
	if o.DenormalizePlaceholders != nil {
		v.Set("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	return v, len(v) > 0
}

func (o *StringCorrectionsDeleteOptions) Values() (url.Values, bool) {
	v := url.Values{}
	if o == nil {
		return v, false
	}
	if o.StringID > 0 {
		v.Set("stringId", fmt.Sprintf("%d", o.StringID))
	}
	return v, len(v) > 0
}

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

