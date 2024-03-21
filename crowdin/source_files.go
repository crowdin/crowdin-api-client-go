package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// Source files are resources for translation. You can keep files structure using
// folders or manage different versions of the content via branches.
// Use API to keep the source files up to date, check on file revisions, and manage project branches.
// Before adding source files to the project, upload each file to the Storage first.
//
// Note: If you use branches, make sure your master branch is the first one you integrate with Crowdin.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/Source-Files
type SourceFilesService struct {
	client *Client
}

// ListBranches returns a list of project branches.
//
// Query parameters:
// - name: Filter branches by name.
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.getMany
func (s *SourceFilesService) ListBranches(ctx context.Context, projectID int64, opts *model.BranchListOptions) (
	[]*model.Branch, *Response, error,
) {
	res := new(model.BranchListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/branches", projectID), nil, res)

	branches := make([]*model.Branch, 0, len(res.Data))
	for _, b := range res.Data {
		branches = append(branches, b.Data)
	}

	return branches, resp, err
}

// GetBranch returns a single project branch.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.get
func (s *SourceFilesService) GetBranch(ctx context.Context, projectID, branchID int64) (*model.Branch, *Response, error) {
	res := new(model.BranchGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/branches/%d", projectID, branchID), nil, res)

	return res.Data, resp, err
}

// AddBranch creates a new project branch.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.post
func (s *SourceFilesService) AddBranch(ctx context.Context, projectID int64, req *model.BranchAddRequest) (
	*model.Branch, *Response, error,
) {
	res := new(model.BranchGetResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/branches", projectID), req, res)

	return res.Data, resp, err
}

// EditBranch updates a project branch.
//
// Request body:
// - op: The operation to perform. Enum: replace, test.
// - path: A JSON Pointer as defined in RFC 6901.  Enum: "/name", "/title", "/exportPattern", "/priority".
// - value: The value to be used within the operations. The value must be one of string.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.patch
func (s *SourceFilesService) EditBranch(ctx context.Context, projectID, branchID int64, req []*model.UpdateRequest) (
	*model.Branch, *Response, error,
) {
	res := new(model.BranchGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/branches/%d", projectID, branchID), req, res)

	return res.Data, resp, err
}

// DeleteBranch deletes a project branch.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.branches.delete
func (s *SourceFilesService) DeleteBranch(ctx context.Context, projectID, branchID int64) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/branches/%d", projectID, branchID))
}

// ListDirectories returns a list of directories in the project.
//
// Query parameters:
// - branchId: The identifier of the branch (filter by branch).
// - directoryId: The identifier of the directory (filter by directory).
// - filter: Filter directories by name.
// - recursion: List directories recursively.
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.getMany
func (s *SourceFilesService) ListDirectories(ctx context.Context, projectID int64, opts *model.DirectoryListOptions) (
	[]*model.Directory, *Response, error,
) {
	res := new(model.DirectoryListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/directories", projectID), opts, res)

	dir := make([]*model.Directory, 0, len(res.Data))
	for _, d := range res.Data {
		dir = append(dir, d.Data)
	}

	return dir, resp, err
}

// GetDirectory returns a single directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.get
func (s *SourceFilesService) GetDirectory(ctx context.Context, projectID, directoryID int64) (*model.Directory, *Response, error) {
	res := new(model.DirectoryGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID), nil, res)

	return res.Data, resp, err
}

// AddDirectory creates a new directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.post
func (s *SourceFilesService) AddDirectory(ctx context.Context, projectID int64, req *model.DirectoryAddRequest) (
	*model.Directory, *Response, error,
) {
	res := new(model.DirectoryGetResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/directories", projectID), req, res)

	return res.Data, resp, err
}

// EditDirectory updates a directory in the project.
//
// Request body:
//   - op: The operation to perform. Enum: replace, test.
//   - path: A JSON Pointer as defined in RFC 6901.
//     Enum: "/branchId", "/directoryId", "/name", "/title", "/exportPattern", "/priority".
//   - value: The value to be used within the operations. The value must be one of string or integer.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.patch
func (s *SourceFilesService) EditDirectory(ctx context.Context, projectID, directoryID int64, req []*model.UpdateRequest) (
	*model.Directory, *Response, error,
) {
	res := new(model.DirectoryGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID), req, res)

	return res.Data, resp, err
}

// DeleteDirectory deletes a directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.delete
func (s *SourceFilesService) DeleteDirectory(ctx context.Context, projectID, directoryID int64) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID))
}

// ListFiles returns a list of files in the project.
//
// Query parameters:
// - branchId: The identifier of the branch (filter by branch).
// - directoryId: The identifier of the directory (filter by directory).
// - filter: Filter files by name.
// - recursion: List files recursively.
// - limit: A maximum number of items to retrieve (default 25, max 500).
// - offset: A starting offset in the collection of items (default 0).
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.getMany
func (s *SourceFilesService) ListFiles(ctx context.Context, projectID int64, opts *model.FileListOptions) (
	[]*model.File, *Response, error,
) {
	res := new(model.FileListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files", projectID), opts, res)

	files := make([]*model.File, 0, len(res.Data))
	for _, f := range res.Data {
		files = append(files, f.Data)
	}

	return files, resp, err
}

// GetFile returns a single file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.get
func (s *SourceFilesService) GetFile(ctx context.Context, projectID, fileID int64) (*model.File, *Response, error) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// AddFile adds a new file to the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.post
func (s *SourceFilesService) AddFile(ctx context.Context, projectID int64, req *model.FileAddRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/files", projectID), req, res)

	return res.Data, resp, err
}

// UpdateRestoreFile updates a file in the project or restores it to one
// of the previous revisions.
// For updating the file, use the `storageId` body parameter.
// For restoring the file, use the `revisionId` body parameter.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.put
func (s *SourceFilesService) UpdateRestoreFile(ctx context.Context, projectID, fileID int64, req *model.FileUpdateRestoreRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Put(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), req, res)

	return res.Data, resp, err
}

// EditFile updates a file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.patch
func (s *SourceFilesService) EditFile(ctx context.Context, projectID, fileID int64, req []*model.UpdateRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), req, res)

	return res.Data, resp, err
}

// DeleteFile deletes a file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.delete
func (s *SourceFilesService) DeleteFile(ctx context.Context, projectID, fileID int64) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID))
}

// DownloadFilePreview returns a download link for a specific file preview.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.preview.get
func (s *SourceFilesService) DownloadFilePreview(ctx context.Context, projectID, fileID int64) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/preview", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// DownloadFile returns a download link for a specific file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.download.get
func (s *SourceFilesService) DownloadFile(ctx context.Context, projectID, fileID int64) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/download", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// ListFileRevisions returns a list of file revisions.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.revisions.getMany
func (s *SourceFilesService) ListFileRevisions(ctx context.Context, projectID, fileID int64, opts *model.ListOptions) (
	[]*model.FileRevision, *Response, error,
) {
	res := new(model.FileRevisionListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/revisions", projectID, fileID), opts, res)

	revisions := make([]*model.FileRevision, 0, len(res.Data))
	for _, rev := range res.Data {
		revisions = append(revisions, rev.Data)
	}

	return revisions, resp, err
}

// GetFileRevision returns a single file revision.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.revisions.get
func (s *SourceFilesService) GetFileRevision(ctx context.Context, projectID, fileID, revisionID int64) (*model.FileRevision, *Response, error) {
	res := new(model.FileRevisionResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/revisions/%d", projectID, fileID, revisionID), nil, res)

	return res.Data, resp, err
}
