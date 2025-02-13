package crowdin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupService_List(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups")
		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 1,
						"name": "KB materials",
						"description": "KB localization materials",
						"parentId": 2,
						"organizationId": 200000299,
						"userId": 6,
						"subgroupsCount": 0,
						"projectsCount": 1,
						"webUrl": "https://example.crowdin.com/u/groups/123",
						"createdAt": "2023-09-20T11:11:05+00:00",
						"updatedAt": "2023-09-20T12:22:20+00:00"
					}
				}
			],
			"pagination": {
				"offset": 0,
			  	"limit": 25
			}
		}`)
	})

	groups, resp, err := client.Groups.List(context.Background(), nil)
	if err != nil {
		t.Errorf("Groups.List returned error: %v", err)
	}

	want := []*model.Group{
		{
			ID:             1,
			Name:           "KB materials",
			Description:    "KB localization materials",
			ParentID:       2,
			OrganizationID: 200000299,
			UserID:         6,
			SubgroupsCount: 0,
			ProjectsCount:  1,
			WebURL:         "https://example.crowdin.com/u/groups/123",
			CreatedAt:      "2023-09-20T11:11:05+00:00",
			UpdatedAt:      "2023-09-20T12:22:20+00:00",
		},
	}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Groups.List returned %+v, want %+v", groups, want)
	}

	expectedPagination := model.Pagination{Offset: 0, Limit: 25}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("Groups.List returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestGroupService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Groups.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestGroupService_ListWithQueryParams(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups?limit=100&offset=1&parentId=1")
		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 1,
						"name": "KB materials",
						"description": "KB localization materials",
						"parentId": 2,
						"organizationId": 200000299,
						"userId": 6,
						"subgroupsCount": 0,
						"projectsCount": 1,
						"webUrl": "https://example.crowdin.com/u/groups/123",
						"createdAt": "2023-09-20T11:11:05+00:00",
						"updatedAt": "2023-09-20T12:22:20+00:00"
					}
				}
			],
			"pagination": {
				"offset": 0,
			  	"limit": 25
			}
		}`)
	})

	groups, _, err := client.Groups.List(
		context.Background(),
		&model.GroupsListOptions{ParentID: 1, ListOptions: model.ListOptions{Limit: 100, Offset: 1}},
	)
	if err != nil {
		t.Errorf("Groups.List returned error: %v", err)
	}

	want := []*model.Group{
		{
			ID:             1,
			Name:           "KB materials",
			Description:    "KB localization materials",
			ParentID:       2,
			OrganizationID: 200000299,
			UserID:         6,
			SubgroupsCount: 0,
			ProjectsCount:  1,
			WebURL:         "https://example.crowdin.com/u/groups/123",
			CreatedAt:      "2023-09-20T11:11:05+00:00",
			UpdatedAt:      "2023-09-20T12:22:20+00:00",
		},
	}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Groups.List returned %+v, want %+v", groups, want)
	}
}

func TestGroupService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups/1")
		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"name": "KB materials",
				"description": "KB localization materials",
				"parentId": 2,
				"organizationId": 200000299,
				"userId": 6,
				"subgroupsCount": 0,
				"projectsCount": 1,
				"webUrl": "https://example.crowdin.com/u/groups/123",
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	group, _, err := client.Groups.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("Groups.Get returned error: %v", err)
	}

	want := &model.Group{
		ID:             1,
		Name:           "KB materials",
		Description:    "KB localization materials",
		ParentID:       2,
		OrganizationID: 200000299,
		UserID:         6,
		SubgroupsCount: 0,
		ProjectsCount:  1,
		WebURL:         "https://example.crowdin.com/u/groups/123",
		CreatedAt:      "2023-09-20T11:11:05+00:00",
		UpdatedAt:      "2023-09-20T12:22:20+00:00",
	}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.Get returned %+v, want %+v", group, want)
	}
}

func TestGroupService_GetByIDNotFound(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups/1")
		http.Error(w, `{"error": {"code": 404, "message": "Group Not Found"}}`, http.StatusNotFound)
	})

	group, resp, err := client.Groups.Get(context.Background(), 1)
	if err == nil {
		t.Error("Groups.Get expected an error, got nil")
	}
	if group != nil {
		t.Errorf("Groups.Get expected nil, got %+v", group)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Groups.Get expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	var e *model.ErrorResponse
	if !errors.As(err, &e) {
		t.Errorf("Groups.Get expected type *model.ErrorResponse, got %T", err)
	}
}

