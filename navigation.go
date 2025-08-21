package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"net/http"
	"strings"
)

var (
	getMethodNode = html.Span(html.Class("method", "get"), "GET")
	contentClass  = html.Class("content")
)

func (b *Browser) writeNavigation(ctx aitch.ImperativeContext) error {
	var err error
	if err = getMethodNode.Render(ctx.Context()); err == nil {
		if req, ok := getContextRequest(ctx.Context()); ok {
			def, defs, paths, defPaths := b.findRequestDef(req)
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
				if b.docsPathDetector != nil {
					if docsPath := b.docsPathDetector.ResolveDocsPath(req, defPaths); docsPath != "" {
						nodes = append(nodes, html.A(html.Class("info"), html.Href(docsPath), html.Target("_blank"), aitch.Text([]byte("&#9432;"))))
					}
				}
				if m, ok := def.Methods[req.Method]; ok && m.Description != "" {
					nodes = append(nodes, html.Span(html.Class("description"),
						html.Title(m.Description), m.Description))
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
