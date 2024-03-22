package model

import (
	"fmt"
	"net/url"
)

// ListOptions specifies the optional parameters to methods that support pagination.
type ListOptions struct {
	// A maximum number of items to retrieve (default 25, max 500).
	Limit int64 `json:"limit,omitempty"`

	// A starting offset in the collection of items (default 0).
	Offset int64 `json:"offset,omitempty"`
}

// Values is used to encode the query parameters into the URL query string.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ListOptions) Values() url.Values {
	v := url.Values{}
	if o.Limit > 0 {
		v.Add("limit", fmt.Sprintf("%d", o.Limit))
	}
	if o.Offset > 0 {
		v.Add("offset", fmt.Sprintf("%d", o.Offset))
	}
	return v
}

// Pagination represents the pagination information.
type Pagination struct {
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}
