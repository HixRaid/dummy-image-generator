package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/hixraid/dummy-image/pkg/data"
)

const (
	oneNumberSizeRegexpPattern = "^\\d*$"
	twoNumberSizeRegexpPattern = "^\\d*x\\d*$"
	sepTwoNumberSize           = "x"
)

var resolutions = map[string]data.ImageResolution{
	"qvga":  data.QVGA,
	"sif":   data.SIF,
	"cif":   data.CIF,
	"wqvga": data.WQVGA,
	"hvga":  data.HVGA,
	"nhd":   data.NHD,
	"vga":   data.VGA,
	"wvga":  data.WVGA,
	"svga":  data.SVGA,
	"fwvga": data.FWVGA,
	"qhd":   data.QHD,
	"wsvga": data.WSVGA,
	"xga":   data.XGA,
	"wxvga": data.WXVGA,
	"hd":    data.HD,
	"wxga":  data.WXGA,
	"sxga":  data.SXGA,
	"xjxga": data.XJXGA,
	"wsxga": data.WSXGA,
	"uxga":  data.UXGA,
	"fhd":   data.FHD,
	"wuxga": data.WUXGA,
}

var ErrInvalidSize = errors.New("invalid size")
var defaultSize = [2]int{0, 0}

func ParseSize(s string) (data.ImageResolution, error) {
	s = strings.ToLower(s)

	if ok, _ := regexp.MatchString(oneNumberSizeRegexpPattern, s); ok {
		number, _ := strconv.ParseInt(s, 10, 64)

		var intNumber = int(number)

		return [2]int{intNumber, intNumber}, nil

	} else if ok, _ := regexp.MatchString(twoNumberSizeRegexpPattern, s); ok {
		s := strings.Split(s, sepTwoNumberSize)

		x, _ := strconv.ParseInt(s[0], 10, 64)
		y, _ := strconv.ParseInt(s[1], 10, 64)

		return [2]int{int(x), int(y)}, nil

	} else if resolution, ok := resolutions[s]; ok {
		return resolution, nil
	}

	return defaultSize, ErrInvalidSize
}
