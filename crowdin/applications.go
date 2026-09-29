package crowdin

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// Crowdin Apps are web applications that can be integrated with Crowdin to extend
// its functionality.
//
// Use the API to manage the necessary app data.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Applications
type ApplicationsService struct {
	client *Client
}

// ListInstallations returns a list of application installations.
//
// Use ListInstallationsWithOptions to filter and sort the installations.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.getMany
func (s *ApplicationsService) ListInstallations(ctx context.Context, opt *model.ListOptions) ([]*model.Installation, *Response, error) {
	return s.listInstallations(ctx, opt)
}

// ListInstallationsWithOptions returns a list of application installations
// filtered and sorted by the given options.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.getMany
func (s *ApplicationsService) ListInstallationsWithOptions(ctx context.Context, opts *model.InstallationsListOptions) (
	[]*model.Installation, *Response, error,
) {
	return s.listInstallations(ctx, opts)
}

func (s *ApplicationsService) listInstallations(ctx context.Context, opts ListOptionsProvider) ([]*model.Installation, *Response, error) {
	res := new(model.InstallationsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/applications/installations", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Installation, 0, len(res.Data))
	for _, installation := range res.Data {
		list = append(list, installation.Data)
	}

	return list, resp, err
}

// GetInstallation returns information about an application installation.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.get
func (s *ApplicationsService) GetInstallation(ctx context.Context, applicationID string) (*model.Installation, *Response, error) {
	res := new(model.InstallationResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/applications/installations/%s", applicationID), nil, res)

	return res.Data, resp, err
}

// Install installs an application from a manifest URL or from inline manifest content.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.post
func (s *ApplicationsService) Install(ctx context.Context, req *model.InstallApplicationRequest) (
	*model.Installation, *Response, error,
) {
	res := new(model.InstallationResponse)
	resp, err := s.client.Post(ctx, "/api/v2/applications/installations", req, res)

	return res.Data, resp, err
}

// EditInstallation updates an application installation.
//
// Request body:
//   - op (string): operation to perform. Enum: replace.
//   - path (string <json-pointer>): path to the field to update.
//     Enum: "/permissions", "/modules/{moduleKey}/permissions", "/manifest".
//     The "/manifest" path is available only for applications installed from manifest content.
//   - value (model.InstallationReplaceValue): object with values to update,
//     or model.ApplicationManifest for the "/manifest" path.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.patch
func (s *ApplicationsService) EditInstallation(ctx context.Context, applicationID string, req []*model.UpdateRequest) (
	*model.Installation, *Response, error,
) {
	res := new(model.InstallationResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/applications/installations/%s", applicationID), req, res)

	return res.Data, resp, err
}

// DeleteInstallation deletes an application installation.
//
//	id: application identifier
//	force: if true, force to delete application installation
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.delete
func (s *ApplicationsService) DeleteInstallation(ctx context.Context, applicationID string, force bool) (*Response, error) {
	path := fmt.Sprintf("/api/v2/applications/installations/%s", applicationID)
	if force {
		path += "?force=true"
	}

	return s.client.Delete(ctx, path, nil)
}

// GetInstallationUpdate returns the diff between the currently installed application
// and the latest cached manifest, tagged with `manifestHash` for optimistic locking on apply.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.update.get
func (s *ApplicationsService) GetInstallationUpdate(ctx context.Context, applicationID string) (
	*model.InstallationUpdate, *Response, error,
) {
	res := new(model.InstallationUpdateResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/applications/installations/%s/update", applicationID), nil, res)

	return res.Data, resp, err
}

// ApplyInstallationUpdate applies the latest cached manifest to an installed application.
// It requires `manifestHash` from a recent GetInstallationUpdate call. If the cached
// manifest has changed since, the API returns 409.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.update.post
func (s *ApplicationsService) ApplyInstallationUpdate(ctx context.Context, applicationID string, req *model.InstallationUpdateApplyRequest) (
	*model.Installation, *Response, error,
) {
	res := new(model.InstallationResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/applications/installations/%s/update", applicationID), req, res)

	return res.Data, resp, err
}

// UploadBundle uploads a bundle archive for a serverless app installed from manifest content.
// The bundle must be a ZIP archive that contains a non-empty `app.js` entry point at its root.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.installations.bundles.post
func (s *ApplicationsService) UploadBundle(ctx context.Context, applicationID string, req *model.ApplicationBundleUploadRequest) (
	*model.Installation, *Response, error,
) {
	res := new(model.InstallationResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/applications/installations/%s/bundles", applicationID), req, res)

	return res.Data, resp, err
}

// ListConsents returns a list of the current user's application consent decisions.
//
// Note: This endpoint is available only in Crowdin.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.consents.getMany
func (s *ApplicationsService) ListConsents(ctx context.Context, opts *model.ApplicationConsentsListOptions) (
	[]*model.ApplicationConsent, *Response, error,
) {
	res := new(model.ApplicationConsentsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/applications/consents", opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ApplicationConsent, 0, len(res.Data))
	for _, consent := range res.Data {
		list = append(list, consent.Data)
	}

	return list, resp, err
}

// AddConsent records the current user's consent decision for an application
// installed by another user.
//
// Note: This endpoint is available only in Crowdin.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.consents.post
func (s *ApplicationsService) AddConsent(ctx context.Context, req *model.ApplicationConsentAddRequest) (
	*model.ApplicationConsent, *Response, error,
) {
	res := new(model.ApplicationConsentResponse)
	resp, err := s.client.Post(ctx, "/api/v2/applications/consents", req, res)

	return res.Data, resp, err
}

// EditConsent changes an existing consent decision.
//
// Request body:
//   - op (string): operation to perform. Enum: replace.
//   - path (string <json-pointer>): path to the field to update. Enum: "/status", "/scopes".
//   - value (string|[]string): a status (`granted`/`denied`) for "/status",
//     or an array of scope identifiers for "/scopes".
//
// Note: This endpoint is available only in Crowdin.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.consents.patch
func (s *ApplicationsService) EditConsent(ctx context.Context, consentID int, req []*model.UpdateRequest) (
	*model.ApplicationConsent, *Response, error,
) {
	res := new(model.ApplicationConsentResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/applications/consents/%d", consentID), req, res)

	return res.Data, resp, err
}

// DeleteConsent forgets the current user's consent decision for an application.
// The next call from that application will prompt for consent again.
//
// Note: This endpoint is available only in Crowdin.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.consents.delete
func (s *ApplicationsService) DeleteConsent(ctx context.Context, consentID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/applications/consents/%d", consentID), nil)
}

// kvRecordPath returns the path of an application KV record. The key is
// path-escaped, and `:` is encoded as `%3A` as required by the API.
func kvRecordPath(applicationID, key string) string {
	return fmt.Sprintf("/api/v2/applications/%s/storage/kv/records/%s",
		applicationID, strings.ReplaceAll(url.PathEscape(key), ":", "%3A"))
}

// ListKVRecords returns a list of the KV records in the application installation's
// storage that are visible to the current user.
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.getMany
func (s *ApplicationsService) ListKVRecords(ctx context.Context, applicationID string, opts *model.ApplicationKVRecordsListOptions) (
	[]*model.ApplicationKVRecord, *Response, error,
) {
	res := new(model.ApplicationKVRecordsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/applications/%s/storage/kv/records", applicationID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ApplicationKVRecord, 0, len(res.Data))
	for _, record := range res.Data {
		list = append(list, record.Data)
	}

	return list, resp, err
}

// AddKVRecord adds a new KV record to the application installation's storage.
// If the key already exists, the API returns 409.
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.post
func (s *ApplicationsService) AddKVRecord(ctx context.Context, applicationID string, req *model.ApplicationKVRecordAddRequest) (
	*model.ApplicationKVRecord, *Response, error,
) {
	res := new(model.ApplicationKVRecordResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/applications/%s/storage/kv/records", applicationID), req, res)

	return res.Data, resp, err
}

// GetKVRecord returns a single KV record.
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.get
func (s *ApplicationsService) GetKVRecord(ctx context.Context, applicationID, key string) (
	*model.ApplicationKVRecord, *Response, error,
) {
	res := new(model.ApplicationKVRecordResponse)
	resp, err := s.client.Get(ctx, kvRecordPath(applicationID, key), nil, res)

	return res.Data, resp, err
}

// EditKVRecord replaces the value, the TTL, or both on an existing KV record.
//
// Request body:
//   - op (string): operation to perform. Enum: replace.
//   - path (string <json-pointer>): path to the field to update. Enum: "/value", "/ttl".
//   - value (any): any JSON value for "/value", or a TTL in seconds (60-31536000) for "/ttl".
//
// Use RemoveKVRecordTTL to remove the TTL and make the record permanent.
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.patch
func (s *ApplicationsService) EditKVRecord(ctx context.Context, applicationID, key string, req []*model.UpdateRequest) (
	*model.ApplicationKVRecord, *Response, error,
) {
	res := new(model.ApplicationKVRecordResponse)
	resp, err := s.client.Patch(ctx, kvRecordPath(applicationID, key), req, res)

	return res.Data, resp, err
}

// RemoveKVRecordTTL removes the TTL of an existing KV record and makes it permanent
// (sends a replace operation for the "/ttl" path with a null value).
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.patch
func (s *ApplicationsService) RemoveKVRecordTTL(ctx context.Context, applicationID, key string) (
	*model.ApplicationKVRecord, *Response, error,
) {
	body := []map[string]any{{"op": model.OpReplace, "path": "/ttl", "value": nil}}

	res := new(model.ApplicationKVRecordResponse)
	resp, err := s.client.Patch(ctx, kvRecordPath(applicationID, key), body, res)

	return res.Data, resp, err
}

// DeleteKVRecord deletes a KV record.
//
// Note: Requires the application's own access token.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.storage.kv.records.delete
func (s *ApplicationsService) DeleteKVRecord(ctx context.Context, applicationID, key string) (*Response, error) {
	return s.client.Delete(ctx, kvRecordPath(applicationID, key), nil)
}

// integrationPath returns the path of an integration application API endpoint
// with the given query parameters.
func integrationPath(applicationID, endpoint string, query url.Values) string {
	path := fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, endpoint)
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	return path
}

// projectQuery returns the query parameters with the given project identifier.
func projectQuery(projectID int) url.Values {
	return url.Values{"projectId": []string{strconv.Itoa(projectID)}}
}

// jobQuery returns the query parameters with the given project identifier
// and job identifier (omitted when empty).
func jobQuery(projectID int, jobID string) url.Values {
	query := projectQuery(projectID)
	if jobID != "" {
		query.Set("jobId", jobID)
	}

	return query
}

// ListIntegrationCrowdinFiles returns a list of Crowdin files of an integration application.
// Each file is a free-form object (e.g. id, name, parentId, type).
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.crowdin.files
func (s *ApplicationsService) ListIntegrationCrowdinFiles(ctx context.Context, applicationID string, projectID int) (
	[]map[string]any, *Response, error,
) {
	res := new(model.IntegrationFilesResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "crowdin-files", projectQuery(projectID)), nil, res)

	return res.Data, resp, err
}

// UpdateIntegrationCrowdinFiles updates Crowdin files from the integration.
// It returns the identifier of the started job.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.crowdin.update
func (s *ApplicationsService) UpdateIntegrationCrowdinFiles(ctx context.Context, applicationID string, req *model.IntegrationCrowdinFilesUpdateRequest) (
	*model.IntegrationJobResult, *Response, error,
) {
	res := new(model.IntegrationJobResultResponse)
	resp, err := s.client.Post(ctx, integrationPath(applicationID, "crowdin-update", nil), req, res)

	return res.Data, resp, err
}

// GetIntegrationFileProgress returns the translation progress of a Crowdin file
// of an integration application.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.file.progress
func (s *ApplicationsService) GetIntegrationFileProgress(ctx context.Context, applicationID string, projectID, fileID int) (
	*model.TranslationProgress, *Response, error,
) {
	query := projectQuery(projectID)
	query.Set("fileId", strconv.Itoa(fileID))

	res := new(model.IntegrationFileProgressResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "file-progress", query), nil, res)

	return res.Data, resp, err
}

// ListIntegrationFiles returns a list of integration files of an integration application.
// Each file is a free-form object (e.g. id, name, parentId, type, node_type).
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.integration.files
func (s *ApplicationsService) ListIntegrationFiles(ctx context.Context, applicationID string, projectID int) (
	[]map[string]any, *Response, error,
) {
	res := new(model.IntegrationFilesResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "integration-files", projectQuery(projectID)), nil, res)

	return res.Data, resp, err
}

// UpdateIntegrationFiles updates integration files from Crowdin.
// It returns the identifier of the started job.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.integration.update
func (s *ApplicationsService) UpdateIntegrationFiles(ctx context.Context, applicationID string, req *model.IntegrationFilesUpdateRequest) (
	*model.IntegrationJobResult, *Response, error,
) {
	res := new(model.IntegrationJobResultResponse)
	resp, err := s.client.Post(ctx, integrationPath(applicationID, "integration-update", nil), req, res)

	return res.Data, resp, err
}

// GetIntegrationJobInfo returns information about an integration job.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.job.info
func (s *ApplicationsService) GetIntegrationJobInfo(ctx context.Context, applicationID string, projectID int, jobID string) (
	[]*model.IntegrationJob, *Response, error,
) {
	res := new(model.IntegrationJobsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "job-info", jobQuery(projectID, jobID)), nil, res)

	return res.Data, resp, err
}

// GetIntegrationJobs returns information about integration jobs.
// The jobID is optional; pass an empty string to omit it.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.job.get
func (s *ApplicationsService) GetIntegrationJobs(ctx context.Context, applicationID string, projectID int, jobID string) (
	[]*model.IntegrationJob, *Response, error,
) {
	res := new(model.IntegrationJobsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "jobs", jobQuery(projectID, jobID)), nil, res)

	return res.Data, resp, err
}

// ListIntegrationJobs returns a list of all jobs of an integration application.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.job.list
func (s *ApplicationsService) ListIntegrationJobs(ctx context.Context, applicationID string, projectID int, opts *model.ListOptions) (
	[]*model.IntegrationJob, *Response, error,
) {
	query, _ := opts.Values()
	query.Set("projectId", strconv.Itoa(projectID))

	res := new(model.IntegrationJobsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "all-jobs", query), nil, res)

	return res.Data, resp, err
}

// CancelIntegrationJob cancels an integration job.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.job.cancel
func (s *ApplicationsService) CancelIntegrationJob(ctx context.Context, applicationID string, projectID int, jobID string) (*Response, error) {
	return s.client.Delete(ctx, integrationPath(applicationID, "jobs", jobQuery(projectID, jobID)), nil)
}

// IntegrationLogin logs in to an integration with the given credentials.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.integration.login
func (s *ApplicationsService) IntegrationLogin(ctx context.Context, applicationID string, req *model.IntegrationLoginRequest) (*Response, error) {
	return s.client.Post(ctx, integrationPath(applicationID, "login", nil), req, nil)
}

// ListIntegrationLoginFields returns the integration login form fields.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.integration.fields
func (s *ApplicationsService) ListIntegrationLoginFields(ctx context.Context, applicationID string) (
	[]*model.IntegrationLoginField, *Response, error,
) {
	res := new(model.IntegrationLoginFieldsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "login-fields", nil), nil, res)

	return res.Data, resp, err
}

// GetIntegrationSettings returns the settings of an integration application.
// The settings are a free-form object.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.settings.get
func (s *ApplicationsService) GetIntegrationSettings(ctx context.Context, applicationID string, projectID int) (
	map[string]any, *Response, error,
) {
	res := new(model.IntegrationSettingsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "settings", projectQuery(projectID)), nil, res)

	return res.Data, resp, err
}

// UpdateIntegrationSettings updates the settings of an integration application.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.settings.update
func (s *ApplicationsService) UpdateIntegrationSettings(ctx context.Context, applicationID string, req *model.IntegrationSettingsUpdateRequest) (
	*Response, error,
) {
	return s.client.Post(ctx, integrationPath(applicationID, "settings", nil), req, nil)
}

// GetIntegrationSyncSettings returns the sync settings of an integration application.
// The result is either an object (file ID to language IDs) or an array of file objects.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.sync.settings.get
func (s *ApplicationsService) GetIntegrationSyncSettings(ctx context.Context, applicationID string, projectID int,
	provider model.IntegrationSyncProvider,
) (any, *Response, error) {
	query := projectQuery(projectID)
	query.Set("provider", string(provider))

	res := new(model.IntegrationSyncSettingsResponse)
	resp, err := s.client.Get(ctx, integrationPath(applicationID, "sync-settings", query), nil, res)

	return res.Data, resp, err
}

// UpdateIntegrationSyncSettings updates the sync settings of an integration application.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.integrations.sync.settings.update
func (s *ApplicationsService) UpdateIntegrationSyncSettings(ctx context.Context, applicationID string, req *model.IntegrationSyncSettingsUpdateRequest) (
	*Response, error,
) {
	return s.client.Post(ctx, integrationPath(applicationID, "sync-settings", nil), req, nil)
}

// GetData returns application data.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.api.get
func (s *ApplicationsService) GetData(ctx context.Context, applicationID, path string) (any, *Response, error) {
	res := new(model.ApplicationDataResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path), nil, res)

	return res.Data, resp, err
}

// AddData adds application data.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.api.post
func (s *ApplicationsService) AddData(ctx context.Context, applicationID, path string, req map[string]any) (
	any, *Response, error,
) {
	res := new(model.ApplicationDataResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path), req, res)

	return res.Data, resp, err
}

// UpdateOrRestoreData updates or restores application data.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.api.put
func (s *ApplicationsService) UpdateOrRestoreData(ctx context.Context, applicationID, path string, req map[string]any) (
	any, *Response, error,
) {
	res := new(model.ApplicationDataResponse)
	resp, err := s.client.Put(ctx, fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path), req, res)

	return res.Data, resp, err
}

// EditData updates application data.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.api.patch
func (s *ApplicationsService) EditData(ctx context.Context, applicationID, path string, req map[string]any) (
	any, *Response, error,
) {
	res := new(model.ApplicationDataResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path), req, res)

	return res.Data, resp, err
}

// DeleteData deletes application data.
//
// https://developer.crowdin.com/api/v2/#operation/api.applications.api.delete
func (s *ApplicationsService) DeleteData(ctx context.Context, applicationID, path string) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/applications/%s/api/%s", applicationID, path), nil)
}
