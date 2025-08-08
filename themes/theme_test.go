package themes

import (
	"bytes"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
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
		node, err := theme.StyleNode()
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
		_, err := theme.StyleNode()
		require.Error(t, err)
	})
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
