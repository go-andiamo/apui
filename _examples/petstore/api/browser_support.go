package api

import (
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui"
	"github.com/go-andiamo/apui/themes"
	"log"
	"net/http"
	"petstore/api/paths"
	"reflect"
	"strings"
)

const logoSvg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128">
	<g fill="#16a34a" stroke="none">
		<circle cx="34" cy="40" r="12"/>
		<circle cx="54" cy="26" r="12"/>
		<circle cx="74" cy="26" r="12"/>
		<circle cx="94" cy="40" r="12"/>
		<ellipse cx="64" cy="90" rx="36" ry="28"/>
	</g>
</svg>`

var (
	logoSvgData = "data:image/svg+xml," + strings.NewReplacer("\n", "", "\t", "", `"`, "'", "<", "%3C", ">", "%3E", "#", "%23").Replace(logoSvg)
	favIcon     = html.Link(
		html.Rel("icon"), html.Type("image/svg+xml"),
		html.Href([]byte(logoSvgData)))
	logoImg = html.Img(
		html.Src([]byte(logoSvgData)),
		html.Style("width:1.1em", "height:1.1em", "vertical-align:text-top"))
)

func (a *api) setupBrowser() {
	apui.SetUriProperty("$uri")
	var err error
	a.browser, err = apui.NewBrowser(
		a, // provides support for apui.ResourceTypeDetector, apui.PagingDetector & apui.DocsPathDetector
		apui.AddHeadNode{favIcon},
		apui.Logo{logoImg},
		definition,
		themes.Dark, themes.Light, themes.HighContrast,
		apui.DefaultTheme("Dark"),
		apui.ShowHeader(true), apui.ShowFooter(true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

var _ apui.ResourceTypeDetector = (*api)(nil)

// var _ apui.PagingDetector = (*api)(nil)
var _ apui.DocsPathDetector = (*api)(nil)

func (a *api) DetectResourceType(response any) apui.ResourceType {
	if reflect.TypeOf(response).Kind() == reflect.Slice {
		return apui.Collection
	}
	return apui.Entity
}

func (a *api) CollectionItems(response any) any {
	return response
}

func (a *api) ResolveDocsPath(r *http.Request, defPaths []string) string {
	if len(defPaths) > 0 {
		var b strings.Builder
		b.WriteString(paths.Root)
		b.WriteString("/docs/index.html#")
		b.WriteString(strings.ToLower(r.Method))
		b.WriteString("-")
		for _, path := range defPaths {
			b.WriteString("/")
			if strings.HasPrefix(path, "{") && strings.HasSuffix(path, "}") {
				b.WriteString("-")
				b.WriteString(strings.TrimSuffix(strings.TrimPrefix(path, "{"), "}"))
				b.WriteString("-")
			} else {
				b.WriteString(path)
			}
		}
		return b.String()
	}
	return ""
}
