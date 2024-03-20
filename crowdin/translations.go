package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// Translators can work with entirely untranslated project or you can pre-translate the files
// to ease the translations process.
// Use API to pre-translate files via Machine Translation (MT) or Translation Memory (TM),
// upload your existing translations, and download translations correspondingly. Pre-translate
// and build are asynchronous operations and shall be completed with sequence of API methods.
//
// Note: If there are no new translations or changes in build parameters, Crowdin will return
// the current build for such requests.
type TranslationsService struct {
	client *Client
}

// PreTranslation returns a pre-translation status for project by its identifier.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.pre-translations.get
func (s *TranslationsService) PreTranslation(ctx context.Context, projectID int64, preTranslationID string) (
	*model.PreTranslation, *Response, error,
) {
	res := new(model.PreTranslationsResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/pre-translations/%s", projectID, preTranslationID), nil, res)

	return res.Data, resp, err
}

// ApplyPreTranslation applies pre-translation to the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.pre-translations.post
func (s *TranslationsService) ApplyPreTranslation(ctx context.Context, projectID int64, req *model.PreTranslationRequest) (
	*model.PreTranslation, *Response, error,
) {
	res := new(model.PreTranslationsResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/pre-translations", projectID), req, res)

	return res.Data, resp, err
}

// BuildProjectDirectoryTranslation builds translations for a specific directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.directories.post
func (s *TranslationsService) BuildProjectDirectoryTranslation(
	ctx context.Context,
	projectID, directoryID int64,
	req *model.BuildProjectDirectoryTranslationRequest,
) (*model.BuildProjectDirectoryTranslation, *Response, error) {
	res := struct {
		Data *model.BuildProjectDirectoryTranslation `json:"data"`
	}{}
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds/directories/%d", projectID, directoryID), req, &res)

	return res.Data, resp, err
}

// BuildProjectFileTranslation builds translations for a specific file in the project.
//
// Note: Pass `etag` identifier to see whether any changes were applied to the file. If etag is not empty,
// it would be added to the If-None-Match request header.
// In case the file was changed it would be built. If not you'll receive a 304 (Not Modified) status code.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.files.post
func (s *TranslationsService) BuildProjectFileTranslation(
	ctx context.Context,
	projectID, fileID int64,
	req *model.BuildProjectFileTranslationRequest,
	etag string,
) (*model.DownloadLink, *Response, error) {
	path := fmt.Sprintf("/api/v2/projects/%d/translations/builds/files/%d", projectID, fileID)
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Post(ctx, path, req, res, Header("If-None-Match", etag))

	return res.Data, resp, err
}

// ListProjectBuilds returns a list of builds for a specific project.
//
// Query parameters:
// - branchId: The identifier of the branch (filter by branch).
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.getMany
func (s *TranslationsService) ListProjectBuilds(ctx context.Context, projectID int64, opts *model.TranslationsBuildsListOptions) (
	[]*model.TranslationsProjectBuild, *Response, error,
) {
	res := new(model.TranslationsProjectBuildsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	builds := make([]*model.TranslationsProjectBuild, 0, len(res.Data))
	for _, build := range res.Data {
		builds = append(builds, build.Data)
	}

	return builds, resp, err
}

// BuildProjectTranslation builds project translations.
// Request body can be either `model.BuildProjectRequest` or `model.PseudoBuildProjectRequest`.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.post
func (s *TranslationsService) BuildProjectTranslation(ctx context.Context, projectID int64, req model.BuildProjectTranslationRequest) (
	*model.TranslationsProjectBuild, *Response, error,
) {
	res := new(model.TranslationsProjectBuildResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds", projectID), req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res.Data, resp, err
}

// UploadTranslations uploads translations for a specific language in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.postOnLanguage
func (s *TranslationsService) UploadTranslations(ctx context.Context, projectID int64, languageID string, req *model.UploadTranslationsRequest) (
	*model.UploadTranslations, *Response, error,
) {
	res := new(model.UploadTranslationsResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/%s", projectID, languageID), req, res)

	return &res.Data, resp, err
}

// DownloadProjectTranslations returns a download link for a specific build.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.download.download
func (s *TranslationsService) DownloadProjectTranslations(ctx context.Context, projectID, buildID int64) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds/%d/download", projectID, buildID), nil, res)

	return res.Data, resp, err
}

// GetBuild checks the status of a project build by its identifier.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.get
func (s *TranslationsService) GetBuild(ctx context.Context, projectID, buildID int64) (*model.TranslationsProjectBuild, *Response, error) {
	res := new(model.TranslationsProjectBuildResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds/%d", projectID, buildID), nil, res)

	return res.Data, resp, err
}

// CancelBuild cancels a build by its identifier.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.builds.delete
func (s *TranslationsService) CancelBuild(ctx context.Context, projectID, buildID int64) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/builds/%d", projectID, buildID))
}

// ExportProjectTranslation exports project translations for a specific language.
//
// Note: For instant translation delivery to your mobile, web, server, or desktop apps,
// it is recommended to use OTA.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.translations.exports.post
func (s *TranslationsService) ExportProjectTranslation(ctx context.Context, projectID int64, req *model.ExportTranslationRequest) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/translations/exports", projectID), req, res)

	return res.Data, resp, err
}
