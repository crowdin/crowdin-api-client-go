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

func TestUsersService_GetProjectMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/members/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 12,
				"username": "john_smith",
				"fullName": "John Smith",
				"role": "translator",
				"permissions": {
					"uk": "translator",
					"it": "proofreader",
					"en": "denied"
				},
				"roles": [
					{
						"name": "translator",
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
				],
				"avatarUrl": "",
				"joinedAt": "2023-07-11T07:40:22+00:00",
				"timezone": "Europe/Kyiv"
			}
		}`)
	})

	member, resp, err := client.Users.GetProjectMember(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectMember{
		ID:       12,
		Username: "john_smith",
		FullName: ToPtr("John Smith"),
		Role:     ToPtr("translator"),
		Permissions: map[string]any{
			"uk": "translator",
			"it": "proofreader",
			"en": "denied",
		},
		Roles: []*model.TranslatorRole{
			{
				Name: "translator",
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
		AvatarURL: ToPtr(""),
		JoinedAt:  ToPtr("2023-07-11T07:40:22+00:00"),
		Timezone:  ToPtr("Europe/Kyiv"),
	}
	assert.Equal(t, expected, member)
	assert.Nil(t, member.FirstName)
	assert.Nil(t, member.LastName)
	assert.Nil(t, member.IsManager)
	assert.Nil(t, member.IsDeveloper)
	assert.Nil(t, member.ManagerOfGroup)
	assert.Nil(t, member.AccessToAllWorkflowSteps)
	assert.Nil(t, member.GivenAccessAt)
}

func TestUsersService_GetProjectMember_EnterpriseAPI(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/members/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 12,
				"username": "john_smith",
				"firstName": "John",
				"lastName": "Smith",
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
				],
				"isManager": false,
				"isDeveloper": false,
				"managerOfGroup": {
					"id": 1,
					"name": "KB materials"
				},
				"accessToAllWorkflowSteps": false,
				"permissions": {
					"it": {
						"workflowStepIds": [313]
					}
				},
				"givenAccessAt": "2023-10-23T11:44:02+00:00"
			}
		}`)
	})

	member, resp, err := client.Users.GetProjectMember(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectMember{
		ID:        12,
		Username:  "john_smith",
		FirstName: ToPtr("John"),
		LastName:  ToPtr("Smith"),
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
		IsManager:   ToPtr(false),
		IsDeveloper: ToPtr(false),
		ManagerOfGroup: &struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   1,
			Name: "KB materials",
		},
		AccessToAllWorkflowSteps: ToPtr(false),
		Permissions: map[string]any{
			"it": map[string]any{
				"workflowStepIds": []any{313.0},
			},
		},
		GivenAccessAt: ToPtr("2023-10-23T11:44:02+00:00"),
	}
	assert.Equal(t, expected, member)
	assert.Nil(t, member.FullName)
	assert.Nil(t, member.Role)
	assert.Nil(t, member.AvatarURL)
	assert.Nil(t, member.JoinedAt)
	assert.Nil(t, member.Timezone)
}

