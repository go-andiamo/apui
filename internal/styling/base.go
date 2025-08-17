package styling

import (
	_ "embed"
	"github.com/go-andiamo/aitch/html"
)

//go:embed base.css
var base []byte

//go:embed methods.css
var methods []byte

//go:embed statuses.css
var statuses []byte

//go:embed header.css
var header []byte

//go:embed navigation.css
var navigation []byte

//go:embed query_params.css
var queryParams []byte

//go:embed pagination.css
var pagination []byte

//go:embed assoc_methods.css
var assocMethods []byte

var BaseCssNode = html.StyleElement(base, methods, statuses, header, navigation, queryParams, pagination, assocMethods)
