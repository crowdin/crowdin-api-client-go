package crowdin

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTranslationStatusService_GetBranchProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/branches/2/languages/progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/projects/1/branches/2/languages/progress?limit=25&offset=1")

		fmt.Fprint(w, getJSONResponseMock())
	})

	opts := &model.ListOptions{
		Limit:  25,
		Offset: 1,
	}
	branchProgress, resp, err := client.TranslationStatus.GetBranchProgress(context.Background(), 1, 2, opts)
	if err != nil {
		t.Errorf("TranslationStatus.GetBranchProgress returned error: %v", err)
	}

	want := []*model.TranslationProgress{
		{
			Words: map[string]int{
				"approved":              3637,
				"preTranslateAppliedTo": 1254,
				"total":                 7249,
				"translated":            3651,
			},
			Phrases: map[string]int{
				"total":                 3041,
				"translated":            2631,
				"preTranslateAppliedTo": 1254,
				"approved":              2622,
			},
			TranslationProgress: 86,
			ApprovalProgress:    86,
			LanguageID:          ToPtr("es"),
			Language: &model.Language{
				ID:                  "es",
				Name:                "Spanish",
				EditorCode:          "es",
				TwoLettersCode:      "es",
				ThreeLettersCode:    "spa",
				Locale:              "es-ES",
				AndroidCode:         "es-rES",
				OSXCode:             "es.lproj",
				OSXLocale:           "es",
				PluralCategoryNames: []string{"one"},
				PluralRules:         "(n != 1)",
				PluralExamples:      []string{"0, 2-999; 1.2, 2.07..."},
				TextDirection:       "ltr",
				DialectOf:           "es",
			},
		},
	}
	expectedPagination := model.Pagination{Offset: 1, Limit: 25}

	if !reflect.DeepEqual(branchProgress, want) {
		t.Errorf("TranslationStatus.GetBranchProgress returned %+v, want %+v", branchProgress, want)
	}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.GetBranchProgress pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_GetDirectoryProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/directories/2/languages/progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/projects/1/directories/2/languages/progress?limit=25&offset=1")

		fmt.Fprint(w, getJSONResponseMock())
	})

	opts := &model.ListOptions{
		Limit:  25,
		Offset: 1,
	}
	dirProgress, resp, err := client.TranslationStatus.GetDirectoryProgress(context.Background(), 1, 2, opts)
	if err != nil {
		t.Errorf("TranslationStatus.GetDirectoryProgress returned error: %v", err)
	}

	want := []*model.TranslationProgress{
		{
			Words: map[string]int{
				"total":                 7249,
				"translated":            3651,
				"preTranslateAppliedTo": 1254,
				"approved":              3637,
			},
			Phrases: map[string]int{
				"total":                 3041,
				"translated":            2631,
				"preTranslateAppliedTo": 1254,
				"approved":              2622,
			},
			TranslationProgress: 86,
			ApprovalProgress:    86,
			LanguageID:          ToPtr("es"),
			Language: &model.Language{
				ID:                  "es",
				Name:                "Spanish",
				EditorCode:          "es",
				TwoLettersCode:      "es",
				ThreeLettersCode:    "spa",
				Locale:              "es-ES",
				AndroidCode:         "es-rES",
				OSXCode:             "es.lproj",
				OSXLocale:           "es",
				PluralCategoryNames: []string{"one"},
				PluralRules:         "(n != 1)",
				PluralExamples:      []string{"0, 2-999; 1.2, 2.07..."},
				TextDirection:       "ltr",
				DialectOf:           "es",
			},
		},
	}
	expectedPagination := model.Pagination{Offset: 1, Limit: 25}

	if !reflect.DeepEqual(dirProgress, want) {
		t.Errorf("TranslationStatus.GetDirectoryProgress returned %+v, want %+v", dirProgress, want)
	}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.GetDirectoryProgress pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_GetFileProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/files/2/languages/progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/projects/1/files/2/languages/progress?limit=25&offset=1")

		fmt.Fprint(w, getJSONResponseMock())
	})

	opts := &model.ListOptions{
		Limit:  25,
		Offset: 1,
	}
	fileProgress, resp, err := client.TranslationStatus.GetFileProgress(context.Background(), 1, 2, opts)
	if err != nil {
		t.Errorf("TranslationStatus.GetFileProgress returned error: %v", err)
	}

	want := []*model.TranslationProgress{
		{
			Words: map[string]int{
				"total":                 7249,
				"translated":            3651,
				"preTranslateAppliedTo": 1254,
				"approved":              3637,
			},
			Phrases: map[string]int{
				"total":                 3041,
				"translated":            2631,
				"preTranslateAppliedTo": 1254,
				"approved":              2622,
			},
			TranslationProgress: 86,
			ApprovalProgress:    86,
			LanguageID:          ToPtr("es"),
			Language: &model.Language{
				ID:                  "es",
				Name:                "Spanish",
				EditorCode:          "es",
				TwoLettersCode:      "es",
				ThreeLettersCode:    "spa",
				Locale:              "es-ES",
				AndroidCode:         "es-rES",
				OSXCode:             "es.lproj",
				OSXLocale:           "es",
				PluralCategoryNames: []string{"one"},
				PluralRules:         "(n != 1)",
				PluralExamples:      []string{"0, 2-999; 1.2, 2.07..."},
				TextDirection:       "ltr",
				DialectOf:           "es",
			},
		},
	}
	expectedPagination := model.Pagination{Offset: 1, Limit: 25}

	if !reflect.DeepEqual(fileProgress, want) {
		t.Errorf("TranslationStatus.GetFileProgress returned %+v, want %+v", fileProgress, want)
	}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.GetFileProgress pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_GetLanguageProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/languages/es/progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/projects/1/languages/es/progress?limit=25&offset=1")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"words": {
							"total": 7249,
							"translated": 3651,
							"preTranslateAppliedTo": 1254,
							"approved": 3637
						},
						"phrases": {
							"total": 3041,
							"translated": 2631,
							"preTranslateAppliedTo": 1254,
							"approved": 2622
						},
						"translationProgress": 86,
						"approvalProgress": 86,
						"fileId": 12,
						"etag": "fd0ea167420ef1687fd16635b9fb67a3"
					}
				}
			],
			"pagination": {
				"offset": 1,
				"limit": 25
			}
		}`)
	})

	opts := &model.ListOptions{
		Limit:  25,
		Offset: 1,
	}
	langProgress, resp, err := client.TranslationStatus.GetLanguageProgress(context.Background(), 1, "es", opts)
	if err != nil {
		t.Errorf("TranslationStatus.GetLanguageProgress returned error: %v", err)
	}

	want := []*model.TranslationProgress{
		{
			Words: map[string]int{
				"total":                 7249,
				"translated":            3651,
				"preTranslateAppliedTo": 1254,
				"approved":              3637,
			},
			Phrases: map[string]int{
				"total":                 3041,
				"translated":            2631,
				"preTranslateAppliedTo": 1254,
				"approved":              2622,
			},
			TranslationProgress: 86,
			ApprovalProgress:    86,
			FileID:              ToPtr(12),
			Etag:                ToPtr("fd0ea167420ef1687fd16635b9fb67a3"),
		},
	}
	expectedPagination := model.Pagination{Offset: 1, Limit: 25}

	if !reflect.DeepEqual(langProgress, want) {
		t.Errorf("TranslationStatus.GetLanguageProgress returned %+v, want %+v", langProgress, want)
	}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.GetLanguageProgress pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_GetProjectProgress(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/languages/progress", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v2/projects/1/languages/progress?languageIds=es%2Cde%2Cfr&limit=25&offset=1")

		fmt.Fprint(w, getJSONResponseMock())
	})

	opts := &model.ProjectProgressListOptions{
		ListOptions: model.ListOptions{
			Limit:  25,
			Offset: 1,
		},
		LanguageIDs: []string{"es", "de", "fr"},
	}
	projectProgress, resp, err := client.TranslationStatus.GetProjectProgress(context.Background(), 1, opts)
	if err != nil {
		t.Errorf("TranslationStatus.GetProjectProgress returned error: %v", err)
	}

	want := []*model.TranslationProgress{
		{
			Words: map[string]int{
				"total":                 7249,
				"translated":            3651,
				"preTranslateAppliedTo": 1254,
				"approved":              3637,
			},
			Phrases: map[string]int{
				"total":                 3041,
				"translated":            2631,
				"preTranslateAppliedTo": 1254,
				"approved":              2622,
			},
			TranslationProgress: 86,
			ApprovalProgress:    86,
			LanguageID:          ToPtr("es"),
			Language: &model.Language{
				ID:                  "es",
				Name:                "Spanish",
				EditorCode:          "es",
				TwoLettersCode:      "es",
				ThreeLettersCode:    "spa",
				Locale:              "es-ES",
				AndroidCode:         "es-rES",
				OSXCode:             "es.lproj",
				OSXLocale:           "es",
				PluralCategoryNames: []string{"one"},
				PluralRules:         "(n != 1)",
				PluralExamples:      []string{"0, 2-999; 1.2, 2.07..."},
				TextDirection:       "ltr",
				DialectOf:           "es",
			},
		},
	}
	expectedPagination := model.Pagination{Offset: 1, Limit: 25}

	if !reflect.DeepEqual(projectProgress, want) {
		t.Errorf("TranslationStatus.GetProjectProgress returned %+v, want %+v", projectProgress, want)
	}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.GetProjectProgress pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_GetProjectProgress_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/languages/progress", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.TranslationStatus.GetProjectProgress(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestTranslationStatusService_ListQAChecks(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	projectID := 1

	mux.HandleFunc(fmt.Sprintf("/api/v2/projects/%d/qa-checks", projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		q := "category=spellcheck&languageIds=uk%2Cde&limit=25&offset=1&validation=empty_string_check%2Cempty_suggestion_check"
		testURL(t, r, "/api/v2/projects/1/qa-checks?"+q)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"stringId": 1,
						"languageId": "uk",
						"category": "spellcheck",
						"categoryDescription": "Spelling",
						"validation": "spellcheck",
						"validationDescription": "Misspelling",
						"pluralId": -1,
						"text": "Spellcheck failed for the following word: 'локалзація'."
					}
				}
			],
			"pagination": {
				"offset": 1,
				"limit": 25
			}
		}`)
	})

	opts := &model.QACheckListOptions{
		ListOptions: model.ListOptions{Limit: 25, Offset: 1},
		Category:    []string{"spellcheck"},
		Validation:  []string{"empty_string_check,empty_suggestion_check"},
		LanguageIDs: []string{"uk", "de"},
	}
	issues, resp, err := client.TranslationStatus.ListQAChecks(context.Background(), projectID, opts)
	if err != nil {
		t.Errorf("TranslationStatus.ListQAChecks returned error: %v", err)
	}

	want := []*model.QACheck{
		{
			StringID:              1,
			LanguageID:            "uk",
			Category:              "spellcheck",
			CategoryDescription:   "Spelling",
			Validation:            "spellcheck",
			ValidationDescription: "Misspelling",
			PluralID:              -1,
			Text:                  "Spellcheck failed for the following word: 'локалзація'.",
		},
	}
	if !reflect.DeepEqual(issues, want) {
		t.Errorf("TranslationStatus.ListQAChecks returned %+v, want %+v", issues, want)
	}

	expectedPagination := model.Pagination{Offset: 1, Limit: 25}
	if !reflect.DeepEqual(resp.Pagination, expectedPagination) {
		t.Errorf("TranslationStatus.ListQAChecks pagination returned %+v, want %+v", resp.Pagination, expectedPagination)
	}
}

