package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// https://developer.crowdin.com/api/v2/#tag/String-Corrections
type StringCorrectionsService struct {
	client *Client
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.getMany
func (s *StringCorrectionsService) ListCorrections(ctx context.Context, projectID int, opts *model.StringCorrectionsListOptions) (
	[]*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.StringCorrection, 0, len(res.Data))
	for _, correction := range res.Data {
		list = append(list, correction.Data)
	}

	return list, resp, nil
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.get
func (s *StringCorrectionsService) GetCorrection(ctx context.Context, projectID, correctionID int, opts *model.StringCorrectionGetOptions) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections/%d", projectID, correctionID), opts, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.post
func (s *StringCorrectionsService) AddCorrection(ctx context.Context, projectID int, req *model.StringCorrectionAddRequest) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionGetResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections", projectID), req, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.deleteMany
func (s *StringCorrectionsService) DeleteCorrections(ctx context.Context, projectID int, opts *model.StringCorrectionsDeleteOptions) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections", projectID), opts)
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.put
func (s *StringCorrectionsService) RestoreCorrection(ctx context.Context, projectID, correctionID int) (
	*model.StringCorrection, *Response, error,
) {
	res := new(model.StringCorrectionGetResponse)
	resp, err := s.client.Put(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections/%d", projectID, correctionID), nil, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.projects.strings.corrections.delete
func (s *StringCorrectionsService) DeleteCorrection(ctx context.Context, projectID, correctionID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/corrections/%d", projectID, correctionID), nil)
}

