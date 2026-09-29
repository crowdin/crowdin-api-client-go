package model

type (
	// CustomSpellchecker represents a custom spellchecker of the organization.
	CustomSpellchecker struct {
		ID        int                       `json:"id"`
		Name      string                    `json:"name"`
		Config    *CustomSpellcheckerConfig `json:"config,omitempty"`
		CreatedAt string                    `json:"createdAt"`
		UpdatedAt *string                   `json:"updatedAt,omitempty"`
	}

	// CustomSpellcheckerConfig represents the configuration of a custom spellchecker.
	CustomSpellcheckerConfig struct {
		// Application identifier.
		Identifier string `json:"identifier"`
		// Application module key.
		Key                  string   `json:"key"`
		RealTimeCheckEnabled bool     `json:"realTimeCheckEnabled"`
		EnabledLanguageIDs   []string `json:"enabledLanguageIds"`
	}
)

// CustomSpellcheckerResponse defines the structure of a response when
// getting a custom spellchecker.
type CustomSpellcheckerResponse struct {
	Data *CustomSpellchecker `json:"data"`
}

// CustomSpellcheckersListResponse defines the structure of a response when
// getting a list of custom spellcheckers.
type CustomSpellcheckersListResponse struct {
	Data []*CustomSpellcheckerResponse `json:"data"`
}
