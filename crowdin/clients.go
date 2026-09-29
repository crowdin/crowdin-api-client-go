package crowdin

import (
	"context"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// ClientsService provides access to the Clients API.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/Clients
type ClientsService struct {
	client *Client
}

// List returns a list of clients of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.clients.getMany
func (s *ClientsService) List(ctx context.Context, opts *model.ListOptions) ([]*model.Client, *Response, error) {
	res := new(model.ClientsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/clients", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Client, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}
