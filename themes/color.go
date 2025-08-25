package themes

import (
	"github.com/go-andiamo/apui/internal/css"
	"strconv"
	"strings"
)

// CssColor is used for colors in themes
//
// it doesn't do any validating - but does normalize '#' and named colors to rgb()/rgba()
type CssColor string

func hex(b byte) int {
	switch {
	case b >= '0' && b <= '9':
		return int(b - '0')
	case b >= 'a' && b <= 'f':
		return int(b-'a') + 10
	case b >= 'A' && b <= 'F':
		return int(b-'A') + 10
	default:
		return -1
	}
}

func (c CssColor) String() string {
	result := string(c)
	isHash := len(result) > 0 && result[0] == '#'
	l := len(result)
	switch {
	case isHash && l == 4: // #rgb
		if r, g, b := hex(result[1]), hex(result[2]), hex(result[3]); r >= 0 && g >= 0 && b >= 0 {
			return buildRGB((r<<4)|r, (g<<4)|g, (b<<4)|b)
		}
	case isHash && l == 5: // #rgba
		if r, g, b, a := hex(result[1]), hex(result[2]), hex(result[3]), hex(result[4]); r >= 0 && g >= 0 && b >= 0 && a >= 0 {
			return buildRGBA((r<<4)|r, (g<<4)|g, (b<<4)|b, (a<<4)|a)
		}
	case isHash && l == 7: // #rrggbb
		if rh, rl, gh, gl, bh, bl := hex(result[1]), hex(result[2]), hex(result[3]), hex(result[4]), hex(result[5]), hex(result[6]); rh >= 0 && rl >= 0 && gh >= 0 && gl >= 0 && bh >= 0 && bl >= 0 {
			return buildRGB((rh<<4)|rl, (gh<<4)|gl, (bh<<4)|bl)
		}
	case isHash && l == 9: // #rrggbbaa
		if rh, rl, gh, gl, bh, bl, ah, al := hex(result[1]), hex(result[2]), hex(result[3]), hex(result[4]), hex(result[5]), hex(result[6]), hex(result[7]), hex(result[8]); rh >= 0 && rl >= 0 && gh >= 0 && gl >= 0 && bh >= 0 && bl >= 0 && ah >= 0 && al >= 0 {
			return buildRGBA((rh<<4)|rl, (gh<<4)|gl, (bh<<4)|bl, (ah<<4)|al)
		}
	default:
		if nc, ok := css.NamedColors[strings.ToLower(result)]; ok {
			return nc
		}
	}
	return result
}

func buildRGB(r, g, b int) string {
	buf := make([]byte, 0, 16)
	buf = append(buf, "rgb("...)
	buf = strconv.AppendInt(buf, int64(r), 10)
	buf = append(buf, ',')
	buf = strconv.AppendInt(buf, int64(g), 10)
	buf = append(buf, ',')
	buf = strconv.AppendInt(buf, int64(b), 10)
	buf = append(buf, ')')
	return string(buf)
}

func buildRGBA(r, g, b, a int) string {
	buf := make([]byte, 0, 23)
	buf = append(buf, "rgba("...)
	buf = strconv.AppendInt(buf, int64(r), 10)
	buf = append(buf, ',')
	buf = strconv.AppendInt(buf, int64(g), 10)
	buf = append(buf, ',')
	buf = strconv.AppendInt(buf, int64(b), 10)
	buf = append(buf, ',')
	buf = strconv.AppendFloat(buf, float64(a)/255.0, 'f', 3, 64)
	buf = append(buf, ')')
	return string(buf)
}
