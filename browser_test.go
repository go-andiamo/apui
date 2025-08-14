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
)

func TestNewBrowser(t *testing.T) {
	b, err := NewBrowser()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.NotNil(t, b.template)
}

func TestBrowser_Write(t *testing.T) {
	//theme := themes.Theme{Name: "Test", Navigation: themes.ThemeItem{BackgroundColor: "red"}}
	//theme := themes.Light
	b, err := NewBrowser(
		petstoreDefinition,
		themes.Light, themes.Dark,
		DefaultTheme("Dark"),
		ShowHeader(true), ShowFooter(true),
		&testPagingDetector{PagingInfo{
			FirstPage:         1,
			LastPage:          10,
			NextPage:          3,
			PreviousPage:      1,
			PageSize:          10,
			PageSizeParamName: "pageSize",
			ShowDisabled:      true,
			PreNode:           html.Span(nbsp, "Paging: "),
		}, true},
	)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost/myapi/pets?search=xxx", nil)
	item := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins ",
			Age:  42,
		},
	}
	b.Write(w, r, item)
	out, err := io.ReadAll(w.Result().Body)
	f, err := os.Create(t.Name() + ".html")
	defer f.Close()
	f.Write(out)
}

var petstoreDefinition = chioas.Definition{
	Info: chioas.Info{
		Title:   "MyAPI",
		Version: "1.0.0",
	},
	Paths: chioas.Paths{
		"/myapi": {
			Methods: chioas.Methods{
				http.MethodGet: {
					Description: "Get root discovery",
				},
			},
			Paths: chioas.Paths{
				"/pets": {
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

type testPagingDetector struct {
	pi PagingInfo
	ok bool
}

func (t *testPagingDetector) IsPaged(response any, req *http.Request, def *chioas.Path) (PagingInfo, bool) {
	return t.pi, t.ok
}

func TestPetstoreYaml(t *testing.T) {
	f, err := os.Create("petstore.yaml")
	require.NoError(t, err)
	defer f.Close()
	err = petstoreDefinition.WriteYaml(f)
	require.NoError(t, err)
}
