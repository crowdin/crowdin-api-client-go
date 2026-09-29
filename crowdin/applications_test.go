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

func TestApplicationsService_GetInstallation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "example-application",
				"name": "Application name",
				"description": "Application description",
				"logo": "/resources/images/logo.png",
				"baseUrl": "https://localhost.dev",
				"manifestUrl": "https://localhost.dev",
				"createdAt": "2023-09-20T11:34:40+00:00",
				"modules": [
					{
						"key": "example-application",
						"type": "module-type",
						"data": {},
						"permissions": {
							"user": {
								"value": "restricted",
								"ids": [1]
							}
						},
						"authenticationType": "none"
					}
				],
				"scopes": [
					"project"
				],
				"permissions": {
					"user": {
						"value": "restricted",
						"ids": [1]
					},
					"project": {
						"value": "restricted",
						"ids": [1]
					}
				},
				"defaultPermissions": {
					"user": "owner",
					"project": "own"
				},
				"limitReached": true
			}
		}`)
	})

	installation, resp, err := client.Applications.GetInstallation(context.Background(), "example-application")
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Installation{
		Identifier:  "example-application",
		Name:        "Application name",
		Description: "Application description",
		Logo:        "/resources/images/logo.png",
		BaseURL:     "https://localhost.dev",
		ManifestURL: "https://localhost.dev",
		CreatedAt:   "2023-09-20T11:34:40+00:00",
		Modules: []*model.Module{
			{
				Key:  "example-application",
				Type: "module-type",
				Data: map[string]any{},
				Permissions: model.UserPermission{
					User: model.Permission{
						Value: "restricted",
						IDs:   []int{1},
					},
				},
				AuthenticationType: "none",
			},
		},
		Scopes: []string{"project"},
		Permissions: model.ProjectPermission{
			Project: model.Permission{
				Value: "restricted",
				IDs:   []int{1},
			},
		},
		DefaultPermissions: struct {
			User    model.PermissionValue `json:"user"`
			Project model.PermissionValue `json:"project"`
		}{
			User:    model.PermissionOwner,
			Project: model.PermissionOwn,
		},
		LimitReached: true,
	}
	assert.Equal(t, expected, installation)
}

func TestApplicationsService_ListInstallation(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name:          "empty options",
			opts:          &model.ListOptions{},
			expectedQuery: "",
		},
		{
			name:          "all options",
			opts:          &model.ListOptions{Limit: 10, Offset: 5},
			expectedQuery: "?limit=10&offset=5",
		},
	}

	for _, tt := range tests {
		client, mux, teardown := setupClient()
		defer teardown()

		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc("/api/v2/applications/installations", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, "/api/v2/applications/installations"+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"identifier": "example-application",
								"name": "Application name"
							}
						}
					],
					"pagination": {
						"offset": 2,
						"limit": 2
					}
				}`)
			})

			installations, resp, err := client.Applications.ListInstallations(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.Installation{
				{
					Identifier: "example-application",
					Name:       "Application name",
				},
			}
			assert.Equal(t, expected, installations)

			assert.Equal(t, 2, resp.Pagination.Offset)
			assert.Equal(t, 2, resp.Pagination.Limit)
		})
	}
}

func TestApplicationsService_ListInstallation_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/applications/installations", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	installations, _, err := client.Applications.ListInstallations(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, installations)
}

func TestApplicationsService_Install(t *testing.T) {
	tests := []struct {
		name string
		req  *model.InstallApplicationRequest
		body string
	}{
		{
			name: "required fields",
			req: &model.InstallApplicationRequest{
				URL: "https://localhost.dev",
			},
			body: `{"url":"https://localhost.dev"}` + "\n",
		},
		{
			name: "with own permissions",
			req: &model.InstallApplicationRequest{
				URL: "https://localhost.dev",
				Permissions: &model.ProjectPermission{
					Project: model.Permission{
						Value: "own",
					},
				},
			},
			body: `{"url":"https://localhost.dev","permissions":{"project":{"value":"own"}}}` + "\n",
		},
		{
			name: "with restricted permissions",
			req: &model.InstallApplicationRequest{
				URL: "https://localhost.dev",
				Permissions: &model.ProjectPermission{
					Project: model.Permission{
						Value: "restricted",
						IDs:   []int{1},
					},
				},
			},
			body: `{"url":"https://localhost.dev","permissions":{"project":{"value":"restricted","ids":[1]}}}` + "\n",
		},
		{
			name: "with modules",
			req: &model.InstallApplicationRequest{
				URL: "https://localhost.dev",
				Modules: []*model.InstallationModule{
					{
						Key: "example-module",
						Permissions: model.UserPermission{
							User: model.Permission{
								Value: "restricted",
								IDs:   []int{2},
							},
						},
					},
				},
			},
			body: `{"url":"https://localhost.dev","modules":[{"key":"example-module","permissions":{"user":{"value":"restricted","ids":[2]}}}]}` + "\n",
		},
	}

	for _, tt := range tests {
		client, mux, teardown := setupClient()
		defer teardown()

		t.Run(tt.name, func(t *testing.T) {
			const path = "/api/v2/applications/installations"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, path)
				testBody(t, r, tt.body)

				fmt.Fprint(w, `{
					"data": {
						"identifier": "example-application",
						"name": "Application name"
					}
				}`)
			})

			installation, resp, err := client.Applications.Install(context.Background(), tt.req)
			require.NoError(t, err)
			assert.NotNil(t, resp)

			expected := &model.Installation{
				Identifier: "example-application",
				Name:       "Application name",
			}
			assert.Equal(t, expected, installation)
		})
	}
}