func TestUsersService_ListProjectMembers(t *testing.T) {
	tests := []struct {
		name          string
		projectID     int
		opts          *model.ProjectMembersListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.ProjectMembersListOptions{},
		},
		{
			name: "with options",
			opts: &model.ProjectMembersListOptions{
				OrderBy:        "createdAt desc,username",
				Search:         "john",
				Role:           "translator",
				LanguageID:     "en",
				WorkflowStepID: 1,
				ListOptions: model.ListOptions{
					Offset: 10,
					Limit:  25,
				},
			},
			expectedQuery: "?languageId=en&limit=25&offset=10&orderBy=createdAt+desc%2Cusername&role=translator&search=john&workflowStepId=1",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for projectID, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v2/projects/%d/members", projectID)
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 12
							}
						},
						{
							"data": {
								"id": 14
							}
						},
						{
							"data": {
								"id": 16
							}
						}
					],
					"pagination": {
						"offset": 10,
						"limit": 25
					}
				}`)
			})

			members, resp, err := client.Users.ListProjectMembers(context.Background(), projectID, tt.opts)
			require.NoError(t, err)

			expected := []*model.ProjectMember{{ID: 12}, {ID: 14}, {ID: 16}}
			assert.Equal(t, expected, members)
			assert.Len(t, members, 3)

			assert.Equal(t, 10, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestUsersService_ListProjectMembers_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/2/members", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Users.ListProjectMembers(context.Background(), 2, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestUsersService_AddProjectMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"userIds": [1],
			"usernames": ["john_smith"],
			"emails": ["john@example.com"],
			"managerAccess": false,
			"developerAccess": false,
			"roles": [
				{
					"name": "translator",
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

		fmt.Fprint(w, `{
			"skipped": [
				{
					"data": {
						"id": 12,
						"username": "john_smith",
						"fullName": "John Smith",
						"role": "translator",
						"permissions": {
							"uk": "translator",
							"it": "proofreader",
							"en": "denied"
						},
						"roles": [
							{
								"name": "translator",
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
						],
						"avatarUrl": "",
						"joinedAt": "2023-07-11T07:40:22+00:00",
						"timezone": "Europe/Kyiv"
						}
				}
			],
			"added": [
				{
					"data": {
						"id": 12,
						"username": "john_smith",
						"fullName": "John Smith",
						"role": "translator",
						"permissions": {
							"uk": "translator",
							"it": "proofreader",
							"en": "denied"
						},
						"roles": [
							{
								"name": "translator",
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
						],
						"avatarUrl": "",
						"joinedAt": "2023-07-11T07:40:22+00:00",
						"timezone": "Europe/Kyiv"
					}
				}
			],
			"pagination": {
				"offset": 0,
				"limit": 25
			}
		}`)
	})

	req := &model.ProjectMemberAddRequest{
		UserIDs:         []int{1},
		Usernames:       []string{"john_smith"},
		Emails:          []string{"john@example.com"},
		ManagerAccess:   ToPtr(false),
		DeveloperAccess: ToPtr(false),
		Roles: []*model.TranslatorRole{
			{
				Name: "translator",
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
			{
				Name: "proofreader",
				Permissions: &model.RolePermissions{
					AllLanguages: ToPtr(true),
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
	member, resp, err := client.Users.AddProjectMember(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := map[string][]*model.ProjectMember{
		"skipped": {
			{
				ID:       12,
				Username: "john_smith",
				FullName: ToPtr("John Smith"),
				Role:     ToPtr("translator"),
				Permissions: map[string]any{
					"uk": "translator",
					"it": "proofreader",
					"en": "denied",
				},
				Roles: []*model.TranslatorRole{
					{
						Name: "translator",
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
				AvatarURL: ToPtr(""),
				JoinedAt:  ToPtr("2023-07-11T07:40:22+00:00"),
				Timezone:  ToPtr("Europe/Kyiv"),
			},
		},
		"added": {
			{
				ID:       12,
				Username: "john_smith",
				FullName: ToPtr("John Smith"),
				Role:     ToPtr("translator"),
				Permissions: map[string]any{
					"uk": "translator",
					"it": "proofreader",
					"en": "denied",
				},
				Roles: []*model.TranslatorRole{
					{
						Name: "translator",
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
				AvatarURL: ToPtr(""),
				JoinedAt:  ToPtr("2023-07-11T07:40:22+00:00"),
				Timezone:  ToPtr("Europe/Kyiv"),
			},
		},
	}
	assert.Equal(t, expected, member)
	assert.Contains(t, member, "skipped")
	assert.Contains(t, member, "added")
}

func TestUsersService_AddProjectMember_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/members", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Users.AddProjectMember(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestUsersService_ReplaceProjectMemberPermissions(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/members/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"managerAccess": false,
			"developerAccess": false,
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

		fmt.Fprint(w, `{
			"data": {
				"id": 12,
				"username": "john_smith",
				"firstName": "John",
				"lastName": "Smith",
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
				],
				"isManager": false,
				"isDeveloper": false,
				"managerOfGroup": {
					"id": 1,
					"name": "KB materials"
				},
				"accessToAllWorkflowSteps": false,
				"permissions": {
					"it": {
						"workflowStepIds": [313]
					}
				},
				"givenAccessAt": "2023-10-23T11:44:02+00:00"
			}
		}`)
	})

	req := &model.ProjectMemberReplaceRequest{
		ManagerAccess:   ToPtr(false),
		DeveloperAccess: ToPtr(false),
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
					AllLanguages: ToPtr(true),
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
	member, resp, err := client.Users.ReplaceProjectMemberPermissions(context.Background(), 1, 2, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectMember{
		ID:        12,
		Username:  "john_smith",
		FirstName: ToPtr("John"),
		LastName:  ToPtr("Smith"),
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
					AllLanguages: ToPtr(true),
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
		IsManager:   ToPtr(false),
		IsDeveloper: ToPtr(false),
		ManagerOfGroup: &struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   1,
			Name: "KB materials",
		},
		AccessToAllWorkflowSteps: ToPtr(false),
		Permissions: map[string]any{
			"it": map[string]any{
				"workflowStepIds": []any{313.0},
			},
		},
		GivenAccessAt: ToPtr("2023-10-23T11:44:02+00:00"),
	}
	assert.Equal(t, expected, member)
}

func TestUsersService_DeleteProjectMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/members/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Users.DeleteProjectMember(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUsersService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/12"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, userJSONResponse(12))
	})

	user, resp, err := client.Users.Get(context.Background(), 12)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.User{
		ID:        12,
		Username:  "john_smith",
		Email:     "jsmith@example.com",
		FirstName: ToPtr("John"),
		LastName:  ToPtr("Smith"),
		Status:    ToPtr("active"),
		AvatarURL: "",
		CreatedAt: "2023-07-11T07:40:22+00:00",
		LastSeen:  "2023-10-23T11:44:02+00:00",
		TwoFactor: "enabled",
		IsAdmin:   ToPtr(true),
		Timezone:  "Europe/Kyiv",
		Fields: map[string]interface{}{
			"fieldSlug": "fieldValue",
		},
	}
	assert.Equal(t, expected, user)
}

func TestUsersService_List(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.UsersListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name:          "empty options",
			opts:          &model.UsersListOptions{},
			expectedQuery: "",
		},
		{
			name: "with options",
			opts: &model.UsersListOptions{
				OrderBy:     "createdAt desc,username",
				Status:      "active",
				Search:      "john",
				TwoFactor:   "enabled",
				ListOptions: model.ListOptions{Offset: 10, Limit: 25},
			},
			expectedQuery: "?limit=25&offset=10&orderBy=createdAt+desc%2Cusername&search=john&status=active&twoFactor=enabled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			path := "/api/v2/users"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 12
							}
						},
						{
							"data": {
								"id": 14
							}
						},
						{
							"data": {
								"id": 16
							}
						}
					],
					"pagination": {
						"offset": 10,
						"limit": 25
					}
				}`)
			})

			users, resp, err := client.Users.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.User{{ID: 12}, {ID: 14}, {ID: 16}}
			assert.Equal(t, expected, users)
			assert.Len(t, users, 3)

			assert.Equal(t, 10, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestUsersService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Users.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestUsersService_GetAuthenticated(t *testing.T) {
	tests := []struct {
		name   string
		resp   string
		expect *model.User
	}{
		{
			name: "authenticated user (crowdin API)",
			resp: `{
				"data": {
					"id": 12,
					"username": "john_smith",
					"email": "jsmith@example.com",
					"fullName": "John Smith",
					"avatarUrl": "",
					"createdAt": "2023-07-11T07:40:22+00:00",
					"lastSeen": "2023-10-23T11:44:02+00:00",
					"twoFactor": "enabled",
					"timezone": "Europe/Kyiv"
				}
			}`,
			expect: &model.User{
				ID:        12,
				Username:  "john_smith",
				Email:     "jsmith@example.com",
				FullName:  ToPtr("John Smith"),
				AvatarURL: "",
				CreatedAt: "2023-07-11T07:40:22+00:00",
				LastSeen:  "2023-10-23T11:44:02+00:00",
				TwoFactor: "enabled",
				Timezone:  "Europe/Kyiv",
			},
		},
		{
			name: "authenticated user (enterprise API)",
			resp: `{
				"data": {
					"id": 12,
					"username": "john_smith",
					"email": "jsmith@example.com",
					"firstName": "John",
					"lastName": "Smith",
					"status": "active",
					"avatarUrl": "",
					"createdAt": "2023-07-11T07:40:22+00:00",
					"lastSeen": "2023-10-23T11:44:02+00:00",
					"twoFactor": "enabled",
					"isAdmin": true,
					"timezone": "Europe/Kyiv",
					"fields": {
						"fieldSlug": "fieldValue"
					}
				}
			}`,
			expect: &model.User{
				ID:        12,
				Username:  "john_smith",
				Email:     "jsmith@example.com",
				FirstName: ToPtr("John"),
				LastName:  ToPtr("Smith"),
				Status:    ToPtr("active"),
				AvatarURL: "",
				CreatedAt: "2023-07-11T07:40:22+00:00",
				LastSeen:  "2023-10-23T11:44:02+00:00",
				TwoFactor: "enabled",
				IsAdmin:   ToPtr(true),
				Timezone:  "Europe/Kyiv",
				Fields: map[string]interface{}{
					"fieldSlug": "fieldValue",
				},
			},
		},
	}

	for _, tt := range tests {
		client, mux, teardown := setupClient()
		defer teardown()

		const path = "/api/v2/user"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path)

			fmt.Fprint(w, tt.resp)
		})

		user, resp, err := client.Users.GetAuthenticated(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, resp)

		assert.Equal(t, tt.expect, user)
	}
}

