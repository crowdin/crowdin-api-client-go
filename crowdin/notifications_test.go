package crowdin

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotificationsService_NotifyProjectMembers(t *testing.T) {
	tests := []struct {
		name string
		req  *model.Notification
		body string
	}{
		{
			name: "by user IDs",
			req: &model.Notification{
				UserIDs: []int{1, 2, 3},
				Message: "notification message",
			},
			body: `{"message":"notification message","userIds":[1,2,3]}` + "\n",
		},
		{
			name: "by role",
			req: &model.Notification{
				Role:    "owner",
				Message: "notification message",
			},
			body: `{"message":"notification message","role":"owner"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc("/api/v2/projects/2/notify", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, "/api/v2/projects/2/notify")
				testBody(t, r, tt.body)

				w.WriteHeader(http.StatusNoContent)
			})

			resp, err := client.Notifications.NotifyProjectMembers(context.Background(), 2, tt.req)
			require.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestNotificationsService_Notify(t *testing.T) {
	tests := []struct {
		name string
		req  *model.Notification
		body string
	}{
		{
			name: "authenticated user",
			req: &model.Notification{
				Message: "notification message",
			},
			body: `{"message":"notification message"}` + "\n",
		},
		{
			name: "organization members by user IDs",
			req: &model.Notification{
				UserIDs: []int{1, 2, 3},
				Message: "notification message",
			},
			body: `{"message":"notification message","userIds":[1,2,3]}` + "\n",
		},
		{
			name: "organization members by role",
			req: &model.Notification{
				Role:    "owner",
				Message: "notification message",
			},
			body: `{"message":"notification message","role":"owner"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc("/api/v2/notify", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, "/api/v2/notify")
				testBody(t, r, tt.body)

				w.WriteHeader(http.StatusNoContent)
			})

			resp, err := client.Notifications.Notify(context.Background(), tt.req)
			require.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestNotificationsService_Notify_validate(t *testing.T) {
	tests := []struct {
		name string
		req  *model.Notification
		err  string
	}{
		{
			name: "nil request",
			err:  "request cannot be nil",
		},
		{
			name: "message too long",
			req: &model.Notification{
				Message: strings.Repeat("a", 10001),
			},
			err: "message can't be longer than 10000 characters",
		},
		{
			name: "both user IDs and role",
			req: &model.Notification{
				UserIDs: []int{1, 2, 3},
				Role:    "owner",
				Message: "notification message",
			},
			err: "can't specify both user IDs and role",
		},
		{
			name: "invalid role",
			req: &model.Notification{
				Role:    "invalid",
				Message: "notification message",
			},
			err: "role must be either owner or manager",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			require.Error(t, err)
			assert.EqualError(t, err, tt.err)
		})
	}
}