func TestApplicationsService_EditInstallation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/permissions","value":{"user":{"value":"managers"},"project":{"value":"restricted","ids":[2]}}}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"identifier": "example-application",
				"name": "Application name",
				"description": "Application description",
				"logo": "/resources/images/logo.png",
				"baseUrl": "https://localhost.dev",
				"manifestUrl": "https://localhost.dev",
				"createdAt": "2023-09-20T11:34:40+00:00",
				"modules": [
					{
						"key": "example-application",
						"type": "module-type",
						"data": {},
						"permissions": {
							"user": {
								"value": "restricted",
								"ids": [1]
							}
						},
						"authenticationType": "none"
					}
				],
				"scopes": [
					"project"
				],
				"permissions": {
					"user": {
						"value": "restricted",
						"ids": [1]
					},
					"project": {
						"value": "restricted",
						"ids": [1]
					}
				},
				"defaultPermissions": {
					"user": "owner",
					"project": "own"
				},
				"limitReached": true
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:   "replace",
			Path: "/permissions",
			Value: model.InstallationReplaceValue{
				Project: model.Permission{
					Value: model.PermissionRestricted,
					IDs:   []int{2},
				},
				User: model.Permission{
					Value: model.PermissionManagers,
				},
			},
		},
	}
	installation, resp, err := client.Applications.EditInstallation(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Installation{
		Identifier:  "example-application",
		Name:        "Application name",
		Description: "Application description",
		Logo:        "/resources/images/logo.png",
		BaseURL:     "https://localhost.dev",
		ManifestURL: "https://localhost.dev",
		CreatedAt:   "2023-09-20T11:34:40+00:00",
		Modules: []*model.Module{
			{
				Key:  "example-application",
				Type: "module-type",
				Data: map[string]any{},
				Permissions: model.UserPermission{
					User: model.Permission{
						Value: "restricted",
						IDs:   []int{1},
					},
				},
				AuthenticationType: "none",
			},
		},
		Scopes: []string{"project"},
		Permissions: model.ProjectPermission{
			Project: model.Permission{
				Value: "restricted",
				IDs:   []int{1},
			},
		},
		DefaultPermissions: struct {
			User    model.PermissionValue `json:"user"`
			Project model.PermissionValue `json:"project"`
		}{
			User:    model.PermissionOwner,
			Project: model.PermissionOwn,
		},
		LimitReached: true,
	}
	assert.Equal(t, expected, installation)
}

func TestApplicationsService_DeleteInstallation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	t.Run("force delete", func(t *testing.T) {
		path := "/api/v2/applications/installations/example-application-1"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path+"?force=true")

			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.Applications.DeleteInstallation(context.Background(), "example-application-1", true)
		require.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("delete", func(t *testing.T) {
		path := "/api/v2/applications/installations/example-application-2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)

			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.Applications.DeleteInstallation(context.Background(), "example-application-2", false)
		require.NoError(t, err)
		assert.NotNil(t, resp)
	})
}

func TestApplicationsService_GetData(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var (
		applicationID = "example-application"
		path          = "example-path"
	)

	endpoint := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, endpoint)

		fmt.Fprint(w, `{
			"data": {
				"key1": "value1",
				"key2": "value2",
				"key3": {
					"foo": "bar"
				}
			}
		}`)
	})

	data, resp, err := client.Applications.GetData(context.Background(), applicationID, path)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": map[string]any{
			"foo": "bar",
		},
	}
	assert.Equal(t, expected, data)
}

func TestApplicationsService_AddData(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var (
		applicationID = "example-application"
		path          = "example-path"

		req = map[string]any{
			"key1": "value1",
			"key2": "value2",
			"key3": map[string]any{
				"foo": "bar",
			},
		}
	)

	endpoint := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, endpoint)
		testBody(t, r, `{"key1":"value1","key2":"value2","key3":{"foo":"bar"}}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"key1": "value1",
				"key2": "value2",
				"key3": {
					"foo": "bar"
				}
			}
		}`)
	})

	data, resp, err := client.Applications.AddData(context.Background(), applicationID, path, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": map[string]any{
			"foo": "bar",
		},
	}
	assert.Equal(t, expected, data)
}

func TestApplicationsService_UpdateOrRestoreData(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var (
		applicationID = "example-application"
		path          = "example-path"

		req = map[string]any{
			"key1": "value1",
			"key3": map[string]any{
				"foo": "bar",
			},
		}
	)

	endpoint := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testURL(t, r, endpoint)
		testBody(t, r, `{"key1":"value1","key3":{"foo":"bar"}}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"key1": "value1",
				"key3": {
					"foo": "bar"
				}
			}
		}`)
	})

	data, resp, err := client.Applications.UpdateOrRestoreData(context.Background(), applicationID, path, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := map[string]any{
		"key1": "value1",
		"key3": map[string]any{
			"foo": "bar",
		},
	}
	assert.Equal(t, expected, data)
}

