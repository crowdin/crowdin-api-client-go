package model

import (
	"errors"
	"fmt"
	"net/url"
)

type (
	// ProjectMember represents a member of a project.
	ProjectMember struct {
		ID             int               `json:"id"`
		Username       string            `json:"username"`
		FirstName      *string           `json:"firstName,omitempty"`
		LastName       *string           `json:"lastName,omitempty"`
		FullName       *string           `json:"fullName,omitempty"`
		Role           *string           `json:"role,omitempty"`
		Permissions    map[string]any    `json:"permissions,omitempty"`
		Roles          []*TranslatorRole `json:"roles"`
		IsManager      *bool             `json:"isManager,omitempty"`
		IsDeveloper    *bool             `json:"isDeveloper,omitempty"`
		ManagerOfGroup *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"managerOfGroup,omitempty"`
		AccessToAllWorkflowSteps *bool   `json:"accessToAllWorkflowSteps,omitempty"`
		GivenAccessAt            *string `json:"givenAccessAt,omitempty"`
		AvatarURL                *string `json:"avatarUrl,omitempty"`
		JoinedAt                 *string `json:"joinedAt,omitempty"`
		Timezone                 *string `json:"timezone,omitempty"`
	}

	// TranslatorRole represents a role of a translator.
	TranslatorRole struct {
		// Roles name.
		// Enum: translator, proofreader, language_coordinator.
		Name Role `json:"name,omitempty"`
		// Role permission configuration.
		Permissions *RolePermissions `json:"permissions,omitempty"`
	}

	// RolePermissions represents permissions of a role.
	RolePermissions struct {
		// Access to all languages and workflow steps.
		// Default: true - means that user will have access to all
		// languages and workflow steps.
		AllLanguages *bool `json:"allLanguages,omitempty"`
		// Access to specific languages.
		// Needed if `allLanguages` is set to false.
		LanguagesAccess map[string]*LanguageAccess `json:"languagesAccess,omitempty"`
	}

	// LanguageAccess represents access to a language.
	LanguageAccess struct {
		// Access to all workflow steps for given language.
		// Default: false.
		AllContent *bool `json:"allContent,omitempty"`
		// Workflow Step Identifiers.
		WorkflowStepIDs []int `json:"workflowStepIds,omitempty"`
	}
)

// ProjectMemberResponse defines the structure of the response
// when getting a single project member.
type ProjectMemberResponse struct {
	Data *ProjectMember `json:"data"`
}

// ProjectMembersListResponse defines the structure of the response
// when getting a list of project members.
type ProjectMembersListResponse struct {
	Data []*ProjectMemberResponse `json:"data"`
}

// ProjectMembersListOptions specifies the optional parameters to the
// ProjectMembersService.List method.
type ProjectMembersListOptions struct {
	// Sort project members by a specific field.
	// Enum: id, username, firstName, lastName, fullName. Default: id.
	// Example: orderBy=lastName desc,username
	OrderBy string `json:"orderBy,omitempty"`
	// Search users by `firstName`, `lastName` or `username`.
	Search string `json:"search,omitempty"`
	// Defines role type. Enum: all, manager, developer, language_coordinator,
	// proofreader, translator, blocked, pending.
	Role string `json:"role,omitempty"`
	// Language Identifier.
	LanguageID string `json:"languageId,omitempty"`
	// Workflow Step Identifier.
	WorkflowStepID int `json:"workflowStepId,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of ProjectMembersListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ProjectMembersListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.Search != "" {
		v.Add("search", o.Search)
	}
	if o.Role != "" {
		v.Add("role", o.Role)
	}
	if o.LanguageID != "" {
		v.Add("languageId", o.LanguageID)
	}
	if o.WorkflowStepID > 0 {
		v.Add("workflowStepId", fmt.Sprintf("%d", o.WorkflowStepID))
	}

	return v, len(v) > 0
}

// Role represents a type of a role in the system.
type Role string

const (
	RoleTranslator          Role = "translator"
	RoleProofreader         Role = "proofreader"
	RoleLanguageCoordinator Role = "language_coordinator"
)

// ProjectMemberDeleteResponse defines the structure of the response
// when adding a new member to a project.
type ProjectMemberAddResponse struct {
	Skipped []*ProjectMemberResponse `json:"skipped"`
	Added   []*ProjectMemberResponse `json:"added"`
}

