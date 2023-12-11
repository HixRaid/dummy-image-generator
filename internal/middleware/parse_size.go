package middleware

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	oneNumberSizeRegexpPattern = "^\\d{1,4}$"
	twoNumberSizeRegexpPattern = "^\\d*x\\d*$"
	sepTwoNumberSize           = "x"
)

func parseSize(path []string) (size [2]int) {
	sizeStr := path[0]

	if ok, _ := regexp.MatchString(oneNumberSizeRegexpPattern, sizeStr); ok {
		number, _ := strconv.ParseInt(sizeStr, 10, 64)

		var intNumber = int(number)

		size = [2]int{intNumber, intNumber}

	} else if ok, _ := regexp.MatchString(twoNumberSizeRegexpPattern, sizeStr); ok {
		s := strings.Split(sizeStr, sepTwoNumberSize)

		x, _ := strconv.ParseInt(s[0], 10, 64)
		y, _ := strconv.ParseInt(s[1], 10, 64)

		size = [2]int{int(x), int(y)}
	}

	return
}