func TestGroupService_Add(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, "/api/v2/groups")
		testBody(t, r, `{"name":"KB materials","parentId":2,"description":"KB localization materials"}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"name": "KB materials",
				"description": "KB localization materials",
				"parentId": 2,
				"organizationId": 200000299,
				"userId": 6,
				"subgroupsCount": 0,
				"projectsCount": 1,
				"webUrl": "https://example.crowdin.com/u/groups/123",
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	group, _, err := client.Groups.Add(
		context.Background(),
		&model.GroupsAddRequest{Name: "KB materials", ParentID: 2, Description: "KB localization materials"},
	)
	if err != nil {
		t.Errorf("Groups.Add returned error: %v", err)
	}

	want := &model.Group{
		ID:             1,
		Name:           "KB materials",
		Description:    "KB localization materials",
		ParentID:       2,
		OrganizationID: 200000299,
		UserID:         6,
		SubgroupsCount: 0,
		ProjectsCount:  1,
		WebURL:         "https://example.crowdin.com/u/groups/123",
		CreatedAt:      "2023-09-20T11:11:05+00:00",
		UpdatedAt:      "2023-09-20T12:22:20+00:00",
	}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.Add returned %+v, want %+v", group, want)
	}
}

func TestGroupService_AddWithEmptyRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Groups.Add(context.Background(), nil)
	if !errors.Is(err, model.ErrNilRequest) {
		t.Errorf("Groups.Add expected error: %v, got: %v", model.ErrNilRequest, err)
	}
}

func TestGroupssService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, "/api/v2/groups/1")

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"name": "Test materials",
				"description": "KB localization materials",
				"parentId": 2,
				"organizationId": 200000299,
				"userId": 6,
				"subgroupsCount": 0,
				"projectsCount": 1,
				"webUrl": "https://example.crowdin.com/u/groups/123",
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "Test materials",
		},
	}
	group, _, err := client.Groups.Edit(context.Background(), 1, req)
	if err != nil {
		t.Errorf("Groups.Edit returned error: %v", err)
	}

	want := "Test materials"
	if group.Name != want {
		t.Errorf("Groups.Edit returned %+v, want %+v", group.Name, want)
	}
}

func TestGroupService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testURL(t, r, "/api/v2/groups/2")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Groups.Delete(context.Background(), 2)
	if err != nil {
		t.Errorf("Groups.Delete returned error: %v", err)
	}
}

func TestManagerService_List(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/2/managers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups/2/managers")
		fmt.Fprint(w, `{
			"data": [
			  {
				"data": {
				  "id": 27,
				  "user": {
					"id": 12,
					"username": "john_smith",
					"email": "jsmith@example.com",
					"firstName": "John",
					"lastName": "Smith",
					"status": "active",
					"avatarUrl": "",
					"createdAt": "2019-07-11T07:40:22+00:00",
					"lastSeen": "2019-10-23T11:44:02+00:00",
					"twoFactor": "enabled",
					"isAdmin": true,
					"timezone": "Europe/Kyiv",
					"fields": {
					  "some-field-1": "some value 1",
					  "some-field-2": 12,
					  "some-field-3": true,
					  "some-field-4": []
					}
				  },
				  "teams": [
					{
					  "id": 2,
					  "name": "Translators Team",
					  "totalMembers": 8,
					  "webUrl": "https://example.crowdin.com/u/teams/1",
					  "createdAt": "2019-09-23T09:04:29+00:00",
					  "updatedAt": "2019-09-23T09:04:29+00:00"
					}
				  ]
				}
			  }
			],
			"pagination": {
			  "offset": 0,
			  "limit": 25
			}
		  }`)
	})

	managers, resp, err := client.Groups.ListManagers(context.Background(), "2", nil)
	if err != nil {
		t.Errorf("Managers.List returned error: %v", err)
	}

	firstname := "John"
	lastname := "Smith"
	status := "active"
	isAdmin := true

	want := []*model.Manager{
		{
			ID: 27,
			User: model.User{
				ID:        12,
				Username:  "john_smith",
				Email:     "jsmith@example.com",
				FirstName: &firstname,
				LastName:  &lastname,
				Status:    &status,
				AvatarURL: "",
				CreatedAt: "2019-07-11T07:40:22+00:00",
				LastSeen:  "2019-10-23T11:44:02+00:00",
				TwoFactor: "enabled",
				IsAdmin:   &isAdmin,
				Timezone:  "Europe/Kyiv",
				Fields: map[string]interface{}{
					"some-field-1": "some value 1",
					"some-field-2": 12,
					"some-field-3": true,
					"some-field-4": []interface{}{},
				},
			},
			Teams: []model.Team{
				{
					ID:           2,
					Name:         "Translators Team",
					TotalMembers: 8,
					WebURL:       "https://example.crowdin.com/u/teams/1",
					CreatedAt:    "2019-09-23T09:04:29+00:00",
					UpdatedAt:    "2019-09-23T09:04:29+00:00",
				},
			},
		},
	}

	if managers[0].ID != want[0].ID {
		t.Errorf("Managers.List returned ID %v, want %v", managers[0].ID, want[0].ID)
	}

	if *managers[0].User.FirstName != *want[0].User.FirstName {
		t.Errorf("Managers.List returned FirstName %v, want %v", *managers[0].User.FirstName, *want[0].User.FirstName)
	}

	expectedPagination := model.Pagination{Offset: 0, Limit: 25}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("Managers.List returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestManagersService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/2/managers", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Groups.ListManagers(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestManagersService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1/managers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, "/api/v2/groups/1/managers")

		fmt.Fprint(w, `{
			"data": [
			  {
				"data": {
				  "id": 27,
				  "user": {
					"id": 18,
					"username": "john_smith",
					"email": "jsmith@example.com",
					"firstName": "John",
					"lastName": "Smith",
					"status": "active",
					"avatarUrl": "",
					"createdAt": "2019-07-11T07:40:22+00:00",
					"lastSeen": "2019-10-23T11:44:02+00:00",
					"twoFactor": "enabled",
					"isAdmin": true,
					"timezone": "Europe/Kyiv",
					"fields": {
					  "some-field-1": "some value 1",
					  "some-field-2": 12,
					  "some-field-3": true,
					  "some-field-4": []
					}
				  },
				  "teams": [
					{
					  "id": 2,
					  "name": "Translators Team",
					  "totalMembers": 8,
					  "webUrl": "https://example.crowdin.com/u/teams/1",
					  "createdAt": "2019-09-23T09:04:29+00:00",
					  "updatedAt": "2019-09-23T09:04:29+00:00"
					}
				  ]
				}
			  }
			]
		  }`)
	})

	req := []*model.UpdateRequest{
		{
			Op:   "add",
			Path: "/userId",
			Value: `{
				"userId": 18,
			}`,
		},
	}
	managers, _, err := client.Groups.EditManagers(context.Background(), "1", req)
	if err != nil {
		t.Errorf("Managers.Edit returned error: %v", err)
	}

	want := 18

	if managers[0].User.ID != want {
		t.Errorf("Managers.Edit returned %+v, want %+v", managers[0].User.ID, want)
	}
}

func TestManagersService_Edit_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1/managers", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Groups.EditManagers(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestManagersService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/fields/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/fields/1")
		fmt.Fprint(w, `{
				"data": {
				  "id": 27,
				  "user": {
					"id": 12,
					"username": "john_smith",
					"email": "jsmith@example.com",
					"firstName": "John",
					"lastName": "Smith",
					"status": "active",
					"avatarUrl": "",
					"createdAt": "2019-07-11T07:40:22+00:00",
					"lastSeen": "2019-10-23T11:44:02+00:00",
					"twoFactor": "enabled",
					"isAdmin": true,
					"timezone": "Europe/Kyiv",
					"fields": {
					  "some-field-1": "some value 1",
					  "some-field-2": 12,
					  "some-field-3": true,
					  "some-field-4": []
					}
				  },
				  "teams": [
					{
					  "id": 2,
					  "name": "Translators Team",
					  "totalMembers": 8,
					  "webUrl": "https://example.crowdin.com/u/teams/1",
					  "createdAt": "2019-09-23T09:04:29+00:00",
					  "updatedAt": "2019-09-23T09:04:29+00:00"
					}
				  ]
				}
		  }`)
	})

	managers, _, err := client.Groups.GetManagers(context.Background(), "1")
	if err != nil {
		t.Errorf("Managers.Get returned error: %v", err)
	}

	firstname := "John"
	lastname := "Smith"
	status := "active"
	isAdmin := true

	want := &model.Manager{
		ID: 27,
		User: model.User{
			ID:        12,
			Username:  "john_smith",
			Email:     "jsmith@example.com",
			FirstName: &firstname,
			LastName:  &lastname,
			Status:    &status,
			AvatarURL: "",
			CreatedAt: "2019-07-11T07:40:22+00:00",
			LastSeen:  "2019-10-23T11:44:02+00:00",
			TwoFactor: "enabled",
			IsAdmin:   &isAdmin,
			Timezone:  "Europe/Kyiv",
			Fields: map[string]interface{}{
				"some-field-1": "some value 1",
				"some-field-2": 12,
				"some-field-3": true,
				"some-field-4": []interface{}{},
			},
		},
		Teams: []model.Team{
			{
				ID:           2,
				Name:         "Translators Team",
				TotalMembers: 8,
				WebURL:       "https://example.crowdin.com/u/teams/1",
				CreatedAt:    "2019-09-23T09:04:29+00:00",
				UpdatedAt:    "2019-09-23T09:04:29+00:00",
			},
		},
	}

	if managers.ID != want.ID {
		t.Errorf("Managers.Get returned %+v, want %+v", managers, want)
	}
}

func TestGroupsTeamsService_List(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/2/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups/2/teams")
		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
							"id": 27,
							"user": {
							"id": 2,
							"name": "Translators Team",
							"totalMembers": 8,
							"webUrl": "https://example.crowdin.com/u/teams/1",
							"createdAt": "2019-09-23T09:04:29+00:00",
							"updatedAt": "2019-09-23T09:04:29+00:00"
						}
					}
				}
			],
			"pagination": {
				"offset": 0,
				"limit": 25
			}
		  }
		`)
	})

	teams, resp, err := client.Groups.ListTeams(context.Background(), "2", nil)
	if err != nil {
		t.Errorf("Group.Teams.List returned error: %v", err.Error())
	}

	want := []*model.TeamsGetResponse{
		{
			Data: &model.GroupsTeams{
				ID: 27,
				User: &model.Team{
					ID:           2,
					Name:         "Translators Team",
					TotalMembers: 8,
					WebURL:       "https://example.crowdin.com/u/teams/1",
					CreatedAt:    "2019-09-23T09:04:29+00:00",
					UpdatedAt:    "2019-09-23T09:04:29+00:00",
				},
			},
		},
	}

	if teams[0].User.Name != want[0].Data.User.Name {
		t.Errorf("Managers.List returned ID %v, want %v", teams[0].ID, want[0].Data.ID)
	}

	expectedPagination := model.Pagination{Offset: 0, Limit: 25}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("Group.Teams.List returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestGroupsTeamsService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/2/teams", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Groups.ListTeams(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestGroupTeamsService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/groups/1/teams/1")
		fmt.Fprint(w, `{
			"data": {
				"id": 27,
				"user": {
				"id": 2,
				"name": "Translators Team",
				"totalMembers": 8,
				"webUrl": "https://example.crowdin.com/u/teams/1",
				"createdAt": "2019-09-23T09:04:29+00:00",
				"updatedAt": "2019-09-23T09:04:29+00:00"
				}
			}
		}`)
	})

	teams, _, err := client.Groups.GetTeams(context.Background(), "1", "1")
	if err != nil {
		t.Errorf("Managers.Get returned error: %v", err)
	}

	want := &model.TeamsGetResponse{
		Data: &model.GroupsTeams{
			ID: 18,
			User: &model.Team{
				ID:           2,
				Name:         "Translators Team",
				TotalMembers: 8,
				WebURL:       "https://example.crowdin.com/u/teams/1",
				CreatedAt:    "2019-09-23T09:04:29+00:00",
				UpdatedAt:    "2019-09-23T09:04:29+00:00",
			},
		},
	}

	if teams.User.ID != want.Data.User.ID {
		t.Errorf("Managers.Get returned %+v, want %+v", teams, want)
	}
}

func TestGroupTeamsService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, "/api/v2/groups/1/teams")

		fmt.Fprint(w, `{
			"data": [
				{
				"data": {
					"id": 2,
					"user": {
					"id": 18,
					"name": "Translators Team",
					"totalMembers": 8,
					"webUrl": "https://example.crowdin.com/u/teams/1",
					"createdAt": "2019-09-23T09:04:29+00:00",
					"updatedAt": "2019-09-23T09:04:29+00:00"
					}
				}
				}
			]
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:   "add",
			Path: "/id",
			Value: `{
				"id": 18,
			}`,
		},
	}
	teams, _, err := client.Groups.EditTeams(context.Background(), "1", req)
	if err != nil {
		t.Errorf("Groups.Teams.Edit returned error: %v", err)
	}

	want := 18

	if teams[0].User.ID != want {
		t.Errorf("Managers.Edit returned %+v, want %+v", teams[0].User.ID, want)
	}
}

func TestGroupsTeamsService_Edit_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/groups/1/teams", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Groups.EditTeams(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, res)
}