func TestUsersService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Users.Delete(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUsersService_Invite(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"email":"jsmith@example.com","firstName":"John","lastName":"Smith","timezone":"Europe/Kyiv","adminAccess":false}`+"\n")

		fmt.Fprint(w, userJSONResponse(12))
	})

	req := &model.InviteUserRequest{
		Email:       "jsmith@example.com",
		FirstName:   "John",
		LastName:    "Smith",
		Timezone:    "Europe/Kyiv",
		AdminAccess: ToPtr(false),
	}
	user, resp, err := client.Users.Invite(context.Background(), req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.User{
		ID:        12,
		Username:  "john_smith",
		Email:     "jsmith@example.com",
		FirstName: ToPtr("John"),
		LastName:  ToPtr("Smith"),
		Status:    ToPtr("active"),
		AvatarURL: "",
		CreatedAt: "2023-07-11T07:40:22+00:00",
		LastSeen:  "2023-10-23T11:44:02+00:00",
		TwoFactor: "enabled",
		IsAdmin:   ToPtr(true),
		Timezone:  "Europe/Kyiv",
		Fields: map[string]interface{}{
			"fieldSlug": "fieldValue",
		},
	}
	assert.Equal(t, expected, user)
}

func TestUsersService_Invite_WithRequiredFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users", func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"email":"john@example.com"}`+"\n")

		fmt.Fprint(w, `{}`)
	})

	req := &model.InviteUserRequest{
		Email: "john@example.com",
	}
	_, _, err := client.Users.Invite(context.Background(), req)
	require.NoError(t, err)
}

