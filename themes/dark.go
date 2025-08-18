package themes

var Dark = Theme{
	Name: "Dark",
	Body: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#1e1e1e",
		BorderColor:     "#444",
		FontFamily:      "Open Sans",
		LinkTextColor:   "lightblue",
	},
	Header: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#2d2d2d",
		BorderColor:     "#444",
		FontFamily:      "Open Sans",
		LinkTextColor:   "lightblue",
		Opened: Coloring{
			TextColor:       "white",
			BackgroundColor: "#47AFE8",
		},
	},
	Navigation: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#1e1e1e",
		BorderColor:     "#555",
		FontFamily:      "Open Sans",
		LinkTextColor:   "lightblue",
		Opened: Coloring{
			TextColor:       "white",
			BackgroundColor: "#47AFE8",
			BorderColor:     "#555",
		},
		Dropdown: Coloring{
			TextColor:       "white",
			BackgroundColor: "#222",
			BorderColor:     "#888",
		},
	},
	Footer: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#2d2d2d",
		BorderColor:     "#555",
		FontFamily:      "Open Sans",
		FontSize:        "75%",
		LinkTextColor:   "lightblue",
	},
	Main: ThemeItem{
		TextColor:       "white",
		BackgroundColor: "#1e1e1e",
		BorderColor:     "#666",
		FontFamily:      "Open Sans",
		LinkTextColor:   "lightblue",
	},
	Json: JsonTheme{
		TextColor:       "white",
		BackgroundColor: "#2d2d2d",
		BorderColor:     "#666",
		FontFamily:      "Google Sans Code",
		FontSize:        "85%",
		CollapsedMarker: Coloring{
			TextColor:       "white",
			BackgroundColor: "#555",
		},
	},
	Methods: MethodsTheme{
		TextColor:       "white",
		BackgroundColor: "#333",
		BorderColor:     "#111",
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
			BackgroundColor: "#aaa",
			BorderColor:     "transparent",
		},
	},
	Links: []Link{
		{
			Href: "https://fonts.googleapis.com/css2?family=Google+Sans+Code:ital,wght@0,300..800;1,300..800&family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap",
		},
	},
}
