package themes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor_String(t *testing.T) {
	testCases := []struct {
		color  CssColor
		expect string
	}{
		{},
		{
			color:  "black",
			expect: "rgb(0,0,0)",
		},
		{
			color:  "Black",
			expect: "rgb(0,0,0)",
		},
		{
			color:  "white",
			expect: "rgb(255,255,255)",
		},
		{
			color:  "WhItE",
			expect: "rgb(255,255,255)",
		},
		{
			color:  "lightblue",
			expect: "rgb(173,216,230)",
		},
		{
			color:  "LightBlue",
			expect: "rgb(173,216,230)",
		},
		{
			color:  "#000",
			expect: "rgb(0,0,0)",
		},
		{
			color:  "#fff",
			expect: "rgb(255,255,255)",
		},
		{
			color:  "#617",
			expect: "rgb(102,17,119)",
		},
		{
			color:  "#0001",
			expect: "rgba(0,0,0,0.067)",
		},
		{
			color:  "#fff1",
			expect: "rgba(255,255,255,0.067)",
		},
		{
			color:  "#ffff",
			expect: "rgba(255,255,255,1.000)",
		},
		{
			color:  "#000000",
			expect: "rgb(0,0,0)",
		},
		{
			color:  "#ffffff",
			expect: "rgb(255,255,255)",
		},
		{
			color:  "#00000010",
			expect: "rgba(0,0,0,0.063)",
		},
		{
			color:  "#ffffff10",
			expect: "rgba(255,255,255,0.063)",
		},
		{
			color:  "#AAA",
			expect: "rgb(170,170,170)",
		},
		{
			color:  "#AAAB",
			expect: "rgba(170,170,170,0.733)",
		},
		{
			color:  "#AAAAAAB1",
			expect: "rgba(170,170,170,0.694)",
		},
		{
			color:  "#xyz",
			expect: "#xyz",
		},
		{
			color:  "#xyzz",
			expect: "#xyzz",
		},
		{
			color:  "#xxyyzz",
			expect: "#xxyyzz",
		},
		{
			color:  "#xxyyzzzz",
			expect: "#xxyyzzzz",
		},
		{
			color:  "unknown",
			expect: "unknown",
		},
	}
	for _, tc := range testCases {
		t.Run(string(tc.color), func(t *testing.T) {
			require.Equal(t, tc.expect, tc.color.String())
		})
	}
}
