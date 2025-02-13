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

	managers, resp, err := client.Managers.List(context.Background(), "2", nil)
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

	res, _, err := client.Groups.List(context.Background(), nil)
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
	managers, _, err := client.Managers.Edit(context.Background(), "1", req)
	if err != nil {
		t.Errorf("Managers.Edit returned error: %v", err)
	}

	want := 18

	if managers[0].User.ID != want {
		t.Errorf("Managers.Edit returned %+v, want %+v", managers[0].User.ID, want)
	}
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

	managers, _, err := client.Managers.Get(context.Background(), 1)
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
