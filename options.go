package apui

import (
	"github.com/go-andiamo/aitch"
)

type HtmlTemplate string

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

type HeaderRenderer aitch.Node

type FooterRenderer aitch.Node

type JsonRenderer aitch.Node
