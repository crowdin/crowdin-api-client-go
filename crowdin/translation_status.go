package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// Status represents the general localization progress on both translations and proofreading.
// Use API to check translation and proofreading progress on different levels: file, language, branch, directory.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Translation-Status
type TranslationStatusService struct {
	client *Client
}

// GetBranchProgress returns the translation and proofreading progress on a branch level.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.languages.progress.getMany
func (s *TranslationStatusService) GetBranchProgress(ctx context.Context, projectID, branchID int64, opts *model.ListOptions) (
	[]*model.TranslationProgress, *Response, error,
) {
	return s.progress(ctx, fmt.Sprintf("projects/%d/branches/%d/languages/progress", projectID, branchID), opts)
}

// GetDirectoryProgress returns the translation and proofreading progress on a directory level.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.languages.progress.getMany
func (s *TranslationStatusService) GetDirectoryProgress(ctx context.Context, projectID, directoryID int64, opts *model.ListOptions) (
	[]*model.TranslationProgress, *Response, error,
) {
	return s.progress(ctx, fmt.Sprintf("projects/%d/directories/%d/languages/progress", projectID, directoryID), opts)
}

// GetFileProgress returns the translation and proofreading progress on a file level.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.languages.progress.getMany
func (s *TranslationStatusService) GetFileProgress(ctx context.Context, projectID, fileID int64, opts *model.ListOptions) (
	[]*model.TranslationProgress, *Response, error,
) {
	return s.progress(ctx, fmt.Sprintf("projects/%d/files/%d/languages/progress", projectID, fileID), opts)
}

// GetLanguageProgress returns the translation and proofreading progress on a language level.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.languages.files.progress.getMany
func (s *TranslationStatusService) GetLanguageProgress(ctx context.Context, projectID int64, languageID string, opts *model.ListOptions) (
	[]*model.TranslationProgress, *Response, error,
) {
	return s.progress(ctx, fmt.Sprintf("projects/%d/languages/%s/progress", projectID, languageID), opts)
}

// GetProjectProgress returns the translation and proofreading progress on a project level.
//
// Query parameters:
// - languageIds: Filter progress by Language Identifier.
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.languages.progress.getMany
func (s *TranslationStatusService) GetProjectProgress(ctx context.Context, projectID int64, opts *model.ProjectProgressListOptions) (
	[]*model.TranslationProgress, *Response, error,
) {
	return s.progress(ctx, fmt.Sprintf("projects/%d/languages/progress", projectID), opts)
}

// ListQAChecks returns a list of QA check issues.
//
// Query parameters:
// - category: Filter progress by Category.
// - validation: Filter progress by Validation type.
// - languageId: Filter progress by Language Identifier.
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.qa-checks.getMany
func (s *TranslationStatusService) ListQAChecks(ctx context.Context, projectID int64, opts *model.QACheckListOptions) (
	[]*model.QACheck, *Response, error,
) {
	res := new(model.QAChecksResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("projects/%d/qa-checks", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	issues := make([]*model.QACheck, len(res.Data))
	for _, i := range res.Data {
		issues = append(issues, i.Data)
	}

	return issues, resp, nil
}

func (s *TranslationStatusService) progress(ctx context.Context, path string, opts ListOptionsProvider) (
	[]*model.TranslationProgress, *Response, error,
) {
	res := new(model.TranslationProgressResponse)
	resp, err := s.client.Get(ctx, path, opts, res)
	if err != nil {
		return nil, resp, err
	}

	progress := make([]*model.TranslationProgress, len(res.Data))
	for _, p := range res.Data {
		progress = append(progress, p.Data)
	}

	return progress, resp, nil
}
