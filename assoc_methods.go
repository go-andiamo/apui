package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/scripts"
	"github.com/go-andiamo/chioas"
	"net/http"
	"slices"
	"strings"
)

var (
	assocMethodsScripNode    = html.Script(scripts.AssocMethodsScript)
	assocMethodsSummary      = html.Summary("Methods")
	assocMethodsClass        = html.Class("ams")
	assocMethodChange        = html.OnChange([]byte("(e => methodSelect(e))(event)"))
	assocMethodClass         = html.Class("ams-method")
	assocMethodSelectedClass = html.Class("ams-method selected")
	assocMethodSendClass     = html.Class("send")
	assocMethodFetchingClass = html.Class("fetching")
	assocMethodResponseClass = html.Class("response")
)

type associatedMethod struct {
	method      string
	def         chioas.Method
	path        string
	displayPath string
}

func (b *Browser) writeAssociatedMethods(ctx aitch.ImperativeContext, req *http.Request, def *chioas.Path) {
	if others := b.associatedMethods(req, def); len(others) > 0 {
		ctx.Start(elemSpan, false, classInlineDropdown).
			Start(elemDetails, false, detailsOnToggle, assocMethodsClass).
			WriteNodes(assocMethodsScripNode, assocMethodsSummary).
			Start(elemDiv, false, contentClass)
		// write method select...
		ctx.Start(elemSelect, false, assocMethodChange)
		for i, m := range others {
			ctx.WriteElement(elemOption, m.method+" "+m.path, html.Value(i), html.Title(m.def.Description))
		}
		ctx.End() //select
		ctx.Start(elemHr, true)
		// write each method panel...
		cls := assocMethodSelectedClass
		for _, m := range others {
			b.writeAssociatedMethod(ctx, m, cls)
			cls = assocMethodClass
		}
		ctx.EndAll()
	}
}

func (b *Browser) writeAssociatedMethod(ctx aitch.ImperativeContext, am associatedMethod, cls aitch.Node) {
	marker := ctx.Marker()
	ctx.Start(elemDiv, false, cls, assocMethodSendClass)
	// write description...
	ctx.Start(elemDiv, false).
		WriteElement(elemEm, am.def.Description).
		End()
	// request (or link - for GET)...
	if am.method == http.MethodGet {
		// link...
		ctx.WriteElement(elemA, am.path, html.Href(am.path))
	} else {
		// write request body input...
		if sample, ok := b.methodRequestSample(am.def); ok {
			ctx.WriteElement(elemPre, sample, contentEditable)
		}
		// write button...
		ctx.Start(elemDiv, false, classLr, assocMethodSendClass).
			WriteElement(elemButton, am.method,
				html.Class("method", strings.ToLower(am.method)),
				html.OnClick([]byte("methodExec('"+am.method+"','"+am.path+"')")),
				html.Title(am.def.Description)).
			End()
		// states & statuses...
		ctx.Start(elemDiv, false, classLl, assocMethodFetchingClass).
			Start(elemDiv, false, html.Class("spinner")).End().
			WriteElement(elemEm, []byte("&nbsp;Executing...")).
			End()
		ctx.Start(elemDiv, false, classLl, assocMethodResponseClass, html.OnClick("methodReset()")).
			WriteElement(elemSpan, "200", html.Class("status")).
			WriteElement(elemSpan, "", html.Class("status-text")).
			End()
	}
	ctx.EndToMark(marker)
}

func (b *Browser) associatedMethods(req *http.Request, def *chioas.Path) (result []associatedMethod) {
	if def != nil {
		currPath := strings.TrimSuffix(req.URL.Path, "/")
		// methods on this path (exc. request method)...
		for k, v := range def.Methods {
			if k != req.Method {
				result = append(result, associatedMethod{
					method: k,
					def:    v,
					path:   currPath,
				})
			}
		}
		// methods on immediate sub-paths...
		for path, pDef := range def.Paths {
			if !strings.Contains(path, "{") {
				for k, v := range pDef.Methods {
					result = append(result, associatedMethod{
						method:      k,
						def:         v,
						path:        currPath + path,
						displayPath: path,
					})
				}
			}
		}
		slices.SortStableFunc(result, func(a, b associatedMethod) int {
			if c := strings.Compare(a.displayPath, b.displayPath); c == 0 {
				return strings.Compare(a.method, b.method)
			} else {
				return c
			}
		})
	}
	return result
}
