package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/apui/internal/styling"
	"net/http"
)

// HtmlTemplate is an option that can be passed to NewBrowser
// and replaces the default html template
//
// Note: the supplied template must include the following markers;
//
//	{{head .}}
//	{{styles .}}
//	{{headScripts .}}
//	{{header .}}
//	{{navigation .}}
//	{{main .}}
//	{{footer .}}
//	{{bodyScripts .}}
type HtmlTemplate struct {
	Template string
}

type TemplateNode struct {
	Name string
	Node aitch.Node
}

// HeadScript is an option that can be passed to NewBrowser
// and contains <script> that is added to <head>
type HeadScript struct {
	Type   string // optional, defaults to "text/javascript"
	Script string
}

// BodyScript is an option that can be passed to NewBrowser
// and contains <script> that is added to <body>
type BodyScript struct {
	Type   string // optional, defaults to "text/javascript"
	Script string
}

// AddStyling is an option that can be passed to NewBrowser
// and contains additional css styling to be added to <head>
type AddStyling struct {
	Media   string
	Content string
}

// HeaderRenderer is an option that can be passed to NewBrowser
// and contains an aitch.Node that is used to override the
// default header
type HeaderRenderer struct {
	Node aitch.Node
}

// FooterRenderer is an option that can be passed to NewBrowser
// and contains an aitch.Node that is used to override the
// default footer
type FooterRenderer struct {
	Node aitch.Node
}

// AddHeadNode is an option that can be passed to NewBrowser
// and contains an aitch.Node that is added to <head>
type AddHeadNode struct {
	Node aitch.Node
}

// ShowHeader is an option that can be passed to NewBrowser
// and determines whether the browser output shows header (defaults to true)
type ShowHeader bool

// ShowFooter is an option that can be passed to NewBrowser
// and determines whether the browser output shows footer (defaults to true)
type ShowFooter bool

// DefaultTheme is an option that can be passed to NewBrowser
// and contains the name of the default theme
type DefaultTheme struct {
	Name string
}

// Logo is an option that can be passed to NewBrowser
// and contains an aitch.Node that is used for the logo in header
type Logo struct {
	Node aitch.Node
}

// JsonIndent is the default indent to use when rendering json
var JsonIndent = 2

type UriPropertyDetector interface {
	IsUriProperty(ptyName string) bool
}

// DefaultUriPropertyDetector is the global setting used when rendering
// json to determine whether a property contains a uri
var DefaultUriPropertyDetector UriPropertyDetector = nil

func SetUriProperty(ptyName string) {
	DefaultUriPropertyDetector = &uriPropertyDetector{ptyName: ptyName}
}

type uriPropertyDetector struct {
	ptyName string
}

func (u *uriPropertyDetector) IsUriProperty(ptyName string) bool {
	return u.ptyName == ptyName
}

// DocsPathDetector is an interface option that can be passed to NewBrowser
// and is used to determine the docs path for a request
type DocsPathDetector interface {
	ResolveDocsPath(r *http.Request, defPath []string) string
}

// CookieJar is an interface option that can be passed to NewBrowser
// and is used to write response cookies from the browser
type CookieJar interface {
	// HtmlResponseCookies supplies cookies to be written for html responses
	HtmlResponseCookies(w http.ResponseWriter, r *http.Request) []*http.Cookie
}

// DefinitionYaml is an option that can be passed to NewBrowser
// and contains the yaml api spec to use (as definition)
type DefinitionYaml struct {
	Data []byte
}

// DefinitionJson is an option that can be passed to NewBrowser
// and contains the json api spec to use (as definition)
type DefinitionJson struct {
	Data []byte
}

// MobileFriendly is an option that can be passed to NewBrowser
// and makes UI display more mobile friendly (i.e. adds MobileViewport & MobileStyling)
type MobileFriendly bool

var MobileViewport = AddHeadNode{
	Node: html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
}

var MobileStyling = AddStyling{
	Media:   "(max-width: 767px)",
	Content: styling.Mobile,
}
