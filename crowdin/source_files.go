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

// ListDirectories returns a list of directories in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.getMany
func (s *SourceFilesService) ListDirectories(ctx context.Context, projectID int, opts *model.DirectoryListOptions) (
	[]*model.Directory, *Response, error,
) {
	res := new(model.DirectoryListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/directories", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	dir := make([]*model.Directory, 0, len(res.Data))
	for _, d := range res.Data {
		dir = append(dir, d.Data)
	}

	return dir, resp, err
}

// GetDirectory returns a single directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.get
func (s *SourceFilesService) GetDirectory(ctx context.Context, projectID, directoryID int) (*model.Directory, *Response, error) {
	res := new(model.DirectoryGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID), nil, res)

	return res.Data, resp, err
}

// AddDirectory creates a new directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.post
func (s *SourceFilesService) AddDirectory(ctx context.Context, projectID int, req *model.DirectoryAddRequest) (
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
func (s *SourceFilesService) EditDirectory(ctx context.Context, projectID, directoryID int, req []*model.UpdateRequest) (
	*model.Directory, *Response, error,
) {
	res := new(model.DirectoryGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID), req, res)

	return res.Data, resp, err
}

// DeleteDirectory deletes a directory in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.delete
func (s *SourceFilesService) DeleteDirectory(ctx context.Context, projectID, directoryID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID), nil)
}

// DeleteDirectoryAsync deletes a directory in the project asynchronously.
// It sends the `Prefer: respond-async` header and returns a delete job.
// Use CheckDirectoryDeleteStatus to poll the status of the job.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.delete
func (s *SourceFilesService) DeleteDirectoryAsync(ctx context.Context, projectID, directoryID int) (
	*model.NodeDeleteJob, *Response, error,
) {
	return s.client.deleteAsync(ctx, fmt.Sprintf("/api/v2/projects/%d/directories/%d", projectID, directoryID))
}

// CheckDirectoryDeleteStatus checks the status of an asynchronous directory deletion.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.directories.jobs.get
func (s *SourceFilesService) CheckDirectoryDeleteStatus(ctx context.Context, projectID, directoryID int, jobID string) (
	*model.NodeDeleteJob, *Response, error,
) {
	path := fmt.Sprintf("/api/v2/projects/%d/directories/%d/jobs/%s", projectID, directoryID, jobID)
	res := new(model.NodeDeleteJobResponse)
	resp, err := s.client.Get(ctx, path, nil, res)

	return res.Data, resp, err
}

// SearchDirectories searches directories by name or title across multiple projects.
// The `filter` option is required.
//
// https://developer.crowdin.com/api/v2/#operation/api.directories.getMany
func (s *SourceFilesService) SearchDirectories(ctx context.Context, opts *model.DirectoriesSearchOptions) (
	[]*model.Directory, *Response, error,
) {
	if err := opts.Validate(); err != nil {
		return nil, nil, err
	}

	res := new(model.DirectoryListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/directories", opts, res)
	if err != nil {
		return nil, resp, err
	}

	dirs := make([]*model.Directory, 0, len(res.Data))
	for _, d := range res.Data {
		dirs = append(dirs, d.Data)
	}

	return dirs, resp, nil
}

// ListFiles returns a list of files in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.getMany
func (s *SourceFilesService) ListFiles(ctx context.Context, projectID int, opts *model.FileListOptions) (
	[]*model.File, *Response, error,
) {
	res := new(model.FileListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	files := make([]*model.File, 0, len(res.Data))
	for _, f := range res.Data {
		files = append(files, f.Data)
	}

	return files, resp, err
}

// GetFile returns a single file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.get
func (s *SourceFilesService) GetFile(ctx context.Context, projectID, fileID int) (*model.File, *Response, error) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// AddFile adds a new file to the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.post
func (s *SourceFilesService) AddFile(ctx context.Context, projectID int, req *model.FileAddRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/files", projectID), req, res)

	return res.Data, resp, err
}

// UpdateOrRestoreFile updates a file in the project or restores it to one
// of the previous revisions.
// For updating the file, use the `storageId` body parameter.
// For restoring the file, use the `revisionId` body parameter.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.put
func (s *SourceFilesService) UpdateOrRestoreFile(ctx context.Context, projectID, fileID int, req *model.FileUpdateRestoreRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Put(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), req, res)

	return res.Data, resp, err
}

// EditFile updates a file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.patch
func (s *SourceFilesService) EditFile(ctx context.Context, projectID, fileID int, req []*model.UpdateRequest) (
	*model.File, *Response, error,
) {
	res := new(model.FileGetResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), req, res)

	return res.Data, resp, err
}

// DeleteFile deletes a file in the project.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.delete
func (s *SourceFilesService) DeleteFile(ctx context.Context, projectID, fileID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID), nil)
}

