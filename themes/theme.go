package themes

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"strings"
)

type Theme struct {
	root       bool
	Name       string
	Body       ThemeItem
	Header     ThemeItem
	Navigation ThemeItem
	Footer     ThemeItem
	Main       ThemeItem
	Json       JsonTheme
	Methods    MethodsTheme
	Links      []Link
	AddCss     string
}

type Link struct {
	Href string
	Rel  string // defaults to "stylesheet"
}

type ThemeItem struct {
	TextColor       string
	BackgroundColor string
	BorderColor     string
	FontFamily      string
	FontSize        string
	Opened          Coloring // only used for Navigation
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
	TextColor       string
	BackgroundColor string
	BorderColor     string
	FontFamily      string
	Get             Coloring
	Delete          Coloring
	Put             Coloring
	Post            Coloring
	Patch           Coloring
	Options         Coloring
}

type Coloring struct {
	TextColor       string
	BackgroundColor string
	BorderColor     string
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
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (t ThemeItem) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
	if part == "nav" {
		writeVar(part, "opened-text-color", t.Opened.TextColor, buf)
		writeVar(part, "opened-bg-color", t.Opened.BackgroundColor, buf)
		writeVar(part, "opened-border-color", t.Opened.TextColor, buf)
	}
}

func (t JsonTheme) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
	writeVar(part, "collapse-fg-color", t.CollapsedMarker.TextColor, buf)
	writeVar(part, "collapse-bg-color", t.CollapsedMarker.BackgroundColor, buf)
}

func (t MethodsTheme) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part+"-get", "text-color", t.Get.TextColor, buf)
	writeVar(part+"-get", "bg-color", t.Get.BackgroundColor, buf)
	writeVar(part+"-get", "border-color", t.Get.BorderColor, buf)
	writeVar(part+"-delete", "text-color", t.Delete.TextColor, buf)
	writeVar(part+"-delete", "bg-color", t.Delete.BackgroundColor, buf)
	writeVar(part+"-delete", "border-color", t.Delete.BorderColor, buf)
	writeVar(part+"-put", "text-color", t.Put.TextColor, buf)
	writeVar(part+"-put", "bg-color", t.Put.BackgroundColor, buf)
	writeVar(part+"-put", "border-color", t.Put.BorderColor, buf)
	writeVar(part+"-post", "text-color", t.Post.TextColor, buf)
	writeVar(part+"-post", "bg-color", t.Post.BackgroundColor, buf)
	writeVar(part+"-post", "border-color", t.Post.BorderColor, buf)
	writeVar(part+"-patch", "text-color", t.Patch.TextColor, buf)
	writeVar(part+"-patch", "bg-color", t.Patch.BackgroundColor, buf)
	writeVar(part+"-patch", "border-color", t.Patch.BorderColor, buf)
	writeVar(part+"-options", "text-color", t.Options.TextColor, buf)
	writeVar(part+"-options", "bg-color", t.Options.BackgroundColor, buf)
	writeVar(part+"-options", "border-color", t.Options.BorderColor, buf)
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
