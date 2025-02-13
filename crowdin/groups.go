package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// Groups allow you to organize your projects based on specific characteristics.
// Use API to manage projects and groups, change their settings, or remove them
// from organization if required.
//
// Crowdin API docs:
// https://developer.crowdin.com/enterprise/api/v2/#tag/Projects-and-Groups
type GroupsService struct {
	client *Client
}

// List returns a list of groups.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.groups.getMany
func (s *GroupsService) List(ctx context.Context, opts *model.GroupsListOptions) ([]*model.Group, *Response, error) {
	res := new(model.GroupsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/groups", opts, res)
	if err != nil {
		return nil, resp, err
	}

	groups := make([]*model.Group, 0, len(res.Data))
	for _, group := range res.Data {
		groups = append(groups, group.Data)
	}

	return groups, resp, nil
}

// Get returns a group by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.groups.get
func (s *GroupsService) Get(ctx context.Context, id int) (*model.Group, *Response, error) {
	res := new(model.GroupsGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/groups/%d", id), nil, res)

	return res.Data, resp, err
}

// Add creates a new group.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.groups.post
func (s *GroupsService) Add(ctx context.Context, req *model.GroupsAddRequest) (*model.Group, *Response, error) {
	res := new(model.GroupsGetResponse)
	resp, err := s.client.Post(ctx, "/api/v2/groups", req, res)

	return res.Data, resp, err
}

// Edit updates a group.
//
// Request body:
//   - op: The operation to perform. Enum: replace, test.
//   - path: A JSON Pointer as defined in RFC 6901. Enum: "/name", "/description", "/parentId".
//   - value: The value to be used within the operations. The value must be one of string or integer.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.groups.patch
func (s *GroupsService) Edit(ctx context.Context, id int, req []*model.UpdateRequest) (*model.Group, *Response, error) {
	res := new(model.GroupsGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/groups/%d", id), req, res)

	return res.Data, resp, err
}

// Delete removes a group from the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.groups.delete
func (s *GroupsService) Delete(ctx context.Context, id int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/groups/%d", id), nil)
}

// List returns a list of managers.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.getMany
func (s *GroupsService) ListManagers(ctx context.Context, groupID string, opts *model.ManagerListOptions) ([]*model.Manager, *Response, error) {
	res := new(model.ManagerListResponse)
	url := fmt.Sprintf("/api/v2/groups/%s/managers", groupID)

	resp, err := s.client.Get(ctx, url, opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Manager, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns a manager by its identifier.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.get
func (s *GroupsService) GetManagers(ctx context.Context, fieldID string) (*model.Manager, *Response, error) {
	res := new(model.ManagerGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/fields/%s", fieldID), nil, res)

	return res.Data, resp, err
}

//	Update a manager.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.patch
func (s *GroupsService) EditManagers(ctx context.Context, groupID string, req []*model.UpdateRequest) ([]*model.Manager, *Response, error) {
	res := new(model.ManagerEditResponse)
	url := fmt.Sprintf("/api/v2/groups/%s/managers", groupID)
	resp, err := s.client.Patch(ctx, url, req, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Manager, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// List returns a list of teams.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Teams/operation/api.groups.teams.getMany
func (s *GroupsService) ListTeams(ctx context.Context, groupID string, opts *model.TeamsListOptions) ([]*model.GroupsTeams, *Response, error) {
	res := new(model.GroupsTeamsData)
	url := fmt.Sprintf("/api/v2/groups/%s/teams", groupID)

	resp, err := s.client.Get(ctx, url, opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.GroupsTeams, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns a teams by its identifier.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Teams/operation/api.groups.teams.get
func (s *GroupsService) GetTeams(ctx context.Context, groupID, teamID string) (*model.GroupsTeams, *Response, error) {
	res := new(model.TeamsGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/groups/%s/teams/%s", groupID, teamID), nil, res)

	return res.Data, resp, err
}

//	Update a teams.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Teams/operation/api.groups.teams.patch
func (s *GroupsService) EditTeams(ctx context.Context, groupID string, req []*model.UpdateRequest) ([]*model.GroupsTeams, *Response, error) {
	res := new(model.GroupsTeamsDataEdit)
	url := fmt.Sprintf("/api/v2/groups/%s/teams", groupID)
	resp, err := s.client.Patch(ctx, url, req, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.GroupsTeams, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}
