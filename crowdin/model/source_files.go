package model

import (
	"errors"
	"fmt"
	"net/url"
)

// Directory represents a project directory.
type Directory struct {
	ID            int    `json:"id"`
	ProjectID     int    `json:"projectId"`
	BranchID      *int   `json:"branchId,omitempty"`
	DirectoryID   *int   `json:"directoryId,omitempty"`
	Name          string `json:"name"`
	Title         string `json:"title"`
	ExportPattern string `json:"exportPattern"`
	Path          string `json:"path"`
	Priority      string `json:"priority"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

// DirectoryGetResponse describes a response with a single directory.
type DirectoryGetResponse struct {
	Data *Directory `json:"data"`
}

// DirectoryListResponse describes a response with a list of directories.
type DirectoryListResponse struct {
	Data []*DirectoryGetResponse `json:"data"`
}

// DirectoryListOptions specifies the optional parameters to the
// SourceFilesService.ListDirectories method.
type DirectoryListOptions struct {
	// OrderBy is used to sort directories.
	// Enum: id, name, title, createdAt, updatedAt, exportPattern, priority.
	// Default: id.
	// Example: orderBy=createdAt desc,name,priority.
	OrderBy string `json:"orderBy,omitempty"`
	// BranchID is the ID of the branch to filter directories by.
	// Note: Can't be used with `directoryID` in the same request.
	// To list the directories from all the nested levels within the branch,
	// ensure to use the `recursion` parameter with the `branchID` parameter.
	BranchID int `json:"branchId,omitempty"`
	// DirectoryID is the ID of the directory to filter directories by.
	// Note: Can't be used with `branchID` in the same request.
	// To list the directories from all the nested levels within the directory,
	// ensure to use the `recursion` parameter with the `directoryID` parameter.
	DirectoryID int `json:"directoryId,omitempty"`
	// Filter directories by name.
	Filter string `json:"filter,omitempty"`
	// Recursion is used to list directories recursively.
	// Note: Works only when `directoryID` or `branchID` parameter is specified.
	Recursion any `json:"recursion,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of DirectoryListOptions.
func (o *DirectoryListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}
	if o.DirectoryID > 0 {
		v.Add("directoryId", fmt.Sprintf("%d", o.DirectoryID))
	}
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if recursion, ok := o.Recursion.(string); ok {
		v.Add("recursion", recursion)
	}

	return v, len(v) > 0
}

