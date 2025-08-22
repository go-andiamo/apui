module petstore

go 1.24

require (
	github.com/go-andiamo/aitch v1.4.5
	github.com/go-andiamo/apui v0.0.0
	github.com/go-andiamo/chioas v1.16.6
	github.com/go-chi/chi/v5 v5.2.2
	github.com/google/uuid v1.6.0
	github.com/go-andiamo/httperr v1.1.0
)

require golang.org/x/net v0.37.0 // indirect

replace github.com/go-andiamo/apui => ../..

require (
	github.com/go-andiamo/splitter v1.2.5 // indirect
	github.com/go-andiamo/urit v1.2.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
