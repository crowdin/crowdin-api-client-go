package model

import (
	"fmt"
	"net/url"
)

// WorkflowStep represents a workflow step in a project.
type WorkflowStep struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Type      string         `json:"type"`
	Languages []string       `json:"languages"`
	Config    map[string]any `json:"config"`
}

// WorkflowStepResponse defines the structure of the response when
// getting a workflow step.
type WorkflowStepResponse struct {
	Data *WorkflowStep `json:"data"`
}

// WorkflowStepsResponse defines the structure of the response when
// getting a list of workflow steps.
type WorkflowStepsResponse struct {
	Data []*WorkflowStepResponse `json:"data"`
}

type (
	// WorkflowTemplate represents a workflow template in the organization.
	WorkflowTemplate struct {
		ID          int                     `json:"id"`
		Title       string                  `json:"title"`
		Description string                  `json:"description"`
		GroupID     int                     `json:"groupId"`
		IsDefault   bool                    `json:"isDefault"`
		Steps       []*WorkflowTemplateStep `json:"steps"`
		WebURL      string                  `json:"webUrl"`
	}

	// WorkflowTemplateStep represents a workflow step in a workflow template.
	WorkflowTemplateStep struct {
		// Workflow template step identifier.
		ID int `json:"id,omitempty"`
		// Target languages identifiers.
		Languages []int `json:"languages,omitempty"`
		// User identifiers.
		Assignees []int `json:"assignees,omitempty"`
		// Vendor identifier.
		VendorID int `json:"vendorId,omitempty"`
		// Workflow template step configuration.
		Config WorkflowTemplateStepConfig `json:"config,omitempty"`
		// Machine translation identifier.
		MTID int `json:"mtId,omitempty"`
	}

	// WorkflowTemplateStepConfig represents a workflow template step configuration.
	WorkflowTemplateStepConfig struct {
		// Minimum match for TM suggestions.
		MinRelevant *int `json:"minRelevant,omitempty"`
		// Improves TM suggestions.
		AutoSubstitution *bool `json:"autoSubstitution,omitempty"`
	}
)

// WorkflowTemplateResponse defines the structure of the response when
// getting a workflow template.
type WorkflowTemplateResponse struct {
	Data *WorkflowTemplate `json:"data"`
}

// WorkflowTemplatesListResponse defines the structure of the response when
// getting a list of workflow templates.
type WorkflowTemplatesListResponse struct {
	Data []*WorkflowTemplateResponse `json:"data"`
}

// WorkflowTemplatesListOptions specifies the optional parameters to the
// WorkflowsService.ListTemplates method.
type WorkflowTemplatesListOptions struct {
	// Group identifier.
	// Note: Set 0 to see items of root group.
	GroupID *int `json:"groupId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the options.
// It implements the crowdin.ListOptionsProvider interface.
func (o *WorkflowTemplatesListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.GroupID != nil {
		v.Add("groupId", fmt.Sprintf("%d", *o.GroupID))
	}

	return v, len(v) > 0
}