// DirectoryAddRequest defines the structure of a request
// to create a new directory.
type DirectoryAddRequest struct {
	// Directory name.
	// Note: Can't contain \ / : * ? " < > | symbols.
	Name string `json:"name"`
	// Branch identifier.
	// Note: Can't be used with `directoryId` in same request.
	BranchID int `json:"branchId,omitempty"`
	// Parent Directory Identifier.
	// Note: Can't be used with `branchId` in same request.
	DirectoryID int `json:"directoryId,omitempty"`
	// Title is used to provide more details for translators.
	// It is available in UI only.
	Title string `json:"title,omitempty"`
	// Directory export pattern. Defines directory name and path in resulting
	// translations bundle.
	// Note: Can't contain : * ? " < > | symbols.
	ExportPattern string `json:"exportPattern,omitempty"`
	// Defines priority level for each branch.
	// Enum: low, normal, high. Default: normal.
	Priority string `json:"priority,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *DirectoryAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.BranchID != 0 && r.DirectoryID != 0 {
		return errors.New("branchId and directoryId cannot be used in the same request")
	}
	return nil
}

// File represents a project file.
type File struct {
	ID          int     `json:"id"`
	ProjectID   int     `json:"projectId"`
	BranchID    *int    `json:"branchId,omitempty"`
	DirectoryID *int    `json:"directoryId,omitempty"`
	Name        string  `json:"name"`
	Title       *string `json:"title,omitempty"`
	Context     *string `json:"context,omitempty"`
	Type        string  `json:"type"`
	Path        string  `json:"path"`
	Status      string  `json:"status"`
	Fields      any     `json:"fields,omitempty"`

	RevisionID             int            `json:"revisionId"`
	Priority               string         `json:"priority"`
	ImportOptions          map[string]any `json:"importOptions,omitempty"`
	ExportOptions          map[string]any `json:"exportOptions,omitempty"`
	ExcludeTargetLanguages []string       `json:"excludedTargetLanguages,omitempty"`
	ParserVersion          *int           `json:"parserVersion,omitempty"`
	CreatedAt              string         `json:"createdAt,omitempty"`
	UpdatedAt              string         `json:"updatedAt,omitempty"`
}

// FileGetResponse describes a response with a single file.
type FileGetResponse struct {
	Data *File `json:"data"`
}

// FileListResponse describes a response with a list of files.
type FileListResponse struct {
	Data []*FileGetResponse `json:"data"`
}

// FileListOptions specifies the optional parameters to the
// SourceFilesService.ListFiles method.
type FileListOptions struct {
	// OrderBy is used to sort files.
	// Enum: id, name, title, status, exportPattern, priority, createdAt, updatedAt.
	// Default: id.
	// Example: orderBy=createdAt desc,name,priority.
	OrderBy string `json:"orderBy,omitempty"`
	// BranchID is the ID of the branch to filter files by.
	// Note: Can't be used with `directoryId` in the same request.
	// To list the files from all the nested levels within the branch,
	// ensure to use the `recursion` parameter with the `branchId` parameter.
	BranchID int `json:"branchId,omitempty"`
	// DirectoryID is the ID of the directory to filter files by.
	// Note: Can't be used with `branchId` in the same request.
	// To list the files from all the nested levels within the directory,
	// ensure to use the `recursion` parameter with the `directoryId` parameter.
	DirectoryID int `json:"directoryId,omitempty"`
	// Filter files by name.
	Filter string `json:"filter,omitempty"`
	// Recursion is used to list files recursively.
	// Note: Works only when `directoryID` or `branchID` parameter is specified.
	Recursion any `json:"recursion,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of FileListOptions.
func (o *FileListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}
	if o.DirectoryID > 0 {
		v.Add("directoryId", fmt.Sprintf("%d", o.DirectoryID))
	}
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if recursion, ok := o.Recursion.(string); ok {
		v.Add("recursion", recursion)
	}

	return v, len(v) > 0
}

