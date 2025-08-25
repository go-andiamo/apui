package themes

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"strings"
)

// Theme describes a styling theme for the apui browser
// and can be passed as an option to apui.NewBrowser
//
// Only values set to non-empty strings will appear as theme vars
// in the output css
type Theme struct {
	root bool
	// Name is the name for the theme
	Name string
	// Body is the overall styling for the html <body>
	Body ThemeItem
	// Header is the styling for the header area
	Header ThemeItem
	// Navigation is the styling for the navigation area
	Navigation ThemeItem
	// Main is the styling for main display area
	Main ThemeItem
	// Footer is the styling for the footer area
	Footer ThemeItem
	// Json is the styling for json rendering
	Json JsonTheme
	// Methods is the styling for displaying http methods (e.g. "GET","PUT" etc.)
	Methods MethodsTheme
	// Statuses is the styling for displaying http response status codes
	Statuses StatusTheme
	// Links is any additional <link> tags (e.g. for loading external stylesheets for fonts)
	Links []Link
	// AddCss is any additional css to be added to the theme <style>
	AddCss string
}

type Link struct {
	Href string
	Rel  string // defaults to "stylesheet"
}

type ThemeItem struct {
	TextColor       CssColor
	BackgroundColor CssColor
	BorderColor     CssColor
	FontFamily      string
	FontSize        string
	LinkTextColor   CssColor
	DisabledColor   CssColor
	Opened          Coloring
	Dropdown        Coloring
}

type JsonTheme struct {
	TextColor       string
	BackgroundColor string
	BorderColor     string
	FontFamily      string
	FontSize        string
	CollapsedMarker Coloring
}

type MethodsTheme struct {
	TextColor       CssColor
	BackgroundColor CssColor
	BorderColor     CssColor
	FontFamily      string
	Get             Coloring
	Delete          Coloring
	Put             Coloring
	Post            Coloring
	Patch           Coloring
	Options         Coloring
}

type StatusTheme struct {
	TextColor       CssColor
	BackgroundColor CssColor
	BorderColor     CssColor
	FontFamily      string
	OneXX           Coloring
	TwoXX           Coloring
	ThreeXX         Coloring
	FourXX          Coloring
	FiveXX          Coloring
}

type Coloring struct {
	TextColor       CssColor
	BackgroundColor CssColor
	BorderColor     CssColor
}

func (t Theme) StyleNode() (aitch.Node, error) {
	if v, err := t.buildVars(); err == nil {
		return html.StyleElement([]byte{'\n'}, v, []byte{'\n'}, []byte(t.AddCss)), nil
	} else {
		return nil, err
	}
}

func NormalizeName(name string) (string, error) {
	result := strings.Map(func(r rune) rune {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-':
			return r
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r == '_' || r == ' ':
			return '-'
		}
		return -1
	}, name)
	if len(name) != len(result) {
		return "", fmt.Errorf("invalid theme name %q", name)
	}
	return result, nil
}

func (t Theme) buildVars() ([]byte, error) {
	buf := new(bytes.Buffer)
	if t.root {
		buf.WriteString(":root")
	} else {
		if t.Name == "" {
			return nil, errors.New("theme must have a name")
		}
		name, err := NormalizeName(t.Name)
		if err != nil {
			return nil, err
		}
		buf.WriteString(".theme-")
		buf.WriteString(name)
	}
	buf.WriteString(" {\n")
	t.Body.buildVars("body", buf)
	t.Header.buildVars("hdr", buf)
	t.Navigation.buildVars("nav", buf)
	t.Footer.buildVars("ftr", buf)
	t.Main.buildVars("main", buf)
	t.Json.buildVars("json", buf)
	t.Methods.buildVars("methods", buf)
	t.Statuses.buildVars("statuses", buf)
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (t ThemeItem) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor.String(), buf)
	writeVar(part, "bg-color", t.BackgroundColor.String(), buf)
	writeVar(part, "border-color", t.BorderColor.String(), buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
	writeVar(part, "link-text-color", t.LinkTextColor.String(), buf)
	writeVar(part, "opened-text-color", t.Opened.TextColor.String(), buf)
	writeVar(part, "opened-bg-color", t.Opened.BackgroundColor.String(), buf)
	writeVar(part, "opened-border-color", t.Opened.BorderColor.String(), buf)
	writeVar(part, "dropdown-text-color", t.Dropdown.TextColor.String(), buf)
	writeVar(part, "dropdown-bg-color", t.Dropdown.BackgroundColor.String(), buf)
	writeVar(part, "dropdown-border-color", t.Dropdown.BorderColor.String(), buf)
}

