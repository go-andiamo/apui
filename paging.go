package apui

import (
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/aitch/svg"
	"github.com/go-andiamo/chioas"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type PagingInfo struct {
	FirstPage         int    // not shown if negative
	LastPage          int    // not shown if negative
	NextPage          int    // not shown if negative
	PreviousPage      int    // not shown if negative
	PageSize          int    // only used if > 0 and PageSizeParamName is non-empty string
	PageParamName     string // defaults to "page" if an empty string
	PageSizeParamName string
	ShowDisabled      bool // unavailable page links shown as disabled
	PreNode           aitch.Node
	PostNode          aitch.Node
}

type PagingDetector interface {
	IsPaged(response any, req *http.Request, def *chioas.Path) (PagingInfo, bool)
}

func (b *Browser) writePagination(ctx aitch.ImperativeContext, req *http.Request, def *chioas.Path) {
	if b.pagingDetector != nil {
		if res, ok := getContextResponse(ctx.Context()); ok {
			if pi, ok := b.pagingDetector.IsPaged(res, req, def); ok {
				if pi.PreNode != nil {
					ctx.WriteNodes(pi.PreNode)
				}
				b.writePagingLinks(ctx, req, pi)
				if pi.PostNode != nil {
					ctx.WriteNodes(pi.PostNode)
				}
			}
		}
	}
}

func (b *Browser) writePagingLinks(ctx aitch.ImperativeContext, req *http.Request, pi PagingInfo) {
	path := req.URL.Path
	if pi.FirstPage >= 0 {
		ctx.WriteElement(elemA, firstSvg, html.Class("paging-btn"), html.Href(path, pagingParams(req, pi, pi.FirstPage)))
	} else if pi.ShowDisabled {
		ctx.WriteElement(elemA, firstSvg, html.Class("paging-btn"), html.Disabled())
	}
	if pi.PreviousPage >= 0 {
		ctx.WriteElement(elemA, prevSvg, html.Class("paging-btn"), html.Href(path, pagingParams(req, pi, pi.PreviousPage)))
	} else if pi.ShowDisabled {
		ctx.WriteElement(elemA, prevSvg, html.Class("paging-btn"), html.Disabled())
	}
	if pi.NextPage >= 0 {
		ctx.WriteElement(elemA, nextSvg, html.Class("paging-btn"), html.Href(path, pagingParams(req, pi, pi.NextPage)))
	} else if pi.ShowDisabled {
		ctx.WriteElement(elemA, nextSvg, html.Class("paging-btn"), html.Disabled())
	}
	if pi.LastPage >= 0 {
		ctx.WriteElement(elemA, lastSvg, html.Class("paging-btn"), html.Href(path, pagingParams(req, pi, pi.LastPage)))
	} else if pi.ShowDisabled {
		ctx.WriteElement(elemA, lastSvg, html.Class("paging-btn"), html.Disabled())
	}
}

func pagingParams(r *http.Request, pi PagingInfo, pg int) string {
	params := r.URL.Query()
	if pi.PageParamName != "" {
		params[pi.PageParamName] = []string{strconv.Itoa(pg)}
	} else {
		params["page"] = []string{strconv.Itoa(pg)}
	}
	if pi.PageSize > 0 && pi.PageSizeParamName == "" {
		params[pi.PageSizeParamName] = []string{strconv.Itoa(pi.PageSize)}
	}
	return paramsBuild(params)
}

func paramsBuild(p url.Values) string {
	if len(p) == 0 {
		return ""
	}
	var ps strings.Builder
	first := true
	for k, v := range p {
		if len(v) > 1 {
			for _, iv := range v {
				if !first {
					ps.WriteString("&")
				}
				if iv == "" {
					ps.WriteString(k)
				} else {
					ps.WriteString(fmt.Sprintf("%s=%s", k, urlEscape(iv)))
				}
				first = false
			}
		} else {
			if !first {
				ps.WriteString("&")
			}
			if v[0] == "" {
				ps.WriteString(k)
			} else {
				ps.WriteString(fmt.Sprintf("%s=%s", k, urlEscape(v[0])))
			}
		}
		first = false
	}
	return "?" + ps.String()
}

func urlEscape(s string) string {
	s = url.QueryEscape(s)
	return strings.ReplaceAll(strings.ReplaceAll(s, "%2A", "*"), "%2C", ",")
}

var (
	nextSvg = svg.Svg(
		aitch.Attribute("viewBox", "0 0 24 24"), aitch.Attribute("version", "1.1"),
		svg.Path(svg.D("M0 19v-14l12 7-12 7zm12 0v-14l12 7-12 7z")))
	prevSvg = svg.Svg(
		aitch.Attribute("viewBox", "0 0 24 24"), aitch.Attribute("version", "1.1"),
		svg.Path(svg.D("M12 12l12-7v14l-12-7zm-12 0l12-7v14l-12-7z")))
	firstSvg = svg.Svg(
		aitch.Attribute("viewBox", "0 0 24 24"), aitch.Attribute("version", "1.1"),
		svg.Path(svg.D("M13 12l11-7v14l-11-7zm-11 0l11-7v14l-11-7zm-2 6h2v-12h-2v12z")))
	lastSvg = svg.Svg(
		aitch.Attribute("viewBox", "0 0 24 24"), aitch.Attribute("version", "1.1"),
		svg.Path(svg.D("M0 19v-14l11 7-11 7zm11 0v-14l11 7-11 7zm13-13h-2v12h2v-12z")))
)