// FileAddRequest defines the structure of a request to create a new file.
type FileAddRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// File name.
	// Note: Can't contain \ / : * ? " < > | symbols. ZIP files are not allowed.
	Name string `json:"name"`
	// Branch Identifier — defines branch to which file will be added.
	// Note: Can't be used with directoryId in same request.
	BranchID int `json:"branchId,omitempty"`
	// Directory Identifier — defines directory to which file will be added.
	// Note: Can't be used with branchId in same request.
	DirectoryID int `json:"directoryId,omitempty"`
	// Title is used to provide more details for translators.
	// It is available in UI only.
	Title string `json:"title,omitempty"`
	// Context is used to provide context about whole file.
	Context string `json:"context,omitempty"`
	// Type of the file. Default: auto.
	// Enum: auto, android, macosx, resx, properties, gettext, yaml, php, json,
	//       xml, ini, rc, resw, resjson, qtts, joomla, chrome, dtd, dklang,
	//       flex, nsh, wxl, xliff, xliff_two, html, haml, txt, csv, md, mdx_v1,
	//       mdx_v2, flsnp, fm_html, fm_md, mediawiki, docx, xlsx, sbv, properties_play,
	//       properties_xml, maxthon, go_json, dita, idml, mif, stringsdict, plist, vtt,
	//       vdf, srt, stf, toml, contentful_rt, svg, js, coffee, ts, i18next_json, xaml,
	//       arb, adoc, fbt, webxml, nestjs_i18n.
	Type string `json:"type,omitempty"`
	// Using latest parser version by default.
	// Note: Must be used together with type.
	ParserVersion int `json:"parserVersion,omitempty"`
	// File import options.
	ImportOptions FileImportOptions `json:"importOptions,omitempty"`
	// File export options.
	ExportOptions FileExportOptions `json:"exportOptions,omitempty"`
	// Set Target Languages the file should not be translated into.
	// Do not use this option if the file should be available for all project languages.
	ExcludedTargetLanguages []string `json:"excludedTargetLanguages,omitempty"`
	// Attach labels to strings.
	AttachLabelIDs []int `json:"attachLabelIds,omitempty"`
	// Fields.
	Fields map[string]any `json:"fields,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *FileAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.BranchID > 0 && r.DirectoryID > 0 {
		return errors.New("branchId and directoryId cannot be used in the same request")
	}

	return nil
}

type (
	FileImportOptions interface{ ValidateFileImportOptions() error }

	// SpreadsheetsFileImportOptions implements the FileImportOptions interface.
	SpreadsheetFileImportOptions struct {
		// Defines whether the file includes a first-row header that should not be imported.
		// Default: false.
		FirstLineContainsHeader *bool `json:"firstLineContainsHeader,omitempty"`
		// Defines whether hidden sheets that should be imported. Default: true.
		ImportHiddenSheets *bool `json:"importHiddenSheets,omitempty"`
		// Defines whether hidden rows that should be imported. Default: true.
		ImportHiddenRows *bool `json:"importHiddenRows,omitempty"`
		// Defines whether to import translations identical to sources. Default: false.
		ImportEqSuggestions *bool `json:"importEqSuggestions,omitempty"`
		// Defines whether to auto-approve imported translations. Default: false.
		AutoApproveImported *bool `json:"autoApproveImported,omitempty"`
		// Defines whether to import translations for hidden strings. Default: false.
		TranslateHidden *bool `json:"translateHidden,omitempty"`
		// Defines whether to add imported translations to TM. Default: true.
		AddToTM *bool `json:"addToTm,omitempty"`
		// Defines whether to import translations from the file. Default: false.
		ImportTranslations *bool `json:"importTranslations,omitempty"`
		// Defines data columns mapping. The column numbering starts at 0.
		// Acceptable values are: none, identifier, sourcePhrase, sourceOrTranslation,
		// translation, context, maxLength, labels and specified languages (ex. "en", "uk").
		Scheme map[string]int `json:"scheme,omitempty"`

		// Important: ContentSegmentation option disables the possibility to upload existing translations
		// for Spreadsheet files when enabled.
		CommonFileImportOptions
	}

	// XMLFileImportOptions implements the FileImportOptions interface.
	XMLFileImportOptions struct {
		// Defines whether to translate texts placed inside the tags. Default: true.
		TranslateContent *bool `json:"translateContent,omitempty"`
		// Defines whether to translate tags attributes. Default: true.
		TranslateAttributes *bool `json:"translateAttributes,omitempty"`
		// Specify tags that should be considered inline.
		InlineTags []string `json:"inlineTags,omitempty"`
		// This is an array of strings, where each item is the XPaths to DOM element that should be imported.
		TranslatableElements []string `json:"translatableElements,omitempty"`

		// Important: ContentSegmentation option disables the possibility to upload existing translations
		// for XML files when enabled.
		CommonFileImportOptions
	}

	// DOCXFileImportOptions implements the FileImportOptions interface.
	DOCXFileImportOptions struct {
		// When checked, strips additional formatting tags related to text spacing. Default: false
		// Note: Works only for files with the following extensions: *.docx, *.dotx, *.docm,
		//       *.dotm, *.xlsx, *.xltx, *.xlsm, *.xltm, *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		CleanTagsAggressively *bool `json:"cleanTagsAggressively,omitempty"`
		// When checked, exposes hidden text for translation. Default: false
		// Note: Works only for files with the following extensions: *.docx, *.dotx, *.docm, *.dotm.
		TranslateHiddenText *bool `json:"translateHiddenText,omitempty"`
		// When checked, exposes hidden hyperlinks for translation. Default: false
		// Note: Works only for files with the following extensions: *.docx, *.dotx, *.docm, *.dotm,
		//       *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		TranslateHyperlinkURLs *bool `json:"translateHyperlinkUrls,omitempty"`
		// When checked, exposes hidden rows and columns for translation. Default: false
		// Note: Works only for files with the following extensions: *.xlsx, *.xltx, *.xlsm, *.xltm.
		TranslateHiddenRowsAndColumns *bool `json:"translateHiddenRowsAndColumns,omitempty"`
		// When checked, expose slide notes for translation. Default: true
		// Note: Works only for files with the following extensions: *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		ImportNotes *bool `json:"importNotes,omitempty"`
		// When checked, exposes hidden slides for translation. Default: false
		// Note: Works only for files with the following extensions: *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		ImportHiddenSlides *bool `json:"importHiddenSlides,omitempty"`
		// When checked, exposes document properties such as title, subject, and creator
		// for translation. Default: false.
		TranslateDocProperties *bool `json:"translateDocProperties,omitempty"`
		// When checked, exposes comments for translation. Default: true.
		TranslateComments *bool `json:"translateComments,omitempty"`
		// When checked, formatting that differs only in whitespace no longer produces
		// separate tags. Default: false.
		// Note: Can be set only when `cleanTagsAggressively` is enabled.
		IgnoreWhitespaceStyles *bool `json:"ignoreWhitespaceStyles,omitempty"`
		// When checked, a tab inside a text run is exposed as a character instead of a tag.
		// Default: false.
		AddTabAsCharacter *bool `json:"addTabAsCharacter,omitempty"`
		// When checked, a line break inside a text run is exposed as `lineSeparatorReplacement`
		// instead of a tag. Default: false.
		AddLineSeparatorAsCharacter *bool `json:"addLineSeparatorAsCharacter,omitempty"`
		// Character that replaces a line break when `addLineSeparatorAsCharacter` is enabled.
		// Note: Must be exactly one character.
		LineSeparatorReplacement string `json:"lineSeparatorReplacement,omitempty"`
		// When checked, a non-breaking hyphen is exposed as a character instead of a tag.
		// Default: false.
		ReplaceNoBreakHyphenTag *bool `json:"replaceNoBreakHyphenTag,omitempty"`
		// When checked, soft hyphens are removed instead of being exposed as tags.
		// Default: false.
		IgnoreSoftHyphenTag *bool `json:"ignoreSoftHyphenTag,omitempty"`
		// Word complex field types whose text is exposed for translation.
		// Note: `HYPERLINK` is always extracted and can't be removed.
		ComplexFieldDefinitionsToExtract []string `json:"complexFieldDefinitionsToExtract,omitempty"`
		// When checked, exposes headers and footers for translation. Default: true.
		TranslateWordHeadersFooters *bool `json:"translateWordHeadersFooters,omitempty"`
		// When checked, exposes the name of a shape or image for translation. Default: true.
		TranslateWordGraphicName *bool `json:"translateWordGraphicName,omitempty"`
		// When checked, exposes the alternative text of a shape or image for translation.
		// Default: false.
		TranslateWordGraphicDescription *bool `json:"translateWordGraphicDescription,omitempty"`
		// When checked, excludes text whose font color falls between
		// `wordFontColorsMinIgnoranceThreshold` and `wordFontColorsMaxIgnoranceThreshold`
		// from translation. Default: false.
		IgnoreWordFontColors *bool `json:"ignoreWordFontColors,omitempty"`
		// Darkest font color ignored when `ignoreWordFontColors` is enabled, as six hexadecimal
		// digits with an optional leading hash. Empty means black.
		WordFontColorsMinIgnoranceThreshold string `json:"wordFontColorsMinIgnoranceThreshold,omitempty"`
		// Lightest font color ignored when `ignoreWordFontColors` is enabled, as six hexadecimal
		// digits with an optional leading hash. Empty means white.
		WordFontColorsMaxIgnoranceThreshold string `json:"wordFontColorsMaxIgnoranceThreshold,omitempty"`
		// Paragraph style names singled out by `translateWordInExcludeStyleMode`,
		// for example `Heading1`.
		ExcludeWordStyles []string `json:"excludeWordStyles,omitempty"`
		// How `excludeWordStyles` is read. If true, the listed styles are excluded from
		// translation; if false, only they are translated. Default: true.
		TranslateWordInExcludeStyleMode *bool `json:"translateWordInExcludeStyleMode,omitempty"`
		// Word highlight colors singled out by `translateWordInExcludeHighlightMode`.
		WordHighlightColors []string `json:"wordHighlightColors,omitempty"`
		// How `wordHighlightColors` is read. If true, text with the listed highlights is
		// excluded from translation; if false, only it is translated. Default: true.
		TranslateWordInExcludeHighlightMode *bool `json:"translateWordInExcludeHighlightMode,omitempty"`
		// When checked, excludes text whose color is listed in `wordExcludedColors`
		// from translation. Default: false.
		TranslateWordExcludeColors *bool `json:"translateWordExcludeColors,omitempty"`
		// Text colors excluded from translation when `translateWordExcludeColors` is enabled,
		// each six hexadecimal digits with an optional leading hash.
		WordExcludedColors []string `json:"wordExcludedColors,omitempty"`
		// When checked, a cell whose text is shared with another cell is exposed once
		// for translation. Default: true.
		TranslateExcelCellsCopied *bool `json:"translateExcelCellsCopied,omitempty"`
		// When checked, exposes worksheet names for translation. Default: false.
		TranslateExcelSheetNames *bool `json:"translateExcelSheetNames,omitempty"`
		// Cell fill colors whose cells are excluded from translation, each the six-digit RGB
		// with an optional leading hash.
		ExcelExcludedColors []string `json:"excelExcludedColors,omitempty"`
		// When checked, exposes text inside SmartArt and diagram data for translation.
		// Default: false.
		TranslateExcelDiagramData *bool `json:"translateExcelDiagramData,omitempty"`
		// When checked, exposes text inside drawings and shapes on a worksheet for translation.
		// Default: false.
		TranslateExcelDrawings *bool `json:"translateExcelDrawings,omitempty"`

		// Important: ContentSegmentation option disables the possibility to upload existing translations
		// for XML files when enabled.
		CommonFileImportOptions
	}

	// WebXMLFileImportOptions implements the FileImportOptions interface.
	WebXMLFileImportOptions struct {
		// Specify tags that should be considered inline.
		InlineTags []string `json:"inlineTags,omitempty"`
		// HTML attributes (e.g., src, href, data) will be imported as hidden strings.
		// Default: false.
		HideAttributeValues *bool `json:"hideAttributeValues,omitempty"`

		CommonFileImportOptions
	}

	// VSDXFileImportOptions implements the FileImportOptions interface.
	VSDXFileImportOptions struct {
		// When checked, strips additional formatting tags related to text spacing.
		// Default: false.
		CleanTagsAggressively *bool `json:"cleanTagsAggressively,omitempty"`
		// When checked, exposes hidden hyperlinks for translation. Default: false.
		TranslateHyperlinkURLs *bool `json:"translateHyperlinkUrls,omitempty"`

		CommonFileImportOptions
	}

	// IDMLFileImportOptions implements the FileImportOptions interface.
	IDMLFileImportOptions struct {
		// When checked, a hyperlink's text is extracted as part of the surrounding
		// sentence so it can be translated in context. Default: false.
		InlineHyperlinkText *bool `json:"inlineHyperlinkText,omitempty"`

		CommonFileImportOptions
	}

	// HTMLFileImportOptions implements the FileImportOptions interface.
	HTMLFileImportOptions struct {
		// Specify CSS selectors for elements that should not be imported.
		ExcludedElements []string `json:"excludedElements,omitempty"`
		// Specify tags that should be considered inline.
		InlineTags []string `json:"inlineTags,omitempty"`
		// HTML attributes (e.g., src, href, data) will be imported as hidden strings.
		// Default: false.
		HideAttributeValues *bool `json:"hideAttributeValues,omitempty"`

		CommonFileImportOptions
	}

	// HTMLWithFrontMatterFileImportOptions implements the FileImportOptions interface.
	HTMLWithFrontMatterFileImportOptions struct {
		// Specify CSS selectors for elements that should not be imported.
		ExcludedElements []string `json:"excludedElements,omitempty"`
		// Specify elements that should not be imported.
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Specify tags that should be considered inline.
		InlineTags []string `json:"inlineTags,omitempty"`
		// HTML attributes (e.g., src, href, data) will be imported as hidden strings.
		// Default: false.
		HideAttributeValues *bool `json:"hideAttributeValues,omitempty"`

		CommonFileImportOptions
	}

	// MDXV1FileImportOptions implements the FileImportOptions interface.
	MDXV1FileImportOptions struct {
		// Specify elements that should not be imported
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Defines whether to import code blocks. Default: false.
		ExcludeCodeBlocks *bool `json:"excludeCodeBlocks,omitempty"`

		CommonFileImportOptions
	}

	// MDXV2FileImportOptions implements the FileImportOptions interface.
	MDXV2FileImportOptions struct {
		// Specify elements that should not be imported
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Defines whether to import code blocks. Default: false.
		ExcludeCodeBlocks *bool `json:"excludeCodeBlocks,omitempty"`

		CommonFileImportOptions
	}

	// MDFileImportOptions implements the FileImportOptions interface.
	MDFileImportOptions struct {
		// Specify elements that should not be imported.
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Defines whether to import code blocks. Default: false.
		ExcludeCodeBlocks *bool `json:"excludeCodeBlocks,omitempty"`
		// Specify tags that should be considered inline.
		InlineTags []string `json:"inlineTags,omitempty"`

		CommonFileImportOptions
	}

	// StringCatalogFileImportOptions implements the FileImportOptions interface.
	StringCatalogFileImportOptions struct {
		// Determines whether to import the key as source string if it does not exist.
		// Default: true.
		ImportKeyAsSource *bool `json:"importKeyAsSource,omitempty"`
		// Defines whether to import translations from the file. Default: false.
		ImportTranslations *bool `json:"importTranslations,omitempty"`
	}

	// VDFFileImportOptions implements the FileImportOptions interface.
	VDFFileImportOptions struct {
		// Defines whether to convert placeholders to the ICU MessageFormat syntax.
		// Default: true.
		ConvertICU *bool `json:"convertIcu,omitempty"`
		// Defines whether to append _gender suffix to variable names when converting
		// gender ICU to ICU MessageFormat. Default: false.
		// Note: Can be enabled only when `convertIcu` is enabled.
		AddGenderArgument *bool `json:"addGenderArgument,omitempty"`
	}

	// AdocFileImportOptions implements the FileImportOptions interface.
	AdocFileImportOptions struct {
		// Skip Include Directives. Default: false.
		ExcludeIncludeDirectives *bool `json:"excludeIncludeDirectives,omitempty"`
	}

	// OtherFileImportOptions implements the FileImportOptions interface.
	OtherFileImportOptions struct {
		// Only for xml, md, flsnp, docx, mif, idml, dita, android8 files.
		//
		// Note: When Content segmentation is enabled, the translation upload is handled by an
		// experimental machine learning technology. To achieve the best results, we recommend
		// uploading translation files with the same or as close as possible file structure
		// as in source files.
		CommonFileImportOptions
	}

	// CommonFileImportOptions implements the FileImportOptions interface.
	CommonFileImportOptions struct {
		// Defines whether to split long texts into smaller text segments. Default: true.
		ContentSegmentation *bool `json:"contentSegmentation,omitempty"`
		// Storage identifier of the SRX segmentation rules file. Default: null.
		SRXStorageID *int `json:"srxStorageId,omitempty"`
	}
)

func (o *CommonFileImportOptions) ValidateFileImportOptions() error        { return nil }
func (o *AdocFileImportOptions) ValidateFileImportOptions() error          { return nil }
func (o *StringCatalogFileImportOptions) ValidateFileImportOptions() error { return nil }
func (o *VDFFileImportOptions) ValidateFileImportOptions() error           { return nil }

type (
	FileExportOptions interface{ ValidateFileExportOptions() error }

	// SpreadsheetsFileExportOptions implements the FileExportOptions interface.
	GeneralFileExportOptions struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern string `json:"exportPattern,omitempty"`
	}

	// PropertiesFileExportOptions implements the FileExportOptions interface.
	PropertyFileExportOptions struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern string `json:"exportPattern,omitempty"`
		// Values available:
		// 0 - Do not escape single quote.
		// 1 - Escape single quote by another single quote.
		// 2 - Escape single quote by a backslash.
		// 3 - Escape single quote by another single quote only in strings containing variables ({0}).
		EscapeQuotes *int `json:"escapeQuotes,omitempty"`
		// Defines whether any special characters (=, :, ! and #) should be escaped by
		// backslash in exported translations. You can add escape_special_characters per-file option.
		// Acceptable values are: 0, 1. Default is 0.
		// 0 - Do not escape special characters.
		// 1 - Escape special characters by a backslash.
		EscapeSpecialCharacters *int `json:"escapeSpecialCharacters,omitempty"`
	}

	// JavaScriptFileExportOptions implements the FileExportOptions interface.
	JavaScriptFileExportOptions struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern string `json:"exportPattern,omitempty"`
		// Acceptable values are: `single`, `double`. Default is `single`.
		// `single` - Output will be enclosed in single quotes.
		// `double` - Output will be enclosed in double quotes.
		ExportQuotes string `json:"exportQuotes,omitempty"`
	}

	// MDFileExportOptions implements the FileExportOptions interface.
	// It is also used for MDX v1 and MDX v2 files (`frontMatterQuotes`
	// is supported for Markdown files only).
	MDFileExportOptions struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern string `json:"exportPattern,omitempty"`
		// Marker to use for strong. Enum: asterisk, underscore. Default: asterisk.
		StrongMarker string `json:"strongMarker,omitempty"`
		// Marker to use for emphasis. Enum: asterisk, underscore. Default: underscore.
		EmphasisMarker string `json:"emphasisMarker,omitempty"`
		// Marker to use for unordered list bullet. Enum: asterisks, plus, dash. Default: dash.
		UnorderedListBullet string `json:"unorderedListBullet,omitempty"`
		// Table formatting. Enum: consolidate, evenly_distribute_cells.
		// Default: evenly_distribute_cells.
		TableColumnWidth string `json:"tableColumnWidth,omitempty"`
		// Export front matter values in quotes. Enum: auto, single, double. Default: auto.
		// Note: Markdown files only.
		FrontMatterQuotes string `json:"frontMatterQuotes,omitempty"`
	}

	// DOCXFileExportOptions implements the FileExportOptions interface.
	DOCXFileExportOptions struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern string `json:"exportPattern,omitempty"`
		// When checked, run properties in the exported document are minified against
		// the style definitions instead of being kept verbatim. Default: true.
		AllowWordStyleOptimization *bool `json:"allowWordStyleOptimization,omitempty"`
		// When checked, excludes colored text runs inside a spreadsheet cell from
		// translation. Default: false.
		// Note: Applied when the file is parsed, even though it is set through `exportOptions`.
		TranslateExcelExcludeColors *bool `json:"translateExcelExcludeColors,omitempty"`
	}
)

func (o *GeneralFileExportOptions) ValidateFileExportOptions() error    { return nil }
func (o *PropertyFileExportOptions) ValidateFileExportOptions() error   { return nil }
func (o *JavaScriptFileExportOptions) ValidateFileExportOptions() error { return nil }
func (o *MDFileExportOptions) ValidateFileExportOptions() error         { return nil }
func (o *DOCXFileExportOptions) ValidateFileExportOptions() error       { return nil }

// FileUpdateRestoreRequest defines the structure of a request
// to update or restore a file.
type FileUpdateRestoreRequest struct {
	// Revision Identifier.
	RevisionID int `json:"revisionId,omitempty"`
	// Storage Identifier.
	StorageID int `json:"storageId,omitempty"`
	// File name.
	// Note: Can't contain \ / : * ? " < > | symbols.
	Name string `json:"name,omitempty"`
	// Update Option defines whether to keep existing translations and
	// approvals for updated strings. Default: `clear_translations_and_approvals`.
	// Enum: `clear_translations_and_approvals`, `keep_translations`, `keep_translations_and_approvals`.
	UpdateOption string `json:"updateOption,omitempty"`
	// File import options.
	ImportOptions FileImportOptions `json:"importOptions,omitempty"`
	// File export options.
	ExportOptions FileExportOptions `json:"exportOptions,omitempty"`
	// Attach labels to updated strings.
	AttachLabelIDs []int `json:"attachLabelIds,omitempty"`
	// Detach labels from updated strings.
	DetachLabelIDs []int `json:"detachLabelIds,omitempty"`
	// Enable to replace context, that have been modified in Crowdin.
	// Default: false.
	ReplaceModifiedContext *bool `json:"replaceModifiedContext,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *FileUpdateRestoreRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.RevisionID == 0 && r.StorageID == 0 {
		return errors.New("one of revisionId or storageId is required")
	}
	if r.RevisionID != 0 && r.StorageID != 0 {
		return errors.New("use only one of revisionId or storageId")
	}
	return nil
}

