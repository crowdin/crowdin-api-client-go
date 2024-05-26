package crowdin

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeamsService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Translators Team",
				"totalMembers": 8,
				"webUrl": "https://example.crowdin.com/u/teams/1",
				"createdAt": "2023-09-23T09:04:29+00:00",
				"updatedAt": "2023-09-23T09:04:29+00:00"
			}
		}`)
	})

	team, resp, err := client.Teams.Get(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Team{
		ID:           2,
		Name:         "Translators Team",
		TotalMembers: 8,
		WebURL:       "https://example.crowdin.com/u/teams/1",
		CreatedAt:    "2023-09-23T09:04:29+00:00",
		UpdatedAt:    "2023-09-23T09:04:29+00:00",
	}
	assert.Equal(t, expected, team)
}

func TestTeamsService_List(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.TeamsListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.TeamsListOptions{},
		},
		{
			name: "with options",
			opts: &model.TeamsListOptions{
				OrderBy: "name",
				ListOptions: model.ListOptions{
					Limit:  10,
					Offset: 5,
				},
			},
			expectedQuery: "?limit=10&offset=5&orderBy=name",
		},
	}

	for _, tt := range tests {
		client, mux, teardown := setupClient()
		defer teardown()

		t.Run(tt.name, func(t *testing.T) {
			const path = "/api/v2/teams"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 1,
								"name": "Translators Team 1",
								"totalMembers": 8,
								"webUrl": "https://example.crowdin.com/u/teams/1",
								"createdAt": "2023-09-23T09:04:29+00:00",
								"updatedAt": "2023-09-23T09:04:29+00:00"
							}
						},
						{
							"data": {
								"id": 2,
								"name": "Translators Team 2",
								"totalMembers": 8,
								"webUrl": "https://example.crowdin.com/u/teams/1",
								"createdAt": "2023-09-23T09:04:29+00:00",
								"updatedAt": "2023-09-23T09:04:29+00:00"
							}
						}
					],
					"pagination": {
						"offset": 10,
						"limit": 25
					}
				}`)
			})

			teams, resp, err := client.Teams.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.Team{
				{
					ID:           1,
					Name:         "Translators Team 1",
					TotalMembers: 8,
					WebURL:       "https://example.crowdin.com/u/teams/1",
					CreatedAt:    "2023-09-23T09:04:29+00:00",
					UpdatedAt:    "2023-09-23T09:04:29+00:00",
				},
				{
					ID:           2,
					Name:         "Translators Team 2",
					TotalMembers: 8,
					WebURL:       "https://example.crowdin.com/u/teams/1",
					CreatedAt:    "2023-09-23T09:04:29+00:00",
					UpdatedAt:    "2023-09-23T09:04:29+00:00",
				},
			}
			assert.Equal(t, expected, teams)
			assert.Len(t, teams, 2)

			assert.Equal(t, 10, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestTeamsService_List_Error(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		w.WriteHeader(http.StatusInternalServerError)
	})

	teams, resp, err := client.Teams.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, teams)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.EqualError(t, err, "client: server returned 500 status code")
}

