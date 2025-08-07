package apui

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTheme_buildVars(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		theme := Theme{
			Name: "Dark",
		}
		output, err := theme.buildVars()
		require.NoError(t, err)
		const expect = `.theme-dark {
}`
		require.Equal(t, expect, string(output))
	})
	t.Run("with vars", func(t *testing.T) {
		theme := Theme{
			Name: "Dark",
			Footer: ThemeItem{
				TextColor:       "#fff",
				BackgroundColor: "#333",
				BorderColor:     "#fff",
			},
		}
		output, err := theme.buildVars()
		require.NoError(t, err)
		const expect = `.theme-dark {
	--ftr-text-color: #fff;
	--ftr-bg-color: #333;
	--ftr-border-color: #fff;
}`
		require.Equal(t, expect, string(output))
	})
	t.Run("root", func(t *testing.T) {
		theme := Theme{
			root: true,
			Footer: ThemeItem{
				TextColor:       "#fff",
				BackgroundColor: "#333",
				BorderColor:     "#fff",
			},
		}
		output, err := theme.buildVars()
		require.NoError(t, err)
		const expect = `:root {
	--ftr-text-color: #fff;
	--ftr-bg-color: #333;
	--ftr-border-color: #fff;
}`
		require.Equal(t, expect, string(output))
	})
	t.Run("bad name - empty", func(t *testing.T) {
		theme := Theme{}
		_, err := theme.buildVars()
		require.Error(t, err)
		require.Equal(t, "theme must have a name", err.Error())
	})
	t.Run("bad name - invalid chars", func(t *testing.T) {
		theme := Theme{
			Name: "Dark .",
		}
		_, err := theme.buildVars()
		require.Error(t, err)
		require.Equal(t, "invalid theme name \"Dark .\"", err.Error())
	})
}

func TestTheme_styleNode(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		theme := Theme{
			Name: "Dark",
			Footer: ThemeItem{
				TextColor:       "#fff",
				BackgroundColor: "#333",
				BorderColor:     "#fff",
			},
		}
		node, err := theme.styleNode()
		require.NoError(t, err)
		output, err := testRender(node, nil, nil)
		require.NoError(t, err)
		const expect = `<style>
.theme-dark {
	--ftr-text-color: #fff;
	--ftr-bg-color: #333;
	--ftr-border-color: #fff;
}
</style>`
		require.Equal(t, expect, output)
	})
	t.Run("bad name", func(t *testing.T) {
		theme := Theme{}
		_, err := theme.styleNode()
		require.Error(t, err)
	})
}

func TestRootThemeVars(t *testing.T) {
	output, err := rootThemeVars.buildVars()
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
	--hdr-bg-color: #333;
	--hdr-border-color: #eee;
	--hdr-font-family: sans-serif;
	--hdr-font-size: initial;
	--nav-text-color: black;
	--nav-bg-color: white;
	--nav-border-color: #eee;
	--nav-font-family: sans-serif;
	--nav-font-size: initial;
	--ftr-text-color: black;
	--ftr-bg-color: white;
	--ftr-border-color: #eee;
	--ftr-font-family: sans-serif;
	--ftr-font-size: initial;
	--main-text-color: black;
	--main-bg-color: white;
	--main-border-color: #eee;
	--main-font-family: sans-serif;
	--main-font-size: initial;
	--json-text-color: black;
	--json-bg-color: #eee;
	--json-border-color: #ddd;
	--json-font-family: monospace;
	--json-font-size: initial;
	--json-collapse-fg-color: black;
	--json-collapse-bg-color: #aaa;
}`