func TestApplicationsService_EditData(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var (
		applicationID = "example-application"
		path          = "example-path"

		req = map[string]any{
			"key1": "value1",
			"key2": "value2",
		}
	)

	endpoint := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, endpoint)
		testBody(t, r, `{"key1":"value1","key2":"value2"}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"key1": "value1",
				"key2": "value2"
			}
		}`)
	})

	data, resp, err := client.Applications.EditData(context.Background(), applicationID, path, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}
	assert.Equal(t, expected, data)
}

func TestApplicationsService_DeleteData(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var (
		applicationID = "example-application"
		path          = "example-path"
	)

	endpoint := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, endpoint)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Applications.DeleteData(context.Background(), applicationID, path)
	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestApplicationsService_GetInstallation_withNewFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "example-application",
				"name": "Application name",
				"installedBy": {
					"id": 12,
					"username": "john_smith",
					"fullName": "John Smith",
					"avatarUrl": ""
				},
				"description": null,
				"logo": "/resources/images/logo.png",
				"logoUrl": "https://localhost.dev/resources/images/logo.png",
				"agent": null,
				"baseUrl": null,
				"manifestUrl": null,
				"manifest": {
					"name": "My App"
				},
				"createdAt": "2019-09-20T11:34:40+00:00",
				"manifestUpdatedAt": "2019-09-20T11:34:40+00:00",
				"isManifestOutdated": true,
				"modules": [
					{
						"key": "example-asset-panel",
						"type": "editor-asset-panel",
						"name": "Asset Panel",
						"modes": {"translate": true},
						"fileNamePattern": "\\.(png|svg)$"
					},
					{
						"key": "example-button",
						"type": "editor-button",
						"modes": ["translate", "review"],
						"multilingual": false
					}
				],
				"scopes": ["project"],
				"permissions": {
					"project": {
						"value": "own"
					}
				},
				"defaultPermissions": {
					"user": "owner",
					"project": "own"
				},
				"limitReached": false,
				"stringBasedAvailable": true,
				"bundle": {
					"mode": "external",
					"url": "http://localhost:8080/"
				}
			}
		}`)
	})

	installation, resp, err := client.Applications.GetInstallation(context.Background(), "example-application")
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Installation{
		Identifier: "example-application",
		Name:       "Application name",
		InstalledBy: &model.ShortUser{
			ID:        12,
			Username:  "john_smith",
			FullName:  "John Smith",
			AvatarURL: "",
		},
		Logo:               "/resources/images/logo.png",
		LogoURL:            ToPtr("https://localhost.dev/resources/images/logo.png"),
		Manifest:           map[string]any{"name": "My App"},
		CreatedAt:          "2019-09-20T11:34:40+00:00",
		ManifestUpdatedAt:  ToPtr("2019-09-20T11:34:40+00:00"),
		IsManifestOutdated: true,
		Modules: []*model.Module{
			{
				Key:             "example-asset-panel",
				Type:            "editor-asset-panel",
				Name:            ToPtr("Asset Panel"),
				Modes:           map[string]any{"translate": true},
				FileNamePattern: ToPtr(`\.(png|svg)$`),
			},
			{
				Key:          "example-button",
				Type:         "editor-button",
				Modes:        []any{"translate", "review"},
				Multilingual: ToPtr(false),
			},
		},
		Scopes: []string{"project"},
		Permissions: model.ProjectPermission{
			Project: model.Permission{Value: model.PermissionOwn},
		},
		DefaultPermissions: struct {
			User    model.PermissionValue `json:"user"`
			Project model.PermissionValue `json:"project"`
		}{
			User:    model.PermissionOwner,
			Project: model.PermissionOwn,
		},
		StringBasedAvailable: true,
		Bundle: &model.ApplicationBundle{
			Mode: "external",
			URL:  "http://localhost:8080/",
		},
	}
	assert.Equal(t, expected, installation)
}

func TestApplicationsService_ListInstallationsWithOptions(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.InstallationsListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name:          "empty options",
			opts:          &model.InstallationsListOptions{},
			expectedQuery: "",
		},
		{
			name: "all options",
			opts: &model.InstallationsListOptions{
				InstalledBy: 12,
				OrderBy:     "createdAt desc",
				ListOptions: model.ListOptions{Limit: 10, Offset: 5},
			},
			expectedQuery: "?installedBy=12&limit=10&offset=5&orderBy=createdAt+desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc("/api/v2/applications/installations", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, "/api/v2/applications/installations"+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"identifier": "example-application",
								"name": "Application name"
							}
						}
					],
					"pagination": {
						"offset": 5,
						"limit": 10
					}
				}`)
			})

			installations, resp, err := client.Applications.ListInstallationsWithOptions(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.Installation{
				{
					Identifier: "example-application",
					Name:       "Application name",
				},
			}
			assert.Equal(t, expected, installations)
			assert.Equal(t, 5, resp.Pagination.Offset)
			assert.Equal(t, 10, resp.Pagination.Limit)
		})
	}
}

