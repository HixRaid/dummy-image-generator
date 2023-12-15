package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/hixraid/dummy-image/pkg/data"
)

const (
	oneNumberSizeRegexpPattern = "^\\d{1,4}$"
	twoNumberSizeRegexpPattern = "^\\d*x\\d*$"
	sepTwoNumberSize           = "x"
)

var resolutions = map[string]data.ImageResolution{
	"QVGA":  data.QVGA,
	"SIF":   data.SIF,
	"CIF":   data.CIF,
	"WQVGA": data.WQVGA,
	"HVGA":  data.HVGA,
	"NHD":   data.NHD,
	"VGA":   data.VGA,
	"WVGA":  data.WVGA,
	"SVGA":  data.SVGA,
	"FWVGA": data.FWVGA,
	"QHD":   data.QHD,
	"WSVGA": data.WSVGA,
	"XGA":   data.XGA,
	"WXVGA": data.WXVGA,
	"HD":    data.HD,
	"WXGA":  data.WXGA,
	"SXGA":  data.SXGA,
	"XJXGA": data.XJXGA,
	"WSXGA": data.WSXGA,
	"UXGA":  data.UXGA,
	"FHD":   data.FHD,
	"WUXGA": data.WUXGA,
}

func ParseSize(s string) (size data.ImageResolution) {
	if ok, _ := regexp.MatchString(oneNumberSizeRegexpPattern, s); ok {
		number, _ := strconv.ParseInt(s, 10, 64)

		var intNumber = int(number)

		size = [2]int{intNumber, intNumber}

	} else if ok, _ := regexp.MatchString(twoNumberSizeRegexpPattern, s); ok {
		s := strings.Split(s, sepTwoNumberSize)

		x, _ := strconv.ParseInt(s[0], 10, 64)
		y, _ := strconv.ParseInt(s[1], 10, 64)

		size = [2]int{int(x), int(y)}
	} else {
		size, _ = ParseResolution(s)
	}

	return
}

func ParseResolution(s string) (data.ImageResolution, error) {
	if resolution, ok := resolutions[s]; ok {
		return resolution, nil
	}

	return [2]int{0, 0}, errors.New("invalid image resolution")
}
