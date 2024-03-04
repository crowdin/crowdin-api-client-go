package model

import (
	"errors"
	"fmt"
)

// Language defines the structure of a language.
type Language struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	EditorCode          string   `json:"editorCode"`
	TwoLettersCode      string   `json:"twoLettersCode"`
	ThreeLettersCode    string   `json:"threeLettersCode"`
	Locale              string   `json:"locale"`
	AndroidCode         string   `json:"androidCode"`
	OsxCode             string   `json:"osxCode"`
	OsxLocale           string   `json:"osxLocale"`
	PluralCategoryNames []string `json:"pluralCategoryNames"`
	PluralRules         string   `json:"pluralRules"`
	PluralExamples      []string `json:"pluralExamples"`
	TextDirection       string   `json:"textDirection"`
	DialectOf           string   `json:"dialectOf"`
}

// LanguagesListResponse defines the structure of a response
// when getting a list of languages.
type LanguagesListResponse struct {
	Data       []*LanguagesGetResponse `json:"data"`
	Pagination *Pagination             `json:"pagination"`
}

// LanguagesGetResponse defines the structure of a response
// when getting a single language.
type LanguagesGetResponse struct {
	Data *Language `json:"data"`
}

// EditLanguageRequest defines the structure of a request
// to edit a custome language.
type EditLanguageRequest struct {
	Op   string `json:"op"`
	Path string `json:"path"`
	// string or array of strings
	Value any `json:"value"`
}

// Validate checks if the edit request is valid.
func (r *EditLanguageRequest) Validate() error {
	const (
		opReplace = "replace"
		opTest    = "test"
	)
	if r.Op == "" {
		return errors.New("op is required")
	}
	if r.Op != opReplace && r.Op != opTest {
		return fmt.Errorf("op must be %q or %q", opReplace, opTest)
	}
	if r.Path == "" {
		return errors.New("path is required")
	}
	if r.Value == nil {
		return errors.New("value is required")
	}
	if _, ok := r.Value.(string); !ok {
		if _, ok := r.Value.([]string); !ok {
			return errors.New("value must be a string or an array of strings")
		}
	}
	return nil
}

// LanguagesEditResponse defines the structure of a response
// when updating a custom language.
type LanguagesEditResponse struct {
	Data *Language `json:"data"`
}

// AddLanguageRequest defines the structure of a request
// to add a new custom language.
type AddLanguageRequest struct {
	Name                string   `json:"name"`
	Code                string   `json:"code"`
	LocaleCode          string   `json:"localeCode"`
	TextDirection       string   `json:"textDirection"`
	PluralCategoryNames []string `json:"pluralCategoryNames"`
	ThreeLettersCode    string   `json:"threeLettersCode"`
	TwoLettersCode      string   `json:"twoLettersCode"`
	DialectOf           string   `json:"dialectOf"`
}

// Validate checks if the add request is valid.
func (r *AddLanguageRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Code == "" {
		return errors.New("code is required")
	}
	if r.LocaleCode == "" {
		return errors.New("localeCode is required")
	}
	if r.ThreeLettersCode == "" {
		return errors.New("threeLettersCode is required")
	}
	if len(r.PluralCategoryNames) == 0 {
		return errors.New("pluralCategoryNames is required")
	}

	const (
		ltr = "ltr"
		rtl = "rtl"
	)
	if r.TextDirection == "" {
		return errors.New("textDirection is required")
	}
	if r.TextDirection != ltr && r.TextDirection != rtl {
		return fmt.Errorf("textDirection must be %q or %q", ltr, rtl)
	}

	return nil
}

// LanguagesAddResponse defines the structure of a response
// when adding a new custom language.
type LanguagesAddResponse struct {
	Data *Language `json:"data"`
}
