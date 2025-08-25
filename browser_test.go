package apui

import (
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/themes"
	"github.com/go-andiamo/chioas"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewBrowser(t *testing.T) {
	b, err := NewBrowser()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.NotNil(t, b.template)
}

func TestBrowser_Write(t *testing.T) {
	DefaultUriPropertyDetector = &testUriPropertyDetector{properties: []string{"$uri"}}
	defer func() {
		DefaultUriPropertyDetector = nil
	}()

	menu := Menu{
		Show:            true,
		ShowThemeSelect: true,
		ShowEndpoints:   true,
		AuthorizationNode: html.Div(
			html.H1("This is an authorize page"),
			html.Div(
				html.Span("Api Key:"),
				html.Input(html.Class("auth"), html.Value("foo")),
			)),
	}
	b, err := NewBrowser(
		petstoreDefinition,
		themes.Dark, themes.Light, themes.HighContrast,
		menu,
		DefaultTheme{"Dark"},
		ShowHeader(true), ShowFooter(true),
		&testDocsPathDetector{},
		&testPagingDetector{PagingInfo{
			FirstPage:         1,
			LastPage:          10,
			NextPage:          3,
			PreviousPage:      1,
			PageSize:          10,
			PageSizeParamName: "pageSize",
			ShowDisabled:      true,
		}, true},
		&testCookieJar{},
	)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost/myapi/pets?search=xxx", nil)
	type obj struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}
	item := []struct {
		Uri       string   `json:"$uri"`
		Name      string   `json:"name"`
		Age       int      `json:"age"`
		Something *string  `json:"something"`
		Object    *obj     `json:"obj,omitempty"`
		Array     []string `json:"arr"`
	}{
		{
			Uri:  "/myapi/pets/123",
			Name: "Bilbo Baggins",
			Age:  42,
			Object: &obj{
				Foo: "foo",
				Bar: 2,
			},
		},
		{
			Uri:   "/myapi/pets/456",
			Name:  "Frodo Baggins",
			Age:   42,
			Array: []string{"foo", "bar"},
		},
	}
	/*
		item := struct {
			Uri       string   `json:"$uri"`
			Name      string   `json:"name"`
			Age       int      `json:"age"`
			Something *string  `json:"something"`
			Object    *obj     `json:"obj,omitempty"`
			Array     []string `json:"arr"`
		}{
			Uri:  "/myapi/pets/123",
			Name: "Bilbo Baggins",
			Age:  42,
			Object: &obj{
				Foo: "foo",
				Bar: 2,
			},
			Array: []string{"foo", "bar"},
		}
	*/
	b.Write(w, r, item)
	out, err := io.ReadAll(w.Result().Body)
	f, err := os.Create(t.Name() + ".html")
	defer f.Close()
	f.Write(out)
}

type testDocsPathDetector struct{}

var _ DocsPathDetector = &testDocsPathDetector{}

func (t *testDocsPathDetector) ResolveDocsPath(r *http.Request, defPath []string) string {
	return "/docs"
}

type testCookieJar struct{}

var _ CookieJar = &testCookieJar{}

func (*testCookieJar) HtmlResponseCookies(w http.ResponseWriter, r *http.Request) []*http.Cookie {
	return []*http.Cookie{
		{
			Name:     "test-cookie",
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
			Expires:  time.Now().Add(15 * time.Minute),
			Value:    "some value",
		},
	}
}

var petstoreDefinition = chioas.Definition{
	Info: chioas.Info{
		Title:   "MyAPI",
		Version: "1.0.0",
	},
	Paths: chioas.Paths{
		"/myapi": {
			Tag: "Root",
			Methods: chioas.Methods{
				http.MethodGet: {
					Description: "Get root discovery",
				},
			},
			Paths: chioas.Paths{
				"/pets": {
					Tag: "Pets",
					Methods: chioas.Methods{
						http.MethodGet: {
							Description: "Get pets",
							QueryParams: chioas.QueryParams{
								{
									Name:        "search",
									Description: "Search for pet(s)",
								},
							},
						},
						http.MethodPost: {
							Description: "Add pet",
							Request: &chioas.Request{
								Required: true,
								Ref:      "addPet",
							},
						},
					},
					Paths: chioas.Paths{
						"/{petId}": {
							Methods: chioas.Methods{
								http.MethodGet: {
									Description: "Get specific pet",
								},
								http.MethodPut: {
									Description: "Update specific pet",
									Request: &chioas.Request{
										Required: true,
										Schema: struct {
											Name string `json:"name"`
											Age  int    `json:"age"`
											Type string `json:"type"`
										}{},
									},
								},
								http.MethodDelete: {
									Description: "Delete specific pet",
								},
							},
							Paths: chioas.Paths{
								"/name": {
									Methods: chioas.Methods{
										http.MethodPut: {
											Description: "Update pet name",
											Request: &chioas.Request{
												Required:  true,
												SchemaRef: "updatePetName",
											},
										},
										http.MethodGet: {
											Description: "Get specific pet name",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	Components: &chioas.Components{
		Schemas: chioas.Schemas{
			{
				Name: "updatePetName",
				Properties: chioas.Properties{
					{
						Name:     "name",
						Type:     "string",
						Required: true,
					},
				},
			},
		},
		Requests: chioas.CommonRequests{
			"addPet": {
				Description: "Add pet body",
				Schema: struct {
					Name string `json:"name"`
					Age  int    `json:"age"`
					Type string `json:"type"`
				}{},
			},
		},
	},
}
