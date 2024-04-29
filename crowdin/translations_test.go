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

func TestTranslationsService_PreTranslationStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/pre-translations/9e7de270-4f83-41cb-b606-2f90631f26e2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "9e7de270-4f83-41cb-b606-2f90631f26e2",
				"status": "created",
				"progress": 90,
				"attributes": {
					"languageIds": ["uk"],
					"fileIds": [742],
					"method": "tm",
					"autoApproveOption": "all",
					"duplicateTranslations": true,
					"skipApprovedTranslations": true,
					"translateUntranslatedOnly": true,
					"translateWithPerfectMatchOnly": true
				},
				"createdAt": "2023-09-20T14:05:50+00:00",
				"updatedAt": "2023-09-20T14:05:50+00:00",
				"startedAt": "2023-08-24T14:15:22Z",
				"finishedAt": "2023-08-24T14:15:22Z"
			}
		}`)
	})

	status, resp, err := client.Translations.PreTranslationStatus(context.Background(), 1, "9e7de270-4f83-41cb-b606-2f90631f26e2")
	require.NoError(t, err)

	expected := &model.PreTranslation{
		Identifier: "9e7de270-4f83-41cb-b606-2f90631f26e2",
		Status:     "created",
		Progress:   90,
		Attributes: &model.PreTranslationAttributes{
			LanguageIDs:                   []string{"uk"},
			FileIDs:                       []int{742},
			Method:                        ToPtr("tm"),
			AutoApproveOption:             ToPtr("all"),
			DuplicateTranslations:         ToPtr(true),
			SkipApprovedTranslations:      ToPtr(true),
			TranslateUntranslatedOnly:     ToPtr(true),
			TranslateWithPerfectMatchOnly: ToPtr(true),
		},
		CreatedAt:  "2023-09-20T14:05:50+00:00",
		UpdatedAt:  "2023-09-20T14:05:50+00:00",
		StartedAt:  "2023-08-24T14:15:22Z",
		FinishedAt: "2023-08-24T14:15:22Z",
	}
	assert.Equal(t, expected, status)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_ApplyPreTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/pre-translations"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)

		expectedReqBody := `{
			"languageIds": ["uk"],
			"fileIds": [742],
			"method": "tm",
			"engineId": 1,
			"autoApproveOption": "all",
			"duplicateTranslations": true,
			"skipApprovedTranslations": true,
			"translateUntranslatedOnly": false,
			"translateWithPerfectMatchOnly": true,
			"fallbackLanguages": {
			  	"languageId": ["uk"]
			},
			"labelIds": [1],
			"excludeLabelIds": [2]
		}`
		testJSONBody(t, r, expectedReqBody)

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, `{
			"data": {
				"identifier": "9e7de270-4f83-41cb-b606-2f90631f26e2",
				"status": "created",
				"progress": 90,
				"attributes": {
					"languageIds": ["uk"],
					"fileIds": [742],
					"method": "tm",
					"autoApproveOption": "all",
					"duplicateTranslations": true,
					"skipApprovedTranslations": true,
					"translateUntranslatedOnly": false,
					"translateWithPerfectMatchOnly": true
				},
				"createdAt": "2023-09-20T14:05:50+00:00",
				"updatedAt": "2023-09-20T14:05:50+00:00",
				"startedAt": "2023-08-24T14:15:22Z",
				"finishedAt": "2023-08-24T14:15:22Z"
			}
		}`)
	})

	req := &model.PreTranslationRequest{
		LanguageIDs:                   []string{"uk"},
		FileIDs:                       []int{742},
		Method:                        "tm",
		EngineID:                      1,
		AutoApproveOption:             "all",
		DuplicateTranslations:         ToPtr(true),
		SkipApprovedTranslations:      ToPtr(true),
		TranslateUntranslatedOnly:     ToPtr(false),
		TranslateWithPerfectMatchOnly: ToPtr(true),
		FallbackLanguages: map[string][]string{
			"languageId": {"uk"},
		},
		LabelIDs:        []int{1},
		ExcludeLabelIDs: []int{2},
	}
	preTranslation, resp, err := client.Translations.ApplyPreTranslation(context.Background(), 1, req)
	require.NoError(t, err)

	expected := &model.PreTranslation{
		Identifier: "9e7de270-4f83-41cb-b606-2f90631f26e2",
		Status:     "created",
		Progress:   90,
		Attributes: &model.PreTranslationAttributes{
			LanguageIDs:                   []string{"uk"},
			FileIDs:                       []int{742},
			Method:                        ToPtr("tm"),
			AutoApproveOption:             ToPtr("all"),
			DuplicateTranslations:         ToPtr(true),
			SkipApprovedTranslations:      ToPtr(true),
			TranslateUntranslatedOnly:     ToPtr(false),
			TranslateWithPerfectMatchOnly: ToPtr(true),
		},
		CreatedAt:  "2023-09-20T14:05:50+00:00",
		UpdatedAt:  "2023-09-20T14:05:50+00:00",
		StartedAt:  "2023-08-24T14:15:22Z",
		FinishedAt: "2023-08-24T14:15:22Z",
	}
	assert.Equal(t, expected, preTranslation)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func TestTranslationsService_ApplyPreTranslation_WithRequiredFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/pre-translations", func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"languageIds":["uk"],"fileIds":[742]}`+"\n")
		fmt.Fprint(w, `{}`)
	})

	req := &model.PreTranslationRequest{
		LanguageIDs: []string{"uk"},
		FileIDs:     []int{742},
	}
	_, _, err := client.Translations.ApplyPreTranslation(context.Background(), 1, req)
	require.NoError(t, err)
}

