package scripts

import (
	_ "embed"
)

//go:embed head.js
var HeadScript []byte

//go:embed navigation.js
var NavigationScript []byte

//go:embed query_params.js
var QueryParamsScript []byte

//go:embed assoc_methods.js
var AssocMethodsScript []byte
