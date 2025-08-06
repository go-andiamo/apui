package apui

import (
	"encoding/json"
	"github.com/go-andiamo/aitch"
	. "github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/splitter"
	"strings"
)

const (
	collapsedFuncName        = "collapsed"
	jsonExpandCollapseScript = `function ` + collapsedFuncName + `(evt) {
  if (evt.currentTarget === evt.target) {
    evt.currentTarget.classList.toggle("collapsed");
  } else if (evt.target.classList.contains("expand")) {
    evt.target.parentElement.classList.toggle("collapsed");
  }
  evt.stopPropagation();
}`
	jsonCss = `div.json {
  	border: 1px solid #ddd;
	background-color: #eee;
	padding: 2px;
	max-height: 400px;
	overflow: hidden;
	overflow-y: auto;
}
div.json div * {
	cursor: text;
}
div.json a {
	font-family: monospace;
	cursor: pointer;
}
div.json div {
	display: inline;
	font-family: monospace;
	cursor: pointer;
}
div.json div.collapsed * {
	display: none;
}
div.json span.expand {
	display: none;
}
div.json div.collapsed > span.expand {
	display: inline;
	background-color: #aaa;
	cursor: pointer;
}`
)

var (
	collapseMarker = aitch.Collection(Span(Class("expand"), "..."), Br())
	//collapseAtt    = OnClick(collapsedFuncName, "(event)")
	collapseAtt = OnClick("(e => ", collapsedFuncName, "(e))(event)")
	objStart    = []byte{'{'}
	arrStart    = []byte{'['}
)

var jsonRenderNode = Div(
	Class("json"),
	ContentEditable("true"),
	OnBeforeInput("return false"),
	aitch.Imperative(func(ctx aitch.ImperativeContext) error {
		data, err := json.MarshalIndent(ctx.Context().Cargo, "", "    ")
		if err != nil {
			return err
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			trimmed := strings.TrimLeft(line, " ")
			if spaces := len(line) - len(trimmed); spaces > 0 {
				ctx.Start(elemSpan, false)
				for i := spaces; i > 0; i-- {
					ctx.WriteRaw(nbsp)
				}
				ctx.End()
			}
			if ptyName, uri, ok := isUriProperty(trimmed); ok {
				ctx.Start(elemCode, false)
				ctx.WriteRaw([]byte(`"` + ptyName + `":`))
				ctx.End()
				ctx.WriteRaw(nbsp)
				ctx.WriteRaw(quote)
				ctx.Start(elemA, false, Href(uri), ContentEditable("false"))
				ctx.WriteRaw([]byte(uri))
				ctx.End()
				ctx.WriteRaw(quote)
				if strings.HasSuffix(trimmed, `,`) {
					ctx.WriteRaw(comma)
				}
				ctx.Start(elemBr, true)
			} else if strings.HasSuffix(trimmed, "{") {
				if pre := strings.TrimSuffix(trimmed, "{"); pre != "" {
					ctx.Start(elemCode, false)
					ctx.WriteRaw([]byte(pre))
					ctx.End()
				}
				ctx.Start(elemDiv, false, collapseAtt)
				ctx.WriteRaw(objStart)
				_ = collapseMarker.Render(ctx.Context())
			} else if strings.HasSuffix(trimmed, "[") {
				if pre := strings.TrimSuffix(trimmed, "["); pre != "" {
					ctx.Start(elemCode, false)
					ctx.WriteRaw([]byte(pre))
					ctx.End()
				}
				ctx.Start(elemDiv, false, collapseAtt)
				ctx.WriteRaw(arrStart)
				_ = collapseMarker.Render(ctx.Context())
			} else if ((strings.HasSuffix(trimmed, "}") || strings.HasSuffix(trimmed, "},")) && !strings.Contains(trimmed, "{}")) || ((strings.HasSuffix(trimmed, "]") || strings.HasSuffix(trimmed, "],")) && !strings.Contains(trimmed, "[]")) {
				ctx.End()
				ctx.Start(elemCode, false)
				ctx.WriteRaw([]byte(trimmed))
				ctx.End()
				ctx.Start(elemBr, true)
			} else {
				ctx.Start(elemCode, false)
				ctx.WriteRaw([]byte(trimmed))
				ctx.End()
				ctx.Start(elemBr, true)
			}
		}
		ctx.EndAll()
		return nil
	}),
)

var propertyLineSplitter = splitter.MustCreateSplitter(':', splitter.DoubleQuotesBackSlashEscaped)

func isUriProperty(line string) (ptyName string, uri string, ok bool) {
	if DefaultUriPropertyDetector != nil && strings.HasPrefix(line, `"`) {
		if parts, err := propertyLineSplitter.Split(line); err == nil && len(parts) == 2 {
			ptyName = strings.Trim(parts[0], `"`)
			uri = strings.TrimSuffix(strings.TrimPrefix(parts[1], " "), ",")
			if strings.HasPrefix(uri, `"`) && strings.HasSuffix(uri, `"`) && len(uri) > 2 {
				uri = uri[1 : len(uri)-1]
				ok = DefaultUriPropertyDetector.IsUriProperty(ptyName)
			}
		}
	}
	return
}

type UriPropertyDetector interface {
	IsUriProperty(ptyName string) bool
}

var DefaultUriPropertyDetector UriPropertyDetector = nil
