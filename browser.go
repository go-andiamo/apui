package apui

import (
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/styling"
	"github.com/go-andiamo/apui/internal/templates"
	"github.com/go-andiamo/apui/themes"
	"github.com/go-andiamo/chioas"
	"net/http"
	"strings"
)

type Browser struct {
	template     *aitch.Template
	definition   *chioas.Definition
	jsonRenderer aitch.Node
	showHeader   bool
	showFooter   bool
	headNodes    []aitch.Node
	defaultTheme string
}

func NewBrowser(options ...any) (*Browser, error) {
	return (&Browser{
		jsonRenderer: jsonRenderNode,
		showHeader:   true,
		showFooter:   true,
	}).initialise(options...)
}

func (b *Browser) initialise(options ...any) (*Browser, error) {
	var htmlTemplate string
	htmlSet := false
	headScripts := make([]aitch.Node, 0)
	bodyScripts := []aitch.Node{jsonExpandCollapseScriptNode}
	rootVarsNode, _ := themes.RootTheme.StyleNode()
	styles := []aitch.Node{rootVarsNode, styling.BaseCssNode, jsonCssNode}
	overrideNodeMap := aitch.NodeMap{}
	var headerRenderer aitch.Node
	var footerRenderer aitch.Node
	for _, o := range options {
		switch option := o.(type) {
		case HtmlTemplate:
			htmlSet = true
			htmlTemplate = string(option)
		case *chioas.Definition:
			b.definition = option
		case chioas.Definition:
			b.definition = &option
		case HeadScript:
			if option.Script != "" {
				if option.Type != "" {
					headScripts = append(headScripts, html.Script(html.Type(option.Type), []byte("\n"+option.Script)))
				} else {
					headScripts = append(headScripts, html.Script(html.Type("text/javascript"), []byte("\n"+option.Script)))
				}
			}
		case BodyScript:
			if option.Script != "" {
				if option.Type != "" {
					bodyScripts = append(bodyScripts, html.Script(html.Type(option.Type), []byte("\n"+option.Script)))
				} else {
					bodyScripts = append(bodyScripts, html.Script(html.Type("text/javascript"), []byte("\n"+option.Script)))
				}
			}
		case AddStyling:
			if option.Content != "" {
				if option.Media != "" {
					styles = append(styles, html.StyleElement(aitch.Attribute("media", option.Media), []byte("\n"+option.Content)))
				} else {
					styles = append(styles, html.StyleElement([]byte("\n"+option.Content)))
				}
			}
		case themes.Theme:
			if ts, err := option.StyleNode(); err == nil {
				styles = append(styles, ts)
				for _, link := range option.Links {
					if link.Href != "" {
						if link.Rel == "" {
							b.headNodes = append(b.headNodes, html.Link(html.Href(link.Href), html.Rel("stylesheet")))
						} else {
							b.headNodes = append(b.headNodes, html.Link(html.Href(link.Href), html.Rel(link.Rel)))
						}
					}
				}
			} else {
				return nil, err
			}
		case HeaderRenderer:
			headerRenderer = option
		case FooterRenderer:
			footerRenderer = option
		case JsonRenderer:
			b.jsonRenderer = option
		case TemplateNode:
			if option.Node == nil {
				return nil, fmt.Errorf("invalid template node (nil Node)")
			}
			overrideNodeMap[option.Name] = option.Node
		case ShowHeader:
			b.showHeader = bool(option)
		case ShowFooter:
			b.showFooter = bool(option)
		case DefaultTheme:
			b.defaultTheme, _ = themes.NormalizeName(string(option))
		}
	}
	if headerRenderer == nil {
		headerRenderer = b.buildHeaderNode()
	}
	if footerRenderer == nil {
		footerRenderer = html.Span("Powered by ", html.Span(html.Class("github")), nbsp, html.A(html.Target("_blank"), html.Href("https://github.com/go-andiamo/apui"), "apui"))
	}
	nodeMap := aitch.NodeMap{
		"head":        aitch.Imperative(b.writeHead),
		"styles":      aitch.Collection(styles...),
		"headScripts": aitch.Collection(headScripts...),
		"bodyScripts": aitch.Collection(bodyScripts...),
		"header":      aitch.When(keyShowHeader, html.Header(html.Class("header"), headerRenderer)),
		"navigation":  html.Header(html.Class("navigation"), aitch.Imperative(b.writeNavigation)),
		"main":        html.Main(aitch.Imperative(b.writeMain)),
		"footer":      aitch.When(keyShowFooter, html.Footer(html.Class("footer"), footerRenderer)),
	}
	for k, v := range overrideNodeMap {
		switch k {
		case "header", "navigation":
			nodeMap[k] = html.Header(html.Class(k), v)
		case "footer":
			nodeMap[k] = html.Footer(html.Class(k), v)
		default:
			if _, has := nodeMap[k]; has {
				return nil, fmt.Errorf("invalid override node: %s", k)
			}
			nodeMap[k] = v
		}
	}
	if b.jsonRenderer == nil {
		b.jsonRenderer = jsonRenderNode
	}
	if !htmlSet {
		htmlTemplate = templates.DefaultTemplate
	}
	template, err := aitch.NewTemplate("index", htmlTemplate, nodeMap)
	if err != nil {
		return nil, err
	}
	b.template = template
	return b, nil
}

