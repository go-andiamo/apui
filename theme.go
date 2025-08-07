package apui

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"strings"
)

var rootThemeVars = Theme{
	root: true,
	Body: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Header: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#333",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Navigation: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Footer: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Main: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Json: JsonThemeItem{
		TextColor:       "black",
		BackgroundColor: "#eee",
		BorderColor:     "#ddd",
		FontFamily:      "monospace",
		FontSize:        "initial",
		CollapsedMarker: CollapseMarker{
			TextColor:       "black",
			BackgroundColor: "#aaa",
		},
	},
}

type Theme struct {
	root       bool
	Name       string
	Body       ThemeItem
	Header     ThemeItem
	Navigation ThemeItem
	Footer     ThemeItem
	Main       ThemeItem
	Json       JsonThemeItem
}

type ThemeItem struct {
	TextColor       string
	BackgroundColor string
	BorderColor     string
	FontFamily      string
	FontSize        string
}

type JsonThemeItem struct {
	TextColor       string
	BackgroundColor string
	BorderColor     string
	FontFamily      string
	FontSize        string
	CollapsedMarker CollapseMarker
}

type CollapseMarker struct {
	TextColor       string
	BackgroundColor string
}

func (t Theme) styleNode() (aitch.Node, error) {
	if v, err := t.buildVars(); err == nil {
		return html.StyleElement([]byte{'\n'}, v, []byte{'\n'}), nil
	} else {
		return nil, err
	}
}

func (t Theme) buildVars() ([]byte, error) {
	buf := new(bytes.Buffer)
	if t.root {
		buf.WriteString(":root")
	} else {
		if t.Name == "" {
			return nil, errors.New("theme must have a name")
		}
		name := strings.Map(func(r rune) rune {
			switch {
			case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-':
				return r
			case r >= 'A' && r <= 'Z':
				return r + 32
			case r == '_' || r == ' ':
				return '-'
			}
			return -1
		}, t.Name)
		if len(name) != len(t.Name) {
			return nil, fmt.Errorf("invalid theme name %q", t.Name)
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
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (t ThemeItem) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
}

func (t JsonThemeItem) buildVars(part string, buf *bytes.Buffer) {
	writeVar(part, "text-color", t.TextColor, buf)
	writeVar(part, "bg-color", t.BackgroundColor, buf)
	writeVar(part, "border-color", t.BorderColor, buf)
	writeVar(part, "font-family", t.FontFamily, buf)
	writeVar(part, "font-size", t.FontSize, buf)
	writeVar(part, "collapse-fg-color", t.CollapsedMarker.TextColor, buf)
	writeVar(part, "collapse-bg-color", t.CollapsedMarker.BackgroundColor, buf)
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
