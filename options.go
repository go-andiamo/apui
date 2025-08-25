package apui

import (
	"github.com/go-andiamo/aitch"
	"net/http"
)

type HtmlTemplate struct {
	Template string
}

type TemplateNode struct {
	Name string
	Node aitch.Node
}

type HeadScript struct {
	Type   string
	Script string
}

type BodyScript struct {
	Type   string
	Script string
}

type AddStyling struct {
	Media   string
	Content string
}

type HeaderRenderer struct {
	Node aitch.Node
}

type FooterRenderer struct {
	Node aitch.Node
}

type JsonRenderer struct {
	Node aitch.Node
}

type AddHeadNode struct {
	Node aitch.Node
}

type ShowHeader bool

type ShowFooter bool

type DefaultTheme struct {
	Name string
}

type Logo struct {
	Node aitch.Node
}

var JsonIndent = 2

type UriPropertyDetector interface {
	IsUriProperty(ptyName string) bool
}

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

type DocsPathDetector interface {
	ResolveDocsPath(r *http.Request, defPath []string) string
}

type CookieJar interface {
	// HtmlResponseCookies supplies cookies to be written for html responses
	HtmlResponseCookies(w http.ResponseWriter, r *http.Request) []*http.Cookie
}
