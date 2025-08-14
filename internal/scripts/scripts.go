package scripts

import (
	_ "embed"
)

//go:embed navigation.js
var NavigationScript []byte

//go:embed query_params.js
var QueryParamsScript []byte

//go:embed assoc_methods.js
var AssocMethodsScript []byte