func TestTranslationStatusService_ListQAChecks_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/qa-checks", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.TranslationStatus.ListQAChecks(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestTranslationStatusService_GetFileProgress_QAChecksStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/files/2/languages/progress"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"words": {"total": 10, "translated": 5, "preTranslateAppliedTo": 2, "approved": 1},
						"phrases": {"total": 4, "translated": 2, "preTranslateAppliedTo": 1, "approved": 1},
						"translationProgress": 50,
						"approvalProgress": 10,
						"qaChecksStatus": {"total": 5, "inProgress": 1, "passed": 3, "failed": 1},
						"languageId": "uk",
						"eTag": "fd0ea167420ef1687fd16635b9fb67a3"
					}
				}
			],
			"pagination": {"offset": 0, "limit": 25}
		}`)
	})

	progress, _, err := client.TranslationStatus.GetFileProgress(context.Background(), 1, 2, nil)
	require.NoError(t, err)

	expected := []*model.TranslationProgress{
		{
			Words:               map[string]int{"total": 10, "translated": 5, "preTranslateAppliedTo": 2, "approved": 1},
			Phrases:             map[string]int{"total": 4, "translated": 2, "preTranslateAppliedTo": 1, "approved": 1},
			TranslationProgress: 50,
			ApprovalProgress:    10,
			QAChecksStatus:      &model.QAChecksStatus{Total: 5, InProgress: 1, Passed: 3, Failed: 1},
			LanguageID:          ToPtr("uk"),
			Etag:                ToPtr("fd0ea167420ef1687fd16635b9fb67a3"),
		},
	}
	assert.Equal(t, expected, progress)
}

func TestTranslationStatusService_ListQAChecks_WithTaskID(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/qa-checks"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?fileId=7&taskId=5")

		fmt.Fprint(w, `{"data": [], "pagination": {"offset": 0, "limit": 25}}`)
	})

	issues, _, err := client.TranslationStatus.ListQAChecks(context.Background(), 1, &model.QACheckListOptions{TaskID: 5, FileID: 7})
	require.NoError(t, err)
	assert.Empty(t, issues)
}

func TestTranslationStatusService_RevalidateQAChecks(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/qa-checks/revalidate"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"qaCheckCategories": ["ai"],
			"languageIds": ["uk", "fr"],
			"failedOnly": true,
			"externalQaCheckIds": [1, 2]
		}`)

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, `{
			"data": {
				"identifier": "b5215a34-1305-4b21-8054-fc2eb252842f",
				"status": "created",
				"progress": 0,
				"attributes": {
					"languageIds": ["uk", "fr"],
					"qaCheckCategories": ["ai"],
					"failedOnly": true,
					"externalQaCheckIds": [1, 2]
				},
				"createdAt": "2025-09-23T11:51:08+00:00",
				"updatedAt": "2025-09-23T11:51:08+00:00",
				"startedAt": null,
				"finishedAt": null
			}
		}`)
	})

	req := &model.QACheckRevalidationRequest{
		QACheckCategories:  []string{"ai"},
		LanguageIDs:        []string{"uk", "fr"},
		FailedOnly:         ToPtr(true),
		ExternalQACheckIDs: []int{1, 2},
	}
	revalidation, resp, err := client.TranslationStatus.RevalidateQAChecks(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	expected := &model.QACheckRevalidation{
		Identifier: "b5215a34-1305-4b21-8054-fc2eb252842f",
		Status:     "created",
		Progress:   0,
		Attributes: &model.QACheckRevalidationAttributes{
			LanguageIDs:        []string{"uk", "fr"},
			QACheckCategories:  []string{"ai"},
			FailedOnly:         true,
			ExternalQACheckIDs: []int{1, 2},
		},
		CreatedAt: "2025-09-23T11:51:08+00:00",
		UpdatedAt: "2025-09-23T11:51:08+00:00",
	}
	assert.Equal(t, expected, revalidation)
}