func TestApplicationsService_ListInstallationsWithOptions_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/applications/installations", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	installations, _, err := client.Applications.ListInstallationsWithOptions(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, installations)
}

func TestApplicationsService_Install_withManifest(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"manifest": {
				"name": "My App",
				"description": "App description",
				"scopes": ["project"],
				"stringBasedAvailable": false,
				"modules": {
					"editor-right-panel": [
						{
							"key": "my-module",
							"name": "My Module",
							"environments": ["crowdin"],
							"permissions": {"user": {"value": "all"}}
						}
					]
				},
				"default_permissions": {"user": "owner", "project": "own"},
				"bundle": {"mode": "internal"}
			},
			"permissions": {"project": {"value": "own"}}
		}`)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "a1b2c3",
				"name": "My App",
				"bundle": {"mode": "internal"}
			}
		}`)
	})

	req := &model.InstallApplicationRequest{
		Manifest: &model.ApplicationManifest{
			Name:                 "My App",
			Description:          "App description",
			Scopes:               []string{"project"},
			StringBasedAvailable: ToPtr(false),
			Modules: map[string][]*model.ApplicationManifestModule{
				"editor-right-panel": {
					{
						Key:          "my-module",
						Name:         "My Module",
						Environments: []string{"crowdin"},
						Permissions: &model.UserPermission{
							User: model.Permission{Value: model.PermissionAll},
						},
					},
				},
			},
			DefaultPermissions: &model.ApplicationDefaultPermissions{
				User:    model.PermissionOwner,
				Project: model.PermissionOwn,
			},
			Bundle: &model.ApplicationBundle{Mode: "internal"},
		},
		Permissions: &model.ProjectPermission{
			Project: model.Permission{Value: model.PermissionOwn},
		},
	}
	installation, resp, err := client.Applications.Install(context.Background(), req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Installation{
		Identifier: "a1b2c3",
		Name:       "My App",
		Bundle:     &model.ApplicationBundle{Mode: "internal"},
	}
	assert.Equal(t, expected, installation)
}

func TestApplicationsService_Install_withAssignAgent(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"url":"https://localhost.dev","assignAgent":false}`+"\n")

		fmt.Fprint(w, `{"data": {"identifier": "example-application"}}`)
	})

	req := &model.InstallApplicationRequest{URL: "https://localhost.dev", AssignAgent: ToPtr(false)}
	installation, _, err := client.Applications.Install(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, "example-application", installation.Identifier)
}

func TestApplicationsService_Install_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Applications.Install(context.Background(), &model.InstallApplicationRequest{})
	require.EqualError(t, err, "url is required")
}

func TestApplicationsService_GetInstallationUpdate(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application/update"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"manifestHash": "5536b952d54e0846",
				"latestManifest": {"name": "My App"},
				"addedScopes": ["tm"],
				"removedScopes": ["project"],
				"addedModules": [{"key": "new-module"}],
				"removedModules": [],
				"changedModules": [{"key": "changed-module"}],
				"changedEvents": {
					"installed": {"from": "/installed", "to": "/v2/installed"}
				},
				"baseUrlChanged": {"from": "https://old.dev", "to": "https://new.dev"},
				"authenticationTypeChanged": null,
				"hasChanges": true
			}
		}`)
	})

	update, resp, err := client.Applications.GetInstallationUpdate(context.Background(), "example-application")
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.InstallationUpdate{
		ManifestHash:   ToPtr("5536b952d54e0846"),
		LatestManifest: map[string]any{"name": "My App"},
		AddedScopes:    []string{"tm"},
		RemovedScopes:  []string{"project"},
		AddedModules:   []map[string]any{{"key": "new-module"}},
		RemovedModules: []map[string]any{},
		ChangedModules: []map[string]any{{"key": "changed-module"}},
		ChangedEvents: map[string]*model.InstallationUpdateChange{
			"installed": {From: ToPtr("/installed"), To: ToPtr("/v2/installed")},
		},
		BaseURLChanged: &model.InstallationUpdateChange{From: ToPtr("https://old.dev"), To: ToPtr("https://new.dev")},
		HasChanges:     true,
	}
	assert.Equal(t, expected, update)
}

func TestApplicationsService_ApplyInstallationUpdate(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application/update"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"manifestHash":"5536b952d54e0846"}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"identifier": "example-application",
				"name": "Application name",
				"isManifestOutdated": false
			}
		}`)
	})

	req := &model.InstallationUpdateApplyRequest{ManifestHash: "5536b952d54e0846"}
	installation, resp, err := client.Applications.ApplyInstallationUpdate(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Installation{
		Identifier: "example-application",
		Name:       "Application name",
	}
	assert.Equal(t, expected, installation)
}

func TestApplicationsService_ApplyInstallationUpdate_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Applications.ApplyInstallationUpdate(context.Background(), "example-application",
		&model.InstallationUpdateApplyRequest{})
	require.EqualError(t, err, "manifestHash is required")
}

func TestApplicationsService_UploadBundle(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/installations/example-application/bundles"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"storageId":12}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"data": {
				"identifier": "example-application",
				"name": "Application name",
				"bundle": {"mode": "internal"}
			}
		}`)
	})

	req := &model.ApplicationBundleUploadRequest{StorageID: 12}
	installation, resp, err := client.Applications.UploadBundle(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := &model.Installation{
		Identifier: "example-application",
		Name:       "Application name",
		Bundle:     &model.ApplicationBundle{Mode: "internal"},
	}
	assert.Equal(t, expected, installation)
}