func TestTranslationsService_ApplyPreTranslation_WithValidateError(t *testing.T) {
	cases := []struct {
		req           *model.PreTranslationRequest
		expectedError string
	}{
		{
			req:           nil,
			expectedError: "request cannot be nil",
		},
		{
			req:           &model.PreTranslationRequest{},
			expectedError: "languageIds is required",
		},
		{
			req:           &model.PreTranslationRequest{LanguageIDs: []string{"uk"}},
			expectedError: "fileIds is required",
		},
	}
	for _, tt := range cases {
		assert.EqualError(t, tt.req.Validate(), tt.expectedError)
	}
}

func TestTranslationsService_BuildProjectDirectoryTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/builds/directories/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)

		expectedReqBody := `{
			"targetLanguageIds": ["uk"],
			"skipUntranslatedStrings": false,
			"skipUntranslatedFiles": false,
			"exportApprovedOnly": false,
			"exportWithMinApprovalsCount": 0,
			"exportStringsThatPassedWorkflow": true,
			"preserveFolderHierarchy": false
		}`
		testJSONBody(t, r, expectedReqBody)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"projectId": 2,
				"status": "finished",
				"progress": 100,
				"createdAt": "2023-09-19T15:10:43+00:00",
				"updatedAt": "2023-09-19T15:10:46+00:00",
				"finishedAt": "2023-09-19T15:10:46+00:00"
			}
		}`)
	})

	req := &model.BuildProjectDirectoryTranslationRequest{
		TargetLanguageIDs:               []string{"uk"},
		SkipUntranslatedStrings:         ToPtr(false),
		SkipUntranslatedFiles:           ToPtr(false),
		ExportApprovedOnly:              ToPtr(false),
		ExportWithMinApprovalsCount:     ToPtr(0),
		ExportStringsThatPassedWorkflow: ToPtr(true),
		PreserveFolderHierarchy:         ToPtr(false),
	}
	buildTranslation, resp, err := client.Translations.BuildProjectDirectoryTranslation(context.Background(), 1, 2, req)
	require.NoError(t, err)

	expected := &model.BuildProjectDirectoryTranslation{
		ID:         2,
		ProjectID:  2,
		Status:     "finished",
		Progress:   100,
		CreatedAt:  "2023-09-19T15:10:43+00:00",
		UpdatedAt:  "2023-09-19T15:10:46+00:00",
		FinishedAt: "2023-09-19T15:10:46+00:00",
	}
	assert.Equal(t, expected, buildTranslation)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestTranslationsService_BuildProjectDirectoryTranslation_WithValidationError(t *testing.T) {
	cases := []struct {
		req           *model.BuildProjectDirectoryTranslationRequest
		expectedError string
	}{
		{
			req:           nil,
			expectedError: "request cannot be nil",
		},
		{
			req: &model.BuildProjectDirectoryTranslationRequest{
				SkipUntranslatedStrings: ToPtr(true),
				SkipUntranslatedFiles:   ToPtr(true),
			},
			expectedError: "skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request",
		},
		{
			req: &model.BuildProjectDirectoryTranslationRequest{
				ExportWithMinApprovalsCount:     ToPtr(1),
				ExportStringsThatPassedWorkflow: ToPtr(true),
			},
			expectedError: "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request",
		},
	}
	for _, tt := range cases {
		assert.EqualError(t, tt.req.Validate(), tt.expectedError)
	}
}

func TestTranslationsService_BuildProjectFileTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const (
		etag = "ebd69a1b7e4c23e6d17891a491c94f832e0c82e4692dedb35a6cd1e624b62"
		path = "/api/v2/projects/1/translations/builds/files/2"
	)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testHeader(t, r, "If-None-Match", etag)

		expectedReqBody := `{
			"targetLanguageId": "uk",
			"skipUntranslatedStrings": true,
			"skipUntranslatedFiles": false,
			"exportApprovedOnly": false,
			"exportWithMinApprovalsCount": 0,
			"exportStringsThatPassedWorkflow": true
		}`
		testJSONBody(t, r, expectedReqBody)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"data": {
				"url": "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition=attachment",
				"expireIn": "2023-09-20T10:31:21+00:00",
				"etag": "ebd69a1b7e4c23e6d17891a491c94f832e0c82e4692dedb35a6cd1e624b62"
			}
		}`)
	})

	req := &model.BuildProjectFileTranslationRequest{
		TargetLanguageID:                "uk",
		SkipUntranslatedStrings:         ToPtr(true),
		SkipUntranslatedFiles:           ToPtr(false),
		ExportApprovedOnly:              ToPtr(false),
		ExportWithMinApprovalsCount:     ToPtr(0),
		ExportStringsThatPassedWorkflow: ToPtr(true),
	}
	downloadLink, resp, err := client.Translations.BuildProjectFileTranslation(context.Background(), 1, 2, req, etag)
	require.NoError(t, err)

	expected := &model.DownloadLink{
		URL:      "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition=attachment",
		ExpireIn: "2023-09-20T10:31:21+00:00",
		Etag:     ToPtr(etag),
	}
	assert.Equal(t, expected, downloadLink)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_BuildProjectFileTranslation_WithRequiredFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/translations/builds/files/2", func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"targetLanguageId":"uk"}`+"\n")
		testHeader(t, r, "If-None-Match", "")

		fmt.Fprint(w, `{}`)
	})

	req := &model.BuildProjectFileTranslationRequest{TargetLanguageID: "uk"}
	_, _, err := client.Translations.BuildProjectFileTranslation(context.Background(), 1, 2, req, "")
	require.NoError(t, err)
}

