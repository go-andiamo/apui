package themes

var RootTheme = Theme{
	root: true,
	Body: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Header: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#464748",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Navigation: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#aaa",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Footer: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "#eee",
		BorderColor:     "#aaa",
		FontFamily:      "sans-serif",
		FontSize:        "75%",
	},
	Main: ThemeItem{
		TextColor:       "black",
		BackgroundColor: "white",
		BorderColor:     "#eee",
		FontFamily:      "sans-serif",
		FontSize:        "initial",
	},
	Json: JsonTheme{
		TextColor:       "black",
		BackgroundColor: "#eee",
		BorderColor:     "#ddd",
		FontFamily:      "monospace",
		FontSize:        "90%",
		CollapsedMarker: Coloring{
			TextColor:       "black",
			BackgroundColor: "#aaa",
		},
	},
	Methods: MethodsTheme{
		TextColor:       "white",
		BackgroundColor: "#333",
		BorderColor:     "black",
		FontFamily:      "sans-serif",
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
}
