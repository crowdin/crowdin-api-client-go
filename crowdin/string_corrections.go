package crowdin

import (
	"context"
	"errors"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// StringCorrectionsService provides access to the String Corrections API methods.
// Corrections allow you to suggest changes to source strings.
//
// Note: Crowdin Enterprise only.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/String-Corrections
type StringCorrectionsService struct {
	client *Client
}

// List returns a list of corrections of a string.
// The `StringID` option is required.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.getMany
func (s *StringCorrectionsService) List(ctx context.Context, projectID int, opts *model.StringCorrectionsListOptions) (
	[]*model.StringCorrection, *Response, error,
) {
	if opts == nil || opts.StringID == 0 {
		return nil, nil, errors.New("stringId is required")
	}

	res := new(model.StringCorrectionsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.StringCorrection, 0, len(res.Data))
	for _, correction := range res.Data {
		list = append(list, correction.Data)
	}

	return list, resp, err
}

// Add adds a new correction to a string.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.post
func (s *StringCorrectionsService) Add(ctx context.Context, projectID int, req *model.StringCorrectionAddRequest) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections", projectID), req, res)

	return res.Data, resp, err
}

// DeleteMany deletes all corrections of a string.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.deleteMany
func (s *StringCorrectionsService) DeleteMany(ctx context.Context, projectID, stringID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections?stringId=%d", projectID, stringID), nil)
}

// Get returns a string correction by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.get
func (s *StringCorrectionsService) Get(ctx context.Context, projectID, correctionID int, opts *model.TranslationGetOptions) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections/%d", projectID, correctionID), opts, res)

	return res.Data, resp, err
}

// Restore restores a string correction by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.put
func (s *StringCorrectionsService) Restore(ctx context.Context, projectID, correctionID int) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionResponse)
	resp, err := s.client.Put(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections/%d", projectID, correctionID), nil, res)

	return res.Data, resp, err
}

// Delete deletes a string correction by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.corrections.delete
func (s *StringCorrectionsService) Delete(ctx context.Context, projectID, correctionID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/corrections/%d", projectID, correctionID), nil)
}
