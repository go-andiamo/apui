package themes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRootThemeVars(t *testing.T) {
	output, err := RootTheme.buildVars()
	require.NoError(t, err)
	require.Equal(t, expectedRootVars, string(output))
}

const expectedRootVars = `:root {
	--body-text-color: black;
	--body-bg-color: white;
	--body-border-color: #eee;
	--body-font-family: sans-serif;
	--body-font-size: initial;
	--hdr-text-color: white;
	--hdr-bg-color: #464748;
	--hdr-border-color: #eee;
	--hdr-font-family: sans-serif;
	--hdr-font-size: initial;
	--nav-text-color: black;
	--nav-bg-color: white;
	--nav-border-color: #eee;
	--nav-font-family: sans-serif;
	--nav-font-size: initial;
	--ftr-text-color: black;
	--ftr-bg-color: #eee;
	--ftr-border-color: #aaa;
	--ftr-font-family: sans-serif;
	--ftr-font-size: 75%;
	--main-text-color: black;
	--main-bg-color: white;
	--main-border-color: #eee;
	--main-font-family: sans-serif;
	--main-font-size: initial;
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
}`