func TestUsersService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/12"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/firstName","value":"John"}]`+"\n")

		fmt.Fprint(w, userJSONResponse(12))
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/firstName",
			Value: "John",
		},
	}
	user, resp, err := client.Users.Edit(context.Background(), 12, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.User{
		ID:        12,
		Username:  "john_smith",
		Email:     "jsmith@example.com",
		FirstName: ToPtr("John"),
		LastName:  ToPtr("Smith"),
		Status:    ToPtr("active"),
		AvatarURL: "",
		CreatedAt: "2023-07-11T07:40:22+00:00",
		LastSeen:  "2023-10-23T11:44:02+00:00",
		TwoFactor: "enabled",
		IsAdmin:   ToPtr(true),
		Timezone:  "Europe/Kyiv",
		Fields: map[string]interface{}{
			"fieldSlug": "fieldValue",
		},
	}
	assert.Equal(t, expected, user)
}

func userJSONResponse(id int) string {
	return fmt.Sprintf(`{
		"data": {
			"id": %d,
			"username": "john_smith",
			"email": "jsmith@example.com",
			"firstName": "John",
			"lastName": "Smith",
			"status": "active",
			"avatarUrl": "",
			"createdAt": "2023-07-11T07:40:22+00:00",
			"lastSeen": "2023-10-23T11:44:02+00:00",
			"twoFactor": "enabled",
			"isAdmin": true,
			"timezone": "Europe/Kyiv",
			"fields": {
				"fieldSlug": "fieldValue"
			}
		}
	}`, id)
}
