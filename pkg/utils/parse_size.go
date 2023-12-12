package utils

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

func ParseSize(s string) (size [2]int) {
	if ok, _ := regexp.MatchString(oneNumberSizeRegexpPattern, s); ok {
		number, _ := strconv.ParseInt(s, 10, 64)

		var intNumber = int(number)

		size = [2]int{intNumber, intNumber}

	} else if ok, _ := regexp.MatchString(twoNumberSizeRegexpPattern, s); ok {
		s := strings.Split(s, sepTwoNumberSize)

		x, _ := strconv.ParseInt(s[0], 10, 64)
		y, _ := strconv.ParseInt(s[1], 10, 64)

		size = [2]int{int(x), int(y)}
	}

	return
}
