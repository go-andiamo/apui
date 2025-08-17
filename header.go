package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/scripts"
	"github.com/go-andiamo/apui/themes"
	"github.com/go-andiamo/chioas"
	"net/http"
	"sort"
	"strings"
)

var (
	headerScript      = html.Script(scripts.HeaderScript)
	headerThemeChange = html.OnChange([]byte("(e => themeSelect(e))(event)"))
)

func (b *Browser) writeHeader(ctx aitch.ImperativeContext) error {
	ctx.WriteNodes(headerScript)
	var title string
	var version string
	if b.definition != nil {
		title, version = b.definition.Info.Title, b.definition.Info.Version
	}
	if title == "" {
		title = "API"
	}
	ctx.Start(elemDiv, false, html.Class("title"))
	if version == "" {
		ctx.WriteNodes(html.H2(title))
	} else {
		ctx.WriteNodes(html.H2(title, html.Sup(html.Class("small"), " Version: ", version)))
	}
	ctx.End()
	b.writeHeaderDropdown(ctx)
	return nil
}

func (b *Browser) writeHeaderDropdown(ctx aitch.ImperativeContext) {
	endpoints := showableEndpoints(b.definition)
	// only show menu if there's something to put in it...
	if len(endpoints) > 0 || len(b.themes) > 0 {
		marker := ctx.Marker()
		ctx.Start(elemDiv, false, html.Class("menus")).
			Start(elemDetails, false, html.Class("dropdown"), detailsOnToggle).
			Start(elemSummary, false).End().
			Start(elemDiv, false, html.Class("dropdown-menu"))
		if len(b.themes) > 0 {
			ctx.Start(elemDiv, false).
				WriteElement(elemSpan, "Theme:").
				WriteRaw(nbsp).
				Start(elemSelect, false, headerThemeChange)
			for _, theme := range b.themes {
				name, _ := themes.NormalizeName(theme.Name)
				ctx.Start(elemOption, false, html.Value("theme-"+name)).
					WriteString(theme.Name).End()
			}
			ctx.End().End()
			if len(endpoints) > 0 {
				ctx.Start(elemHr, true)
			}
		}
		if emptyTag, ok := endpoints[""]; ok {
			ctx.Start(elemH4, false).WriteString("Endpoints").End()
			for _, e := range emptyTag {
				ctx.Start(elemDiv, false, html.Class("ll")).
					Start(elemA, false, html.Href(e.path), html.Title(e.method.Description)).
					WriteString(e.path).End().End()
			}
		}

		sortedKeys := make([]string, 0, len(endpoints))
		for k := range endpoints {
			if k != "" {
				sortedKeys = append(sortedKeys, k)
			}
		}
		sort.Strings(sortedKeys)
		for _, k := range sortedKeys {
			ctx.Start(elemH4, false).WriteString(k).End()
			for _, e := range endpoints[k] {
				ctx.Start(elemDiv, false).
					Start(elemA, false, html.Href(e.path), html.Title(e.method.Description)).
					WriteString(e.path).End().End()
			}
		}
		ctx.EndToMark(marker)
	}
}

func showableEndpoints(d *chioas.Definition) map[string][]showableEndpoint {
	endpoints := make(map[string][]showableEndpoint)
	if d != nil {
		collectShowables(endpoints, d.Paths, "", "")
	}
	return endpoints
}

func collectShowables(endpoints map[string][]showableEndpoint, paths chioas.Paths, currPath string, tag string) {
	for path, pathDef := range paths {
		if !strings.Contains(path, "{") {
			if pathDef.Tag != "" {
				tag = pathDef.Tag
			}
			if m, ok := pathDef.Methods[http.MethodGet]; ok {
				addShowable(endpoints, m, strings.TrimSuffix(currPath, "/")+path, tag)
			}
			collectShowables(endpoints, pathDef.Paths, strings.TrimSuffix(currPath, "/")+path, tag)
		}
	}
}

func addShowable(endpoints map[string][]showableEndpoint, method chioas.Method, path string, tag string) {
	if _, ok := endpoints[tag]; ok {
		endpoints[tag] = append(endpoints[tag], showableEndpoint{
			method: method,
			path:   path,
		})
	} else {
		endpoints[tag] = []showableEndpoint{{
			method: method,
			path:   path,
		}}
	}
}

type showableEndpoint struct {
	path   string
	method chioas.Method
}