func TestTeamsService_Add(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"name":"Translators Team"}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Translators Team",
				"totalMembers": 8,
				"webUrl": "https://example.crowdin.com/u/teams/1",
				"createdAt": "2023-09-23T09:04:29+00:00",
				"updatedAt": "2023-09-23T09:04:29+00:00"
			}
		}`)
	})

	team, resp, err := client.Teams.Add(context.Background(), &model.TeamAddRequest{Name: "Translators Team"})
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := &model.Team{
		ID:           2,
		Name:         "Translators Team",
		TotalMembers: 8,
		WebURL:       "https://example.crowdin.com/u/teams/1",
		CreatedAt:    "2023-09-23T09:04:29+00:00",
		UpdatedAt:    "2023-09-23T09:04:29+00:00",
	}
	assert.Equal(t, expected, team)
}

func TestTeamsService_Add_requestValidation(t *testing.T) {
	tests := []struct {
		name string
		req  *model.TeamAddRequest
		err  string
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &model.TeamAddRequest{},
			err:  "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, tt.req.Validate(), tt.err)
		})
	}
}

func TestTeamsService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/name","value":"Translators Team"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Translators Team",
				"totalMembers": 8,
				"webUrl": "https://example.crowdin.com/u/teams/1",
				"createdAt": "2023-09-23T09:04:29+00:00",
				"updatedAt": "2023-09-23T09:04:29+00:00"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "Translators Team",
		},
	}
	team, resp, err := client.Teams.Edit(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Team{
		ID:           2,
		Name:         "Translators Team",
		TotalMembers: 8,
		WebURL:       "https://example.crowdin.com/u/teams/1",
		CreatedAt:    "2023-09-23T09:04:29+00:00",
		UpdatedAt:    "2023-09-23T09:04:29+00:00",
	}
	assert.Equal(t, expected, team)
}

func TestTeamsService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Teams.Delete(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTeamsService_ListMembers(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?limit=25&offset=1")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 1,
						"username": "john.doe",
						"firstName": "John",
						"lastName": "Doe",
						"avatarUrl": "",
						"addedAt": "2023-09-23T09:04:29+00:00"
					}
				}
			],
			"pagination": {
				"offset": 10,
				"limit": 25
			}
		}`)
	})

	members, resp, err := client.Teams.ListMembers(context.Background(), 1, &model.ListOptions{Limit: 25, Offset: 1})
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.TeamMember{
		{
			ID:        1,
			Username:  "john.doe",
			FirstName: "John",
			LastName:  "Doe",
			AvatarURL: "",
			AddedAt:   "2023-09-23T09:04:29+00:00",
		},
	}
	assert.Equal(t, expected, members)
	assert.Len(t, members, 1)
}

func TestTeamsService_ListMembers_error(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		w.WriteHeader(http.StatusInternalServerError)
	})

	teams, resp, err := client.Teams.ListMembers(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, teams)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.EqualError(t, err, "client: server returned 500 status code")
}

