package apui

import (
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/scripts"
	"github.com/go-andiamo/apui/internal/styling"
	"github.com/go-andiamo/apui/internal/templates"
	"github.com/go-andiamo/apui/themes"
	"github.com/go-andiamo/chioas"
	"net/http"
)

var (
	detailsOnToggle = html.OnToggle([]byte("(e => toggleDetails(e))(event)"))
)

type Browser struct {
	template             *aitch.Template
	definition           *chioas.Definition
	showHeader           bool
	showFooter           bool
	headNodes            []aitch.Node
	themes               []themes.Theme
	defaultTheme         string
	pagingDetector       PagingDetector
	resourceTypeDetector ResourceTypeDetector
	docsPathDetector     DocsPathDetector
	logo                 aitch.Node
}

func NewBrowser(options ...any) (*Browser, error) {
	return (&Browser{
		showHeader:           true,
		showFooter:           true,
		resourceTypeDetector: defaultResourceTypeDetector,
		logo:                 logoSpan,
	}).initialise(options...)
}

func (b *Browser) initialise(options ...any) (*Browser, error) {
	var htmlTemplate string
	htmlSet := false
	headScripts := []aitch.Node{html.Script(scripts.HeadScript)}
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
		case AddHeadNode:
			if option.Node != nil {
				b.headNodes = append(b.headNodes, option.Node)
			}
		case themes.Theme:
			if ts, err := option.StyleNode(); err == nil {
				b.themes = append(b.themes, option)
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
			headerRenderer = option.Node
		case FooterRenderer:
			footerRenderer = option.Node
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
		case Logo:
			b.logo = option.Node
		default:
			// interfaces...
			if intf, ok := o.(ResourceTypeDetector); ok {
				b.resourceTypeDetector = intf
			}
			if intf, ok := o.(PagingDetector); ok {
				b.pagingDetector = intf
			}
			if intf, ok := o.(DocsPathDetector); ok {
				b.docsPathDetector = intf
			}
		}
	}
	if headerRenderer == nil {
		headerRenderer = aitch.Imperative(b.writeHeader)
	}
	if footerRenderer == nil {
		footerRenderer = defaultFooterNode
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
		case "navigation":
			nodeMap[k] = html.Header(html.Class(k), v)
		case "header":
			nodeMap[k] = aitch.When(keyShowHeader, html.Header(html.Class(k), v))
		case "footer":
			nodeMap[k] = aitch.When(keyShowFooter, html.Footer(html.Class(k), v))
		case "main":
			nodeMap[k] = html.Main(v)
		default:
			if _, has := nodeMap[k]; has {
				return nil, fmt.Errorf("invalid override node: %s", k)
			}
			nodeMap[k] = v
		}
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

func (b *Browser) writeHead(ctx aitch.ImperativeContext) error {
	for _, node := range b.headNodes {
		if err := node.Render(ctx.Context()); err != nil {
			return err
		}
	}
	// todo more?
	return nil
}

func getContextRequest(ctx *context.Context) (*http.Request, bool) {
	if r, ok := ctx.Data[keyRequest]; ok {
		if req, ok := r.(*http.Request); ok {
			return req, true
		}
	}
	return nil, false
}

func getContextResponse(ctx *context.Context) (any, bool) {
	r, ok := ctx.Data[keyResponse]
	return r, ok
}

func (b *Browser) Write(w http.ResponseWriter, r *http.Request, response any, addCargo ...map[string]any) {
	cargo := map[string]any{
		keyTitle: "API",
	}
	if b.definition != nil && b.definition.Info.Title != "" {
		cargo[keyTitle] = b.definition.Info.Title
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
