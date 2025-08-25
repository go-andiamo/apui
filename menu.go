package apui

import "github.com/go-andiamo/aitch"

// Menu is an option that can be passed to NewBrowser
// and controls the appearance of the context menu (in the header)
type Menu struct {
	Show              bool
	ShowThemeSelect   bool
	ShowEndpoints     bool
	AuthorizationNode aitch.Node
	Additional        []aitch.Node
	Links             []Link
	AddCss            string
}

type Link struct {
	Href string
	Rel  string // defaults to "stylesheet"
}
