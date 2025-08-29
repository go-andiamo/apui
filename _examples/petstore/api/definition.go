package api

import (
	"github.com/go-andiamo/chioas"
	"net/http"
	"petstore/api/paths"
	"petstore/models"
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
					Tag:              "Pets",
					ApplyMiddlewares: applyAuthMiddlewares,
					Methods: chioas.Methods{
						http.MethodGet: {
							Handler:     (*api).GetPets,
							Description: "List/search pets",
							QueryParams: (params.PetFilter{}).ToQueryParams(),
							Responses: chioas.Responses{
								http.StatusOK: {
									IsArray: true,
									Schema:  (&chioas.Schema{}).MustFrom(models.Pet{}),
								},
							},
						},
						http.MethodPost: {
							Handler:     (*api).PostPets,
							Description: "Add pet",
							Request: &chioas.Request{
								Description: "Pet to add",
								Required:    true,
								Schema:      (&chioas.Schema{}).MustFrom(requests.AddPet{}),
							},
							Responses: chioas.Responses{
								http.StatusCreated: {
									Schema: (&chioas.Schema{}).MustFrom(models.Pet{}),
								},
							},
						},
					},
					Paths: chioas.Paths{
						paths.UuidPath: {
							PathParams: chioas.PathParams{
								"id": {
									Description: "Pet ID",
									Schema: &chioas.Schema{
										Type:   "string",
										Format: "uuid",
									},
								},
							},
							Methods: chioas.Methods{
								http.MethodGet: {
									Handler:     (*api).GetPet,
									Description: "Get pet",
									Responses: chioas.Responses{
										http.StatusOK: {
											Schema: (&chioas.Schema{}).MustFrom(models.Pet{}),
										},
									},
								},
								http.MethodPut: {
									Handler:     (*api).PutPet,
									Description: "Update pet",
									Request: &chioas.Request{
										Description: "Pet update",
										Required:    true,
										Schema:      (&chioas.Schema{}).MustFrom(requests.UpdatePet{}),
									},
									Responses: chioas.Responses{
										http.StatusOK: {
											Schema: (&chioas.Schema{}).MustFrom(models.Pet{}),
										},
									},
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
					Tag:              "Categories",
					ApplyMiddlewares: applyAuthMiddlewares,
					Methods: chioas.Methods{
						http.MethodGet: {
							Handler:     (*api).GetCategories,
							Description: "List categories",
							Responses: chioas.Responses{
								http.StatusOK: {
									IsArray: true,
									Schema:  (&chioas.Schema{}).MustFrom(models.Category{}),
								},
							},
						},
					},
					Paths: chioas.Paths{
						paths.UuidPath: {
							PathParams: chioas.PathParams{
								"id": {
									Description: "Category ID",
									Schema: &chioas.Schema{
										Type:   "string",
										Format: "uuid",
									},
								},
							},
							Methods: chioas.Methods{
								http.MethodGet: {
									Handler:     (*api).GetCategory,
									Description: "Get category",
									Responses: chioas.Responses{
										http.StatusOK: {
											Schema: (&chioas.Schema{}).MustFrom(models.Category{}),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	Components: &chioas.Components{
		SecuritySchemes: chioas.SecuritySchemes{
			{
				Name:        "ApiKey",
				Description: "Authorize using API key",
				In:          "header",
				Type:        "apiKey",
				ParamName:   ApiKeyHdr,
			},
		},
	},
	Security: chioas.SecuritySchemes{
		{Name: "ApiKey"},
	},
}
