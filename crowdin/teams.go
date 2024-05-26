package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// TeamsService provides access to the organizaion teams API.
//
// CrowdIn API docs: https://developer.crowdin.com/enterprise/api/v2/#tag/Teams
type TeamsService struct {
	client *Client
}

// List returns a list of teams in the organization.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.getMany
func (s *TeamsService) List(ctx context.Context, opts *model.TeamsListOptions) ([]*model.Team, *Response, error) {
	res := new(model.TeamsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/teams", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Team, 0, len(res.Data))
	for _, team := range res.Data {
		list = append(list, team.Data)
	}

	return list, resp, nil
}

// Get returns a team by its identifier.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.get
func (s *TeamsService) Get(ctx context.Context, teamID int) (*model.Team, *Response, error) {
	res := new(model.TeamResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/teams/%d", teamID), nil, res)

	return res.Data, resp, err
}

// Add creates a new team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.post
func (s *TeamsService) Add(ctx context.Context, req *model.TeamAddRequest) (*model.Team, *Response, error) {
	res := new(model.TeamResponse)
	resp, err := s.client.Post(ctx, "/api/v2/teams", req, res)

	return res.Data, resp, err
}

// Edit updates a team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.patch
func (s *TeamsService) Edit(ctx context.Context, teamID int, req []*model.UpdateRequest) (*model.Team, *Response, error) {
	res := new(model.TeamResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/teams/%d", teamID), req, res)

	return res.Data, resp, err
}

// Delete removes a team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.delete
func (s *TeamsService) Delete(ctx context.Context, teamID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/teams/%d", teamID))
}

// ListMembers returns a list of team members.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.members.getMany
func (s *TeamsService) ListMembers(ctx context.Context, teamID int, opts *model.ListOptions) (
	[]*model.TeamMember, *Response, error,
) {
	res := new(model.TeamMembersListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/teams/%d/members", teamID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.TeamMember, 0, len(res.Data))
	for _, member := range res.Data {
		list = append(list, member.Data)
	}

	return list, resp, nil
}

// AddMember adds a new member to the team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.members.post
func (s *TeamsService) AddMember(ctx context.Context, teamID int, req *model.TeamMemberAddRequest) (
	map[string][]*model.TeamMember, *Response, error,
) {
	res := new(model.TeamMemberAddResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/teams/%d/members", teamID), req, res)
	if err != nil {
		return nil, resp, err
	}

	skipped := make([]*model.TeamMember, 0, len(res.Skipped))
	for _, member := range res.Skipped {
		skipped = append(skipped, member.Data)
	}

	added := make([]*model.TeamMember, 0, len(res.Added))
	for _, member := range res.Added {
		added = append(added, member.Data)
	}

	return map[string][]*model.TeamMember{
		"skipped": skipped,
		"added":   added,
	}, resp, err
}

// DeleteMember removes a member from the team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.members.delete
func (s *TeamsService) DeleteMember(ctx context.Context, teamID, memberID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/teams/%d/members/%d", teamID, memberID))
}

// DeleteMembers deletes all members from the team.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.teams.members.deleteMany
func (s *TeamsService) DeleteMembers(ctx context.Context, teamID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/teams/%d/members", teamID))
}

// AddToProject adds a team to the project.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.teams.post
func (s *TeamsService) AddToProject(ctx context.Context, projectID int, req *model.ProjectTeamAddRequest) (
	map[string]*model.ProjectTeam, *Response, error,
) {
	res := new(model.ProjectTeamAddResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/teams", projectID), req, res)
	if err != nil {
		return nil, resp, err
	}

	return map[string]*model.ProjectTeam{
		"skipped": res.Skipped,
		"added":   res.Added,
	}, resp, nil
}