func TestTeamsService_AddMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"userIds":[1,2,5,10,99]}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"skipped": [
				{
					"data": {
						"id": 1,
						"username": "john.doe",
						"firstName": "John",
						"lastName": "Doe",
						"avatarUrl": "",
						"addedAt": "2023-09-23T09:04:29+00:00"
					}
				}
			],
			"added": [
				{
					"data": {
						"id": 1,
						"username": "john.doe",
						"firstName": "John",
						"lastName": "Doe",
						"avatarUrl": "",
						"addedAt": "2023-09-23T09:04:29+00:00"
					}
				}
			],
			"pagination": {
				"offset": 10,
				"limit": 25
			}
		}`)
	})

	req := &model.TeamMemberAddRequest{
		UserIDs: []int{1, 2, 5, 10, 99},
	}
	member, resp, err := client.Teams.AddMember(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := map[string][]*model.TeamMember{
		"skipped": {
			{
				ID:        1,
				Username:  "john.doe",
				FirstName: "John",
				LastName:  "Doe",
				AvatarURL: "",
				AddedAt:   "2023-09-23T09:04:29+00:00",
			},
		},
		"added": {
			{
				ID:        1,
				Username:  "john.doe",
				FirstName: "John",
				LastName:  "Doe",
				AvatarURL: "",
				AddedAt:   "2023-09-23T09:04:29+00:00",
			},
		},
	}
	assert.Equal(t, expected, member)
}

func TestTeamsService_AddMember_requestValidation(t *testing.T) {
	tests := []struct {
		name string
		req  *model.TeamMemberAddRequest
		err  string
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &model.TeamMemberAddRequest{},
			err:  "userIds is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, tt.req.Validate(), tt.err)
		})
	}
}

func TestTeamsService_AddMember_error(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members"
	mux.HandleFunc(path, func(_ http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
	})

	teams, resp, err := client.Teams.AddMember(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, teams)
	assert.Nil(t, resp)
	assert.EqualError(t, err, "request cannot be nil")
}

func TestTeamsService_DeleteMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Teams.DeleteMember(context.Background(), 1, 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTeamsService_DeleteAllMembers(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/teams/1/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Teams.DeleteMembers(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTeamsService_AddToProject(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/teams"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"teamId": 1,
			"managerAccess": false,
			"developerAccess": true,
			"roles": [
				{
					"name": "translator",
					"permissions": {
						"allLanguages": false,
						"languagesAccess": {
							"uk": {
								"allContent": false,
								"workflowStepIds": [882]
							},
							"it": {
								"allContent": true
							}
						}
					}
				},
				{
					"name": "proofreader",
					"permissions": {
						"allLanguages": true
					}
				},
				{
					"name": "language_coordinator",
					"permissions": {
						"allLanguages": false,
						"languagesAccess": {
							"uk": {
								"allContent": true
							},
							"it": {
								"allContent": true
							}
						}
					}
				}
			]
		}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"skipped": {
				"id": 1,
				"hasManagerAccess": false,
				"hasDeveloperAccess": true,
				"hasAccessToAllWorkflowSteps": false,
				"permissions": {
					"it": {
						"workflowStepIds": [313]
					}
				},
				"roles": [
					{
						"name": "translator",
						"permissions": {
							"allLanguages": false,
							"languagesAccess": {
								"uk": {
									"allContent": false,
									"workflowStepIds": [882]
								},
								"it": {
									"allContent": true
								}
							}
						}
					},
					{
						"name": "proofreader",
						"permissions": {
							"allLanguages": true,
							"languagesAccess": {}
						}
					},
					{
						"name": "language_coordinator",
						"permissions": {
							"allLanguages": false,
							"languagesAccess": {
								"uk": {
									"allContent": true
								},
								"it": {
									"allContent": true
								}
							}
						}
					}
				]
			},
			"added": {}
		}`)
	})

	req := &model.ProjectTeamAddRequest{
		TeamID:          1,
		ManagerAccess:   ToPtr(false),
		DeveloperAccess: ToPtr(true),
		Roles: []*model.TranslatorRole{
			{
				Name: "translator",
				Permissions: &model.RolePermissions{
					AllLanguages: ToPtr(false),
					LanguagesAccess: map[string]*model.LanguageAccess{
						"uk": {
							AllContent:      ToPtr(false),
							WorkflowStepIDs: []int{882},
						},
						"it": {
							AllContent: ToPtr(true),
						},
					},
				},
			},
			{
				Name: "proofreader",
				Permissions: &model.RolePermissions{
					AllLanguages:    ToPtr(true),
					LanguagesAccess: map[string]*model.LanguageAccess{},
				},
			},
			{
				Name: "language_coordinator",
				Permissions: &model.RolePermissions{
					AllLanguages: ToPtr(false),
					LanguagesAccess: map[string]*model.LanguageAccess{
						"uk": {
							AllContent: ToPtr(true),
						},
						"it": {
							AllContent: ToPtr(true),
						},
					},
				},
			},
		},
	}
	team, resp, err := client.Teams.AddToProject(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := map[string]*model.ProjectTeam{
		"skipped": {
			ID:                          1,
			HasManagerAccess:            false,
			HasDeveloperAccess:          true,
			HasAccessToAllWorkflowSteps: false,
			Permissions: map[string]any{
				"it": map[string]any{
					"workflowStepIds": []any{float64(313)},
				},
			},
			Roles: []*model.TranslatorRole{
				{
					Name: "translator",
					Permissions: &model.RolePermissions{
						AllLanguages: ToPtr(false),
						LanguagesAccess: map[string]*model.LanguageAccess{
							"uk": {
								AllContent:      ToPtr(false),
								WorkflowStepIDs: []int{882},
							},
							"it": {
								AllContent: ToPtr(true),
							},
						},
					},
				},
				{
					Name: "proofreader",
					Permissions: &model.RolePermissions{
						AllLanguages:    ToPtr(true),
						LanguagesAccess: map[string]*model.LanguageAccess{},
					},
				},
				{
					Name: "language_coordinator",
					Permissions: &model.RolePermissions{
						AllLanguages: ToPtr(false),
						LanguagesAccess: map[string]*model.LanguageAccess{
							"uk": {
								AllContent: ToPtr(true),
							},
							"it": {
								AllContent: ToPtr(true),
							},
						},
					},
				},
			},
		},
		"added": {},
	}
	assert.Equal(t, expected, team)
}

func TestTeamsService_AddToProject_requestValidation(t *testing.T) {
	tests := []struct {
		name string
		req  *model.ProjectTeamAddRequest
		err  string
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &model.ProjectTeamAddRequest{},
			err:  "teamId is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, tt.req.Validate(), tt.err)
		})
	}
}

func TestTeamsService_AddToProject_error(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/teams"
	mux.HandleFunc(path, func(_ http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
	})

	teams, resp, err := client.Teams.AddToProject(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.ErrorIs(t, err, model.ErrNilRequest)
	assert.Nil(t, teams)
	assert.Nil(t, resp)
}
