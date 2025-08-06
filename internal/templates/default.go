package templates

import (
	_ "embed"
)

//go:embed default.html
var DefaultTemplate string