func TestTranslationsService_BuildProjectFileTranslation_WithValidationError(t *testing.T) {
	cases := []struct {
		req           *model.BuildProjectFileTranslationRequest
		expectedError string
	}{
		{
			req:           nil,
			expectedError: "request cannot be nil",
		},
		{
			req:           &model.BuildProjectFileTranslationRequest{},
			expectedError: "targetLanguageId is required",
		},
		{
			req: &model.BuildProjectFileTranslationRequest{
				TargetLanguageID:        "uk",
				SkipUntranslatedStrings: ToPtr(true),
				SkipUntranslatedFiles:   ToPtr(true),
			},
			expectedError: "skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request",
		},
		{
			req: &model.BuildProjectFileTranslationRequest{
				TargetLanguageID:                "uk",
				ExportWithMinApprovalsCount:     ToPtr(1),
				ExportStringsThatPassedWorkflow: ToPtr(true),
			},
			expectedError: "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request",
		},
	}
	for _, tt := range cases {
		assert.EqualError(t, tt.req.Validate(), tt.expectedError)
	}
}

func TestTranslationsService_ListProjectBuilds(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	cases := []struct {
		name     string
		opts     *model.TranslationsBuildsListOptions
		expected string
	}{
		{
			name:     "nil options",
			opts:     nil,
			expected: "",
		},
		{
			name:     "empty options",
			opts:     &model.TranslationsBuildsListOptions{},
			expected: "",
		},
		{
			name: "with options",
			opts: &model.TranslationsBuildsListOptions{
				BranchID:    1,
				ListOptions: model.ListOptions{Limit: 10, Offset: 5},
			},
			expected: "?branchId=1&limit=10&offset=5",
		},
	}

	for projectID, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v2/projects/%d/translations/builds", projectID)
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expected)

				fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": 2,
							"projectId": 2,
							"status": "finished",
							"progress": 100,
							"createdAt": "2023-09-19T15:10:43+00:00",
							"updatedAt": "2023-09-19T15:10:46+00:00",
							"finishedAt": "2023-09-19T15:10:46+00:00",
							"attributes": {
								"branchId": 1,
								"targetLanguageIds": ["en"],
								"skipUntranslatedStrings": false,
								"skipUntranslatedFiles": false,
								"exportWithMinApprovalsCount": 0,
								"exportStringsThatPassedWorkflow": true
							}
						}
					}
				],
				"pagination": {
					"offset": 5,
					"limit": 10
				}
			}`)
			})

			builds, resp, err := client.Translations.ListProjectBuilds(context.Background(), projectID, tt.opts)
			require.NoError(t, err)

			expected := []*model.TranslationsProjectBuild{
				{
					ID:         2,
					ProjectID:  2,
					Status:     "finished",
					Progress:   100,
					CreatedAt:  "2023-09-19T15:10:43+00:00",
					UpdatedAt:  "2023-09-19T15:10:46+00:00",
					FinishedAt: "2023-09-19T15:10:46+00:00",
					Attributes: &model.BuildAttributes{
						BranchID:                        ToPtr(1),
						TargetLanguageIDs:               []string{"en"},
						SkipUntranslatedStrings:         ToPtr(false),
						SkipUntranslatedFiles:           ToPtr(false),
						ExportWithMinApprovalsCount:     ToPtr(0),
						ExportStringsThatPassedWorkflow: ToPtr(true),
					},
				},
			}
			assert.Equal(t, expected, builds)

			expectedPagination := model.Pagination{Offset: 5, Limit: 10}
			assert.Equal(t, expectedPagination, resp.Pagination)
		})
	}
}

func TestTranslationsService_BuildProjectTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	cases := []struct {
		name     string
		req      model.BuildProjectTranslationRequest
		expected string
	}{
		{
			name: "CrowdinTranslationCreateProjectBuildForm",
			req: &model.BuildProjectRequest{
				BranchID:                1,
				TargetLanguageIDs:       []string{"uk"},
				SkipUntranslatedStrings: ToPtr(true),
				SkipUntranslatedFiles:   ToPtr(false),
				ExportApprovedOnly:      ToPtr(false),

				ExportWithMinApprovalsCount:     ToPtr(0),
				ExportStringsThatPassedWorkflow: ToPtr(true),
			},
			expected: `{
				"branchId": 1,
				"targetLanguageIds": ["uk"],
				"skipUntranslatedStrings": true,
				"skipUntranslatedFiles": false,
				"exportApprovedOnly": false,
				"exportWithMinApprovalsCount": 0,
				"exportStringsThatPassedWorkflow": true
			}`,
		},
		{
			name: "TranslationCreateProjectPseudoBuildForm",
			req: &model.PseudoBuildProjectRequest{
				Pseudo:               ToPtr(true),
				BranchID:             1,
				Prefix:               "pseudo",
				Suffix:               "pseudo",
				LengthTransformation: ToPtr(0),
				CharTransformation:   "european",
			},
			expected: `{
				"pseudo": true,
				"branchId": 1,
				"prefix": "pseudo",
				"suffix": "pseudo",
				"lengthTransformation": 0,
				"charTransformation": "european"
			}`,
		},
	}

	for projectID, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v2/projects/%d/translations/builds", projectID)
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, path)
				testJSONBody(t, r, tt.expected)

				w.WriteHeader(http.StatusCreated)
				fmt.Fprint(w, `{
					"data": {
						"id": 2,
						"projectId": 2,
						"status": "finished",
						"progress": 100,
						"createdAt": "2023-09-19T15:10:43+00:00",
						"updatedAt": "2023-09-19T15:10:46+00:00",
						"finishedAt": "2023-09-19T15:10:46+00:00",
						"attributes": {
							"branchId": 1,
							"targetLanguageIds": ["en"],
							"skipUntranslatedStrings": false,
							"skipUntranslatedFiles": false,
							"exportWithMinApprovalsCount": 0,
							"exportStringsThatPassedWorkflow": true
						}
					}
				}`)
			})

			build, resp, err := client.Translations.BuildProjectTranslation(context.Background(), projectID, tt.req)
			require.NoError(t, err)

			expected := &model.TranslationsProjectBuild{
				ID:         2,
				ProjectID:  2,
				Status:     "finished",
				Progress:   100,
				CreatedAt:  "2023-09-19T15:10:43+00:00",
				UpdatedAt:  "2023-09-19T15:10:46+00:00",
				FinishedAt: "2023-09-19T15:10:46+00:00",

				Attributes: &model.BuildAttributes{
					BranchID:                        ToPtr(1),
					TargetLanguageIDs:               []string{"en"},
					SkipUntranslatedStrings:         ToPtr(false),
					SkipUntranslatedFiles:           ToPtr(false),
					ExportWithMinApprovalsCount:     ToPtr(0),
					ExportStringsThatPassedWorkflow: ToPtr(true),
				},
			}
			assert.Equal(t, expected, build)
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
		})
	}
}

func TestTranslationsService_BuildProjectTranslation_WithValidationError(t *testing.T) {
	cases := []struct {
		name          string
		req           model.BuildProjectTranslationRequest
		expectedError string
	}{
		{
			name:          "nil request",
			req:           nil,
			expectedError: "body cannot be nil",
		},
		{
			name: "skipUntranslatedStrings and skipUntranslatedFiles are true",
			req: &model.BuildProjectRequest{
				SkipUntranslatedStrings: ToPtr(true),
				SkipUntranslatedFiles:   ToPtr(true),
			},
			expectedError: "`skipUntranslatedStrings` and `skipUntranslatedFiles` must not be true at the same request",
		},
		{
			name: "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow are true",
			req: &model.BuildProjectRequest{
				ExportWithMinApprovalsCount:     ToPtr(1),
				ExportStringsThatPassedWorkflow: ToPtr(true),
			},
			expectedError: "`exportWithMinApprovalsCount` and `exportStringsThatPassedWorkflow` must not be true at the same request",
		},
		{
			name: "lengthTransformation is out of range",
			req: &model.PseudoBuildProjectRequest{
				LengthTransformation: ToPtr(1000),
			},
			expectedError: "lengthTransformation must be from -50 to 100",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for projectID, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v2/projects/%d/translations/builds", projectID)
			mux.HandleFunc(path, func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, `{}`)
			})

			_, _, err := client.Translations.BuildProjectTranslation(context.Background(), projectID, tt.req)
			assert.EqualError(t, err, tt.expectedError)
		})
	}
}

func TestTranslationsService_BuildProjectTranslation_WithNilRequest(t *testing.T) {
	var req *model.BuildProjectRequest
	assert.EqualError(t, req.Validate(), "request cannot be nil")

	var reqPseudo *model.PseudoBuildProjectRequest
	assert.EqualError(t, reqPseudo.Validate(), "request cannot be nil")
}

func TestTranslationsService_UploadTranslations(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/uk"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"storageId":34,"fileId":56,"importEqSuggestions":true,"autoApproveImported":false,"translateHidden":false}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"projectId": 8,
				"storageId": 34,
				"languageId": "uk",
				"fileId": 56
			}
		}`)
	})

	req := &model.UploadTranslationsRequest{
		StorageID:           34,
		FileID:              56,
		ImportEqSuggestions: ToPtr(true),
		AutoApproveImported: ToPtr(false),
		TranslateHidden:     ToPtr(false),
	}
	uploadTranslations, resp, err := client.Translations.UploadTranslations(context.Background(), 1, "uk", req)
	require.NoError(t, err)

	expected := &model.UploadTranslations{
		ProjectID:  8,
		StorageID:  34,
		LanguageID: "uk",
		FileID:     56,
	}
	assert.Equal(t, expected, uploadTranslations)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_UploadTranslations_WithValidationError(t *testing.T) {
	cases := []struct {
		req           *model.UploadTranslationsRequest
		expectedError string
	}{
		{
			req:           nil,
			expectedError: "request cannot be nil",
		},
		{
			req:           &model.UploadTranslationsRequest{},
			expectedError: "storageId is required",
		},
		{
			req:           &model.UploadTranslationsRequest{StorageID: 1, FileID: 2, BranchID: 3},
			expectedError: "fileId and branchId can not be used at the same request",
		},
	}
	for _, tt := range cases {
		assert.EqualError(t, tt.req.Validate(), tt.expectedError)
	}
}

