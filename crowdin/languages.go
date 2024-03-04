package crowdin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// LanguagesService handle communication with Languages methods.
type LanguagesService struct {
	client *Client
}

// List returns a list of all supported languages.
// https://developer.crowdin.com/api/v2/#operation/api.languages.getMany
func (s *LanguagesService) List(ctx context.Context, opts *model.ListOptions) ([]*model.Language, *Response, error) {
	path := "languages?" + opts.Values().Encode()
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := new(model.LanguagesListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, resp, err
	}

	langs := make([]*model.Language, 0, len(body.Data))
	for _, lang := range body.Data {
		langs = append(langs, lang.Data)
	}
	return langs, resp, nil
}

// Get returns a language by its identifier.
// https://developer.crowdin.com/api/v2/#operation/api.languages.get
func (s *LanguagesService) Get(ctx context.Context, langID string) (*model.Language, *Response, error) {
	path := fmt.Sprintf("languages/%s", langID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := new(model.LanguagesGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, resp, err
	}
	return body.Data, resp, nil
}

// Add adds a new custom language.
// https://developer.crowdin.com/api/v2/#operation/api.languages.post
func (s *LanguagesService) Add(ctx context.Context, r *model.AddLanguageRequest) (*model.Language, *Response, error) {
	if err := r.Validate(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, "languages", r)
	if err != nil {
		return nil, nil, err
	}

	body := new(model.LanguagesAddResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, resp, err
	}
	return body.Data, resp, nil
}

// Edit updates a custom language by its identifier.
// https://developer.crowdin.com/api/v2/#operation/api.languages.patch
func (s *LanguagesService) Edit(ctx context.Context, langID string, r *model.EditLanguageRequest) (*model.Language, *Response, error) {
	if err := r.Validate(); err != nil {
		return nil, nil, err
	}

	path := fmt.Sprintf("languages/%s", langID)
	req, err := s.client.NewRequest(ctx, http.MethodPatch, path, r)
	if err != nil {
		return nil, nil, err
	}

	body := new(model.LanguagesEditResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, resp, err
	}
	return body.Data, resp, nil
}

// Delete deletes a custom language by its identifier.
// https://developer.crowdin.com/api/v2/#operation/api.languages.delete
func (s *LanguagesService) Delete(ctx context.Context, langID string) (*Response, error) {
	path := fmt.Sprintf("languages/%s", langID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