// ProjectMemberAddRequest defines the structure of the request
// when adding a new member to a project.
type ProjectMemberAddRequest struct {
	// User Identifier.
	// Note: One of fields `userIds`, `usernames` or `emails` is required.
	UserIDs []int `json:"userIds,omitempty"`
	// User Names.
	// Note: One of fields `userIds`, `usernames` or `emails` is required.
	Usernames []string `json:"usernames,omitempty"`
	// User Emails.
	// Note: One of fields `userIds`, `usernames` or `emails` is required.
	Emails []string `json:"emails,omitempty"`
	// Grand manager access to a project. Default: false.
	ManagerAccess *bool `json:"managerAccess,omitempty"`
	// Grand developer access to a project. Default: false.
	DeveloperAccess *bool `json:"developerAccess,omitempty"`
	// Note: `managerAccess`, `developerAccess` and `roles` parameters
	// are mutually exclusive.
	Roles []*TranslatorRole `json:"roles,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.Validator interface.
func (r *ProjectMemberAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if len(r.UserIDs) == 0 && len(r.Usernames) == 0 && len(r.Emails) == 0 {
		return fmt.Errorf("one of fields `userIds`, `usernames` or `emails` is required")
	}

	return nil
}

// ProjectMemberReplaceRequest defines the structure of the request
// when replacing permissions of a project member.
type ProjectMemberReplaceRequest struct {
	// Grand manager access to a project. Default: false.
	ManagerAccess *bool `json:"managerAccess,omitempty"`
	// Grand developer access to a project. Default: false.
	DeveloperAccess *bool `json:"developerAccess,omitempty"`
	// Translator roles.
	// Note: `managerAccess`, `developerAccess` and `roles` parameters
	// are mutually exclusive.
	Roles []*TranslatorRole `json:"roles,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.Validator interface.
func (r *ProjectMemberReplaceRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

// User represents a user in the system.
type User struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	FullName  *string `json:"fullName,omitempty"`
	Status    *string `json:"status,omitempty"` // Enum: active, pending, blocked
	AvatarURL string  `json:"avatarUrl"`
	CreatedAt string  `json:"createdAt"`
	LastSeen  string  `json:"lastSeen,omitempty"`
	TwoFactor string  `json:"twoFactor"` // Enum: enabled, disabled
	IsAdmin   *bool   `json:"isAdmin,omitempty"`
	Timezone  string  `json:"timezone,omitempty"`
	Fields    any     `json:"fields,omitempty"`
}

// ShortUser is a simplified version of the User model.
// It is used in the responses where only a few fields are required.
type ShortUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"avatarUrl"`
}

// UserResponse defines the structure of the response
// when getting a single user.
type UserResponse struct {
	Data *User `json:"data"`
}

// UsersListResponse defines the structure of the response
// when getting a list of users.
type UsersListResponse struct {
	Data []*UserResponse `json:"data"`
}

// UsersListOptions specifies the optional parameters to the
// UsersService.List method.
type UsersListOptions struct {
	// Sort users by a specific field.
	// Enum: id, username, firstName, lastName, email, status, createdAt, lastSeen.
	// Default: id.
	// Example: orderBy=createdAt desc,username
	OrderBy string `json:"orderBy,omitempty"`
	// Filter users by status.
	// Enum: active, pending, blocked.
	Status string `json:"status,omitempty"`
	// Search users by `firstName`, `lastName`, `username`, or `email`.
	Search string `json:"search,omitempty"`
	// Filter users by two-factor authentication status.
	// Enum: enabled, disabled.
	TwoFactor string `json:"twoFactor,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of UsersListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *UsersListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Search != "" {
		v.Add("search", o.Search)
	}
	if o.TwoFactor != "" {
		v.Add("twoFactor", o.TwoFactor)
	}

	return v, len(v) > 0
}

// InviteUserRequest defines the structure of the request
// when inviting a new user.
type InviteUserRequest struct {
	// Invited user email.
	Email string `json:"email"`
	// Invited user first name.
	FirstName string `json:"firstName,omitempty"`
	// Invited user last name.
	LastName string `json:"lastName,omitempty"`
	// Invited user timezone.
	Timezone string `json:"timezone,omitempty"`
	// Grant admin access to an organization.
	// Default: false.
	AdminAccess *bool `json:"adminAccess,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.Validator interface.
func (r *InviteUserRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
