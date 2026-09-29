package model

type (
	// OrganizationInfo represents the information about the organization.
	OrganizationInfo struct {
		ID                  int     `json:"id"`
		Domain              string  `json:"domain"`
		Name                string  `json:"name"`
		Logo                *string `json:"logo,omitempty"`
		DefaultLogo         string  `json:"defaultLogo"`
		Description         *string `json:"description,omitempty"`
		InternalDescription *string `json:"internalDescription,omitempty"`
		CName               *string `json:"cname,omitempty"`
		IsVendor            bool    `json:"isVendor"`
		// Default view of the public projects. Enum: grid, list.
		DefaultPublicProjectsView string                 `json:"defaultPublicProjectsView"`
		Plan                      *OrganizationPlan      `json:"plan,omitempty"`
		Defaults                  []*OrganizationDefault `json:"defaults,omitempty"`
	}

	// OrganizationPlan represents the subscription plan of the organization.
	OrganizationPlan struct {
		Name          string `json:"name"`
		WordsLimit    *int   `json:"wordsLimit,omitempty"`
		ManagersLimit *int   `json:"managersLimit,omitempty"`
	}

	// OrganizationDefault represents a default setting of the organization.
	OrganizationDefault struct {
		Name         string `json:"name"`
		IsLocked     bool   `json:"isLocked"`
		DefaultValue string `json:"defaultValue"`
	}
)

// OrganizationInfoResponse defines the structure of a response when
// getting the organization info.
type OrganizationInfoResponse struct {
	Data *OrganizationInfo `json:"data"`
}

type (
	// OrganizationAuthSettings represents the authentication settings of the organization.
	OrganizationAuthSettings struct {
		AllowSignUp             bool                      `json:"allowSignUp"`
		TwoFactorAuthentication bool                      `json:"twoFactorAuthentication"`
		AuthMethods             []*OrganizationAuthMethod `json:"authMethods"`
	}

	// OrganizationAuthMethod represents an authentication method of the organization.
	OrganizationAuthMethod struct {
		Name      string `json:"name"`
		IsEnabled bool   `json:"isEnabled"`
		IsDefault bool   `json:"isDefault"`
	}
)

// OrganizationAuthSettingsResponse defines the structure of a response when
// getting the organization authentication settings.
type OrganizationAuthSettingsResponse struct {
	Data *OrganizationAuthSettings `json:"data"`
}
