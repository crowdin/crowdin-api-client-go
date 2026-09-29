package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstallApplicationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *InstallApplicationRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &InstallApplicationRequest{},
			err:  "url is required",
		},
		{
			name:  "valid request",
			req:   &InstallApplicationRequest{URL: "https://example.com/app/install"},
			valid: true,
		},
		{
			name: "url and manifest",
			req: &InstallApplicationRequest{
				URL:      "https://example.com/app/install",
				Manifest: &ApplicationManifest{Name: "My App"},
			},
			err: "url and manifest cannot be used together",
		},
		{
			name: "manifest without name",
			req:  &InstallApplicationRequest{Manifest: &ApplicationManifest{}},
			err:  "manifest name is required",
		},
		{
			name: "manifest without modules",
			req:  &InstallApplicationRequest{Manifest: &ApplicationManifest{Name: "My App"}},
			err:  "manifest modules are required",
		},
		{
			name: "valid manifest request",
			req: &InstallApplicationRequest{
				Manifest: &ApplicationManifest{
					Name: "My App",
					Modules: map[string][]*ApplicationManifestModule{
						"editor-right-panel": {{Key: "my-module", Name: "My Module"}},
					},
				},
				AssignAgent: toPtr(true),
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.valid {
				assert.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestInstallationsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *InstallationsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &InstallationsListOptions{},
		},
		{
			name: "with installed by",
			opt:  &InstallationsListOptions{InstalledBy: 12},
			out:  "installedBy=12",
		},
		{
			name: "with all options",
			opt: &InstallationsListOptions{
				InstalledBy: 12,
				OrderBy:     "createdAt desc,name",
				ListOptions: ListOptions{Limit: 10, Offset: 5},
			},
			out: "installedBy=12&limit=10&offset=5&orderBy=createdAt+desc%2Cname",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestApplicationConsentsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *ApplicationConsentsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &ApplicationConsentsListOptions{},
		},
		{
			name: "with identifier",
			opt:  &ApplicationConsentsListOptions{Identifier: "example-application"},
			out:  "identifier=example-application",
		},
		{
			name: "with all options",
			opt: &ApplicationConsentsListOptions{
				Identifier:  "example-application",
				OrderBy:     "createdAt desc",
				ListOptions: ListOptions{Limit: 10, Offset: 5},
			},
			out: "identifier=example-application&limit=10&offset=5&orderBy=createdAt+desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestApplicationKVRecordsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *ApplicationKVRecordsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &ApplicationKVRecordsListOptions{},
		},
		{
			name: "with prefix",
			opt:  &ApplicationKVRecordsListOptions{Prefix: "user:1:"},
			out:  "prefix=user%3A1%3A",
		},
		{
			name: "with all options",
			opt: &ApplicationKVRecordsListOptions{
				Prefix:      "settings.",
				OrderBy:     "key",
				ListOptions: ListOptions{Limit: 10, Offset: 5},
			},
			out: "limit=10&offset=5&orderBy=key&prefix=settings.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

type applicationsRequestValidator interface {
	Validate() error
}

func TestApplicationsRequestsValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   applicationsRequestValidator
		err   string
		valid bool
	}{
		// InstallationUpdateApplyRequest.
		{
			name: "update apply: nil request",
			req:  (*InstallationUpdateApplyRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "update apply: empty request",
			req:  &InstallationUpdateApplyRequest{},
			err:  "manifestHash is required",
		},
		{
			name:  "update apply: valid request",
			req:   &InstallationUpdateApplyRequest{ManifestHash: "5536b952"},
			valid: true,
		},
		// ApplicationBundleUploadRequest.
		{
			name: "bundle upload: nil request",
			req:  (*ApplicationBundleUploadRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "bundle upload: empty request",
			req:  &ApplicationBundleUploadRequest{},
			err:  "storageId is required",
		},
		{
			name:  "bundle upload: valid request",
			req:   &ApplicationBundleUploadRequest{StorageID: 12},
			valid: true,
		},
		// ApplicationConsentAddRequest.
		{
			name: "consent add: nil request",
			req:  (*ApplicationConsentAddRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "consent add: missing identifier",
			req:  &ApplicationConsentAddRequest{InstalledBy: 12, Status: ApplicationConsentGranted},
			err:  "identifier is required",
		},
		{
			name: "consent add: missing installedBy",
			req:  &ApplicationConsentAddRequest{Identifier: "app", Status: ApplicationConsentGranted},
			err:  "installedBy is required",
		},
		{
			name: "consent add: invalid status",
			req:  &ApplicationConsentAddRequest{Identifier: "app", InstalledBy: 12, Status: "unknown"},
			err:  `invalid status: "unknown", must be one of granted, denied`,
		},
		{
			name:  "consent add: valid request",
			req:   &ApplicationConsentAddRequest{Identifier: "app", InstalledBy: 12, Status: ApplicationConsentDenied},
			valid: true,
		},
		// ApplicationKVRecordAddRequest.
		{
			name: "kv add: nil request",
			req:  (*ApplicationKVRecordAddRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "kv add: missing key",
			req:  &ApplicationKVRecordAddRequest{Value: "dark"},
			err:  "key is required",
		},
		{
			name: "kv add: missing value",
			req:  &ApplicationKVRecordAddRequest{Key: "settings.theme"},
			err:  "value is required",
		},
		{
			name:  "kv add: valid request",
			req:   &ApplicationKVRecordAddRequest{Key: "settings.theme", Value: false, Secret: toPtr(false), TTL: 60},
			valid: true,
		},
		// IntegrationCrowdinFilesUpdateRequest.
		{
			name: "crowdin files update: nil request",
			req:  (*IntegrationCrowdinFilesUpdateRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "crowdin files update: missing projectId",
			req:  &IntegrationCrowdinFilesUpdateRequest{Files: []map[string]any{{"id": "1"}}},
			err:  "projectId is required",
		},
		{
			name: "crowdin files update: missing files",
			req:  &IntegrationCrowdinFilesUpdateRequest{ProjectID: 1},
			err:  "files are required",
		},
		{
			name:  "crowdin files update: valid request",
			req:   &IntegrationCrowdinFilesUpdateRequest{ProjectID: 1, Files: []map[string]any{{"id": "1"}}},
			valid: true,
		},
		// IntegrationFilesUpdateRequest.
		{
			name: "integration files update: nil request",
			req:  (*IntegrationFilesUpdateRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "integration files update: missing projectId",
			req:  &IntegrationFilesUpdateRequest{Files: map[string][]string{"102": {"de"}}},
			err:  "projectId is required",
		},
		{
			name: "integration files update: missing files",
			req:  &IntegrationFilesUpdateRequest{ProjectID: 1},
			err:  "files are required",
		},
		{
			name:  "integration files update: valid request",
			req:   &IntegrationFilesUpdateRequest{ProjectID: 1, Files: map[string][]string{"102": {"de"}}},
			valid: true,
		},
		// IntegrationLoginRequest.
		{
			name: "login: nil request",
			req:  (*IntegrationLoginRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "login: missing projectId",
			req:  &IntegrationLoginRequest{Credentials: map[string]any{"apiKey": "key"}},
			err:  "projectId is required",
		},
		{
			name: "login: missing credentials",
			req:  &IntegrationLoginRequest{ProjectID: 1},
			err:  "credentials are required",
		},
		{
			name:  "login: valid request",
			req:   &IntegrationLoginRequest{ProjectID: 1, Credentials: map[string]any{"apiKey": "key"}},
			valid: true,
		},
		// IntegrationSettingsUpdateRequest.
		{
			name: "settings update: nil request",
			req:  (*IntegrationSettingsUpdateRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "settings update: missing projectId",
			req:  &IntegrationSettingsUpdateRequest{Config: map[string]any{"schedule": 0}},
			err:  "projectId is required",
		},
		{
			name: "settings update: missing config",
			req:  &IntegrationSettingsUpdateRequest{ProjectID: 1},
			err:  "config is required",
		},
		{
			name:  "settings update: valid request",
			req:   &IntegrationSettingsUpdateRequest{ProjectID: 1, Config: map[string]any{"schedule": 0}},
			valid: true,
		},
		// IntegrationSyncSettingsUpdateRequest.
		{
			name: "sync settings update: nil request",
			req:  (*IntegrationSyncSettingsUpdateRequest)(nil),
			err:  "request cannot be nil",
		},
		{
			name: "sync settings update: missing projectId",
			req:  &IntegrationSyncSettingsUpdateRequest{Provider: IntegrationSyncProviderCrowdin, Files: map[string][]string{}},
			err:  "projectId is required",
		},
		{
			name: "sync settings update: invalid provider",
			req:  &IntegrationSyncSettingsUpdateRequest{ProjectID: 1, Provider: "unknown", Files: map[string][]string{}},
			err:  `invalid provider: "unknown", must be one of crowdin, integration`,
		},
		{
			name: "sync settings update: missing files",
			req:  &IntegrationSyncSettingsUpdateRequest{ProjectID: 1, Provider: IntegrationSyncProviderIntegration},
			err:  "files are required",
		},
		{
			name: "sync settings update: valid request",
			req: &IntegrationSyncSettingsUpdateRequest{
				ProjectID: 1,
				Provider:  IntegrationSyncProviderCrowdin,
				Files:     map[string][]string{"102": {"uk", "de"}},
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.valid {
				assert.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.err)
			}
		})
	}
}
