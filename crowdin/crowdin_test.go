package crowdin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testURL(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.RequestURI; got != want {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to read request body: %v", err)
	}
	if got := buf.String(); got != want {
		t.Errorf("Request body: %v, want %v", got, want)
	}
}

// testJSONBody checks if the request body is equal to the expected JSON (beautified)
// by comparing the unmarshalled maps.
func testJSONBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to read request body: %v", err)
	}

	var wantMap, gotMap map[string]any
	if err := json.Unmarshal([]byte(want), &wantMap); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}

	got := buf.String()
	if err := json.Unmarshal([]byte(got), &gotMap); err != nil {
		t.Fatalf("Failed to unmarshal result JSON: %v", err)
	}

	if !reflect.DeepEqual(wantMap, gotMap) {
		t.Errorf("JSON does not match:\nExpected: %v\nGot: %v", wantMap, gotMap)
	}
}

func testHeader(t *testing.T, r *http.Request, header, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Request header %s: %v, want %v", header, got, want)
	}
}

func testClientServices(t *testing.T, c *Client) {
	services := []string{
		"Storages",
		"Languages",
		"Groups",
		"Projects",
		"Branches",
		"SourceFiles",
		"SourceStrings",
		"StringTranslations",
		"StringComments",
		"Translations",
		"TranslationStatus",
		"MachineTranslationEngines",
		"Screenshots",
	}

	ptr := reflect.ValueOf(c)
	val := reflect.Indirect(ptr)

	for _, s := range services {
		if val.FieldByName(s).IsNil() {
			t.Errorf("c.%s should not be nil", s)
		}
	}
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
	testClientServices(t, c)
}

func TestNewClient_emptyToken(t *testing.T) {
	if _, err := NewClient(""); err == nil {
		t.Error("NewClient with empty token returned nil, want error")
	}
}

func TestNewEnterpriseClient(t *testing.T) {
	var (
		token        = "access_token"
		organization = "demo"
		apiURL       = "https://demo.api.crowdin.com/"
	)
	c, _ := NewClient(token, WithOrganization(organization))
	if c.token != token {
		t.Errorf("Enterprise client token is %v, want %v", c.token, token)
	}
	if c.userAgent != userAgent {
		t.Errorf("Enterprise client userAgent is %v, want %v", c.userAgent, userAgent)
	}
	if c.baseURL.String() != apiURL {
		t.Errorf("Enterprise client baseURL is %v, want %v", c.baseURL.String(), apiURL)
	}
	testClientServices(t, c)
}

func TestWithCustomHTTPClient(t *testing.T) {
	c, err := NewClient("token", WithHTTPClient(http.DefaultClient))
	if err != nil {
		t.Errorf("NewClient error: %v", err)
	}
	if c.httpClient != http.DefaultClient {
		t.Errorf("NewClient httpClient is %v, want %v", c.httpClient, http.DefaultClient)
	}
}

func TestGet(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	type test struct {
		Hello string `json:"hello"`
	}

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Request method is %v, want %v", r.Method, http.MethodGet)
		}
		fmt.Fprint(w, `{"hello":"world"}`)
	})

	res := new(test)
	_, _ = client.Get(context.Background(), "/get", nil, res)

	want := &test{"world"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Response is %v, want %v", res, want)
	}
}

func TestPost(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	type test struct {
		Hello string `json:"hello"`
	}

	type body struct {
		Foo string `json:"foo"`
	}

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Request method is %v, want %v", r.Method, http.MethodGet)
		}
		fmt.Fprint(w, `{"hello":"world"}`)
	})

	res := new(test)
	_, _ = client.Post(context.Background(), "/post", &body{"bar"}, res)

	want := &test{"world"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Response is %v, want %v", res, want)
	}
}

func TestPut(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	type test struct {
		Hello string `json:"hello"`
	}

	type body struct {
		Foo string `json:"foo"`
	}

	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Request method is %v, want %v", r.Method, http.MethodGet)
		}
		fmt.Fprint(w, `{"hello":"world"}`)
	})

	res := new(test)
	_, _ = client.Put(context.Background(), "/put", &body{"bar"}, res)

	want := &test{"world"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Response is %v, want %v", res, want)
	}
}

func TestPatch(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	type test struct {
		Hello string `json:"hello"`
	}

	type body struct {
		Foo string `json:"foo"`
	}

	mux.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Request method is %v, want %v", r.Method, http.MethodGet)
		}
		fmt.Fprint(w, `{"hello":"world"}`)
	})

	res := new(test)
	_, _ = client.Patch(context.Background(), "/patch", &body{"bar"}, res)

	want := &test{"world"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Response is %v, want %v", res, want)
	}
}

func TestDelete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/delete", func(_ http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Request method is %v, want %v", r.Method, http.MethodGet)
		}
	})

	_, _ = client.Delete(context.Background(), "/delete")
}
