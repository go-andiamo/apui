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

func TestBrowser_writeQueryParams(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writeQueryParams(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, nil)
		require.NoError(t, err)
		require.Empty(t, s)
	})
	t.Run("existing params", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example?foo=bar&bar=baz", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writeQueryParams(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, nil)
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		require.Len(t, queryAll(doc, "tr"), 2)
		require.Len(t, queryAll(doc, "select"), 0)
	})
	t.Run("defined params", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writeQueryParams(ctx, r, &chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: chioas.Method{
						QueryParams: chioas.QueryParams{
							{
								Name:        "foo",
								Description: "foo description",
							},
							{
								Name:        "bar",
								Description: "bar description",
							},
						},
					},
				},
			})
			return nil
		})
		s, err := testRender(n, nil, nil)
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		require.Len(t, queryAll(doc, "tr"), 0)
		require.Len(t, queryAll(doc, "select option"), 2)
	})
}
