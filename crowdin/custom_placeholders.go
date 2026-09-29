package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// CustomPlaceholdersService provides access to the Custom Placeholders API.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/Custom-Placeholders
type CustomPlaceholdersService struct {
	client *Client
}

// List returns a list of the custom placeholders of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-placeholders.getMany
func (s *CustomPlaceholdersService) List(ctx context.Context, opts *model.ListOptions) (
	[]*model.CustomPlaceholder, *Response, error,
) {
	res := new(model.CustomPlaceholdersListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/custom-placeholders", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.CustomPlaceholder, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns a custom placeholder of the organization by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-placeholders.get
func (s *CustomPlaceholdersService) Get(ctx context.Context, customPlaceholderID int) (
	*model.CustomPlaceholder, *Response, error,
) {
	res := new(model.CustomPlaceholderResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/custom-placeholders/%d", customPlaceholderID), nil, res)

	return res.Data, resp, err
}

// Add creates a custom placeholder for the organization.
// Assign it to a project with ProjectPlaceholdersService.Add.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-placeholders.post
func (s *CustomPlaceholdersService) Add(ctx context.Context, req *model.CustomPlaceholderAddRequest) (
	*model.CustomPlaceholder, *Response, error,
) {
	res := new(model.CustomPlaceholderResponse)
	resp, err := s.client.Post(ctx, "/api/v2/custom-placeholders", req, res)

	return res.Data, resp, err
}

// Edit updates a custom placeholder of the organization.
//
// Request body:
//   - op: The operation to perform. Enum: replace, test.
//   - path: A JSON Pointer as defined in RFC 6901. Enum: "/description", "/definition", "/argumentDelimiter".
//   - value: The value to be used within the operations. The value must be a string.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-placeholders.patch
func (s *CustomPlaceholdersService) Edit(ctx context.Context, customPlaceholderID int, req []*model.UpdateRequest) (
	*model.CustomPlaceholder, *Response, error,
) {
	res := new(model.CustomPlaceholderResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/custom-placeholders/%d", customPlaceholderID), req, res)

	return res.Data, resp, err
}

// Delete deletes a custom placeholder of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-placeholders.delete
func (s *CustomPlaceholdersService) Delete(ctx context.Context, customPlaceholderID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/custom-placeholders/%d", customPlaceholderID), nil)
}
