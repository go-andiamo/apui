package apui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_renderColumn(t *testing.T) {
	testCases := []struct {
		value      any
		notPresent bool
		name       string
		noAlign    bool
		options    []any
		expect     string
	}{
		{
			notPresent: true,
			expect:     `<tr><td></td></tr>`,
		},
		{
			expect: `<tr><td><span class="null">null</span></td></tr>`,
		},
		{
			value:  "foo",
			expect: `<tr><td>foo</td></tr>`,
		},
		{
			value:  true,
			expect: `<tr><td>true</td></tr>`,
		},
		{
			value:  json.Number("1"),
			expect: `<tr><td class="right">1</td></tr>`,
		},
		{
			value:  1,
			expect: `<tr><td class="right">1</td></tr>`,
		},
		{
			value:  time.Date(2026, 1, 16, 0, 0, 0, 0, time.UTC),
			expect: `<tr><td>2026-01-16T00:00:00Z</td></tr>`,
		},
		{
			value:  map[string]any{},
			expect: `<tr><td><details class="keep"><summary>Object</summary><div class="json" contenteditable="true" onbeforeinput="return false"><code>{}</code><br></div></details></td></tr>`,
		},
		{
			value:  []any{},
			expect: `<tr><td><details class="keep"><summary>Array</summary><div class="json" contenteditable="true" onbeforeinput="return false"><code>[]</code><br></div></details></td></tr>`,
		},
		{
			value: "/foo",
			name:  "$uri",
			options: []any{
				&testUriPropertyDetector{properties: []string{"$uri"}},
			},
			expect: `<tr><td><a href="/foo">/foo</a></td></tr>`,
		},
		{
			value: time.Date(2026, 1, 16, 0, 0, 0, 0, time.UTC),
			options: []any{
				&testValueFormatter{
					fn: func(value any) any {
						if dt, ok := value.(time.Time); ok {
							return dt.Format(time.RFC1123Z)
						}
						return value
					},
				},
			},
			expect: `<tr><td>Fri, 16 Jan 2026 00:00:00 +0000</td></tr>`,
		},
		{
			value:  testStringer{"foo"},
			expect: `<tr><td>&lt;foo&gt;</td></tr>`,
		},
		{
			value: time.Date(2026, 1, 16, 0, 0, 0, 0, time.UTC),
			options: []any{
				&testValueFormatter{
					fn: func(value any) any {
						if dt, ok := value.(time.Time); ok {
							return html.Span(html.Class("datetime"), dt.Format(time.RFC1123Z))
						}
						return value
					},
				},
			},
			expect: `<tr><td><span class="datetime">Fri, 16 Jan 2026 00:00:00 +0000</span></td></tr>`,
		},
		{
			value: time.Date(2026, 1, 16, 0, 0, 0, 0, time.UTC),
			options: []any{
				&testValueFormatter{
					fn: func(value any) any {
						if dt, ok := value.(time.Time); ok {
							return []aitch.Node{
								html.Span(html.Text("Date/time:", []byte("&nbsp;"))),
								html.Span(html.Class("datetime"), dt.Format(time.RFC1123Z)),
							}
						}
						return value
					},
				},
			},
			expect: `<tr><td><span>Date/time:&nbsp;</span><span class="datetime">Fri, 16 Jan 2026 00:00:00 +0000</span></td></tr>`,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			b, err := NewBrowser(tc.options...)
			require.NoError(t, err)
			s, err := testColumnRender(b, tc.value, !tc.notPresent, tc.name, tc.noAlign)
			require.NoError(t, err)
			assert.Equal(t, tc.expect, s)
		})
	}
}

func testColumnRender(b *Browser, value any, present bool, name string, noAlign bool) (string, error) {
	var w bytes.Buffer
	node := html.Tr(
		aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			b.renderColumn(ctx, present, name, value, noAlign)
			return nil
		}),
	)
	err := node.Render(&context.Context{Writer: &w})
	return w.String(), err
}

type testValueFormatter struct {
	fn func(value any) any
}

func (t *testValueFormatter) FormatValue(value any) (formatted any) {
	return t.fn(value)
}

type testStringer struct {
	s string
}

func (t testStringer) String() string {
	return "<" + t.s + ">"
}
