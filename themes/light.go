package themes

var Light = Theme{
	Name: "Light",
	Body: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "Open Sans",
	},
	Header: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#464748",
		BorderColor:     "#eee",
		FontFamily:      "Open Sans",
	},
	Navigation: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#aaa",
		FontFamily:      "Open Sans",
		Opened: Coloring{
			TextColor:       "black",
			BackgroundColor: "#47AFE8",
			BorderColor:     "#aaa",
		},
	},
	Footer: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "#eee",
		BorderColor:     "#aaa",
		FontFamily:      "Open Sans",
		FontSize:        "75%",
	},
	Main: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "Open Sans",
	},
	Json: JsonTheme{
		TextColor:       "black",
		BackgroundColor: "#eee",
		BorderColor:     "#ddd",
		FontFamily:      "Google Sans Code",
		FontSize:        "85%",
		CollapsedMarker: Coloring{
			TextColor:       "black",
			BackgroundColor: "#aaa",
		},
	},
	Methods: MethodsTheme{
		TextColor:       "white",
		BackgroundColor: "#333",
		BorderColor:     "black",
		FontFamily:      "Open Sans",
		Get: Coloring{
			TextColor:       "black",
			BackgroundColor: "#47AFE8",
			BorderColor:     "transparent",
		},
		Delete: Coloring{
			TextColor:       "black",
			BackgroundColor: "#F06560",
			BorderColor:     "transparent",
		},
		Put: Coloring{
			TextColor:       "black",
			BackgroundColor: "#FF9900",
			BorderColor:     "transparent",
		},
		Post: Coloring{
			TextColor:       "black",
			BackgroundColor: "#690",
			BorderColor:     "transparent",
		},
		Patch: Coloring{
			TextColor:       "black",
			BackgroundColor: "#827717",
			BorderColor:     "transparent",
		},
		Options: Coloring{
			TextColor:       "black",
			BackgroundColor: "#ddd",
			BorderColor:     "transparent",
		},
	},
	Links: []Link{
		{
			Href: "https://fonts.googleapis.com/css2?family=Google+Sans+Code:ital,wght@0,300..800;1,300..800&family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap",
		},
	},
}
