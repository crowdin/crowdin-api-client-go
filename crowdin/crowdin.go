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
	baseURL    = "https://api.crowdin.com/"
	apiVersion = "api/v2"

	userAgent = "crowdin-api-client-go/0.0.1"
)

// Client is a Crowdin API client.
type Client struct {
	baseURL    *url.URL
	token      string
	userAgent  string
	httpClient *http.Client
}

// EnterpriseClient is a Crowdin Enterprise API client.
type EnterpriseClient struct {
	Client
}

// NewClient creates a new Crowdin API client with provided options (ex. WithHTTPClient).
// `token` is a personal access token.
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

	return c, nil
}

// NewEnterpriseClient creates a new Crowdin Enterprise API client with provided options.
// `token` is a personal access token and `organization` is the name of the organization.
func NewEnterpriseClient(token, organization string, opts ...ClientOption) (*EnterpriseClient, error) {
	if organization == "" {
		return nil, errors.New("organization name cannot be empty")
	}
	c, err := NewClient(token, opts...)
	if err != nil {
		return nil, err
	}
	c.baseURL.Host = fmt.Sprintf("%s.%s", organization, c.baseURL.Host)

	ec := &EnterpriseClient{Client: *c}

	return ec, nil
}

// ClientOption is a client functional option.
type ClientOption func(*Client) error

// WithHTTPClient sets the custom HTTP client.
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

// NewRequest creates a new HTTP request with the provided method, path and body (if any).
func (c *Client) NewRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%s/%s", apiVersion, path))
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
	req = req.WithContext(ctx)

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(r *http.Request, v any) (*Response, error) {
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

	if err = handleErrorResponse(resp); err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
}

// handleErrorResponse checks the API response for errors and returns
// them if they are found.
func handleErrorResponse(r *http.Response) error {
	if code := r.StatusCode; http.StatusOK <= code && code <= 299 {
		return nil
	}

	var respBody = &ErrorResponse{Response: r}

	data, err := io.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal(data, respBody)
		if err != nil {
			respBody.Err = Error{
				Code:    "",
				Message: string(data),
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
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse is the error response structure from the API.
type ErrorResponse struct {
	Response *http.Response `json:"-"`

	Err Error `json:"error"`
}

// Error implements the Error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s", r.Err.Code, r.Err.Message)
}

// ListOptions specifies the optional parameters to methods that support pagination.
type ListOptions struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// Values returns the ListOptions as url.Values for use in query strings.
func (o *ListOptions) Values() url.Values {
	v := url.Values{}
	if o.Limit > 0 {
		v.Add("limit", fmt.Sprintf("%d", o.Limit))
	}
	if o.Offset > 0 {
		v.Add("offset", fmt.Sprintf("%d", o.Offset))
	}
	return v
}