type (
	// FileRevision represents a file revision.
	FileRevision struct {
		ID                int  `json:"id"`
		ProjectID         int  `json:"projectId"`
		FileID            int  `json:"fileId"`
		RestoreToRevision *int `json:"restoreToRevision,omitempty"`
		Info              struct {
			Added   RevisionInfo `json:"added"`
			Deleted RevisionInfo `json:"deleted"`
			Updated RevisionInfo `json:"updated"`
		} `json:"info"`
		Date string `json:"date"`
	}

	// RevisionInfo contains the number of strings and words
	// in a file revision.
	RevisionInfo struct {
		Strings int `json:"strings"`
		Words   int `json:"words"`
	}
)

// FileRevisionResponse describes a response with
// a single file revision.
type FileRevisionResponse struct {
	Data *FileRevision `json:"data"`
}

// FileRevisionListResponse describes a response with
// a list of file revisions.
type FileRevisionListResponse struct {
	Data []*FileRevisionResponse `json:"data"`
}

// ReviewedBuild represents a reviewed source file build.
type ReviewedBuild struct {
	ID         int    `json:"id"`
	ProjectID  int    `json:"projectId"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
	Attributes struct {
		BranchID         *int   `json:"branchId,omitempty"`
		TargetLanguageID string `json:"targetLanguageId"`
	} `json:"attributes"`
}

// ReviewedBuildResponse describes a response with a single reviewed build.
type ReviewedBuildResponse struct {
	Data *ReviewedBuild `json:"data"`
}

// ReviewedBuildListResponse describes a response with a list of reviewed builds.
type ReviewedBuildListResponse struct {
	Data []*ReviewedBuildResponse `json:"data"`
}

// ReviewedBuildListOptions specifies the optional parameters to the
// SourceFilesService.ListReviewedBuilds method.
type ReviewedBuildListOptions struct {
	// BranchID is the ID of the branch to filter reviewed builds by.
	BranchID int `json:"branchId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of ReviewedBuildListOptions.
func (o *ReviewedBuildListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}

	return v, len(v) > 0
}

