package apui

import "github.com/go-andiamo/aitch/html"

var (
	nbsp        = []byte("&nbsp;")
	quote       = []byte{'"'}
	comma       = []byte{','}
	elemSpan    = []byte("span")
	elemCode    = []byte("code")
	elemA       = []byte("a")
	elemBr      = []byte("br")
	elemDiv     = []byte("div")
	elemDetails = []byte("details")
	elemSummary = []byte("summary")
	elemTable   = []byte("table")
	elemTHead   = []byte("thead")
	elemTBody   = []byte("tbody")
	elemTr      = []byte("tr")
	elemTd      = []byte("td")
	elemTh      = []byte("th")
	elemInput   = []byte("input")
	elemSelect  = []byte("select")
	elemOption  = []byte("option")
	elemButton  = []byte("button")
	elemEm      = []byte("em")
	elemPre     = []byte("pre")
	elemHr      = []byte("hr")
	elemH3      = []byte("h3")
	elemH4      = []byte("h4")
)

var (
	classInlineDropdown = html.Class("inline-dropdown")
	classLl             = html.Class("ll")
	classLr             = html.Class("lr")
	contentEditable     = html.ContentEditable(true)
)
