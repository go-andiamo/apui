package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/chioas"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBrowser_writeAssociatedMethods(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writeAssociatedMethods(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, nil)
		require.NoError(t, err)
		require.Empty(t, s)
	})
	t.Run("with others", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/pets", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writeAssociatedMethods(ctx, r, &chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost:   {},
					http.MethodDelete: {},
				},
			})
			return nil
		})
		s, err := testRender(n, nil, nil)
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		require.Len(t, queryAll(doc, "select option"), 2)
		require.Len(t, queryAll(doc, "div.content div.ams-method"), 2)
		require.Len(t, queryAll(doc, "div.content div.ams-method.selected"), 1)
	})
}
