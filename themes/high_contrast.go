package themes

var HighContrast = Theme{
	Name: "High Contrast",
	Body: ThemeItem{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		FontSize:        "100%",
		LinkTextColor:   "#FFD700",
	},
	Header: ThemeItem{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		LinkTextColor:   "#FFD700",
		Opened: Coloring{
			TextColor:       "#000000",
			BackgroundColor: "#FFD700",
		},
	},
	Navigation: ThemeItem{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		LinkTextColor:   "#FFD700",
		Opened: Coloring{
			TextColor:       "#000000",
			BackgroundColor: "#FFD700",
			BorderColor:     "#000000",
		},
	},
	Footer: ThemeItem{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		FontSize:        "90%",
		LinkTextColor:   "#FFD700",
	},
	Main: ThemeItem{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		LinkTextColor:   "#FFD700",
	},
	Json: JsonTheme{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#111111",
		BorderColor:     "#FFFFFF",
		FontFamily:      "JetBrains Mono",
		FontSize:        "95%",
		CollapsedMarker: Coloring{
			TextColor:       "#000000",
			BackgroundColor: "#FFD700",
		},
	},
	Methods: MethodsTheme{
		TextColor:       "#FFFFFF",
		BackgroundColor: "#000000",
		BorderColor:     "#FFFFFF",
		FontFamily:      "Atkinson Hyperlegible",
		Get: Coloring{
			TextColor:       "#FFFFFF",
			BackgroundColor: "#00A651", // green
			BorderColor:     "transparent",
		},
		Delete: Coloring{
			TextColor:       "#FFFFFF",
			BackgroundColor: "#D50000", // red
			BorderColor:     "transparent",
		},
		Put: Coloring{
			TextColor:       "#000000",
			BackgroundColor: "#FFD54F", // amber
			BorderColor:     "transparent",
		},
		Post: Coloring{
			TextColor:       "#FFFFFF",
			BackgroundColor: "#2962FF", // blue
			BorderColor:     "transparent",
		},
		Patch: Coloring{
			TextColor:       "#FFFFFF",
			BackgroundColor: "#7B1FA2", // purple
			BorderColor:     "transparent",
		},
		Options: Coloring{
			TextColor:       "#000000",
			BackgroundColor: "#FFFFFF",
			BorderColor:     "transparent",
		},
	},
	Links: []Link{
		{
			Href: "https://fonts.googleapis.com/css2?family=Atkinson+Hyperlegible:wght@400;700&family=JetBrains+Mono:wght@400;600&display=swap",
		},
	},
}
