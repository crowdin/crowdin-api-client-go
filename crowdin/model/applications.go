package model

import (
	"errors"
	"fmt"
	"net/url"
)

type PermissionValue string

const (
	PermissionOwn        PermissionValue = "own"        // All projects (Enterprise) / Own project
	PermissionOwner      PermissionValue = "owner"      // Only organization admins (Enterprise) / All project members
	PermissionManagers   PermissionValue = "managers"   // Organization admins, project managers and developers
	PermissionAll        PermissionValue = "all"        // All users in the organization projects
	PermissionGuests     PermissionValue = "guests"     // All users, including guests (unauthenticated users)
	PermissionRestricted PermissionValue = "restricted" // Selected projects
)

type (
	// Installation represents an application installation.
	Installation struct {
		Identifier         string            `json:"identifier"`
		Name               string            `json:"name"`
		Description        string            `json:"description"`
		Logo               string            `json:"logo"`
		BaseURL            string            `json:"baseUrl"`
		ManifestURL        string            `json:"manifestUrl"`
		CreatedAt          string            `json:"createdAt"`
		Modules            []*Module         `json:"modules"`
		Scopes             []string          `json:"scopes"`
		Permissions        ProjectPermission `json:"permissions"`
		DefaultPermissions struct {
			User    PermissionValue `json:"user"`
			Project PermissionValue `json:"project"`
		} `json:"defaultPermissions"`
		LimitReached bool `json:"limitReached"`

		// InstalledBy is the user who installed the application.
		InstalledBy *ShortUser `json:"installedBy,omitempty"`
		// LogoURL is the absolute URL the logo is served from.
		LogoURL *string `json:"logoUrl,omitempty"`
		// Agent is the agent user of the application.
		Agent *ShortUser `json:"agent,omitempty"`
		// Manifest is the manifest of the installation as authored.
		// It is nil for applications installed from a manifest URL.
		Manifest map[string]any `json:"manifest,omitempty"`
		// ManifestUpdatedAt is the date when the installation's manifest was last synced.
		ManifestUpdatedAt *string `json:"manifestUpdatedAt,omitempty"`
		// IsManifestOutdated reports whether the installed manifest is outdated
		// compared to the latest available manifest.
		IsManifestOutdated bool `json:"isManifestOutdated"`
		// StringBasedAvailable reports whether the app is offered in string-based projects.
		StringBasedAvailable bool `json:"stringBasedAvailable"`
		// Bundle describes where the app is loaded from.
		Bundle *ApplicationBundle `json:"bundle,omitempty"`
	}

	// Module represents an application module.
	Module struct {
		Key                string         `json:"key"`
		Type               string         `json:"type"`
		Data               any            `json:"data"`
		Permissions        UserPermission `json:"permissions"`
		AuthenticationType string         `json:"authenticationType"`

		// Integration module fields.
		Identifier *string `json:"identifier,omitempty"`
		Scopes     any     `json:"scopes,omitempty"`
		Iframe     any     `json:"iframe,omitempty"`

		// Module type specific fields. Each field is present only
		// for the module types that support it.
		Name                *string `json:"name,omitempty"`
		Description         *string `json:"description,omitempty"`
		URL                 *string `json:"url,omitempty"`
		Logo                *string `json:"logo,omitempty"`
		LogoURL             *string `json:"logoUrl,omitempty"`
		BaseURL             *string `json:"baseUrl,omitempty"`
		Modes               any     `json:"modes,omitempty"`
		Mode                *string `json:"mode,omitempty"`
		Options             any     `json:"options,omitempty"`
		SignaturePatterns   any     `json:"signaturePatterns,omitempty"`
		FileNamePattern     *string `json:"fileNamePattern,omitempty"`
		FileType            *string `json:"fileType,omitempty"`
		Multilingual        *bool   `json:"multilingual,omitempty"`
		Severity            *string `json:"severity,omitempty"`
		Category            *string `json:"category,omitempty"`
		RefreshPolicy       *string `json:"refreshPolicy,omitempty"`
		Title               *string `json:"title,omitempty"`
		Summary             *string `json:"summary,omitempty"`
		RecheckURL          *string `json:"recheckUrl,omitempty"`
		ConfiguratorIframe  any     `json:"configuratorIframe,omitempty"`
		CompileURL          *string `json:"compileUrl,omitempty"`
		ChatCompletionsURL  *string `json:"chatCompletionsUrl,omitempty"`
		ModelsURL           *string `json:"modelsUrl,omitempty"`
		Placement           *string `json:"placement,omitempty"`
		InputSchemaURL      *string `json:"inputSchemaUrl,omitempty"`
		OutputSchemaURL     *string `json:"outputSchemaUrl,omitempty"`
		ExecuteURL          *string `json:"executeUrl,omitempty"`
		ValidateSettingsURL *string `json:"validateSettingsUrl,omitempty"`
		InvocationWaitMode  *string `json:"invocationWaitMode,omitempty"`
		BatchSizeURL        *string `json:"batchSizeUrl,omitempty"`
		RunQaCheckURL       *string `json:"runQaCheckUrl,omitempty"`
		UpdateSettingsURL   *string `json:"updateSettingsUrl,omitempty"`
		DeleteSettingsURL   *string `json:"deleteSettingsUrl,omitempty"`
		EditorMode          *string `json:"editorMode,omitempty"`
		Boundaries          any     `json:"boundaries,omitempty"`
	}

	// ProjectPermission represents a permission for a project where
	// users will be able to use the app.
	ProjectPermission struct {
		// Value enum: own, restricted.
		Project Permission `json:"project,omitempty"`
	}

	// UserPermission represents a permission for a user that will
	// be able to use the app.
	UserPermission struct {
		// Value enum: owner, managers, all, guests, restricted.
		// Note: For exporters, the `all` value will be set.
		User Permission `json:"user,omitempty"`
	}

	// Permission represents a permission value for a project or a user.
	Permission struct {
		// Value of the permission.
		Value PermissionValue `json:"value,omitempty"`
		// IDs is only available for restricted value.
		IDs []int `json:"ids,omitempty"`
	}
)