// ReviewedBuildRequest defines the structure of a request to create a new reviewed build.
type ReviewedBuildRequest struct {
	// Branch Identifier.
	BranchID int `json:"branchId,omitempty"`
}

// Validate checks if the reviewed build request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ReviewedBuildRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	return nil
}

// DirectoriesSearchOptions specifies the parameters to the
// SourceFilesService.SearchDirectories method.
type DirectoriesSearchOptions struct {
	// Search directories by `name` or `title` (required).
	Filter string `json:"filter"`
	// Project identifiers to search across (max 50).
	// Omit to search all accessible projects.
	// Note: On crowdin.com all projects must belong to the same owner.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// Owner (user) whose projects to search when `projectIds` is omitted.
	// Defaults to your own account.
	// Note: Available for crowdin.com only.
	UserID int `json:"userId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of DirectoriesSearchOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *DirectoriesSearchOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if len(o.ProjectIDs) > 0 {
		v.Add("projectIds", JoinSlice(o.ProjectIDs))
	}
	if o.UserID > 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}

	return v, len(v) > 0
}

// Validate checks if the search options are valid.
func (o *DirectoriesSearchOptions) Validate() error {
	if o == nil {
		return ErrNilRequest
	}
	return validateSearchOptions(o.Filter, o.ProjectIDs)
}

// FilesSearchOptions specifies the parameters to the
// SourceFilesService.SearchFiles method.
type FilesSearchOptions struct {
	// Search files by `name` or `title` (required).
	Filter string `json:"filter"`
	// Project identifiers to search across (max 50).
	// Omit to search all accessible projects.
	// Note: On crowdin.com all projects must belong to the same owner.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// Owner (user) whose projects to search when `projectIds` is omitted.
	// Defaults to your own account.
	// Note: Available for crowdin.com only.
	UserID int `json:"userId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of FilesSearchOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *FilesSearchOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if len(o.ProjectIDs) > 0 {
		v.Add("projectIds", JoinSlice(o.ProjectIDs))
	}
	if o.UserID > 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}

	return v, len(v) > 0
}

// Validate checks if the search options are valid.
func (o *FilesSearchOptions) Validate() error {
	if o == nil {
		return ErrNilRequest
	}
	return validateSearchOptions(o.Filter, o.ProjectIDs)
}

// AssetReference represents a reference file attached to an asset file.
type AssetReference struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	User      *ShortUser `json:"user"`
	CreatedAt string     `json:"createdAt"`
	// MIME type of the reference file.
	MimeType string `json:"mimeType"`
}

// AssetReferenceResponse describes a response with a single asset reference.
type AssetReferenceResponse struct {
	Data *AssetReference `json:"data"`
}

// AssetReferenceListResponse describes a response with a list of asset references.
type AssetReferenceListResponse struct {
	Data []*AssetReferenceResponse `json:"data"`
}

// AssetReferenceAddRequest defines the structure of a request
// to add an asset reference.
type AssetReferenceAddRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// Reference file name.
	Name string `json:"name"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AssetReferenceAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
