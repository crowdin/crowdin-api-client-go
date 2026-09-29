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

func TestOrganizationService_GetInfo(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/organization"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"domain": "crowdin",
				"name": "Crowdin",
				"logo": null,
				"defaultLogo": "",
				"description": "Localization platform",
				"internalDescription": null,
				"cname": null,
				"isVendor": false,
				"defaultPublicProjectsView": "grid",
				"plan": {
					"name": "Team+",
					"wordsLimit": 7500000,
					"managersLimit": null
				},
				"defaults": [
					{
						"name": "editorView",
						"isLocked": true,
						"defaultValue": "side-by-side"
					}
				]
			}
		}`)
	})

	info, resp, err := client.Organization.GetInfo(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.OrganizationInfo{
		ID:                        1,
		Domain:                    "crowdin",
		Name:                      "Crowdin",
		DefaultLogo:               "",
		Description:               ToPtr("Localization platform"),
		IsVendor:                  false,
		DefaultPublicProjectsView: "grid",
		Plan: &model.OrganizationPlan{
			Name:       "Team+",
			WordsLimit: ToPtr(7500000),
		},
		Defaults: []*model.OrganizationDefault{
			{
				Name:         "editorView",
				IsLocked:     true,
				DefaultValue: "side-by-side",
			},
		},
	}
	assert.Equal(t, expected, info)
}

func TestOrganizationService_GetInfo_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/organization", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Organization.GetInfo(context.Background())
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestOrganizationService_GetAuthSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/organization/auth-settings"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"allowSignUp": true,
				"twoFactorAuthentication": false,
				"authMethods": [
					{
						"name": "Password",
						"isEnabled": true,
						"isDefault": true
					},
					{
						"name": "SAML",
						"isEnabled": false,
						"isDefault": false
					}
				]
			}
		}`)
	})

	settings, resp, err := client.Organization.GetAuthSettings(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.OrganizationAuthSettings{
		AllowSignUp:             true,
		TwoFactorAuthentication: false,
		AuthMethods: []*model.OrganizationAuthMethod{
			{Name: "Password", IsEnabled: true, IsDefault: true},
			{Name: "SAML", IsEnabled: false, IsDefault: false},
		},
	}
	assert.Equal(t, expected, settings)
}
