package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// AdvisorsService provides access to the Advisors API methods.
// Advisors inspect a project and produce insights about its
// localization health.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Advisors
type AdvisorsService struct {
	client *Client
}

// CreateCheck triggers an asynchronous re-check of advisor inspectors
// for the project. If req is nil or empty, every inspector is re-checked.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.advisors.checks.post
func (s *AdvisorsService) CreateCheck(ctx context.Context, projectID int, req *model.AdvisorCheckCreateRequest) (
	*model.AdvisorCheck, *Response, error,
) {
	if req == nil {
		req = &model.AdvisorCheckCreateRequest{}
	}

	res := new(model.AdvisorCheckResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/advisors/checks", projectID), req, res)

	return res.Data, resp, err
}

// GetCheck returns the status of an advisor check job.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.advisors.checks.get
func (s *AdvisorsService) GetCheck(ctx context.Context, projectID int, checkID string) (
	*model.AdvisorCheck, *Response, error,
) {
	res := new(model.AdvisorCheckResponse)
	path := fmt.Sprintf("/api/v2/projects/%d/advisors/checks/%s", projectID, checkID)
	resp, err := s.client.Get(ctx, path, nil, res)

	return res.Data, resp, err
}

// ListInsights returns a list of advisor insights of the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.advisors.insights.getMany
func (s *AdvisorsService) ListInsights(ctx context.Context, projectID int, opts *model.AdvisorInsightsListOptions) (
	[]*model.AdvisorInsight, *Response, error,
) {
	res := new(model.AdvisorInsightsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/advisors/insights", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.AdvisorInsight, 0, len(res.Data))
	for _, insight := range res.Data {
		list = append(list, insight.Data)
	}

	return list, resp, err
}

// EditInsight updates an advisor insight by its identifier.
//
// Request body:
// - op - operation to perform. Enum: replace, test.
// - path (json-pointer) - path to the field to update. Enum: "/isDismissed".
// - value (bool) - new value for the field.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.advisors.insights.patch
func (s *AdvisorsService) EditInsight(ctx context.Context, projectID, insightID int, req []*model.UpdateRequest) (
	*model.AdvisorInsight, *Response, error,
) {
	res := new(model.AdvisorInsightResponse)
	path := fmt.Sprintf("/api/v2/projects/%d/advisors/insights/%d", projectID, insightID)
	resp, err := s.client.Patch(ctx, path, req, res)

	return res.Data, resp, err
}

// CreateOrUpdateApplicationInsight publishes the result of an installed
// application's `advisor-inspector` module check. Subsequent calls for the
// same string overwrite the previous insight from this module.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.applications.modules.advisors.insights.put
func (s *AdvisorsService) CreateOrUpdateApplicationInsight(
	ctx context.Context,
	projectID int,
	applicationIdentifier, moduleKey string,
	req *model.ApplicationAdvisorInsightRequest,
) (*Response, error) {
	path := fmt.Sprintf("/api/v2/projects/%d/applications/%s/modules/%s/advisors/insights",
		projectID, applicationIdentifier, moduleKey)

	return s.client.Put(ctx, path, req, nil)
}
