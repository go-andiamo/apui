package apui

import (
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/templates"
	"github.com/go-andiamo/chioas"
	"net/http"
)

type Browser struct {
	template    *aitch.Template
	definition  *chioas.Definition
	headScripts aitch.Node
	bodyScripts aitch.Node
	styles      aitch.Node
}

func NewBrowser(options ...any) (*Browser, error) {
	return (&Browser{}).initialise(options...)
}

func (b *Browser) initialise(options ...any) (*Browser, error) {
	var htmlTemplate string
	htmlSet := false
	nodeMap := aitch.NodeMap{
		"head":        aitch.Imperative(b.writeHead),
		"headScripts": aitch.Imperative(b.writeHeadScripts),
		"header":      aitch.Imperative(b.writeHeader),
		"inner":       aitch.Imperative(b.writeInner),
		"footer":      aitch.Imperative(b.writeFooter),
		"bodyScripts": aitch.Imperative(b.writeBodyScripts),
	}
	headScripts := make([]aitch.Node, 0)
	bodyScripts := make([]aitch.Node, 0)
	styles := make([]aitch.Node, 0)
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
		case Css:
			if option.Content != "" {
				if option.Media != "" {
					styles = append(styles, html.StyleElement(aitch.Attribute("media", option.Media), []byte("\n"+option.Content)))
				} else {
					styles = append(styles, html.StyleElement([]byte("\n"+option.Content)))
				}
			}
		case TemplateNode:
			if option.Node == nil {
				return nil, fmt.Errorf("invalid template node (nil Node)")
			}
			nodeMap[option.Name] = option.Node
		}
	}
	b.headScripts = aitch.Collection(headScripts...)
	b.bodyScripts = aitch.Collection(bodyScripts...)
	b.styles = aitch.Collection(styles...)
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
	return b.styles.Render(ctx.Context())
}

func (b *Browser) writeHeadScripts(ctx aitch.ImperativeContext) error {
	return b.headScripts.Render(ctx.Context())
}

func (b *Browser) writeHeader(ctx aitch.ImperativeContext) error {
	//todo
	return nil
}

func (b *Browser) writeInner(ctx aitch.ImperativeContext) error {
	jctx := &context.Context{
		Cargo:  ctx.Context().Data["response"],
		Writer: ctx.Context().Writer,
		Parent: ctx.Context(),
	}
	return jsonRenderNode.Render(jctx)
}

func (b *Browser) writeFooter(ctx aitch.ImperativeContext) error {
	//todo
	return nil
}

func (b *Browser) writeBodyScripts(ctx aitch.ImperativeContext) error {
	return b.bodyScripts.Render(ctx.Context())
}

func (b *Browser) Write(w http.ResponseWriter, r *http.Request, response any, addCargo ...map[string]any) {
	cargo := map[string]any{
		"title": "Test me!",
	}
	for _, a := range addCargo {
		for k, v := range a {
			cargo[k] = v
		}
	}
	data := map[string]any{
		"request":  r,
		"response": response,
	}
	if err := b.template.Execute(w, data, cargo); err != nil {
		panic(err)
	}
}