// ApplicationBundle describes where the app is loaded from.
// Either Crowdin hosts the app in its own storage (`internal`)
// or the app is served from an external URL (`external`).
type ApplicationBundle struct {
	// Mode enum: internal, external.
	Mode string `json:"mode"`
	// URL the app is served from. Required for the `external` mode.
	URL string `json:"url,omitempty"`
}

// InstallationResponse defines the structure of the response
// when getting an installation.
type InstallationResponse struct {
	Data *Installation `json:"data"`
}

// InstallationsListResponse defines the structure of the response
// when getting a list of installations.
type InstallationsListResponse struct {
	Data []*InstallationResponse `json:"data"`
}

// InstallationsListOptions specifies the optional parameters to the
// ApplicationsService.ListInstallationsWithOptions method.
type InstallationsListOptions struct {
	// Filter installations by the identifier of the user who installed the application.
	InstalledBy int `json:"installedBy,omitempty"`
	// Sort installations by the specified fields.
	// Example: orderBy=createdAt desc,name.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of InstallationsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *InstallationsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.InstalledBy > 0 {
		v.Add("installedBy", fmt.Sprintf("%d", o.InstalledBy))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// InstallApplicationRequest defines the structure of the request
// to install an application. Either URL or Manifest must be set.
type InstallApplicationRequest struct {
	// Manifest URL of the application.
	URL string `json:"url,omitempty"`
	// Manifest is the inline manifest content of the application.
	// Supported only for serverless apps (apps without `baseUrl`).
	Manifest *ApplicationManifest `json:"manifest,omitempty"`
	// Permissions to set for the application.
	Permissions *ProjectPermission `json:"permissions,omitempty"`
	// Modules with permissions to set for the application.
	Modules []*InstallationModule `json:"modules,omitempty"`
	// Assign Agent as a manager to all existing projects.
	// Available only when installing from a manifest URL.
	AssignAgent *bool `json:"assignAgent,omitempty"`
}

// ApplicationManifest defines the inline manifest content used
// to install an application.
type ApplicationManifest struct {
	// Display name of the application.
	Name string `json:"name"`
	// Short description of what the application does.
	Description string `json:"description,omitempty"`
	// Relative URL (within the application bundle) of the application logo.
	Logo string `json:"logo,omitempty"`
	// OAuth scopes granted to the app for host-proxied Crowdin API calls.
	Scopes []string `json:"scopes,omitempty"`
	// Whether the app is offered in string-based projects.
	StringBasedAvailable *bool `json:"stringBasedAvailable,omitempty"`
	// UI module definitions: keys are UI module types (e.g. `editor-right-panel`)
	// and values are arrays of module definitions.
	Modules map[string][]*ApplicationManifestModule `json:"modules"`
	// Default permissions of the application.
	DefaultPermissions *ApplicationDefaultPermissions `json:"default_permissions,omitempty"`
	// Where the app is loaded from.
	Bundle *ApplicationBundle `json:"bundle,omitempty"`
}

// ApplicationManifestModule defines a UI module definition of an inline manifest.
type ApplicationManifestModule struct {
	// Unique module key within the app.
	Key string `json:"key"`
	// Display name of the module.
	Name string `json:"name"`
	// Relative URL (within the application bundle) of the module logo.
	Logo string `json:"logo,omitempty"`
	// Short description of the module.
	Description string `json:"description,omitempty"`
	// Crowdin editions the module is enabled in.
	// Enum: crowdin, crowdin-enterprise.
	Environments []string `json:"environments,omitempty"`
	// Module permissions.
	Permissions *UserPermission `json:"permissions,omitempty"`
}

// ApplicationDefaultPermissions defines the default permissions of an application.
type ApplicationDefaultPermissions struct {
	// Which users the app is available to by default.
	// Enum: owner, managers, all, guests.
	User PermissionValue `json:"user,omitempty"`
	// Which projects the app is available in by default.
	// Enum: own, restricted.
	Project PermissionValue `json:"project,omitempty"`
}

// InstallationReplaceValue represents the structure of the values to be replaced.
// Can be used to update permissions or module permissions with
// the replace operation in the ApplicationsService.EditInstallation method.
type InstallationReplaceValue struct {
	User Permission `json:"user,omitempty"`
	// Available only for application permissions.
	Project Permission `json:"project,omitempty"`
}

type InstallationModule struct {
	Key         string         `json:"key,omitempty"`
	Permissions UserPermission `json:"permissions,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *InstallApplicationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.URL == "" && r.Manifest == nil {
		return errors.New("url is required")
	}
	if r.URL != "" && r.Manifest != nil {
		return errors.New("url and manifest cannot be used together")
	}
	if r.Manifest != nil {
		if r.Manifest.Name == "" {
			return errors.New("manifest name is required")
		}
		if len(r.Manifest.Modules) == 0 {
			return errors.New("manifest modules are required")
		}
	}

	return nil
}

// ApplicationDataResponse defines the structure of the response
// with application data. The data field can contain any application-specific data.
type ApplicationDataResponse struct {
	Data any `json:"data"`
}

// InstallationUpdate represents a pending update (diff) between
// the installed application and the latest manifest.
type InstallationUpdate struct {
	// Hash of the latest manifest. Pass it back in the apply request.
	ManifestHash *string `json:"manifestHash"`
	// Full latest manifest payload.
	LatestManifest map[string]any `json:"latestManifest"`
	// Scopes added in the latest manifest.
	AddedScopes []string `json:"addedScopes"`
	// Scopes removed in the latest manifest.
	RemovedScopes []string `json:"removedScopes"`
	// Modules added in the latest manifest.
	AddedModules []map[string]any `json:"addedModules"`
	// Modules removed in the latest manifest.
	RemovedModules []map[string]any `json:"removedModules"`
	// Modules that still exist but whose behavior fields (url, permissions) changed.
	ChangedModules []map[string]any `json:"changedModules"`
	// Lifecycle event URLs that changed (installed/uninstall/status).
	ChangedEvents map[string]*InstallationUpdateChange `json:"changedEvents"`
	// The baseUrl change (host that serves the app).
	BaseURLChanged *InstallationUpdateChange `json:"baseUrlChanged"`
	// Authentication type change. It cannot be applied as an update,
	// the app must be uninstalled and reinstalled.
	AuthenticationTypeChanged *InstallationUpdateChange `json:"authenticationTypeChanged"`
	// Whether the manifest has any changes.
	HasChanges bool `json:"hasChanges"`
}

// InstallationUpdateChange represents a change of a single value.
type InstallationUpdateChange struct {
	From *string `json:"from"`
	To   *string `json:"to"`
}

// InstallationUpdateResponse defines the structure of the response
// when getting an application installation update.
type InstallationUpdateResponse struct {
	Data *InstallationUpdate `json:"data"`
}

// InstallationUpdateApplyRequest defines the structure of the request
// to apply the latest manifest to an installed application.
type InstallationUpdateApplyRequest struct {
	// Hash of the manifest version that was reviewed
	// (from the ApplicationsService.GetInstallationUpdate response).
	ManifestHash string `json:"manifestHash"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *InstallationUpdateApplyRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ManifestHash == "" {
		return errors.New("manifestHash is required")
	}

	return nil
}

// ApplicationBundleUploadRequest defines the structure of the request
// to upload a bundle archive for a serverless app.
type ApplicationBundleUploadRequest struct {
	// Storage Identifier. The storage file must be a ZIP archive
	// containing the application bundle.
	StorageID int `json:"storageId"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ApplicationBundleUploadRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}

	return nil
}

// ApplicationConsentStatus represents the status of an application consent decision.
type ApplicationConsentStatus string

const (
	// ApplicationConsentGranted means the application may act on behalf of the user.
	ApplicationConsentGranted ApplicationConsentStatus = "granted"
	// ApplicationConsentDenied means the application may not act on behalf of the user.
	ApplicationConsentDenied ApplicationConsentStatus = "denied"
)

// ApplicationConsent represents the current user's decision on whether
// an application may act on their behalf.
type ApplicationConsent struct {
	ID int `json:"id"`
	// The user who installed the application the decision applies to.
	InstalledBy *ShortUser `json:"installedBy"`
	// Application identifier the decision applies to.
	Identifier string `json:"identifier"`
	// Name of the application. Nil when the application is no longer installed.
	Name *string `json:"name"`
	// Whether the application may act on behalf of the current user.
	Status ApplicationConsentStatus `json:"status"`
	// Scopes the decision was made for.
	Scopes    []string `json:"scopes"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// ApplicationConsentResponse defines the structure of the response
// when getting an application consent decision.
type ApplicationConsentResponse struct {
	Data *ApplicationConsent `json:"data"`
}

// ApplicationConsentsListResponse defines the structure of the response
// when getting a list of application consent decisions.
type ApplicationConsentsListResponse struct {
	Data []*ApplicationConsentResponse `json:"data"`
}

// ApplicationConsentsListOptions specifies the optional parameters to the
// ApplicationsService.ListConsents method.
type ApplicationConsentsListOptions struct {
	// Filter by application identifier.
	Identifier string `json:"identifier,omitempty"`
	// Sort consent decisions by the specified fields.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of ApplicationConsentsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ApplicationConsentsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.Identifier != "" {
		v.Add("identifier", o.Identifier)
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// ApplicationConsentAddRequest defines the structure of the request
// to record the current user's consent decision for an application.
type ApplicationConsentAddRequest struct {
	// Application identifier the decision applies to.
	Identifier string `json:"identifier"`
	// Identifier of the user who installed the application.
	InstalledBy int `json:"installedBy"`
	// The decision. Enum: granted, denied.
	Status ApplicationConsentStatus `json:"status"`
	// Scopes the decision was made for.
	Scopes []string `json:"scopes,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ApplicationConsentAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Identifier == "" {
		return errors.New("identifier is required")
	}
	if r.InstalledBy == 0 {
		return errors.New("installedBy is required")
	}

	switch r.Status {
	case ApplicationConsentGranted, ApplicationConsentDenied: // valid
	default:
		return fmt.Errorf("invalid status: %q, must be one of %s, %s",
			r.Status, ApplicationConsentGranted, ApplicationConsentDenied)
	}

	return nil
}

// ApplicationKVRecord represents a key-value record in the storage
// of an installed application.
type ApplicationKVRecord struct {
	// Key of the record.
	Key string `json:"key"`
	// Value of the record. Any JSON value except null.
	Value any `json:"value"`
	// If true, the value is encrypted at rest.
	Secret bool `json:"secret"`
	// Date the record was created, in ISO 8601 format.
	CreatedAt string `json:"createdAt"`
	// Date the record was last updated, in ISO 8601 format.
	UpdatedAt string `json:"updatedAt"`
	// Date the record expires, in ISO 8601 format. Nil for permanent records.
	ExpiresAt *string `json:"expiresAt"`
}

// ApplicationKVRecordResponse defines the structure of the response
// when getting an application KV record.
type ApplicationKVRecordResponse struct {
	Data *ApplicationKVRecord `json:"data"`
}

// ApplicationKVRecordsListResponse defines the structure of the response
// when getting a list of application KV records.
type ApplicationKVRecordsListResponse struct {
	Data []*ApplicationKVRecordResponse `json:"data"`
}

// ApplicationKVRecordsListOptions specifies the optional parameters to the
// ApplicationsService.ListKVRecords method.
type ApplicationKVRecordsListOptions struct {
	// Filter results to keys that start with the given prefix.
	Prefix string `json:"prefix,omitempty"`
	// Sort records by the specified fields.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of ApplicationKVRecordsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ApplicationKVRecordsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.Prefix != "" {
		v.Add("prefix", o.Prefix)
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// ApplicationKVRecordAddRequest defines the structure of the request
// to add a new KV record to the application installation's storage.
type ApplicationKVRecordAddRequest struct {
	// Key of the record. 1-500 characters from `a-z`, `A-Z`, `0-9`, `:`, `.`, `_`, `-`.
	// Reserved prefixes `user:{userId}:` and `module:{moduleKey}:` narrow visibility.
	Key string `json:"key"`
	// Value of the record. Any JSON value except null.
	Value any `json:"value"`
	// If true, the value is encrypted at rest. Immutable. Default is false.
	Secret *bool `json:"secret,omitempty"`
	// Time to live in seconds (60-31536000). Omit for a permanent record.
	TTL int `json:"ttl,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ApplicationKVRecordAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Key == "" {
		return errors.New("key is required")
	}
	if r.Value == nil {
		return errors.New("value is required")
	}

	return nil
}

// IntegrationJob represents a job of an integration application.
type IntegrationJob struct {
	// The unique identifier for the job.
	ID string `json:"id"`
	// The progress of the job.
	Progress int `json:"progress"`
	// The current status of the job.
	Status string `json:"status"`
	// The title of the job.
	Title string `json:"title"`
}

// IntegrationJobsResponse defines the structure of the response
// when getting integration jobs.
type IntegrationJobsResponse struct {
	Data []*IntegrationJob `json:"data"`
}

// IntegrationFilesResponse defines the structure of the response
// when listing Crowdin or integration files of an integration application.
// Each file is a free-form object (e.g. id, name, parentId, type, node_type).
type IntegrationFilesResponse struct {
	Data []map[string]any `json:"data"`
}

// IntegrationFileProgressResponse defines the structure of the response
// when getting file progress of an integration application.
type IntegrationFileProgressResponse struct {
	Data *TranslationProgress `json:"data"`
}

// IntegrationJobResult represents the result of starting an integration job.
type IntegrationJobResult struct {
	JobID string `json:"jobId"`
}

// IntegrationJobResultResponse defines the structure of the response
// when starting an integration job.
type IntegrationJobResultResponse struct {
	Data *IntegrationJobResult `json:"data"`
}

// IntegrationLoginField represents an integration login form field.
type IntegrationLoginField struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// IntegrationLoginFieldsResponse defines the structure of the response
// when getting integration login form fields.
type IntegrationLoginFieldsResponse struct {
	Data []*IntegrationLoginField `json:"data"`
}

// IntegrationSettingsResponse defines the structure of the response
// when getting integration application settings.
type IntegrationSettingsResponse struct {
	Data map[string]any `json:"data"`
}

// IntegrationSyncSettingsResponse defines the structure of the response
// when getting integration sync settings. The data is either an object
// (file ID to language IDs) or an array of file objects.
type IntegrationSyncSettingsResponse struct {
	Data any `json:"data"`
}

// IntegrationCrowdinFilesUpdateRequest defines the structure of the request
// to update Crowdin files from the integration.
type IntegrationCrowdinFilesUpdateRequest struct {
	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Integration files to sync. Each file is a free-form object
	// (e.g. id, name, parentId, type, node_type).
	Files []map[string]any `json:"files"`
	// Upload existing translations from the integration.
	UploadTranslations *bool `json:"uploadTranslations,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *IntegrationCrowdinFilesUpdateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ProjectID == 0 {
		return errors.New("projectId is required")
	}
	if len(r.Files) == 0 {
		return errors.New("files are required")
	}

	return nil
}

// IntegrationFilesUpdateRequest defines the structure of the request
// to update integration files from Crowdin.
type IntegrationFilesUpdateRequest struct {
	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Crowdin file ID mapped to the list of language IDs.
	// Example: {"102": ["de", "fr"]}.
	Files map[string][]string `json:"files"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *IntegrationFilesUpdateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ProjectID == 0 {
		return errors.New("projectId is required")
	}
	if len(r.Files) == 0 {
		return errors.New("files are required")
	}

	return nil
}

// IntegrationLoginRequest defines the structure of the request
// to log in to an integration.
type IntegrationLoginRequest struct {
	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Login form fields values. Get the keys via
	// the ApplicationsService.ListIntegrationLoginFields method.
	Credentials map[string]any `json:"credentials"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *IntegrationLoginRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ProjectID == 0 {
		return errors.New("projectId is required")
	}
	if len(r.Credentials) == 0 {
		return errors.New("credentials are required")
	}

	return nil
}

// IntegrationSettingsUpdateRequest defines the structure of the request
// to update integration application settings.
type IntegrationSettingsUpdateRequest struct {
	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Application settings.
	Config map[string]any `json:"config"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *IntegrationSettingsUpdateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ProjectID == 0 {
		return errors.New("projectId is required")
	}
	if r.Config == nil {
		return errors.New("config is required")
	}

	return nil
}

// IntegrationSyncProvider represents the provider of integration sync settings.
type IntegrationSyncProvider string

const (
	// IntegrationSyncProviderCrowdin is the Crowdin provider.
	IntegrationSyncProviderCrowdin IntegrationSyncProvider = "crowdin"
	// IntegrationSyncProviderIntegration is the integration provider.
	IntegrationSyncProviderIntegration IntegrationSyncProvider = "integration"
)

// IntegrationSyncSettingsUpdateRequest defines the structure of the request
// to update integration sync settings.
type IntegrationSyncSettingsUpdateRequest struct {
	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Provider. Enum: crowdin, integration.
	Provider IntegrationSyncProvider `json:"provider"`
	// Files to sync. Either an object (file ID to language IDs,
	// e.g. map[string][]string) or an array of file objects
	// (e.g. []map[string]any).
	Files any `json:"files"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *IntegrationSyncSettingsUpdateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.ProjectID == 0 {
		return errors.New("projectId is required")
	}

	switch r.Provider {
	case IntegrationSyncProviderCrowdin, IntegrationSyncProviderIntegration: // valid
	default:
		return fmt.Errorf("invalid provider: %q, must be one of %s, %s",
			r.Provider, IntegrationSyncProviderCrowdin, IntegrationSyncProviderIntegration)
	}

	if r.Files == nil {
		return errors.New("files are required")
	}

	return nil
}
