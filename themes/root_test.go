package themes

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRootThemeVars(t *testing.T) {
	output, err := RootTheme.buildVars()
	require.NoError(t, err)
	fmt.Println(string(output))
	require.Equal(t, expectedRootVars, string(output))
}

const expectedRootVars = `:root {
	--body-text-color: black;
	--body-bg-color: white;
	--body-border-color: #eee;
	--body-font-family: sans-serif;
	--body-font-size: initial;
	--body-link-text-color: rgb(0,0,238);
	--hdr-text-color: white;
	--hdr-bg-color: #464748;
	--hdr-border-color: #eee;
	--hdr-font-family: sans-serif;
	--hdr-font-size: initial;
	--hdr-link-text-color: rgb(0,0,238);
	--nav-text-color: black;
	--nav-bg-color: white;
	--nav-border-color: #aaa;
	--nav-font-family: sans-serif;
	--nav-font-size: initial;
	--nav-link-text-color: rgb(0,0,238);
	--nav-disabled-text-color: #aaa;
	--ftr-text-color: black;
	--ftr-bg-color: #eee;
	--ftr-border-color: #aaa;
	--ftr-font-family: sans-serif;
	--ftr-font-size: 75%;
	--ftr-link-text-color: rgb(0,0,238);
	--main-text-color: black;
	--main-bg-color: white;
	--main-border-color: #eee;
	--main-font-family: sans-serif;
	--main-font-size: initial;
	--main-link-text-color: rgb(0,0,238);
	--json-text-color: black;
	--json-bg-color: #eee;
	--json-border-color: #ddd;
	--json-font-family: monospace;
	--json-font-size: 90%;
	--json-collapse-fg-color: black;
	--json-collapse-bg-color: #aaa;
	--methods-text-color: white;
	--methods-bg-color: #333;
	--methods-border-color: black;
	--methods-font-family: sans-serif;
	--methods-get-text-color: black;
	--methods-get-bg-color: #47AFE8;
	--methods-get-border-color: transparent;
	--methods-delete-text-color: black;
	--methods-delete-bg-color: #F06560;
	--methods-delete-border-color: transparent;
	--methods-put-text-color: black;
	--methods-put-bg-color: #FF9900;
	--methods-put-border-color: transparent;
	--methods-post-text-color: black;
	--methods-post-bg-color: #690;
	--methods-post-border-color: transparent;
	--methods-patch-text-color: black;
	--methods-patch-bg-color: #827717;
	--methods-patch-border-color: transparent;
	--methods-options-text-color: black;
	--methods-options-bg-color: #ddd;
	--methods-options-border-color: transparent;
	--statuses-text-color: black;
	--statuses-bg-color: white;
	--statuses-border-color: black;
	--statuses-font-family: sans-serif;
	--statuses-1xx-text-color: black;
	--statuses-1xx-bg-color: #d0e8ff;
	--statuses-1xx-border-color: black;
	--statuses-2xx-text-color: black;
	--statuses-2xx-bg-color: #d4f4d3;
	--statuses-2xx-border-color: black;
	--statuses-3xx-text-color: black;
	--statuses-3xx-bg-color: #f9f3a1;
	--statuses-3xx-border-color: black;
	--statuses-4xx-text-color: black;
	--statuses-4xx-bg-color: #ffdb99;
	--statuses-4xx-border-color: black;
	--statuses-5xx-text-color: black;
	--statuses-5xx-bg-color: #ffcccb;
	--statuses-5xx-border-color: black;
}`