// DeleteFileAsync deletes a file in the project asynchronously.
// It sends the `Prefer: respond-async` header and returns a delete job.
// Use CheckFileDeleteStatus to poll the status of the job.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.delete
func (s *SourceFilesService) DeleteFileAsync(ctx context.Context, projectID, fileID int) (*model.NodeDeleteJob, *Response, error) {
	return s.client.deleteAsync(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d", projectID, fileID))
}

// CheckFileDeleteStatus checks the status of an asynchronous file deletion.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.jobs.get
func (s *SourceFilesService) CheckFileDeleteStatus(ctx context.Context, projectID, fileID int, jobID string) (
	*model.NodeDeleteJob, *Response, error,
) {
	path := fmt.Sprintf("/api/v2/projects/%d/files/%d/jobs/%s", projectID, fileID, jobID)
	res := new(model.NodeDeleteJobResponse)
	resp, err := s.client.Get(ctx, path, nil, res)

	return res.Data, resp, err
}

// SearchFiles searches files by name or title across multiple projects.
// The `filter` option is required.
//
// https://developer.crowdin.com/api/v2/#operation/api.files.getMany
func (s *SourceFilesService) SearchFiles(ctx context.Context, opts *model.FilesSearchOptions) (
	[]*model.File, *Response, error,
) {
	if err := opts.Validate(); err != nil {
		return nil, nil, err
	}

	res := new(model.FileListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/files", opts, res)
	if err != nil {
		return nil, resp, err
	}

	files := make([]*model.File, 0, len(res.Data))
	for _, f := range res.Data {
		files = append(files, f.Data)
	}

	return files, resp, nil
}

// ListAssetReferences returns a list of references of an asset file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.references.getMany
func (s *SourceFilesService) ListAssetReferences(ctx context.Context, projectID, fileID int, opts *model.ListOptions) (
	[]*model.AssetReference, *Response, error,
) {
	res := new(model.AssetReferenceListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/references", projectID, fileID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.AssetReference, 0, len(res.Data))
	for _, ref := range res.Data {
		list = append(list, ref.Data)
	}

	return list, resp, nil
}

// GetAssetReference returns a single reference of an asset file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.references.get
func (s *SourceFilesService) GetAssetReference(ctx context.Context, projectID, fileID, referenceID int) (
	*model.AssetReference, *Response, error,
) {
	path := fmt.Sprintf("/api/v2/projects/%d/files/%d/references/%d", projectID, fileID, referenceID)
	res := new(model.AssetReferenceResponse)
	resp, err := s.client.Get(ctx, path, nil, res)

	return res.Data, resp, err
}

// AddAssetReference adds a new reference to an asset file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.references.post
func (s *SourceFilesService) AddAssetReference(ctx context.Context, projectID, fileID int, req *model.AssetReferenceAddRequest) (
	*model.AssetReference, *Response, error,
) {
	res := new(model.AssetReferenceResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/references", projectID, fileID), req, res)

	return res.Data, resp, err
}

// DeleteAssetReference deletes a reference of an asset file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.references.delete
func (s *SourceFilesService) DeleteAssetReference(ctx context.Context, projectID, fileID, referenceID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/references/%d", projectID, fileID, referenceID), nil)
}

// DownloadFilePreview returns a download link for a specific file preview.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.preview.get
func (s *SourceFilesService) DownloadFilePreview(ctx context.Context, projectID, fileID int) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/preview", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// DownloadFile returns a download link for a specific file.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.download.get
func (s *SourceFilesService) DownloadFile(ctx context.Context, projectID, fileID int) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/download", projectID, fileID), nil, res)

	return res.Data, resp, err
}

// ListFileRevisions returns a list of file revisions.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.revisions.getMany
func (s *SourceFilesService) ListFileRevisions(ctx context.Context, projectID, fileID int, opts *model.ListOptions) (
	[]*model.FileRevision, *Response, error,
) {
	res := new(model.FileRevisionListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/revisions", projectID, fileID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	revisions := make([]*model.FileRevision, 0, len(res.Data))
	for _, rev := range res.Data {
		revisions = append(revisions, rev.Data)
	}

	return revisions, resp, err
}

// GetFileRevision returns a single file revision.
//
// https://developer.crowdin.com/api/v2/#operation/api.projects.files.revisions.get
func (s *SourceFilesService) GetFileRevision(ctx context.Context, projectID, fileID, revisionID int) (*model.FileRevision, *Response, error) {
	res := new(model.FileRevisionResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/files/%d/revisions/%d", projectID, fileID, revisionID), nil, res)

	return res.Data, resp, err
}

// ListReviewedBuilds returns a list of reviewed source files builds.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.strings.reviewed-builds.getMany
func (s *SourceFilesService) ListReviewedBuilds(ctx context.Context, projectID int, opts *model.ReviewedBuildListOptions) (
	[]*model.ReviewedBuild, *Response, error,
) {
	res := new(model.ReviewedBuildListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/reviewed-builds", projectID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	builds := make([]*model.ReviewedBuild, 0, len(res.Data))
	for _, b := range res.Data {
		builds = append(builds, b.Data)
	}

	return builds, resp, err
}

// CheckReviewedBuildStatus checks the status of a specific reviewed source files build.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.strings.reviewed-builds.get
func (s *SourceFilesService) CheckReviewedBuildStatus(ctx context.Context, projectID, buildID int) (*model.ReviewedBuild, *Response, error) {
	res := new(model.ReviewedBuildResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/reviewed-builds/%d", projectID, buildID), nil, res)

	return res.Data, resp, err
}

// BuildReviewedFiles starts a new build of reviewed source files.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.strings.reviewed-builds.post
func (s *SourceFilesService) BuildReviewedFiles(ctx context.Context, projectID int, req *model.ReviewedBuildRequest) (
	*model.ReviewedBuild, *Response, error,
) {
	res := new(model.ReviewedBuildResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/reviewed-builds", projectID), req, res)

	return res.Data, resp, err
}

// DownloadReviewedBuild returns a download link for a specific reviewed source files build.
//
// https://developer.crowdin.com/enterprise/api/v2/#operation/api.projects.strings.reviewed-builds.download.download
func (s *SourceFilesService) DownloadReviewedBuild(ctx context.Context, projectID, buildID int) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/strings/reviewed-builds/%d/download", projectID, buildID), nil, res)

	return res.Data, resp, err
}
