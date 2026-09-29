package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// CustomSpellcheckersService provides access to the Custom Spellcheckers API.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/Custom-Spellcheckers
type CustomSpellcheckersService struct {
	client *Client
}

// List returns a list of the custom spellcheckers of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-spellcheckers.getMany
func (s *CustomSpellcheckersService) List(ctx context.Context, opts *model.ListOptions) (
	[]*model.CustomSpellchecker, *Response, error,
) {
	res := new(model.CustomSpellcheckersListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/custom-spellcheckers", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.CustomSpellchecker, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns a custom spellchecker by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.custom-spellcheckers.get
func (s *CustomSpellcheckersService) Get(ctx context.Context, customSpellcheckerID int) (
	*model.CustomSpellchecker, *Response, error,
) {
	res := new(model.CustomSpellcheckerResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/custom-spellcheckers/%d", customSpellcheckerID), nil, res)

	return res.Data, resp, err
}
