package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/aitch/svg"
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
	headerThemeId     = html.Id("theme-select")
	logoSvg           = svg.Svg(
		aitch.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		aitch.Attribute("viewBox", "0 0 128 128"),
		svg.Rect(
			svg.X(8), svg.Y(8), svg.Width(112), svg.Height(112),
			svg.Rx(24), svg.Ry(24),
			svg.Fill("none"), svg.Stroke("currentColor"), svg.StrokeWidth(10),
		),
		svg.G(
			aitch.Attribute("transform", "translate(0, 7)"),
			svg.Path(
				svg.D("M54 28 C46 28, 42 36, 42 44 C42 52, 38 56, 34 56 C38 56, 42 60, 42 68 C42 76, 46 84, 54 84"),
				svg.Fill("none"), svg.Stroke("currentColor"), svg.StrokeWidth(10),
				svg.StrokeLineCap("round"), svg.StrokeLineJoin("round"),
			),
			svg.Path(
				svg.D("M74 28 C82 28, 86 36, 86 44 C86 52, 90 56, 94 56 C90 56, 86 60, 86 68 C86 76, 82 84, 74 84"),
				svg.Fill("none"), svg.Stroke("currentColor"), svg.StrokeWidth(10),
				svg.StrokeLineCap("round"), svg.StrokeLineJoin("round"),
			),
			svg.Circle(
				svg.Cx(64), svg.Cy(56), svg.R(6), svg.Fill("currentColor"),
			),
		),
	)
	logoSpan = html.Span(
		html.Class("logo"),
		logoSvg,
	)
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
		ctx.WriteNodes(html.H2(b.logo, title))
	} else {
		ctx.WriteNodes(html.H2(b.logo, title, html.Sup(html.Class("small"), " Version: ", version)))
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
				Start(elemSelect, false, headerThemeChange, headerThemeId)
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
