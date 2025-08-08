package apui

import (
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
	theme := themes.Light
	b, err := NewBrowser(
		petstoreDefinition,
		theme,
		DefaultTheme("Light"),
		ShowHeader(true), ShowFooter(true),
	)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost/myapi/pets/", nil)
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
						},
						http.MethodPost: {
							Description: "Add pet",
							Request: &chioas.Request{
								Required: true,
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
									},
								},
								http.MethodDelete: {
									Description: "Delete specific pet",
								},
							},
						},
					},
				},
			},
		},
	},
}
