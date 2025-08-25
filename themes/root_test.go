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
	--body-text-color: rgb(0,0,0);
	--body-bg-color: rgb(255,255,255);
	--body-border-color: rgb(238,238,238);
	--body-font-family: sans-serif;
	--body-font-size: initial;
	--body-link-text-color: rgb(0,0,238);
	--hdr-text-color: rgb(255,255,255);
	--hdr-bg-color: rgb(70,71,72);
	--hdr-border-color: rgb(238,238,238);
	--hdr-font-family: sans-serif;
	--hdr-font-size: initial;
	--hdr-link-text-color: rgb(0,0,238);
	--nav-text-color: rgb(0,0,0);
	--nav-bg-color: rgb(255,255,255);
	--nav-border-color: rgb(170,170,170);
	--nav-font-family: sans-serif;
	--nav-font-size: initial;
	--nav-link-text-color: rgb(0,0,238);
	--ftr-text-color: rgb(0,0,0);
	--ftr-bg-color: rgb(238,238,238);
	--ftr-border-color: rgb(170,170,170);
	--ftr-font-family: sans-serif;
	--ftr-font-size: 75%;
	--ftr-link-text-color: rgb(0,0,238);
	--main-text-color: rgb(0,0,0);
	--main-bg-color: rgb(255,255,255);
	--main-border-color: rgb(238,238,238);
	--main-font-family: sans-serif;
	--main-font-size: initial;
	--main-link-text-color: rgb(0,0,238);
	--json-text-color: black;
	--json-bg-color: #eee;
	--json-border-color: #ddd;
	--json-font-family: monospace;
	--json-font-size: 90%;
	--json-collapse-fg-color: rgb(0,0,0);
	--json-collapse-bg-color: rgb(170,170,170);
	--methods-text-color: rgb(255,255,255);
	--methods-bg-color: rgb(51,51,51);
	--methods-border-color: rgb(0,0,0);
	--methods-font-family: sans-serif;
	--methods-get-text-color: rgb(0,0,0);
	--methods-get-bg-color: rgb(71,175,232);
	--methods-get-border-color: transparent;
	--methods-delete-text-color: rgb(0,0,0);
	--methods-delete-bg-color: rgb(240,101,96);
	--methods-delete-border-color: transparent;
	--methods-put-text-color: rgb(0,0,0);
	--methods-put-bg-color: rgb(255,153,0);
	--methods-put-border-color: transparent;
	--methods-post-text-color: rgb(0,0,0);
	--methods-post-bg-color: rgb(102,153,0);
	--methods-post-border-color: transparent;
	--methods-patch-text-color: rgb(0,0,0);
	--methods-patch-bg-color: rgb(130,119,23);
	--methods-patch-border-color: transparent;
	--methods-options-text-color: rgb(0,0,0);
	--methods-options-bg-color: rgb(221,221,221);
	--methods-options-border-color: transparent;
	--statuses-text-color: rgb(0,0,0);
	--statuses-bg-color: rgb(255,255,255);
	--statuses-border-color: rgb(0,0,0);
	--statuses-font-family: sans-serif;
	--statuses-1xx-text-color: rgb(0,0,0);
	--statuses-1xx-bg-color: rgb(208,232,255);
	--statuses-1xx-border-color: rgb(0,0,0);
	--statuses-2xx-text-color: rgb(0,0,0);
	--statuses-2xx-bg-color: rgb(212,244,211);
	--statuses-2xx-border-color: rgb(0,0,0);
	--statuses-3xx-text-color: rgb(0,0,0);
	--statuses-3xx-bg-color: rgb(249,243,161);
	--statuses-3xx-border-color: rgb(0,0,0);
	--statuses-4xx-text-color: rgb(0,0,0);
	--statuses-4xx-bg-color: rgb(255,219,153);
	--statuses-4xx-border-color: rgb(0,0,0);
	--statuses-5xx-text-color: rgb(0,0,0);
	--statuses-5xx-bg-color: rgb(255,204,203);
	--statuses-5xx-border-color: rgb(0,0,0);
}`