func TestTranslationsService_DownloadProjectTranslations(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/builds/2/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition=attachment",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	downloadLink, resp, err := client.Translations.DownloadProjectTranslations(context.Background(), 1, 2)
	require.NoError(t, err)

	expected := &model.DownloadLink{
		URL:      "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition=attachment",
		ExpireIn: "2023-09-20T10:31:21+00:00",
	}
	assert.Equal(t, expected, downloadLink)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_CheckBuildStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/builds/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"projectId": 2,
				"status": "finished",
				"progress": 100,
				"createdAt": "2023-09-19T15:10:43+00:00",
				"updatedAt": "2023-09-19T15:10:46+00:00",
				"finishedAt": "2023-09-19T15:10:46+00:00",
				"attributes": {
					"branchId": 1,
					"targetLanguageIds": ["en"],
					"skipUntranslatedStrings": false,
					"skipUntranslatedFiles": false,
					"exportWithMinApprovalsCount": 0,
					"exportStringsThatPassedWorkflow": true
				}
			}
		}`)
	})

	build, resp, err := client.Translations.CheckBuildStatus(context.Background(), 1, 2)
	require.NoError(t, err)

	expected := &model.TranslationsProjectBuild{
		ID:         2,
		ProjectID:  2,
		Status:     "finished",
		Progress:   100,
		CreatedAt:  "2023-09-19T15:10:43+00:00",
		UpdatedAt:  "2023-09-19T15:10:46+00:00",
		FinishedAt: "2023-09-19T15:10:46+00:00",
		Attributes: &model.BuildAttributes{
			BranchID:                        ToPtr(1),
			TargetLanguageIDs:               []string{"en"},
			SkipUntranslatedStrings:         ToPtr(false),
			SkipUntranslatedFiles:           ToPtr(false),
			ExportWithMinApprovalsCount:     ToPtr(0),
			ExportStringsThatPassedWorkflow: ToPtr(true),
		},
	}
	assert.Equal(t, expected, build)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_CancelBuild(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/builds/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Translations.CancelBuild(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTranslationsService_ExportProjectTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/exports"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)

		expectedReqBody := `{
			"targetLanguageId": "uk",
			"format": "xliff",
			"labelIds": [1],
			"branchIds": [2],
			"directoryIds": [3],
			"fileIds": [4],
			"skipUntranslatedStrings": false,
			"skipUntranslatedFiles": false,
			"exportApprovedOnly": false
		}`
		testJSONBody(t, r, expectedReqBody)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"data": {
				"url": "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	req := &model.ExportTranslationRequest{
		TargetLanguageID:        "uk",
		Format:                  "xliff",
		LabelIDs:                []int{1},
		BranchIDs:               []int{2},
		DirectoryIDs:            []int{3},
		FileIDs:                 []int{4},
		SkipUntranslatedStrings: ToPtr(false),
		SkipUntranslatedFiles:   ToPtr(false),
		ExportApprovedOnly:      ToPtr(false),
	}
	downloadLink, resp, err := client.Translations.ExportProjectTranslation(context.Background(), 1, req)
	require.NoError(t, err)

	expected := &model.DownloadLink{
		URL:      "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff?response-content-disposition",
		ExpireIn: "2023-09-20T10:31:21+00:00",
	}
	assert.Equal(t, expected, downloadLink)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTranslationsService_ExportProjectTranslation_WithValidationError(t *testing.T) {
	cases := []struct {
		req           *model.ExportTranslationRequest
		expectedError string
	}{
		{
			req:           nil,
			expectedError: "request cannot be nil",
		},
		{
			req:           &model.ExportTranslationRequest{},
			expectedError: "targetLanguageId is required",
		},
	}
	for _, tt := range cases {
		assert.EqualError(t, tt.req.Validate(), tt.expectedError)
	}
}