func (b *Browser) buildHeaderNode() aitch.Node {
	var title string
	var version string
	if b.definition != nil {
		title, version = b.definition.Info.Title, b.definition.Info.Version
	}
	if title == "" {
		title = "API"
	}
	if version == "" {
		return html.H2(title)
	}
	return html.H2(title, html.Sup(html.Class("small"), " Version: ", version))
}

func (b *Browser) writeHead(ctx aitch.ImperativeContext) error {
	for _, node := range b.headNodes {
		if err := node.Render(ctx.Context()); err != nil {
			return err
		}
	}
	// todo more?
	return nil
}

var getMethodNode = html.Span(html.Class("method", "get"), "GET")

func getContextRequest(ctx *context.Context) (*http.Request, bool) {
	if r, ok := ctx.Data[keyRequest]; ok {
		if req, ok := r.(*http.Request); ok {
			return req, true
		}
	}
	return nil, false
}

func (b *Browser) writeNavigation(ctx aitch.ImperativeContext) error {
	_ = getMethodNode.Render(ctx.Context())
	if req, ok := getContextRequest(ctx.Context()); ok {
		def, defs, paths := b.findRequestDef(req)
		nodes := []aitch.Node{aitch.Text("/")}
		for i, p := range paths {
			if i == len(paths)-1 {
				nodes = append(nodes, aitch.Text(p))
			} else {
				var partNode aitch.Node
				if pd := defs[i]; pd != nil {
					if m, ok := pd.Methods[http.MethodGet]; ok {
						partNode = html.A(html.Href("/"+strings.Join(paths[:i+1], "/")), html.Title(m.Description), p)
					}
				}
				if partNode == nil {
					partNode = aitch.Text(p)
				}
				nodes = append(nodes, partNode, aitch.Text("/"))
			}
		}
		if def != nil {
			if m, ok := def.Methods[req.Method]; ok && m.Description != "" {
				nodes = append(nodes, html.Span(html.Class("description"), html.Title(m.Description), m.Description))
			}
		}
		ctx.WriteNodes(nodes...)
		// todo does the endpoint support paging? (and how can we tell?)
		// todo does the endpoint have query params?
		// todo does the endpoint have associated methods?
	}
	return nil
}

func (b *Browser) writeMain(ctx aitch.ImperativeContext) error {
	// todo this is just sample code...
	jctx := &context.Context{
		Cargo:  ctx.Context().Data["response"],
		Writer: ctx.Context().Writer,
		Parent: ctx.Context(),
	}
	return b.jsonRenderer.Render(jctx)
}

func (b *Browser) Write(w http.ResponseWriter, r *http.Request, response any, addCargo ...map[string]any) {
	cargo := map[string]any{
		keyTitle: "Test me!",
	}
	if b.defaultTheme != "" {
		cargo[keyTheme] = "theme-" + b.defaultTheme
	}
	for _, a := range addCargo {
		for k, v := range a {
			cargo[k] = v
		}
	}
	data := map[string]any{
		keyShowHeader: b.showHeader,
		keyShowFooter: b.showFooter,
		keyRequest:    r,
		keyResponse:   response,
	}
	if err := b.template.Execute(w, data, cargo); err != nil {
		panic(err)
	}
}

const (
	keyTitle      = "title"
	keyTheme      = "theme"
	keyShowHeader = "show-header"
	keyShowFooter = "show-footer"
	keyRequest    = "request"
	keyResponse   = "response"
)