func TestTranslationStatusService_RevalidateQAChecks_NilRequest(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/qa-checks/revalidate"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, "{}\n")

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, `{"data": {"identifier": "b5215a34-1305-4b21-8054-fc2eb252842f", "status": "created"}}`)
	})

	revalidation, _, err := client.TranslationStatus.RevalidateQAChecks(context.Background(), 1, nil)
	require.NoError(t, err)
	assert.Equal(t, "b5215a34-1305-4b21-8054-fc2eb252842f", revalidation.Identifier)
}

func TestTranslationStatusService_GetQAChecksRevalidationStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/qa-checks/revalidate/b5215a34-1305-4b21-8054-fc2eb252842f"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "b5215a34-1305-4b21-8054-fc2eb252842f",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"languageIds": ["uk"],
					"qaCheckCategories": ["terms"],
					"failedOnly": false
				},
				"createdAt": "2025-09-23T11:51:08+00:00",
				"updatedAt": "2025-09-23T11:52:08+00:00",
				"startedAt": "2025-09-23T11:51:10+00:00",
				"finishedAt": "2025-09-23T11:52:08+00:00"
			}
		}`)
	})

	status, resp, err := client.TranslationStatus.GetQAChecksRevalidationStatus(context.Background(), 1, "b5215a34-1305-4b21-8054-fc2eb252842f")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expected := &model.QACheckRevalidation{
		Identifier: "b5215a34-1305-4b21-8054-fc2eb252842f",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.QACheckRevalidationAttributes{
			LanguageIDs:       []string{"uk"},
			QACheckCategories: []string{"terms"},
		},
		CreatedAt:  "2025-09-23T11:51:08+00:00",
		UpdatedAt:  "2025-09-23T11:52:08+00:00",
		StartedAt:  ToPtr("2025-09-23T11:51:10+00:00"),
		FinishedAt: ToPtr("2025-09-23T11:52:08+00:00"),
	}
	assert.Equal(t, expected, status)
}

func TestTranslationStatusService_CancelQAChecksRevalidation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/qa-checks/revalidate/b5215a34-1305-4b21-8054-fc2eb252842f"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.TranslationStatus.CancelQAChecksRevalidation(context.Background(), 1, "b5215a34-1305-4b21-8054-fc2eb252842f")
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTranslationStatusService_ValidateQAChecks(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/translations/validate-qa-checks"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBodyAny(t, r, `[
			{"stringId": 1, "languageId": "uk", "text": "Перша локалзація."},
			{"stringId": 2, "languageId": "uk", "text": "Рядки", "pluralCategoryName": "few"}
		]`)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"stringId": 1,
						"languageId": "uk",
						"category": "spellcheck",
						"categoryDescription": "Spelling",
						"validation": "spellcheck",
						"validationDescription": "Misspelling",
						"pluralId": -1,
						"pluralCategoryName": "none",
						"text": "Spellcheck failed for the following word: 'локалзація'.",
						"translation": "Перша локалзація."
					}
				}
			],
			"pagination": {"offset": 0, "limit": 25}
		}`)
	})

	req := []*model.QACheckValidateRequest{
		{StringID: 1, LanguageID: "uk", Text: "Перша локалзація."},
		{StringID: 2, LanguageID: "uk", Text: "Рядки", PluralCategoryName: "few"},
	}
	issues, resp, err := client.TranslationStatus.ValidateQAChecks(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expected := []*model.QACheckValidation{
		{
			StringID:              1,
			LanguageID:            "uk",
			Category:              "spellcheck",
			CategoryDescription:   "Spelling",
			Validation:            "spellcheck",
			ValidationDescription: "Misspelling",
			PluralID:              -1,
			PluralCategoryName:    "none",
			Text:                  "Spellcheck failed for the following word: 'локалзація'.",
			Translation:           "Перша локалзація.",
		},
	}
	assert.Equal(t, expected, issues)
}

