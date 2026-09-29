package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// ExternalQAChecksService provides access to the External QA Checks API.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/External-QA-Checks
type ExternalQAChecksService struct {
	client *Client
}

// List returns a list of the external QA checks of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.external-qa-checks.getMany
func (s *ExternalQAChecksService) List(ctx context.Context, opts *model.ListOptions) (
	[]*model.ExternalQACheck, *Response, error,
) {
	res := new(model.ExternalQAChecksListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/external-qa-checks", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ExternalQACheck, 0, len(res.Data))
	for _, item := range res.Data {
		list = append(list, item.Data)
	}

	return list, resp, nil
}

// Get returns an external QA check by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.external-qa-checks.get
func (s *ExternalQAChecksService) Get(ctx context.Context, externalQACheckID int) (
	*model.ExternalQACheck, *Response, error,
) {
	res := new(model.ExternalQACheckResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/external-qa-checks/%d", externalQACheckID), nil, res)

	return res.Data, resp, err
}
