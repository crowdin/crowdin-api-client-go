package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// ProjectPlaceholdersService provides access to the project placeholders
// (custom placeholders assigned to a project and system placeholders) of the Projects API.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Projects
type ProjectPlaceholdersService struct {
	client *Client
}

// List returns a list of the custom placeholders assigned to the project.
// The placeholders Crowdin ships are listed by ListSystemPlaceholders.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.placeholders.getMany
func (s *ProjectPlaceholdersService) List(ctx context.Context, projectID int, opts *model.ListOptions) (
	[]*model.ProjectPlaceholder, *Response, error,
) {
	res := new(model.ProjectPlaceholdersListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/placeholders", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ProjectPlaceholder, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns one custom placeholder assigned to the project, with the settings it has there.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.placeholders.get
func (s *ProjectPlaceholdersService) Get(ctx context.Context, projectID, projectPlaceholderID int) (
	*model.ProjectPlaceholder, *Response, error,
) {
	res := new(model.ProjectPlaceholderResponse)
	path := fmt.Sprintf("/api/v2/projects/%d/placeholders/%d", projectID, projectPlaceholderID)
	resp, err := s.client.Get(ctx, path, nil, res)

	return res.Data, resp, err
}

// Add assigns a custom placeholder of the organization to the project.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.placeholders.post
func (s *ProjectPlaceholdersService) Add(ctx context.Context, projectID int, req *model.ProjectPlaceholderAddRequest) (
	*model.ProjectPlaceholder, *Response, error,
) {
	res := new(model.ProjectPlaceholderResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/placeholders", projectID), req, res)

	return res.Data, resp, err
}

// Edit updates the settings a custom placeholder has in the project.
// The placeholder definition itself is not edited here.
//
// Request body:
//   - op: The operation to perform. Enum: replace, test.
//   - path: A JSON Pointer as defined in RFC 6901. Enum: "/index", "/type", "/isBlocking", "/formats".
//   - value: The value to be used within the operations.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.placeholders.patch
func (s *ProjectPlaceholdersService) Edit(ctx context.Context, projectID, projectPlaceholderID int, req []*model.UpdateRequest) (
	*model.ProjectPlaceholder, *Response, error,
) {
	res := new(model.ProjectPlaceholderResponse)
	path := fmt.Sprintf("/api/v2/projects/%d/placeholders/%d", projectID, projectPlaceholderID)
	resp, err := s.client.Patch(ctx, path, req, res)

	return res.Data, resp, err
}

// Delete unassigns a custom placeholder from the project.
// The placeholder itself stays defined for the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.placeholders.delete
func (s *ProjectPlaceholdersService) Delete(ctx context.Context, projectID, projectPlaceholderID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/placeholders/%d", projectID, projectPlaceholderID), nil)
}

// ListSystemPlaceholders returns every placeholder Crowdin ships,
// with the state it has in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.system-placeholders.getMany
func (s *ProjectPlaceholdersService) ListSystemPlaceholders(ctx context.Context, projectID int, opts *model.ListOptions) (
	[]*model.ProjectSystemPlaceholder, *Response, error,
) {
	res := new(model.ProjectSystemPlaceholdersListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/system-placeholders", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ProjectSystemPlaceholder, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// SystemPlaceholderBatchOperations turns system placeholders on and off in the project.
// It returns the placeholders it changed.
//
// Request body:
//   - op: The operation to perform. Enum: replace, test.
//   - path: A JSON Pointer as defined in RFC 6901. The first segment is the key of the
//     system placeholder, and "/isEnabled" is the only property it accepts
//     (e.g. "/printfSpecifier/isEnabled").
//   - value: Whether Crowdin recognizes the placeholder in the project (boolean).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.system-placeholders.batchPatch
func (s *ProjectPlaceholdersService) SystemPlaceholderBatchOperations(ctx context.Context, projectID int, req []*model.UpdateRequest) (
	[]*model.ProjectSystemPlaceholder, *Response, error,
) {
	res := new(model.ProjectSystemPlaceholdersListResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/system-placeholders", projectID), req, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ProjectSystemPlaceholder, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}