const applicationConsentJSON = `{
	"id": 12,
	"installedBy": {
		"id": 12,
		"username": "john_smith",
		"fullName": "John Smith",
		"avatarUrl": ""
	},
	"identifier": "example-application",
	"name": "My App",
	"status": "granted",
	"scopes": ["project", "tm"],
	"createdAt": "2026-07-24T10:00:00+00:00",
	"updatedAt": "2026-07-24T10:00:00+00:00"
}`

func expectedApplicationConsent() *model.ApplicationConsent {
	return &model.ApplicationConsent{
		ID: 12,
		InstalledBy: &model.ShortUser{
			ID:       12,
			Username: "john_smith",
			FullName: "John Smith",
		},
		Identifier: "example-application",
		Name:       ToPtr("My App"),
		Status:     model.ApplicationConsentGranted,
		Scopes:     []string{"project", "tm"},
		CreatedAt:  "2026-07-24T10:00:00+00:00",
		UpdatedAt:  "2026-07-24T10:00:00+00:00",
	}
}

func TestApplicationsService_ListConsents(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ApplicationConsentsListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name: "all options",
			opts: &model.ApplicationConsentsListOptions{
				Identifier:  "example-application",
				OrderBy:     "createdAt desc",
				ListOptions: model.ListOptions{Limit: 10, Offset: 5},
			},
			expectedQuery: "?identifier=example-application&limit=10&offset=5&orderBy=createdAt+desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/applications/consents"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprintf(w, `{
					"data": [{"data": %s}],
					"pagination": {"offset": 5, "limit": 10}
				}`, applicationConsentJSON)
			})

			consents, resp, err := client.Applications.ListConsents(context.Background(), tt.opts)
			require.NoError(t, err)

			assert.Equal(t, []*model.ApplicationConsent{expectedApplicationConsent()}, consents)
			assert.Equal(t, 5, resp.Pagination.Offset)
			assert.Equal(t, 10, resp.Pagination.Limit)
		})
	}
}

func TestApplicationsService_ListConsents_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/applications/consents", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	consents, _, err := client.Applications.ListConsents(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, consents)
}

func TestApplicationsService_AddConsent(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/consents"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"identifier":"example-application","installedBy":12,"status":"granted","scopes":["project","tm"]}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"data": %s}`, applicationConsentJSON)
	})

	req := &model.ApplicationConsentAddRequest{
		Identifier:  "example-application",
		InstalledBy: 12,
		Status:      model.ApplicationConsentGranted,
		Scopes:      []string{"project", "tm"},
	}
	consent, resp, err := client.Applications.AddConsent(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, expectedApplicationConsent(), consent)
}

func TestApplicationsService_AddConsent_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Applications.AddConsent(context.Background(), &model.ApplicationConsentAddRequest{})
	require.EqualError(t, err, "identifier is required")
}

func TestApplicationsService_EditConsent(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/consents/12"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/status","value":"granted"}]`+"\n")

		fmt.Fprintf(w, `{"data": %s}`, applicationConsentJSON)
	})

	req := []*model.UpdateRequest{
		{
			Op:    model.OpReplace,
			Path:  "/status",
			Value: model.ApplicationConsentGranted,
		},
	}
	consent, resp, err := client.Applications.EditConsent(context.Background(), 12, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedApplicationConsent(), consent)
}

func TestApplicationsService_DeleteConsent(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/consents/12"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Applications.DeleteConsent(context.Background(), 12)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

const applicationKVRecordJSON = `{
	"key": "settings.theme",
	"value": {"color": "dark"},
	"secret": false,
	"createdAt": "2026-08-17T10:00:00+00:00",
	"updatedAt": "2026-08-17T10:00:00+00:00",
	"expiresAt": "2026-09-17T10:00:00+00:00"
}`

func expectedApplicationKVRecord() *model.ApplicationKVRecord {
	return &model.ApplicationKVRecord{
		Key:       "settings.theme",
		Value:     map[string]any{"color": "dark"},
		Secret:    false,
		CreatedAt: "2026-08-17T10:00:00+00:00",
		UpdatedAt: "2026-08-17T10:00:00+00:00",
		ExpiresAt: ToPtr("2026-09-17T10:00:00+00:00"),
	}
}

func TestApplicationsService_ListKVRecords(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ApplicationKVRecordsListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name: "all options",
			opts: &model.ApplicationKVRecordsListOptions{
				Prefix:      "settings.",
				OrderBy:     "key",
				ListOptions: model.ListOptions{Limit: 10, Offset: 5},
			},
			expectedQuery: "?limit=10&offset=5&orderBy=key&prefix=settings.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/applications/example-application/storage/kv/records"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprintf(w, `{
					"data": [{"data": %s}],
					"pagination": {"offset": 5, "limit": 10}
				}`, applicationKVRecordJSON)
			})

			records, resp, err := client.Applications.ListKVRecords(context.Background(), "example-application", tt.opts)
			require.NoError(t, err)

			assert.Equal(t, []*model.ApplicationKVRecord{expectedApplicationKVRecord()}, records)
			assert.Equal(t, 5, resp.Pagination.Offset)
			assert.Equal(t, 10, resp.Pagination.Limit)
		})
	}
}

func TestApplicationsService_ListKVRecords_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/applications/example-application/storage/kv/records", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	records, _, err := client.Applications.ListKVRecords(context.Background(), "example-application", nil)
	require.Error(t, err)
	assert.Nil(t, records)
}

func TestApplicationsService_AddKVRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/example-application/storage/kv/records"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"key":"settings.theme","value":{"color":"dark"},"secret":false,"ttl":2678400}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"data": %s}`, applicationKVRecordJSON)
	})

	req := &model.ApplicationKVRecordAddRequest{
		Key:    "settings.theme",
		Value:  map[string]any{"color": "dark"},
		Secret: ToPtr(false),
		TTL:    2678400,
	}
	record, resp, err := client.Applications.AddKVRecord(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, expectedApplicationKVRecord(), record)
}