func (t JsonTheme) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
	writeVar(part, "collapse-fg-color", t.CollapsedMarker.TextColor.String(), buf)
	writeVar(part, "collapse-bg-color", t.CollapsedMarker.BackgroundColor.String(), buf)
}

func (t MethodsTheme) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor.String(), buf)
	writeVar(part, "bg-color", t.BackgroundColor.String(), buf)
	writeVar(part, "border-color", t.BorderColor.String(), buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part+"-get", "text-color", t.Get.TextColor.String(), buf)
	writeVar(part+"-get", "bg-color", t.Get.BackgroundColor.String(), buf)
	writeVar(part+"-get", "border-color", t.Get.BorderColor.String(), buf)
	writeVar(part+"-delete", "text-color", t.Delete.TextColor.String(), buf)
	writeVar(part+"-delete", "bg-color", t.Delete.BackgroundColor.String(), buf)
	writeVar(part+"-delete", "border-color", t.Delete.BorderColor.String(), buf)
	writeVar(part+"-put", "text-color", t.Put.TextColor.String(), buf)
	writeVar(part+"-put", "bg-color", t.Put.BackgroundColor.String(), buf)
	writeVar(part+"-put", "border-color", t.Put.BorderColor.String(), buf)
	writeVar(part+"-post", "text-color", t.Post.TextColor.String(), buf)
	writeVar(part+"-post", "bg-color", t.Post.BackgroundColor.String(), buf)
	writeVar(part+"-post", "border-color", t.Post.BorderColor.String(), buf)
	writeVar(part+"-patch", "text-color", t.Patch.TextColor.String(), buf)
	writeVar(part+"-patch", "bg-color", t.Patch.BackgroundColor.String(), buf)
	writeVar(part+"-patch", "border-color", t.Patch.BorderColor.String(), buf)
	writeVar(part+"-options", "text-color", t.Options.TextColor.String(), buf)
	writeVar(part+"-options", "bg-color", t.Options.BackgroundColor.String(), buf)
	writeVar(part+"-options", "border-color", t.Options.BorderColor.String(), buf)
}

func (t StatusTheme) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor.String(), buf)
	writeVar(part, "bg-color", t.BackgroundColor.String(), buf)
	writeVar(part, "border-color", t.BorderColor.String(), buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part+"-1xx", "text-color", t.OneXX.TextColor.String(), buf)
	writeVar(part+"-1xx", "bg-color", t.OneXX.BackgroundColor.String(), buf)
	writeVar(part+"-1xx", "border-color", t.OneXX.BorderColor.String(), buf)
	writeVar(part+"-2xx", "text-color", t.TwoXX.TextColor.String(), buf)
	writeVar(part+"-2xx", "bg-color", t.TwoXX.BackgroundColor.String(), buf)
	writeVar(part+"-2xx", "border-color", t.TwoXX.BorderColor.String(), buf)
	writeVar(part+"-3xx", "text-color", t.ThreeXX.TextColor.String(), buf)
	writeVar(part+"-3xx", "bg-color", t.ThreeXX.BackgroundColor.String(), buf)
	writeVar(part+"-3xx", "border-color", t.ThreeXX.BorderColor.String(), buf)
	writeVar(part+"-4xx", "text-color", t.FourXX.TextColor.String(), buf)
	writeVar(part+"-4xx", "bg-color", t.FourXX.BackgroundColor.String(), buf)
	writeVar(part+"-4xx", "border-color", t.FourXX.BorderColor.String(), buf)
	writeVar(part+"-5xx", "text-color", t.FiveXX.TextColor.String(), buf)
	writeVar(part+"-5xx", "bg-color", t.FiveXX.BackgroundColor.String(), buf)
	writeVar(part+"-5xx", "border-color", t.FiveXX.BorderColor.String(), buf)
}

func writeVar(part string, name string, value string, buf *bytes.Buffer) {
	if value != "" {
		buf.WriteString("\t--")
		buf.WriteString(part)
		buf.WriteString("-")
		buf.WriteString(name)
		buf.WriteString(": ")
		buf.WriteString(value)
		buf.WriteString(";\n")
	}
}
