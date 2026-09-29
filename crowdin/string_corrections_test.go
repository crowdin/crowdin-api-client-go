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

const stringCorrectionJSON = `{
	"id": 190695,
	"text": "This string has been corrected",
	"pluralCategoryName": "few",
	"user": {
		"id": 19,
		"username": "john_doe",
		"fullName": "John Smith",
		"avatarUrl": ""
	},
	"createdAt": "2019-09-23T11:26:54+00:00"
}`

func expectedStringCorrection() *model.StringCorrection {
	return &model.StringCorrection{
		ID:                 190695,
		Text:               "This string has been corrected",
		PluralCategoryName: "few",
		User: &model.ShortUser{
			ID:        19,
			Username:  "john_doe",
			FullName:  "John Smith",
			AvatarURL: "",
		},
		CreatedAt: "2019-09-23T11:26:54+00:00",
	}
}

func TestStringCorrectionsService_List(t *testing.T) {
	tests := []struct {
		name     string
		opts     *model.StringCorrectionsListOptions
		expected string
	}{
		{
			name:     "string ID only",
			opts:     &model.StringCorrectionsListOptions{StringID: 35434},
			expected: "?stringId=35434",
		},
		{
			name: "with options",
			opts: &model.StringCorrectionsListOptions{
				StringID:                35434,
				OrderBy:                 "createdAt desc",
				DenormalizePlaceholders: ToPtr(1),
				ListOptions:             model.ListOptions{Offset: 1, Limit: 25},
			},
			expected: "?denormalizePlaceholders=1&limit=25&offset=1&orderBy=createdAt+desc&stringId=35434",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/corrections"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expected)

				fmt.Fprintf(w, `{
					"data": [{"data": %s}],
					"pagination": {"offset": 1, "limit": 25}
				}`, stringCorrectionJSON)
			})

			corrections, resp, err := client.StringCorrections.List(context.Background(), 1, tt.opts)
			require.NoError(t, err)

			assert.Equal(t, []*model.StringCorrection{expectedStringCorrection()}, corrections)
			assert.Equal(t, 1, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestStringCorrectionsService_List_missingStringID(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	for _, opts := range []*model.StringCorrectionsListOptions{nil, {}} {
		res, resp, err := client.StringCorrections.List(context.Background(), 1, opts)
		require.EqualError(t, err, "stringId is required")
		assert.Nil(t, res)
		assert.Nil(t, resp)
	}
}

func TestStringCorrectionsService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/corrections", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	opts := &model.StringCorrectionsListOptions{StringID: 1}
	res, _, err := client.StringCorrections.List(context.Background(), 1, opts)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestStringCorrectionsService_Add(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/corrections"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"stringId": 35434,
			"text": "This string has been corrected",
			"pluralCategoryName": "few"
		}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"data": %s}`, stringCorrectionJSON)
	})

	req := &model.StringCorrectionAddRequest{
		StringID:           35434,
		Text:               "This string has been corrected",
		PluralCategoryName: "few",
	}
	correction, resp, err := client.StringCorrections.Add(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, expectedStringCorrection(), correction)
}

func TestStringCorrectionsService_Add_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.StringCorrections.Add(context.Background(), 1, nil)
	require.ErrorIs(t, err, model.ErrNilRequest)

	_, _, err = client.StringCorrections.Add(context.Background(), 1, &model.StringCorrectionAddRequest{StringID: 1})
	require.EqualError(t, err, "text is required")
}

func TestStringCorrectionsService_DeleteMany(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/corrections"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path+"?stringId=35434")

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.StringCorrections.DeleteMany(context.Background(), 1, 35434)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestStringCorrectionsService_Get(t *testing.T) {
	tests := []struct {
		name     string
		opts     *model.TranslationGetOptions
		expected string
	}{
		{
			name:     "nil options",
			opts:     nil,
			expected: "",
		},
		{
			name:     "with options",
			opts:     &model.TranslationGetOptions{DenormalizePlaceholders: ToPtr(1)},
			expected: "?denormalizePlaceholders=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/corrections/190695"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expected)

				fmt.Fprintf(w, `{"data": %s}`, stringCorrectionJSON)
			})

			correction, resp, err := client.StringCorrections.Get(context.Background(), 1, 190695, tt.opts)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, expectedStringCorrection(), correction)
		})
	}
}

func TestStringCorrectionsService_Get_notFound(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/corrections/404"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"error": {"code": 404, "message": "Correction Not Found"}}`, http.StatusNotFound)
	})

	correction, resp, err := client.StringCorrections.Get(context.Background(), 1, 404, nil)
	require.Error(t, err)
	assert.Nil(t, correction)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var e *model.ErrorResponse
	assert.True(t, errors.As(err, &e))
}

func TestStringCorrectionsService_Restore(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/corrections/190695"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testURL(t, r, path)
		testBody(t, r, "")

		fmt.Fprintf(w, `{"data": %s}`, stringCorrectionJSON)
	})

	correction, resp, err := client.StringCorrections.Restore(context.Background(), 1, 190695)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedStringCorrection(), correction)
}

func TestStringCorrectionsService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/corrections/190695"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.StringCorrections.Delete(context.Background(), 1, 190695)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