func TestApplicationsService_AddKVRecord_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Applications.AddKVRecord(context.Background(), "example-application",
		&model.ApplicationKVRecordAddRequest{Key: "settings.theme"})
	require.EqualError(t, err, "value is required")
}

func TestApplicationsService_GetKVRecord(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		expectedURI string
	}{
		{
			name:        "plain key",
			key:         "settings.theme",
			expectedURI: "/api/v2/applications/example-application/storage/kv/records/settings.theme",
		},
		{
			name:        "key with reserved prefix",
			key:         "user:1:settings.theme",
			expectedURI: "/api/v2/applications/example-application/storage/kv/records/user%3A1%3Asettings.theme",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc("/api/v2/applications/example-application/storage/kv/records/", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, tt.expectedURI)

				fmt.Fprintf(w, `{"data": %s}`, applicationKVRecordJSON)
			})

			record, resp, err := client.Applications.GetKVRecord(context.Background(), "example-application", tt.key)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, expectedApplicationKVRecord(), record)
		})
	}
}

func TestApplicationsService_EditKVRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/example-application/storage/kv/records/settings.theme"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/value","value":{"color":"dark"}},{"op":"replace","path":"/ttl","value":2678400}]`+"\n")

		fmt.Fprintf(w, `{"data": %s}`, applicationKVRecordJSON)
	})

	req := []*model.UpdateRequest{
		{
			Op:    model.OpReplace,
			Path:  "/value",
			Value: map[string]any{"color": "dark"},
		},
		{
			Op:    model.OpReplace,
			Path:  "/ttl",
			Value: 2678400,
		},
	}
	record, resp, err := client.Applications.EditKVRecord(context.Background(), "example-application", "settings.theme", req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedApplicationKVRecord(), record)
}

func TestApplicationsService_RemoveKVRecordTTL(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/applications/example-application/storage/kv/records/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, "/api/v2/applications/example-application/storage/kv/records/module%3Aeditor%3Atheme")
		testBody(t, r, `[{"op":"replace","path":"/ttl","value":null}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"key": "module:editor:theme",
				"value": "dark",
				"secret": true,
				"createdAt": "2026-08-17T10:00:00+00:00",
				"updatedAt": "2026-08-17T10:00:00+00:00",
				"expiresAt": null
			}
		}`)
	})

	record, resp, err := client.Applications.RemoveKVRecordTTL(context.Background(), "example-application", "module:editor:theme")
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ApplicationKVRecord{
		Key:       "module:editor:theme",
		Value:     "dark",
		Secret:    true,
		CreatedAt: "2026-08-17T10:00:00+00:00",
		UpdatedAt: "2026-08-17T10:00:00+00:00",
	}
	assert.Equal(t, expected, record)
}

func TestApplicationsService_DeleteKVRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/applications/example-application/storage/kv/records/settings.theme"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Applications.DeleteKVRecord(context.Background(), "example-application", "settings.theme")
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

const integrationAPIPath = "/api/v2/applications/example-application/api/"

func TestApplicationsService_ListIntegrationCrowdinFiles(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"crowdin-files", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"crowdin-files?projectId=12")

		fmt.Fprint(w, `{
			"data": [
				{"id": 1, "name": "Landing pages"},
				{"id": 2, "parentId": 1, "name": "Home Page", "type": "json"}
			]
		}`)
	})

	files, resp, err := client.Applications.ListIntegrationCrowdinFiles(context.Background(), "example-application", 12)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []map[string]any{
		{"id": float64(1), "name": "Landing pages"},
		{"id": float64(2), "parentId": float64(1), "name": "Home Page", "type": "json"},
	}
	assert.Equal(t, expected, files)
}

func TestApplicationsService_UpdateIntegrationCrowdinFiles(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"crowdin-update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, integrationAPIPath+"crowdin-update")
		testJSONBody(t, r, `{
			"projectId": 12,
			"files": [
				{"id": "1", "name": "Landing pages", "parent_id": "0", "node_type": "0"}
			],
			"uploadTranslations": true
		}`)

		fmt.Fprint(w, `{"data": {"jobId": "067da473-fc0b-43e3-b0a2-09d26af130c1"}}`)
	})

	req := &model.IntegrationCrowdinFilesUpdateRequest{
		ProjectID: 12,
		Files: []map[string]any{
			{"id": "1", "name": "Landing pages", "parent_id": "0", "node_type": "0"},
		},
		UploadTranslations: ToPtr(true),
	}
	result, resp, err := client.Applications.UpdateIntegrationCrowdinFiles(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, &model.IntegrationJobResult{JobID: "067da473-fc0b-43e3-b0a2-09d26af130c1"}, result)
}

func TestApplicationsService_UpdateIntegrationCrowdinFiles_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.Applications.UpdateIntegrationCrowdinFiles(context.Background(), "example-application",
		&model.IntegrationCrowdinFilesUpdateRequest{ProjectID: 12})
	require.EqualError(t, err, "files are required")
}

func TestApplicationsService_GetIntegrationFileProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"file-progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"file-progress?fileId=102&projectId=12")

		fmt.Fprint(w, `{
			"data": {
				"languageId": "af",
				"eTag": "fd0ea167420ef1687fd16635b9fb67a3",
				"words": {"total": 7249, "translated": 3651, "approved": 3637},
				"phrases": {"total": 3041, "translated": 2631, "approved": 2622},
				"translationProgress": 86,
				"approvalProgress": 86
			}
		}`)
	})

	progress, resp, err := client.Applications.GetIntegrationFileProgress(context.Background(), "example-application", 12, 102)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.TranslationProgress{
		LanguageID:          ToPtr("af"),
		Etag:                ToPtr("fd0ea167420ef1687fd16635b9fb67a3"),
		Words:               map[string]int{"total": 7249, "translated": 3651, "approved": 3637},
		Phrases:             map[string]int{"total": 3041, "translated": 2631, "approved": 2622},
		TranslationProgress: 86,
		ApprovalProgress:    86,
	}
	assert.Equal(t, expected, progress)
}

func TestApplicationsService_ListIntegrationFiles(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"integration-files", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"integration-files?projectId=12")

		fmt.Fprint(w, `{
			"data": [
				{"id": 1, "name": "Landing pages", "parent_id": 0, "node_type": 0},
				{"name": "Intro", "id": 73291251883, "parentId": "Landing pages", "type": "json", "node_type": 1}
			]
		}`)
	})

	files, resp, err := client.Applications.ListIntegrationFiles(context.Background(), "example-application", 12)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []map[string]any{
		{"id": float64(1), "name": "Landing pages", "parent_id": float64(0), "node_type": float64(0)},
		{"name": "Intro", "id": float64(73291251883), "parentId": "Landing pages", "type": "json", "node_type": float64(1)},
	}
	assert.Equal(t, expected, files)
}

func TestApplicationsService_UpdateIntegrationFiles(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"integration-update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, integrationAPIPath+"integration-update")
		testBody(t, r, `{"projectId":12,"files":{"102":["de","fr"],"999":["uk"]}}`+"\n")

		fmt.Fprint(w, `{"data": {"jobId": "067da473-fc0b-43e3-b0a2-09d26af130c1"}}`)
	})

	req := &model.IntegrationFilesUpdateRequest{
		ProjectID: 12,
		Files: map[string][]string{
			"102": {"de", "fr"},
			"999": {"uk"},
		},
	}
	result, resp, err := client.Applications.UpdateIntegrationFiles(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, &model.IntegrationJobResult{JobID: "067da473-fc0b-43e3-b0a2-09d26af130c1"}, result)
}

const integrationJobsJSON = `{
	"data": [
		{
			"id": "067da473-fc0b-43e3-b0a2-09d26af130c1",
			"progress": 94,
			"status": "inProgress",
			"title": "Sync files to Crowdin"
		}
	]
}`

func expectedIntegrationJobs() []*model.IntegrationJob {
	return []*model.IntegrationJob{
		{
			ID:       "067da473-fc0b-43e3-b0a2-09d26af130c1",
			Progress: 94,
			Status:   "inProgress",
			Title:    "Sync files to Crowdin",
		},
	}
}

func TestApplicationsService_GetIntegrationJobInfo(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"job-info", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"job-info?jobId=067da473&projectId=12")

		fmt.Fprint(w, integrationJobsJSON)
	})

	jobs, resp, err := client.Applications.GetIntegrationJobInfo(context.Background(), "example-application", 12, "067da473")
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedIntegrationJobs(), jobs)
}

func TestApplicationsService_GetIntegrationJobs(t *testing.T) {
	tests := []struct {
		name          string
		jobID         string
		expectedQuery string
	}{
		{
			name:          "without job id",
			jobID:         "",
			expectedQuery: "?projectId=12",
		},
		{
			name:          "with job id",
			jobID:         "067da473",
			expectedQuery: "?jobId=067da473&projectId=12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(integrationAPIPath+"jobs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, integrationAPIPath+"jobs"+tt.expectedQuery)

				fmt.Fprint(w, integrationJobsJSON)
			})

			jobs, resp, err := client.Applications.GetIntegrationJobs(context.Background(), "example-application", 12, tt.jobID)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, expectedIntegrationJobs(), jobs)
		})
	}
}

func TestApplicationsService_ListIntegrationJobs(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "?projectId=12",
		},
		{
			name:          "with pagination",
			opts:          &model.ListOptions{Limit: 10, Offset: 5},
			expectedQuery: "?limit=10&offset=5&projectId=12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(integrationAPIPath+"all-jobs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, integrationAPIPath+"all-jobs"+tt.expectedQuery)

				fmt.Fprint(w, integrationJobsJSON)
			})

			jobs, resp, err := client.Applications.ListIntegrationJobs(context.Background(), "example-application", 12, tt.opts)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, expectedIntegrationJobs(), jobs)
		})
	}
}

func TestApplicationsService_CancelIntegrationJob(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, integrationAPIPath+"jobs?jobId=067da473&projectId=12")

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Applications.CancelIntegrationJob(context.Background(), "example-application", 12, "067da473")
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestApplicationsService_IntegrationLogin(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"login", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, integrationAPIPath+"login")
		testBody(t, r, `{"projectId":12,"credentials":{"email":"user@crowdin.com","password":"password"}}`+"\n")

		w.WriteHeader(http.StatusNoContent)
	})

	req := &model.IntegrationLoginRequest{
		ProjectID: 12,
		Credentials: map[string]any{
			"email":    "user@crowdin.com",
			"password": "password",
		},
	}
	resp, err := client.Applications.IntegrationLogin(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestApplicationsService_IntegrationLogin_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, err := client.Applications.IntegrationLogin(context.Background(), "example-application",
		&model.IntegrationLoginRequest{ProjectID: 12})
	require.EqualError(t, err, "credentials are required")
}

func TestApplicationsService_ListIntegrationLoginFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"login-fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"login-fields")

		fmt.Fprint(w, `{
			"data": [
				{"key": "email", "name": "User email"},
				{"key": "password", "name": "User password"}
			]
		}`)
	})

	fields, resp, err := client.Applications.ListIntegrationLoginFields(context.Background(), "example-application")
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.IntegrationLoginField{
		{Key: "email", Name: "User email"},
		{Key: "password", Name: "User password"},
	}
	assert.Equal(t, expected, fields)
}

func TestApplicationsService_GetIntegrationSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, integrationAPIPath+"settings?projectId=12")

		fmt.Fprint(w, `{"data": {"schedule": 0, "condition": "1"}}`)
	})

	settings, resp, err := client.Applications.GetIntegrationSettings(context.Background(), "example-application", 12)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, map[string]any{"schedule": float64(0), "condition": "1"}, settings)
}

func TestApplicationsService_UpdateIntegrationSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, integrationAPIPath+"settings")
		testBody(t, r, `{"projectId":12,"config":{"condition":0,"schedule":0}}`+"\n")

		w.WriteHeader(http.StatusNoContent)
	})

	req := &model.IntegrationSettingsUpdateRequest{
		ProjectID: 12,
		Config:    map[string]any{"schedule": 0, "condition": 0},
	}
	resp, err := client.Applications.UpdateIntegrationSettings(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestApplicationsService_GetIntegrationSyncSettings(t *testing.T) {
	tests := []struct {
		name     string
		provider model.IntegrationSyncProvider
		response string
		expected any
	}{
		{
			name:     "object response",
			provider: model.IntegrationSyncProviderCrowdin,
			response: `{"data": {"102": ["uk", "de"]}}`,
			expected: map[string]any{"102": []any{"uk", "de"}},
		},
		{
			name:     "array response",
			provider: model.IntegrationSyncProviderIntegration,
			response: `{"data": [{"id": "1", "name": "Intro", "schedule": true}]}`,
			expected: []any{map[string]any{"id": "1", "name": "Intro", "schedule": true}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(integrationAPIPath+"sync-settings", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, integrationAPIPath+"sync-settings?projectId=12&provider="+string(tt.provider))

				fmt.Fprint(w, tt.response)
			})

			settings, resp, err := client.Applications.GetIntegrationSyncSettings(context.Background(), "example-application", 12, tt.provider)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.expected, settings)
		})
	}
}

func TestApplicationsService_UpdateIntegrationSyncSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc(integrationAPIPath+"sync-settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, integrationAPIPath+"sync-settings")
		testBody(t, r, `{"projectId":12,"provider":"crowdin","files":{"102":["uk","de"]}}`+"\n")

		w.WriteHeader(http.StatusNoContent)
	})

	req := &model.IntegrationSyncSettingsUpdateRequest{
		ProjectID: 12,
		Provider:  model.IntegrationSyncProviderCrowdin,
		Files:     map[string][]string{"102": {"uk", "de"}},
	}
	resp, err := client.Applications.UpdateIntegrationSyncSettings(context.Background(), "example-application", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestApplicationsService_UpdateIntegrationSyncSettings_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, err := client.Applications.UpdateIntegrationSyncSettings(context.Background(), "example-application",
		&model.IntegrationSyncSettingsUpdateRequest{ProjectID: 12, Provider: "unknown"})
	require.EqualError(t, err, `invalid provider: "unknown", must be one of crowdin, integration`)
}
