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
	elemTable   = []byte("table")
	elemTr      = []byte("tr")
	elemTd      = []byte("td")
	elemTh      = []byte("th")
	elemInput   = []byte("input")
	elemSelect  = []byte("select")
	elemOption  = []byte("option")
	elemButton  = []byte("button")
)

var (
	classInlineDropdown = html.Class("inline-dropdown")
	classLr             = html.Class("lr")
)
