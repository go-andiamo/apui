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
	queryParamsToggle      = html.OnToggle([]byte("(e => toggleQueryParams(e))(event)"))
	queryParamsSummary     = html.Summary("Parameters")
	queryParamsClass       = html.Class("qps")
	queryParamContent      = html.Class("content")
	queryParamsTableAtts   = []aitch.Node{html.Class("qps"), html.Id("qps")}
	queryParamsAtts        = []aitch.Node{html.Class("qp")}
	queryParamRemoveButton = html.Button(
		html.OnClick([]byte("(e => removeQueryParam(e))(event)")),
		html.Title("Remove"),
		"-")
	queryParamRemoveTd  = html.Td(queryParamRemoveButton)
	addQueryParamButton = html.Button(
		html.OnClick([]byte("addQueryParam()")),
		html.Title("Add"),
		"+")
	queryParamSelectId  = html.Id("qps-select")
	queryParamsGetClass = html.Class("method get")
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
	if m, ok := def.Methods[req.Method]; ok {
		params = m.QueryParams[:]
		slices.SortStableFunc(params, func(a chioas.QueryParam, b chioas.QueryParam) int {
			return strings.Compare(a.Name, b.Name)
		})
		for _, p := range params {
			titles[p.Name] = p.Description
		}
	}
	// has defined or current query params?...
	if len(params) > 0 || len(curr) > 0 {
		ctx.Start(elemSpan, false, classInlineDropdown)
		ctx.Start(elemDetails, false, queryParamsToggle, queryParamsClass)
		ctx.WriteNodes(queryParamsScripNode, queryParamsSummary)
		ctx.Start(elemDiv, false, queryParamContent)
		// existing params table...
		ctx.Start(elemTable, false, queryParamsTableAtts...)
		for _, c := range curr {
			ctx.Start(elemTr, false, html.Title(titles[c.name]))
			ctx.Start(elemTh, false, queryParamsAtts...)
			ctx.WriteString(c.name)
			ctx.End() //td
			ctx.Start(elemTd, false, queryParamsAtts...)
			ctx.Start(elemInput, true, html.Value(c.value))
			ctx.End() //td
			ctx.WriteNodes(queryParamRemoveTd)
			ctx.End() //tr
		}
		ctx.End() //table
		if len(params) > 0 {
			// add query params select...
			ctx.Start(elemDiv, false, classLr)
			ctx.Start(elemSelect, false, queryParamSelectId)
			for _, p := range params {
				ctx.Start(elemOption, false, html.Value(p.Name), html.Title(p.Description))
				ctx.WriteString(p.Name)
				ctx.End() //option
			}
			ctx.End() //select
			ctx.WriteNodes(addQueryParamButton)
			ctx.End() //div
		}
		// get button...
		ctx.Start(elemDiv, false, classLr)
		ctx.Start(elemButton, false, queryParamsGetClass,
			html.OnClick([]byte("queryParamsGet('"+req.URL.Path+"')")))
		ctx.WriteString(req.Method)
		ctx.EndAll()
	}
}
