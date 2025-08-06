package apui

import (
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
	b, err := NewBrowser(
		BodyScript{
			Script: jsonExpandCollapseScript,
		},
		Css{
			Content: jsonCss,
		},
	)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	item := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Bilbo",
		Age:  42,
	}
	b.Write(w, r, item)
	out, err := io.ReadAll(w.Result().Body)
	f, err := os.Create(t.Name() + ".html")
	defer f.Close()
	f.Write(out)
}
