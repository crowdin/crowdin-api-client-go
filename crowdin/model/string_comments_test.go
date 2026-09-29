package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCommentsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *StringCommentsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
			out:  "",
		},
		{
			name: "empty options",
			opts: &StringCommentsListOptions{},
			out:  "",
		},
		{
			name: "with all options",
			opts: &StringCommentsListOptions{OrderBy: "createdAt desc,text", StringID: 1, Type: "comment",
				IssueType: []string{"general_question", "translation_mistake"}, IssueStatus: "resolved",
				ListOptions: ListOptions{Offset: 1, Limit: 10},
			},
			out: "issueStatus=resolved&issueType=general_question%2Ctranslation_mistake&limit=10&offset=1&orderBy=createdAt+desc%2Ctext&stringId=1&type=comment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opts.Values()
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

func TestStringCommentsAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *StringCommentsAddRequest
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
			req:  &StringCommentsAddRequest{},
			err:  "text is required",
		},
		{
			name: "missing stringId",
			req:  &StringCommentsAddRequest{Text: "test text"},
			err:  "stringId or fileId is required",
		},
		{
			name: "both stringId and fileId",
			req:  &StringCommentsAddRequest{Text: "test text", StringID: 1, FileID: 2},
			err:  "stringId and fileId cannot be used in the same request",
		},
		{
			name: "missing targetLanguageId",
			req:  &StringCommentsAddRequest{Text: "test text", StringID: 1},
			err:  "targetLanguageId is required",
		},
		{
			name: "missing type",
			req:  &StringCommentsAddRequest{Text: "test text", StringID: 1, TargetLanguageID: "en"},
			err:  "type is required",
		},
		{
			name: "attachment without id",
			req: &StringCommentsAddRequest{Text: "test text", StringID: 1, TargetLanguageID: "en", Type: "comment",
				Attachments: []*StringCommentAttachmentRequest{{ID: 1}, {}}},
			err: "attachment id is required",
		},
		{
			name: "nil attachment",
			req: &StringCommentsAddRequest{Text: "test text", StringID: 1, TargetLanguageID: "en", Type: "comment",
				Attachments: []*StringCommentAttachmentRequest{nil}},
			err: "attachment id is required",
		},
		{
			name:  "valid request",
			req:   &StringCommentsAddRequest{Text: "test text", StringID: 1, TargetLanguageID: "en", Type: "comment"},
			valid: true,
		},
		{
			name: "valid request with attachments",
			req: &StringCommentsAddRequest{Text: "test text", StringID: 1, TargetLanguageID: "en", Type: "comment",
				Attachments: []*StringCommentAttachmentRequest{{ID: 1}, {ID: 2}}},
			valid: true,
		},
		{
			name:  "valid asset comment request",
			req:   &StringCommentsAddRequest{Text: "test text", FileID: 2, TargetLanguageID: "en", Type: "issue"},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.valid {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
