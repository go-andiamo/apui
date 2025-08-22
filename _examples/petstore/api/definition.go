package api

import (
	"github.com/go-andiamo/chioas"
	"net/http"
	"petstore/api/paths"
	"petstore/models/params"
	"petstore/models/requests"
)

var definition = chioas.Definition{
	DocOptions: chioas.DocOptions{
		ServeDocs: true,
		Title:     "Pet Store API",
		UIStyle:   chioas.Rapidoc,
		Path:      "/api/docs",
		RapidocOptions: &chioas.RapidocOptions{
			ShowHeader:         true,
			HeadingText:        "Pet Store API",
			ShowMethodInNavBar: "as-colored-block",
			SchemaStyle:        "table",
			UsePathInNavBar:    true,
			UpdateRoute:        true,
			ShowComponents:     true,
			FavIcons: map[int]string{
				16: logoSvgData,
			},
			LogoSrc:   logoSvgData,
			LogoStyle: "width:20px;height:20px;",
		},
	},
	Info: chioas.Info{
		Title:   "Pet Store API",
		Version: "1.0.0",
	},
	Paths: chioas.Paths{
		paths.Root: {
			Tag: "Root",
			Methods: chioas.Methods{
				http.MethodGet: {
					Handler:     (*api).GetRoot,
					Description: "Root discovery",
				},
			},
			Paths: chioas.Paths{
				paths.Pets: {
					Tag: "Pets",
					Methods: chioas.Methods{
						http.MethodGet: {
							Handler:     (*api).GetPets,
							Description: "List/search pets",
							QueryParams: (params.PetFilter{}).ToQueryParams(),
						},
						http.MethodPost: {
							Handler:     (*api).PostPets,
							Description: "Add pet",
							Request: &chioas.Request{
								Description: "Pet to add",
								Required:    true,
								Schema:      (&chioas.Schema{}).MustFrom(requests.AddPet{}),
							},
						},
					},
					Paths: chioas.Paths{
						paths.UuidPath: {
							Methods: chioas.Methods{
								http.MethodGet: {
									Handler:     (*api).GetPet,
									Description: "Get pet",
								},
								http.MethodDelete: {
									Handler:     (*api).DeletePet,
									Description: "Delete pet",
								},
							},
						},
					},
				},
				paths.Categories: {
					Tag: "Categories",
					Methods: chioas.Methods{
						http.MethodGet: {
							Handler:     (*api).GetCategories,
							Description: "List categories",
						},
					},
					Paths: chioas.Paths{
						paths.UuidPath: {
							Methods: chioas.Methods{
								http.MethodGet: {
									Handler:     (*api).GetCategory,
									Description: "Get category",
								},
							},
						},
					},
				},
			},
		},
	},
}
