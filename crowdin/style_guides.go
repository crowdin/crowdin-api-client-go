package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// StyleGuidesService provides access to the Style Guides API methods.
// Style guides define rules that translators and AI models follow
// when translating project content.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Style-Guides
type StyleGuidesService struct {
	client *Client
}

// List returns a list of style guides.
//
// https://developer.crowdin.com/api/v2/#operation/api.style-guides.getMany
func (s *StyleGuidesService) List(ctx context.Context, opts *model.StyleGuidesListOptions) (
	[]*model.StyleGuide, *Response, error,
) {
	res := new(model.StyleGuidesListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/style-guides", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.StyleGuide, 0, len(res.Data))
	for _, guide := range res.Data {
		list = append(list, guide.Data)
	}

	return list, resp, err
}

// Create creates a new style guide.
//
// https://developer.crowdin.com/api/v2/#operation/api.style-guides.post
func (s *StyleGuidesService) Create(ctx context.Context, req *model.StyleGuideCreateRequest) (
	*model.StyleGuide, *Response, error,
) {
	res := new(model.StyleGuideResponse)
	resp, err := s.client.Post(ctx, "/api/v2/style-guides", req, res)

	return res.Data, resp, err
}

// Get returns a style guide by its identifier.
//
// https://developer.crowdin.com/api/v2/#operation/api.style-guides.get
func (s *StyleGuidesService) Get(ctx context.Context, styleGuideID int) (*model.StyleGuide, *Response, error) {
	res := new(model.StyleGuideResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/style-guides/%d", styleGuideID), nil, res)

	return res.Data, resp, err
}

// Edit updates a style guide by its identifier.
//
// Request body:
// - op - operation to perform. Enum: replace, test.
// - path (json-pointer) - path to the field to update.
// Enum: "/name", "/aiInstructions", "/languageIds", "/projectIds", "/isShared", "/storageId",
// "/groupId" (Enterprise only).
// - value - new value for the field.
//
// https://developer.crowdin.com/api/v2/#operation/api.style-guides.patch
func (s *StyleGuidesService) Edit(ctx context.Context, styleGuideID int, req []*model.UpdateRequest) (
	*model.StyleGuide, *Response, error,
) {
	res := new(model.StyleGuideResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/style-guides/%d", styleGuideID), req, res)

	return res.Data, resp, err
}

// Delete removes a style guide by its identifier.
//
// https://developer.crowdin.com/api/v2/#operation/api.style-guides.delete
func (s *StyleGuidesService) Delete(ctx context.Context, styleGuideID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/style-guides/%d", styleGuideID), nil)
}
