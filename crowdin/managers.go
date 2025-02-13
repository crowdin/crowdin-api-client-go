package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// ManagersService provides access to the Managers API.
//
// Crowdin API docs: https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users
type ManagersService struct {
	client *Client
}

// List returns a list of managers.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.getMany
func (s *ManagersService) List(ctx context.Context, groupID string, opts *model.FieldsListOptions) ([]*model.Field, *Response, error) {
	res := new(model.FieldsListResponse)
	url := fmt.Sprintf("/api/v2/groups/%s/managers", groupID)

	resp, err := s.client.Get(ctx, url, opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Field, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns a manager by its identifier.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.get
func (s *ManagersService) Get(ctx context.Context, fieldID int) (*model.Field, *Response, error) {
	res := new(model.FieldResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/fields/%d", fieldID), nil, res)

	return res.Data, resp, err
}

//	Update a manager.
//
// https://support.crowdin.com/developer/enterprise/api/v2/#tag/Users/operation/api.groups.managers.patch
func (s *FieldsService) Update(ctx context.Context, req *model.FieldAddRequest) (*model.Field, *Response, error) {
	res := new(model.FieldResponse)
	resp, err := s.client.Post(ctx, "/api/v2/fields", req, res)

	return res.Data, resp, err
}
