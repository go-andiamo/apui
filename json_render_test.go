package apui

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func TestJsonRenderNode(t *testing.T) {
	testCases := []struct {
		item      any
		expectErr bool
		expected  string
		setup     func()
		reset     func()
	}{
		{
			item:     nil,
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><code>null</code><br></div>`,
		},
		{
			item:     true,
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><code>true</code><br></div>`,
		},
		{
			item:     "some string",
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><code>"some string"</code><br></div>`,
		},
		{
			item:     1.23,
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><code>1.23</code><br></div>`,
		},
		{
			item:     []string{"one", "two", "three"},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">[<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"one",</code><br><span>&nbsp;&nbsp;</span><code>"two",</code><br><span>&nbsp;&nbsp;</span><code>"three"</code><br></div><code>]</code><br></div>`,
		},
		{
			item: struct {
				Foo string `json:"foo"`
				Bar string `json:"bar"`
			}{},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"foo": "",</code><br><span>&nbsp;&nbsp;</span><code>"bar": ""</code><br></div><code>}</code><br></div>`,
		},
		{
			item: struct {
				Foo string `json:"foo"`
				Bar string `json:"bar"`
			}{
				Foo: "some/uri",
			},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"foo":</code>&nbsp;"<a href="some/uri" contenteditable="false">some/uri</a>",<br><span>&nbsp;&nbsp;</span><code>"bar": ""</code><br></div><code>}</code><br></div>`,
			setup: func() {
				DefaultUriPropertyDetector = &testUriPropertyDetector{properties: []string{"foo"}}
			},
			reset: func() {
				DefaultUriPropertyDetector = nil
			},
		},
		{
			item: map[string]any{
				"foo": "bar",
			},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"foo":</code>&nbsp;"<a href="bar" contenteditable="false">bar</a>"<br></div><code>}</code><br></div>`,
			setup: func() {
				DefaultUriPropertyDetector = &testUriPropertyDetector{properties: []string{"foo"}}
			},
			reset: func() {
				DefaultUriPropertyDetector = nil
			},
		},
		{
			item: map[string]any{
				"foo": []string{"bar", "baz"},
			},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"foo": </code><div onclick="(e => collapsed(e))(event)">[<span class="expand">...</span><br><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><code>"bar",</code><br><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><code>"baz"</code><br><span>&nbsp;&nbsp;</span></div><code>]</code><br></div><code>}</code><br></div>`,
		},
		{
			item: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			expected: `<div class="json" contenteditable="true" onbeforeinput="return false"><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;</span><code>"foo": </code><div onclick="(e => collapsed(e))(event)">{<span class="expand">...</span><br><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><code>"bar": "baz"</code><br><span>&nbsp;&nbsp;</span></div><code>}</code><br></div><code>}</code><br></div>`,
		},
		{
			item:      &unmarshallable{},
			expectErr: true,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			if tc.setup != nil {
				tc.setup()
			}
			if tc.reset != nil {
				defer tc.reset()
			}
			s, err := testRender(jsonRenderNode, tc.item)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, s)
			}
		})
	}
}

type testUriPropertyDetector struct {
	properties []string
}

func (t *testUriPropertyDetector) IsUriProperty(ptyName string) bool {
	return slices.Contains(t.properties, ptyName)
}

type unmarshallable struct{}

func (u *unmarshallable) MarshalJSON() ([]byte, error) {
	return nil, errors.New("unmarshallable")
}

func testRender(node aitch.Node, cargo any, data ...map[string]any) (string, error) {
	var w bytes.Buffer
	useData := map[string]any{}
	for _, d := range data {
		for k, v := range d {
			useData[k] = v
		}
	}
	err := node.Render(&context.Context{Writer: &w, Cargo: cargo, Data: useData})
	return w.String(), err
}
