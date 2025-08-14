package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/scripts"
	"net/http"
	"strings"
)

var (
	detailsOnToggle      = html.OnToggle([]byte("(e => toggleDetails(e))(event)"))
	navigationScriptNode = html.Script(scripts.NavigationScript)
	getMethodNode        = html.Span(html.Class("method", "get"), "GET")
	contentClass         = html.Class("content")
)

func (b *Browser) writeNavigation(ctx aitch.ImperativeContext) error {
	var err error
	if err = getMethodNode.Render(ctx.Context()); err == nil {
		if req, ok := getContextRequest(ctx.Context()); ok {
			def, defs, paths := b.findRequestDef(req)
			nodes := []aitch.Node{navigationScriptNode, aitch.Text("/")}
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
			b.writePagination(ctx, req, def)
			b.writeQueryParams(ctx, req, def)
			b.writeAssociatedMethods(ctx, req, def)
		}
	}
	return err
}
