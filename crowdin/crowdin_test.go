package crowdin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setupClient() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ = NewClient("access_token")
	url, _ := url.Parse(server.URL)
	client.baseURL = url

	return client, mux, server.Close
}

func TestNewClient(t *testing.T) {
	var token = "access_token"
	c, _ := NewClient(token)
	if c.token != token {
		t.Errorf("Client token is %v, want %v", c.token, token)
	}
	if c.userAgent != userAgent {
		t.Errorf("Client userAgent is %v, want %v", c.userAgent, userAgent)
	}
	if c.baseURL.String() != baseURL {
		t.Errorf("Client baseURL is %v, want %v", c.baseURL.String(), baseURL)
	}
}

func TestNewEnterpriseClient(t *testing.T) {
	var (
		token        = "access_token"
		organization = "demo"
		apiURL       = "https://demo.api.crowdin.com/api/v2"
	)
	c, _ := NewEnterpriseClient(token, organization)
	if c.token != token {
		t.Errorf("Enterprise client token is %v, want %v", c.token, token)
	}
	if c.userAgent != userAgent {
		t.Errorf("Enterprise client userAgent is %v, want %v", c.userAgent, userAgent)
	}
	if c.baseURL.String() != apiURL {
		t.Errorf("Enterprise client baseURL is %v, want %v", c.baseURL.String(), baseURL)
	}
}
