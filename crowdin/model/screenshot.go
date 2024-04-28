package model

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Screenshot struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	WebURL string `json:"webUrl"`
	Name   string `json:"name"`
	Size   struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"size"`
	TagsCount int    `json:"tagsCount"`
	Tags      []*Tag `json:"tags"`
	LabelIds  []int  `json:"labelIds"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ScreenshotResponse struct {
	Data *Screenshot `json:"data"`
}

type ScreenshotListResponse struct {
	Data []*ScreenshotResponse `json:"data"`
}

type Tag struct {
	ID           int          `json:"id"`
	ScreenshotID int          `json:"screenshotId"`
	StringID     int          `json:"stringId"`
	Position     *TagPosition `json:"position"`
	CreatedAt    string       `json:"createdAt"`
}

// TagPosition represents the position of a tag on a screenshot.
type TagPosition struct {
	X      *int `json:"x,omitempty"`
	Y      *int `json:"y,omitempty"`
	Width  *int `json:"width,omitempty"`
	Height *int `json:"height,omitempty"`
}

type TagResponse struct {
	Data *Tag `json:"data"`
}

type TagListResponse struct {
	Data []*TagResponse `json:"data"`
}

type ScreenshotListOptions struct {
	// Sort screenshots by specified field.
	// Enum: id, name, tagsCount, createdAt, updatedAt. Default: id.
	// Example: orderBy=createdAt desc,name,tagsCount
	OrderBy string `url:"orderBy,omitempty"`
	// String Identifier.
	StringID int `url:"stringId,omitempty"`
	// Label Identifiers.
	// Example: labelIds=1,2,3
	LabelIDs []string `url:"labelIds,omitempty"`
	// Label Identifiers to exclude.
	// Example: excludeLabelIds=1,2,3
	ExcludeLabelIDs []string `url:"excludeLabelIds,omitempty"`

	ListOptions
}

func (o *ScreenshotListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if len(o.OrderBy) > 0 {
		v.Add("orderBy", o.OrderBy)
	}
	if o.StringID > 0 {
		v.Add("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if len(o.LabelIDs) > 0 {
		v.Add("labelIds", strings.Join(o.LabelIDs, ","))
	}
	if len(o.ExcludeLabelIDs) > 0 {
		v.Add("excludeLabelIds", strings.Join(o.ExcludeLabelIDs, ","))
	}

	return v, len(v) > 0
}

type ScreenshotAddRequest struct {
	// Storage Identifier. Storage file must be image in one of the
	// following formats: jpeg, jpg, png, gif
	StorageID int `json:"storageId"`
	// Screenshot name.
	Name string `json:"name"`
	// Automatically tags screenshot.
	AutoTag *bool `json:"autoTag,omitempty"`
	// File Identifier.
	// Note: Must be used together with `autoTag`. Can't be used
	// with `directoryId` or `branchId` in same request.
	FileID int `json:"fileId,omitempty"`
	// Branch identifier.
	// Note: Must be used together with `autoTag`. Can't be used
	// with `fileId` or `directoryId` in the same request.
	BranchID int `json:"branchId,omitempty"`
	// Directory Identifier.
	// Note: Must be used together with `autoTag`. Can't be used
	// with `fileId` or `branchId` in same request.
	DirectoryID int `json:"directoryId,omitempty"`
	// Label Identifiers.
	LabelIDs []int `json:"labelIds,omitempty"`
}

func (r *ScreenshotAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID <= 0 {
		return errors.New("storageId is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.FileID > 0 && (r.BranchID > 0 || r.DirectoryID > 0) {
		return errors.New("must use either branchId, fileId, or directoryId")
	}
	if r.BranchID > 0 && (r.FileID > 0 || r.DirectoryID > 0) {
		return errors.New("must use either branchId, fileId, or directoryId")
	}

	return nil
}

type ScreenshotUpdateRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// Screenshot name.
	Name string `json:"name"`
}

func (r *ScreenshotUpdateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID <= 0 {
		return errors.New("storageId is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

type TagAddRequest struct {
	// String Identifier.
	StringID int `json:"stringId"`
	// Tag position.
	Position *TagPosition `json:"position,omitempty"`
}

func (r *TagAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StringID <= 0 {
		return errors.New("stringId is required")
	}

	return nil
}

type ReplaceTagsRequest struct {
	// String Identifier.
	StringID int `json:"stringId"`
	// Tag position.
	Position *TagPosition `json:"position,omitempty"`
}

func (r *ReplaceTagsRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StringID <= 0 {
		return errors.New("stringId is required")
	}

	return nil
}

type AutoTagRequest struct {
	// Automatically tags screenshot and replaces old tags with new ones.
	AutoTag *bool `json:"autoTag"`
	// File Identifier. Note: Can't be used with `directoryId` or `branchId`
	// in same request.
	FileID int `json:"fileId,omitempty"`
	// Branch identifier. Note: Can't be used with `fileId` or `directoryId`
	// in the same request.
	BranchID int `json:"branchId,omitempty"`
	// Directory Identifier. Note: Can't be used with `fileId` or `branchId`
	// in same request.
	DirectoryID int `json:"directoryId,omitempty"`
}

func (r *AutoTagRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.AutoTag == nil {
		return errors.New("autoTag is required")
	}
	if r.FileID > 0 && (r.BranchID > 0 || r.DirectoryID > 0) {
		return errors.New("must use either branchId, fileId, or directoryId")
	}
	if r.BranchID > 0 && (r.FileID > 0 || r.DirectoryID > 0) {
		return errors.New("must use either branchId, fileId, or directoryId")
	}

	return nil
}
