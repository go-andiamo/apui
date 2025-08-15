package apui

import (
	"github.com/andybalholm/cascadia"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/chioas"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestBrowser_writePagination(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		b := &Browser{
			pagingDetector: &testPagingDetector{
				ok: true,
			},
		}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagination(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{
			keyResponse: nil,
		})
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		links := queryAll(doc, "a")
		require.Len(t, links, 4)
	})
	t.Run("normal with pre and post nodes", func(t *testing.T) {
		b := &Browser{
			pagingDetector: &testPagingDetector{
				ok: true,
				pi: PagingInfo{
					PreNode:  aitch.Element("h1"),
					PostNode: aitch.Element("h2"),
				},
			},
		}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagination(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{
			keyResponse: nil,
		})
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		require.Len(t, queryAll(doc, "h1"), 1)
		require.Len(t, queryAll(doc, "h2"), 1)
	})
	t.Run("no paging detector", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagination(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{})
		require.NoError(t, err)
		require.Empty(t, s)
	})
	t.Run("no context request", func(t *testing.T) {
		b := &Browser{
			pagingDetector: &testPagingDetector{},
		}
		r := httptest.NewRequest("GET", "/example", nil)
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagination(ctx, r, nil)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{})
		require.NoError(t, err)
		require.Empty(t, s)
	})
}

func TestBrowser_writePagingLinks(t *testing.T) {
	b := &Browser{}
	r := httptest.NewRequest("GET", "/example", nil)
	t.Run("none shown", func(t *testing.T) {
		pi := PagingInfo{
			FirstPage:    -1,
			LastPage:     -1,
			NextPage:     -1,
			PreviousPage: -1,
		}
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagingLinks(ctx, r, pi)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{})
		require.NoError(t, err)
		require.Empty(t, s)
	})
	t.Run("all disabled", func(t *testing.T) {
		pi := PagingInfo{
			FirstPage:    -1,
			LastPage:     -1,
			NextPage:     -1,
			PreviousPage: -1,
			ShowDisabled: true,
		}
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagingLinks(ctx, r, pi)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{})
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		links := queryAll(doc, "a[disabled]")
		require.Len(t, links, 4)
	})
	t.Run("all", func(t *testing.T) {
		pi := PagingInfo{
			FirstPage:    1,
			LastPage:     10,
			NextPage:     4,
			PreviousPage: 2,
		}
		n := aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.writePagingLinks(ctx, r, pi)
			return nil
		})
		s, err := testRender(n, nil, map[string]any{})
		require.NoError(t, err)
		doc, err := html.Parse(strings.NewReader(s))
		require.NoError(t, err)
		links := queryAll(doc, "a")
		require.Len(t, links, 4)
		href, ok := getAttr(links[0], "href")
		require.True(t, ok)
		require.Equal(t, "/example?page=1", href.Val)
		href, ok = getAttr(links[1], "href")
		require.True(t, ok)
		require.Equal(t, "/example?page=2", href.Val)
		href, ok = getAttr(links[2], "href")
		require.True(t, ok)
		require.Equal(t, "/example?page=4", href.Val)
		href, ok = getAttr(links[3], "href")
		require.True(t, ok)
		require.Equal(t, "/example?page=10", href.Val)
	})
}

func TestPagingParams(t *testing.T) {
	t.Run("default param name", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/?foo=bar", nil)
		pi := PagingInfo{}
		s := pagingParams(r, pi, 2)
		require.Contains(t, s, "foo=bar")
		require.Contains(t, s, "page=2")
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 1, strings.Count(s, "&"))
	})
	t.Run("with param name", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/?foo=bar", nil)
		pi := PagingInfo{
			PageParamName: "pg",
		}
		s := pagingParams(r, pi, 2)
		require.Contains(t, s, "foo=bar")
		require.Contains(t, s, "pg=2")
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 1, strings.Count(s, "&"))
	})
	t.Run("with page size", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/?foo=bar", nil)
		pi := PagingInfo{
			PageSize:          10,
			PageSizeParamName: "pageSize",
		}
		s := pagingParams(r, pi, 2)
		require.Contains(t, s, "foo=bar")
		require.Contains(t, s, "page=2")
		require.Contains(t, s, "pageSize=10")
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 2, strings.Count(s, "&"))
	})
	t.Run("clear page size", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/?foo=bar&pageSize=10", nil)
		pi := PagingInfo{
			PageSize:          0, // clear it
			PageSizeParamName: "pageSize",
		}
		s := pagingParams(r, pi, 2)
		require.Contains(t, s, "foo=bar")
		require.Contains(t, s, "page=2")
		require.NotContains(t, s, "pageSize=")
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 1, strings.Count(s, "&"))
	})
}

func TestParamsBuild(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		v := url.Values{}
		s := paramsBuild(v)
		require.Equal(t, "", s)
	})
	t.Run("single", func(t *testing.T) {
		v := url.Values{
			"foo": []string{"1"},
		}
		s := paramsBuild(v)
		require.Equal(t, "?foo=1", s)
	})
	t.Run("multiple", func(t *testing.T) {
		v := url.Values{
			"foo": []string{"1"},
			"bar": []string{"2"},
			"baz": []string{""},
		}
		s := paramsBuild(v)
		require.Contains(t, s, "foo=1")
		require.Contains(t, s, "bar=2")
		require.Contains(t, s, "baz")
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 2, strings.Count(s, "&"))
	})
	t.Run("multi-valued", func(t *testing.T) {
		v := url.Values{
			"foo": []string{"1", "2"},
			"bar": []string{"2", "3"},
			"baz": []string{"3", ""},
		}
		v.Encode()
		s := paramsBuild(v)
		require.Contains(t, s, "foo=1")
		require.Contains(t, s, "foo=2")
		require.Contains(t, s, "bar=2")
		require.Contains(t, s, "bar=3")
		require.Contains(t, s, "baz=3")
		require.Equal(t, 2, strings.Count(s, "baz"))
		require.True(t, strings.HasPrefix(s, "?"))
		require.Equal(t, 5, strings.Count(s, "&"))
	})
}

type testPagingDetector struct {
	pi PagingInfo
	ok bool
}

func (t *testPagingDetector) IsPaged(response any, req *http.Request, def *chioas.Path) (PagingInfo, bool) {
	return t.pi, t.ok
}

func queryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, sel)
}

func getAttr(n *html.Node, name string) (*html.Attribute, bool) {
	for _, a := range n.Attr {
		if a.Key == name {
			return &a, true
		}
	}
	return nil, false
}