func TestTranslationStatusService_ValidateQAChecks_ValidationError(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.TranslationStatus.ValidateQAChecks(context.Background(), 1, nil)
	require.EqualError(t, err, "request cannot be empty")

	req := []*model.QACheckValidateRequest{{StringID: 1, LanguageID: "uk"}}
	_, _, err = client.TranslationStatus.ValidateQAChecks(context.Background(), 1, req)
	require.EqualError(t, err, "text is required")
}

func TestTranslationStatusService_ValidateQAChecks_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/translations/validate-qa-checks", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	req := []*model.QACheckValidateRequest{{StringID: 1, LanguageID: "uk", Text: "Text"}}
	res, _, err := client.TranslationStatus.ValidateQAChecks(context.Background(), 1, req)
	require.Error(t, err)
	assert.Nil(t, res)
}

func getJSONResponseMock() string {
	return `{
		"data": [
			{
				"data": {
					"words": {
						"total": 7249,
						"translated": 3651,
						"preTranslateAppliedTo": 1254,
						"approved": 3637
					},
					"phrases": {
						"total": 3041,
						"translated": 2631,
						"preTranslateAppliedTo": 1254,
						"approved": 2622
					},
					"translationProgress": 86,
					"approvalProgress": 86,
					"languageId": "es",
					"language": {
						"id": "es",
						"name": "Spanish",
						"editorCode": "es",
						"twoLettersCode": "es",
						"threeLettersCode": "spa",
						"locale": "es-ES",
						"androidCode": "es-rES",
						"osxCode": "es.lproj",
						"osxLocale": "es",
						"pluralCategoryNames": [
							"one"
						],
						"pluralRules": "(n != 1)",
						"pluralExamples": [
							"0, 2-999; 1.2, 2.07..."
						],
						"textDirection": "ltr",
						"dialectOf": "es"
					}
				}
			}
		],
		"pagination": {
			"offset": 1,
			"limit": 25
		}
	}`
}
