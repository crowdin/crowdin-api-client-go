package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *ProjectsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &ProjectsListOptions{},
		},
		{
			name: "with hasManagerAccess = 0",
			opts: &ProjectsListOptions{HasManagerAccess: toPtr(0)},
			out:  "hasManagerAccess=0",
		},
		{
			name: "with type = 0",
			opts: &ProjectsListOptions{Type: toPtr(0)},
			out:  "type=0",
		},
		{
			name: "with all options",
			opts: &ProjectsListOptions{OrderBy: "createdAt desc,name,id", UserID: 1, HasManagerAccess: toPtr(1),
				Type: toPtr(1), ListOptions: ListOptions{Offset: 1, Limit: 10}},
			out: "hasManagerAccess=1&limit=10&offset=1&orderBy=createdAt+desc%2Cname%2Cid&type=1&userId=1",
		},
		{
			name: "with filter",
			opts: &ProjectsListOptions{Filter: "My Project"},
			out:  "filter=My+Project",
		},
		{
			name: "with groupId = 0",
			opts: &ProjectsListOptions{GroupID: toPtr(0)},
			out:  "groupId=0",
		},
		{
			name: "with filter and groupId",
			opts: &ProjectsListOptions{Filter: "app", GroupID: toPtr(3), HasManagerAccess: toPtr(1)},
			out:  "filter=app&groupId=3&hasManagerAccess=1",
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

func TestProjectsAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ProjectsAddRequest
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
			req:  &ProjectsAddRequest{},
			err:  "name is required",
		},
		{
			name: "empty request",
			req:  &ProjectsAddRequest{SourceLanguageID: "en"},
			err:  "name is required",
		},
		{
			name: "missing sourceLanguageId",
			req:  &ProjectsAddRequest{Name: "Knowledge Base"},
			err:  "sourceLanguageId is required",
		},
		{
			name: "valid request",
			req: &ProjectsAddRequest{
				Name:              "Knowledge Base",
				SourceLanguageID:  "en",
				TargetLanguageIDs: []string{"fr"},
				Fields: map[string]any{
					"key_1": "value_1",
					"key_2": 2,
					"key_3": true,
					"key_4": []string{"en", "uk"},
				},
			},
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

func TestProjectsAddFileFormatSettingsRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ProjectsAddFileFormatSettingsRequest
		err   string
		valid bool
	}{
		{
			name: "Nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "Empty request",
			req:  &ProjectsAddFileFormatSettingsRequest{},
			err:  "format is required",
		},
		{
			name: "Empty format",
			req:  &ProjectsAddFileFormatSettingsRequest{Format: ""},
			err:  "format is required",
		},
		{
			name: "Empty settings",
			req:  &ProjectsAddFileFormatSettingsRequest{Format: "android"},
			err:  "settings is required",
		},
		{
			name: "Valid request (common file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &CommonFileFormatSettings{ContentSegmentation: toPtr(false),
					SRXStorageID: toPtr(1), ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (property file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &PropertyFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (xml file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &XMLFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (media wiki file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &MediaWikiFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (txt file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &TXTFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (javascript file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &JavaScriptFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (string catalog file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &StringCatalogFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (other file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "android",
				Settings: &OtherFileFormatSettings{ExportPattern: toPtr("pattern")}},
			valid: true,
		},
		{
			name: "Valid request (vdf file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "vdf",
				Settings: &VDFFileFormatSettings{ConvertICU: toPtr(false), AddGenderArgument: toPtr(true)}},
			valid: true,
		},
		{
			name: "Valid request (md file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "md",
				Settings: &MDFileFormatSettings{StrongMarker: "underscore", InlineTags: []string{"span"}}},
			valid: true,
		},
		{
			name: "Valid request (fm html file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "fm_html",
				Settings: &FMHTMLFileFormatSettings{ExcludedFrontMatterElements: []string{"title"}}},
			valid: true,
		},
		{
			name: "Valid request (web xml file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "webxml",
				Settings: &WebXMLFileFormatSettings{InlineTags: []string{"b"}}},
			valid: true,
		},
		{
			name: "Valid request (idml file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "idml",
				Settings: &IDMLFileFormatSettings{InlineHyperlinkText: toPtr(true)}},
			valid: true,
		},
		{
			name: "Valid request (xliff file format settings)",
			req: &ProjectsAddFileFormatSettingsRequest{Format: "xliff",
				Settings: &XLIFFFileFormatSettings{
					CommonFileFormatSettings: CommonFileFormatSettings{ContentSegmentation: toPtr(false)},
				}},
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

func TestProjectsStringsExporterSettingsRequestValidate(t *testing.T) {
	cases := []struct {
		name  string
		req   *ProjectsStringsExporterSettingsRequest
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
			req:  &ProjectsStringsExporterSettingsRequest{},
			err:  "format is required",
		},
		{
			name: "empty format",
			req:  &ProjectsStringsExporterSettingsRequest{},
			err:  "format is required",
		},
		{
			name: "empty settings",
			req:  &ProjectsStringsExporterSettingsRequest{Settings: StringsExporterSettings{}},
			err:  "format is required",
		},
		{
			name: "empty language pair mapping",
			req:  &ProjectsStringsExporterSettingsRequest{Format: "macosx"},
			err:  "settings is required",
		},
		{
			name: "empty language pair mapping",
			req: &ProjectsStringsExporterSettingsRequest{Format: "xliff",
				Settings: StringsExporterSettings{LanguagePairMapping: map[string]string{}}},
			err: "settings is required",
		},
		{
			name: "valid request",
			req: &ProjectsStringsExporterSettingsRequest{Format: "xliff",
				Settings: StringsExporterSettings{ConvertPlaceholders: toPtr(false), LanguagePairMapping: map[string]string{"en": "fr"}}},
			valid: true,
		},
		{
			name: "valid request (convertLineBreaks only)",
			req: &ProjectsStringsExporterSettingsRequest{Format: "android",
				Settings: StringsExporterSettings{ConvertLineBreaks: toPtr(true)}},
			valid: true,
		},
		{
			name: "valid request (useCdataForStringsWithTags only)",
			req: &ProjectsStringsExporterSettingsRequest{Format: "android",
				Settings: StringsExporterSettings{UseCdataForStringsWithTags: toPtr(true)}},
			valid: true,
		},
		{
			name: "valid request (exportContext only)",
			req: &ProjectsStringsExporterSettingsRequest{Format: "macosx",
				Settings: StringsExporterSettings{ExportContext: toPtr(false)}},
			valid: true,
		},
		{
			name: "valid request (copySourceToEmptyTarget only)",
			req: &ProjectsStringsExporterSettingsRequest{Format: "xliff",
				Settings: StringsExporterSettings{CopySourceToEmptyTarget: toPtr(false)}},
			valid: true,
		},
		{
			name: "valid request (exportTranslatorsComment only)",
			req: &ProjectsStringsExporterSettingsRequest{Format: "xliff",
				Settings: StringsExporterSettings{ExportTranslatorsComment: toPtr(true)}},
			valid: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.valid {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
