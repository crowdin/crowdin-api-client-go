package crowdin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL = "https://api.crowdin.com/"

	userAgent = "crowdin-api-client-go/0.0.1"
)

// Client is a Crowdin API client.
type Client struct {
	baseURL      *url.URL
	token        string
	organization string
	userAgent    string
	httpClient   *http.Client

	Storages           *StorageService
	Languages          *LanguagesService
	Groups             *GroupsService
	Projects           *ProjectsService
	Branches           *BranchesService
	SourceFiles        *SourceFilesService
	SourceStrings      *SourceStringsService
	StringTranslations *StringTranslationsService
	Translations       *TranslationsService
	TranslationStatus  *TranslationStatusService
}

// NewClient creates a new Crowdin API client with provided options (ex. WithHTTPClient).
// `token` is a personal access token. To create a client, use the following code:
//
//	client, err := crowdin.NewClient("token")
//
// To create an Enterprise client, use the WithOrganization() option as below:
//
//	client, err := crowdin.NewClient("token", crowdin.WithOrganization("organization"))
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}
	u, _ := url.Parse(baseURL)
	c := &Client{
		token:     token,
		baseURL:   u,
		userAgent: userAgent,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	if c.organization != "" {
		c.baseURL.Host = fmt.Sprintf("%s.%s", c.organization, c.baseURL.Host)
	}

	// Initialize services.
	c.Storages = &StorageService{client: c}
	c.Languages = &LanguagesService{client: c}
	c.Groups = &GroupsService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.Branches = &BranchesService{client: c}
	c.SourceFiles = &SourceFilesService{client: c}
	c.Translations = &TranslationsService{client: c}
	c.TranslationStatus = &TranslationStatusService{client: c}
	c.SourceStrings = &SourceStringsService{client: c}
	c.StringTranslations = &StringTranslationsService{client: c}

	return c, nil
}

// ClientOption is a client functional option.
type ClientOption func(*Client) error

// WithOrganization sets the organization name.
func WithOrganization(organization string) ClientOption {
	return func(c *Client) error {
		c.organization = organization
		return nil
	}
}

// WithHTTPClient sets the custom HTTP client. If not set http.DefaultClient will be used.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = hc
		return nil
	}
}

// WithTimeout modifies the default HTTP client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}

// RequestOption represents an option that can be used to modify a http.Request.
type RequestOption func(*http.Request) error

// Header sets a header as an option for the request.
func Header(key, value string) RequestOption {
	return func(r *http.Request) error {
		if value != "" {
			r.Header.Set(key, value)
		}
		return nil
	}
}

// newRequest creates a new HTTP request with the provided method, path and body (if any).
func (c *Client) newRequest(ctx context.Context, method, path string, body any, opts ...RequestOption) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	req = req.WithContext(ctx)

	return req, nil
}

// do sends an API request and returns the API response.
func (c *Client) do(r *http.Request, v any) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := &Response{Response: resp}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("client: error reading response body: %w", err)
	}

	if code := resp.StatusCode; code >= http.StatusBadRequest && code <= 599 {
		err = handleErrorResponse(resp, body)
		return response, err
	}

	if r.Method == http.MethodGet {
		if err = response.populatePagination(body); err != nil {
			return response, fmt.Errorf("client: error parsing pagination: %w", err)
		}
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
}

// RequestValidator is an interface for validating requests.
type RequestValidator interface {
	Validate() error
}

// Post makes a POST request to the specified path.
func (c *Client) Post(ctx context.Context, path string, body, v any, opts ...RequestOption) (*Response, error) {
	if body == nil {
		return nil, errors.New("body cannot be nil")
	}
	if rv, ok := body.(RequestValidator); ok {
		if err := rv.Validate(); err != nil {
			return nil, err
		}
	}
	req, err := c.newRequest(ctx, "POST", path, body, opts...)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// Patch makes a PATCH request to the specified path.
func (c *Client) Patch(ctx context.Context, path string, body, v any) (*Response, error) {
	if body == nil {
		return nil, errors.New("body cannot be nil")
	}
	if rv, ok := body.(RequestValidator); ok {
		if err := rv.Validate(); err != nil {
			return nil, err
		}
	}
	req, err := c.newRequest(ctx, "PATCH", path, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// Put makes a PUT request to the specified path.
func (c *Client) Put(ctx context.Context, path string, body, v any) (*Response, error) {
	if rv, ok := body.(RequestValidator); ok {
		if err := rv.Validate(); err != nil {
			return nil, err
		}
	}
	req, err := c.newRequest(ctx, "PUT", path, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

type ListOptionsProvider interface {
	Values() url.Values
}

// Get makes a GET request to the specified path.
func (c *Client) Get(ctx context.Context, path string, params ListOptionsProvider, v any) (*Response, error) {
	if params != nil {
		path += "?" + params.Values().Encode()
	}
	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// Delete makes a DELETE request to the specified path.
func (c *Client) Delete(ctx context.Context, path string) (*Response, error) {
	req, err := c.newRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, nil)
}

// handleErrorResponse checks the API response for errors and returns
// them if they are found.
func handleErrorResponse(r *http.Response) error {
	if code := r.StatusCode; http.StatusOK <= code && code <= 299 {
		return nil
	}

	var respBody error
	if r.StatusCode == http.StatusBadRequest {
		respBody = &ValidationErrorResponse{Response: r, Status: r.StatusCode}
	} else {
		respBody = &ErrorResponse{Response: r}
	}

	data, err := io.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal(data, respBody)
		if err != nil {
			respBody = &ErrorResponse{
				Response: r,
				Err: Error{
					Code:    r.StatusCode,
					Message: err.Error(),
				},
			}
		}
	}
	return respBody
}

// Response is a Crowdin response that wraps http.Response.
type Response struct {
	*http.Response

	Pagination *Pagination
}

func (r *Response) ParsePagination(body []byte) error {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		var p Pagination
		if err := json.Unmarshal(body, &p); err != nil {
			return err
		}
		r.Pagination = &p
	}
	return nil
}

// Pagination represents the pagination information.
type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse is the error response structure from the API.
type ErrorResponse struct {
	Response *http.Response `json:"-"`

	Err Error `json:"error"`
}

// Error implements the Error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %s", r.Err.Code, r.Err.Message)
}

type ValidationError struct {
	Error struct {
		Key    string `json:"key"`
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"error"`
}

// ValidationErrorResponse
type ValidationErrorResponse struct {
	Response *http.Response `json:"-"`

	Errors []ValidationError `json:"errors"`
	Status int
}

// Error implements the Error interface.
func (r *ValidationErrorResponse) Error() string {
	var sb strings.Builder
	for i, err := range r.Errors {
		if i != 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(fmt.Sprintf("%s: ", err.Error.Key))
		for j, e := range err.Error.Errors {
			if j != 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s (%s)", e.Message, e.Code))
		}
	}
	return sb.String()
}
