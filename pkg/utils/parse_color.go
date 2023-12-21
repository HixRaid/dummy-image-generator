package utils

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strings"

	"github.com/hixraid/dummy-image/pkg/data"
)

const (
	hexColorRegexpPattern = "^(?:[0-9a-fA-F]{1,3}|[0-9a-fA-F]{6})$"
	rgbColorRegexpPattern = "^(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5]),(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5]),(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])$"

	hexFactor = 17

	hexFormat         = "%02x%02x%02x"
	hexShortFormat    = "%1x%1x%1x"
	hexShortTwoFormat = "%02x"
	hexShortOneFormat = "%1x"
	rgbFormat         = "%d,%d,%d"
)

var colors = map[string]color.Color{
	"indianred":            data.IndianRed,
	"lightcoral":           data.LightCoral,
	"salmon":               data.Salmon,
	"darksalmon":           data.DarkSalmon,
	"lightsalmon":          data.LightSalmon,
	"crimson":              data.Crimson,
	"red":                  data.Red,
	"firebrick":            data.FireBrick,
	"darkred":              data.DarkRed,
	"pink":                 data.Pink,
	"lightpink":            data.LightPink,
	"hotpink":              data.HotPink,
	"deeppink":             data.DeepPink,
	"mediumvioletred":      data.MediumVioletRed,
	"palevioletred":        data.PaleVioletRed,
	"coral":                data.Coral,
	"tomato":               data.Tomato,
	"orangered":            data.OrangeRed,
	"darkorange":           data.DarkOrange,
	"orange":               data.Orange,
	"gold":                 data.Gold,
	"yellow":               data.Yellow,
	"lightyellow":          data.LightYellow,
	"lemonchiffon":         data.LemonChiffon,
	"lightgoldenrodyellow": data.LightGoldenrodYellow,
	"papayawhip":           data.PapayaWhip,
	"moccasin":             data.Moccasin,
	"peachpuff":            data.PeachPuff,
	"palegoldenrod":        data.PaleGoldenrod,
	"khaki":                data.Khaki,
	"darkkhaki":            data.DarkKhaki,
	"lavender":             data.Lavender,
	"thistle":              data.Thistle,
	"plum":                 data.Plum,
	"violet":               data.Violet,
	"orchid":               data.Orchid,
	"fuchsia":              data.Fuchsia,
	"magenta":              data.Magenta,
	"mediumorchid":         data.MediumOrchid,
	"mediumpurple":         data.MediumPurple,
	"rebeccapurple":        data.RebeccaPurple,
	"blueviolet":           data.BlueViolet,
	"darkviolet":           data.DarkViolet,
	"darkorchid":           data.DarkOrchid,
	"darkmagenta":          data.DarkMagenta,
	"purple":               data.Purple,
	"indigo":               data.Indigo,
	"slateblue":            data.SlateBlue,
	"darkslateblue":        data.DarkSlateBlue,
	"mediumslateblue":      data.MediumSlateBlue,
	"greenyellow":          data.GreenYellow,
	"chartreuse":           data.Chartreuse,
	"lawngreen":            data.LawnGreen,
	"lime":                 data.Lime,
	"limegreen":            data.LimeGreen,
	"palegreen":            data.PaleGreen,
	"lightgreen":           data.LightGreen,
	"mediumspringgreen":    data.MediumSpringGreen,
	"springgreen":          data.SpringGreen,
	"mediumseagreen":       data.MediumSeaGreen,
	"seagreen":             data.SeaGreen,
	"forestgreen":          data.ForestGreen,
	"green":                data.Green,
	"darkgreen":            data.DarkGreen,
	"yellowgreen":          data.YellowGreen,
	"olivedrab":            data.OliveDrab,
	"olive":                data.Olive,
	"darkolivegreen":       data.DarkOliveGreen,
	"mediumaquamarine":     data.MediumAquamarine,
	"darkseagreen":         data.DarkSeaGreen,
	"lightseagreen":        data.LightSeaGreen,
	"darkcyan":             data.DarkCyan,
	"teal":                 data.Teal,
	"aqua":                 data.Aqua,
	"cyan":                 data.Cyan,
	"lightcyan":            data.LightCyan,
	"paleturquoise":        data.PaleTurquoise,
	"aquamarine":           data.Aquamarine,
	"turquoise":            data.Turquoise,
	"mediumturquoise":      data.MediumTurquoise,
	"darkturquoise":        data.DarkTurquoise,
	"cadetblue":            data.CadetBlue,
	"steelblue":            data.SteelBlue,
	"lightsteelblue":       data.LightSteelBlue,
	"powderblue":           data.PowderBlue,
	"lightblue":            data.LightBlue,
	"skyblue":              data.SkyBlue,
	"lightskyblue":         data.LightSkyBlue,
	"deepskyblue":          data.DeepSkyBlue,
	"dodgerblue":           data.DodgerBlue,
	"cornflowerblue":       data.CornflowerBlue,
	"royalblue":            data.RoyalBlue,
	"blue":                 data.Blue,
	"mediumblue":           data.MediumBlue,
	"darkblue":             data.DarkBlue,
	"navy":                 data.Navy,
	"midnightblue":         data.MidnightBlue,
	"cornsilk":             data.Cornsilk,
	"blanchedalmond":       data.BlanchedAlmond,
	"bisque":               data.Bisque,
	"navajowhite":          data.NavajoWhite,
	"wheat":                data.Wheat,
	"burlywood":            data.BurlyWood,
	"tan":                  data.Tan,
	"rosybrown":            data.RosyBrown,
	"sandybrown":           data.SandyBrown,
	"goldenrod":            data.Goldenrod,
	"darkgoldenrod":        data.DarkGoldenrod,
	"peru":                 data.Peru,
	"chocolate":            data.Chocolate,
	"saddlebrown":          data.SaddleBrown,
	"sienna":               data.Sienna,
	"brown":                data.Brown,
	"maroon":               data.Maroon,
	"white":                data.White,
	"snow":                 data.Snow,
	"honeydew":             data.HoneyDew,
	"mintcream":            data.MintCream,
	"azure":                data.Azure,
	"aliceblue":            data.AliceBlue,
	"ghostwhite":           data.GhostWhite,
	"whitesmoke":           data.WhiteSmoke,
	"seashell":             data.SeaShell,
	"beige":                data.Beige,
	"oldlace":              data.OldLace,
	"floralwhite":          data.FloralWhite,
	"ivory":                data.Ivory,
	"antiquewhite":         data.AntiqueWhite,
	"linen":                data.Linen,
	"lavenderblush":        data.LavenderBlush,
	"mistyrose":            data.MistyRose,
	"gainsboro":            data.Gainsboro,
	"lightgray":            data.LightGray,
	"silver":               data.Silver,
	"darkgray":             data.DarkGray,
	"gray":                 data.Gray,
	"dimgray":              data.DimGray,
	"lightslategray":       data.LightSlateGray,
	"slategray":            data.SlateGray,
	"darkslategray":        data.DarkSlateGray,
	"black":                data.Black,
}

