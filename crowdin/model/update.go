package model

import "errors"

// UpdateRequest defines the structure of a request to update a resource.
type UpdateRequest struct {
	// Patch operation to perform
	Op string `json:"op"`
	// A JSON Pointer as defined by RFC 6901.
	Path string `json:"path"`
	// Value must be one of boolean, integer, string, array of strings,
	// array of integers or object.
	Value any `json:"value"`
}

// Validate checks if the update request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *UpdateRequest) Validate() error {
	if r.Op == "" {
		return errors.New("op is required")
	}
	if r.Path == "" {
		return errors.New("path is required")
	}
	if r.Value == nil {
		return errors.New("value is required")
	}
	return nil
}
