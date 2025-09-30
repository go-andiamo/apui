package api

import (
	_ "embed"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui"
	"github.com/go-andiamo/apui/themes"
	"log"
	"reflect"
	"strings"
)

//go:embed petstore.yaml
var specYaml []byte

const (
	logoSvg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128">
	<g fill="#16a34a" stroke="none">
		<circle cx="34" cy="40" r="12"/>
		<circle cx="54" cy="26" r="12"/>
		<circle cx="74" cy="26" r="12"/>
		<circle cx="94" cy="40" r="12"/>
		<ellipse cx="64" cy="90" rx="36" ry="28"/>
	</g>
</svg>`
)

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
	menu := &apui.Menu{
		Show:            true,
		ShowThemeSelect: true,
		ShowEndpoints:   true,
		Additional: []aitch.Node{
			html.Div(html.Class("example"),
				html.Em(logoImg, "This is an example of PetStore API browser"),
			),
		},
		AddCss: `div.example em {color:green;}
div.example {position:absolute;bottom:0;}`,
	}
	a.browser, err = apui.NewBrowser(
		a, // provides support for apui.ResourceTypeDetector
		apui.AddHeadNode{favIcon},
		apui.Logo{logoImg},
		apui.DefinitionYaml{specYaml},
		themes.Dark, themes.Light, themes.HighContrast,
		apui.DefaultTheme{"Dark"},
		apui.ShowHeader(true), apui.ShowFooter(true),
		menu,
		apui.MobileFriendly(true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

var _ apui.ResourceTypeDetector = (*api)(nil)

func (a *api) DetectResourceType(response any) apui.ResourceType {
	if _, ok := response.(error); ok {
		return apui.Error
	}
	if reflect.TypeOf(response).Kind() == reflect.Slice {
		return apui.Collection
	}
	return apui.Entity
}

func (a *api) CollectionItems(response any) any {
	return response
}
