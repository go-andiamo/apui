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
	queryParamsScripNode   = html.Script(scripts.QueryParamsScript)
	queryParamsSummary     = html.Summary("Parameters")
	queryParamsClass       = html.Class("qps")
	queryParamsId          = html.Id("qps-detail")
	queryParamsTableAtts   = []aitch.Node{html.Class("qps"), html.Id("qps")}
	queryParamsAtts        = []aitch.Node{html.Class("qp")}
	queryParamRemoveButton = html.Button(
		html.OnClick([]byte("(e => removeQueryParam(e))(event)")),
		html.Title("Remove"),
		html.Class("remove"),
		"-")
	queryParamRemoveTd  = html.Td(queryParamRemoveButton)
	queryParamAddButton = html.Button(
		html.OnClick([]byte("addQueryParam()")),
		html.Title("Add"),
		"+")
	queryParamSelectId  = html.Id("qps-select")
	queryParamsGetClass = html.Class("method", "get")
	queryParamsGetId    = html.Id("qps-get")
)

func (b *Browser) writeQueryParams(ctx aitch.ImperativeContext, req *http.Request, def *chioas.Path) {
	type current struct {
		name  string
		value string
	}
	curr := make([]current, 0)
	for name, q := range req.URL.Query() {
		for _, v := range q {
			curr = append(curr, current{name, v})
		}
	}
	slices.SortStableFunc(curr, func(a current, b current) int {
		return strings.Compare(a.name, b.name)
	})
	var params chioas.QueryParams
	titles := make(map[string]string)
	if def != nil {
		if m, ok := def.Methods[req.Method]; ok {
			params = m.QueryParams[:]
			slices.SortStableFunc(params, func(a chioas.QueryParam, b chioas.QueryParam) int {
				return strings.Compare(a.Name, b.Name)
			})
			for _, p := range params {
				titles[p.Name] = p.Description
			}
		}
	}
	// has defined or current query params?...
	if len(params) > 0 || len(curr) > 0 {
		ctx.Start(elemSpan, false, classInlineDropdown).
			Start(elemDetails, false, detailsOnToggle, queryParamsClass, queryParamsId,
				html.Data("path", []byte(req.URL.Path))).
			WriteNodes(queryParamsScripNode, queryParamsSummary).
			Start(elemDiv, false, contentClass,
				html.OnInput([]byte("buildQuery()")),
				html.OnChange([]byte("buildQuery()")))
		// existing params table...
		ctx.Start(elemTable, false, queryParamsTableAtts...)
		for _, c := range curr {
			ctx.Start(elemTr, false, html.Title(titles[c.name])).
				WriteElement(elemTh, c.name, queryParamsAtts...).
				Start(elemTd, false, queryParamsAtts...).
				Start(elemInput, true, html.Value(c.value), html.Name(c.name)).
				End(). //td
				WriteNodes(queryParamRemoveTd).
				End() //tr
		}
		ctx.End() //table
		if len(params) > 0 {
			// add query params select...
			ctx.Start(elemDiv, false, classLr).
				Start(elemSelect, false, queryParamSelectId)
			for _, p := range params {
				ctx.WriteElement(elemOption, p.Name, html.Value(p.Name), html.Title(p.Description))
			}
			ctx.End(). //select
					WriteNodes(queryParamAddButton).
					End() //div
		}
		ctx.Start(elemDiv, false, classLr).
			Start(elemA, false,
				html.Href([]byte(req.URL.Path+paramsBuild(req.URL.Query()))),
				queryParamsGetClass, queryParamsGetId).WriteString("GET")
		ctx.EndAll()
	}
}
