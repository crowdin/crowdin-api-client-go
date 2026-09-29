package crowdin

import (
	"context"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// OrganizationService provides access to the Organization API.
//
// Crowdin API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/Organization
type OrganizationService struct {
	client *Client
}

// GetInfo returns the information about the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.organization.get
func (s *OrganizationService) GetInfo(ctx context.Context) (*model.OrganizationInfo, *Response, error) {
	res := new(model.OrganizationInfoResponse)
	resp, err := s.client.Get(ctx, "/api/v2/organization", nil, res)

	return res.Data, resp, err
}

// GetAuthSettings returns the authentication settings of the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.organization.auth-settings.get
func (s *OrganizationService) GetAuthSettings(ctx context.Context) (*model.OrganizationAuthSettings, *Response, error) {
	res := new(model.OrganizationAuthSettingsResponse)
	resp, err := s.client.Get(ctx, "/api/v2/organization/auth-settings", nil, res)

	return res.Data, resp, err
}