var ErrInvalidColor = errors.New("invalid color")

func ParseColor(s string) (color.Color, error) {
	s = strings.ToLower(s)
	if ok, _ := regexp.MatchString(hexColorRegexpPattern, s); ok {
		return parseHEX(s), nil
	}

	if ok, _ := regexp.MatchString(rgbColorRegexpPattern, s); ok {
		return parseRGB(s), nil
	}

	if color, ok := colors[s]; ok {
		return color, nil
	}

	return nil, ErrInvalidColor
}

func parseHEX(s string) color.Color {
	s = strings.ToLower(s)

	var r, g, b uint8

	switch len(s) {
	case 1:
		fmt.Sscanf(s, hexShortOneFormat, &r)
		r *= hexFactor
		g = r
		b = r
	case 2:
		fmt.Sscanf(s, hexShortTwoFormat, &r)
		g = r
		b = r
	case 3:
		fmt.Sscanf(s, hexShortFormat, &r, &g, &b)
		r *= hexFactor
		g *= hexFactor
		b *= hexFactor
	case 6:
		fmt.Sscanf(s, hexFormat, &r, &g, &b)
	}

	return color.RGBA{r, g, b, 255}
}

func parseRGB(s string) color.Color {
	var r, g, b uint8

	fmt.Sscanf(s, rgbFormat, &r, &g, &b)

	return color.RGBA{r, g, b, 255}
}
