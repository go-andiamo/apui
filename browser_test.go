package apui

import (
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
	def := chioas.Definition{
		Info: chioas.Info{
			Title:   "MyAPI",
			Version: "1.0.0",
		},
	}
	b, err := NewBrowser(
		def,
		Theme{Name: "Test", Navigation: ThemeItem{BackgroundColor: "red"}},
		ShowHeader(true), ShowFooter(true),
	)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	item := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins Bilbo Baggins ",
			Age:  42,
		},
	}
	b.Write(w, r, item, map[string]any{
		"theme": "theme-test",
	})
	out, err := io.ReadAll(w.Result().Body)
	f, err := os.Create(t.Name() + ".html")
	defer f.Close()
	f.Write(out)
}
