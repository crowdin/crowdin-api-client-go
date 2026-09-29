package crowdin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const styleGuideJSON = `{
	"id": 2,
	"name": "Be My Eyes iOS's Style Guide",
	"aiInstructions": "Rules to be used by AI models",
	"userId": 2,
	"groupId": 5,
	"languageIds": ["uk", "fr", "de"],
	"projectIds": [6],
	"isShared": false,
	"webUrl": "https://example.crowdin.com/u/style-guides?id=1",
	"downloadLink": "https://storage.crowdin.com/style-guide/123/456/file.pdf",
	"createdAt": "2019-09-16T13:42:04+00:00",
	"updatedAt": "2019-09-16T13:42:04+00:00"
}`

func expectedStyleGuide() *model.StyleGuide {
	return &model.StyleGuide{
		ID:             2,
		Name:           "Be My Eyes iOS's Style Guide",
		AIInstructions: ToPtr("Rules to be used by AI models"),
		UserID:         2,
		GroupID:        5,
		LanguageIDs:    []string{"uk", "fr", "de"},
		ProjectIDs:     []int{6},
		IsShared:       false,
		WebURL:         "https://example.crowdin.com/u/style-guides?id=1",
		DownloadLink:   "https://storage.crowdin.com/style-guide/123/456/file.pdf",
		CreatedAt:      "2019-09-16T13:42:04+00:00",
		UpdatedAt:      "2019-09-16T13:42:04+00:00",
	}
}

func TestStyleGuidesService_List(t *testing.T) {
	tests := []struct {
		name     string
		opts     *model.StyleGuidesListOptions
		expected string
	}{
		{
			name:     "nil options",
			opts:     nil,
			expected: "",
		},
		{
			name:     "empty options",
			opts:     &model.StyleGuidesListOptions{},
			expected: "",
		},
		{
			name: "with options",
			opts: &model.StyleGuidesListOptions{
				OrderBy:     "createdAt desc",
				UserID:      3,
				ListOptions: model.ListOptions{Offset: 1, Limit: 25},
			},
			expected: "?limit=25&offset=1&orderBy=createdAt+desc&userId=3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/style-guides"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expected)

				fmt.Fprintf(w, `{
					"data": [{"data": %s}],
					"pagination": {"offset": 1, "limit": 25}
				}`, styleGuideJSON)
			})

			guides, resp, err := client.StyleGuides.List(context.Background(), tt.opts)
			require.NoError(t, err)

			assert.Equal(t, []*model.StyleGuide{expectedStyleGuide()}, guides)
			assert.Equal(t, 1, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestStyleGuidesService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/style-guides", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.StyleGuides.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestStyleGuidesService_Create(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"name": "Be My Eyes iOS's Style Guide",
			"storageId": 1,
			"groupId": 0,
			"aiInstructions": "Rules to be used by AI models",
			"languageIds": ["uk", "fr", "de"],
			"projectIds": [6],
			"isShared": false
		}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"data": %s}`, styleGuideJSON)
	})

	req := &model.StyleGuideCreateRequest{
		Name:           "Be My Eyes iOS's Style Guide",
		StorageID:      ToPtr(1),
		GroupID:        ToPtr(0),
		AIInstructions: "Rules to be used by AI models",
		LanguageIDs:    []string{"uk", "fr", "de"},
		ProjectIDs:     []int{6},
		IsShared:       ToPtr(false),
	}
	guide, resp, err := client.StyleGuides.Create(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, expectedStyleGuide(), guide)
}

func TestStyleGuidesService_Create_nullStorageID(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"name":"Guide","storageId":null,"aiInstructions":"Be concise"}`+"\n")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"data": %s}`, styleGuideJSON)
	})

	req := &model.StyleGuideCreateRequest{Name: "Guide", AIInstructions: "Be concise"}
	_, _, err := client.StyleGuides.Create(context.Background(), req)
	require.NoError(t, err)
}

func TestStyleGuidesService_Create_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.StyleGuides.Create(context.Background(), nil)
	require.ErrorIs(t, err, model.ErrNilRequest)

	_, _, err = client.StyleGuides.Create(context.Background(), &model.StyleGuideCreateRequest{})
	require.EqualError(t, err, "name is required")
}

func TestStyleGuidesService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprintf(w, `{"data": %s}`, styleGuideJSON)
	})

	guide, resp, err := client.StyleGuides.Get(context.Background(), 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedStyleGuide(), guide)
}

func TestStyleGuidesService_Get_notFound(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides/404"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"error": {"code": 404, "message": "Style Guide Not Found"}}`, http.StatusNotFound)
	})

	guide, resp, err := client.StyleGuides.Get(context.Background(), 404)
	require.Error(t, err)
	assert.Nil(t, guide)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var e *model.ErrorResponse
	assert.True(t, errors.As(err, &e))
}

func TestStyleGuidesService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testJSONBodyAny(t, r, `[
			{"op": "replace", "path": "/name", "value": "Be My Eyes iOS's Style Guide"},
			{"op": "replace", "path": "/isShared", "value": false}
		]`)

		fmt.Fprintf(w, `{"data": %s}`, styleGuideJSON)
	})

	req := []*model.UpdateRequest{
		{Op: model.OpReplace, Path: "/name", Value: "Be My Eyes iOS's Style Guide"},
		{Op: model.OpReplace, Path: "/isShared", Value: false},
	}
	guide, resp, err := client.StyleGuides.Edit(context.Background(), 2, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedStyleGuide(), guide)
}

func TestStyleGuidesService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/style-guides/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.StyleGuides.Delete(context.Background(), 2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
