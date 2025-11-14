package utils

import (
	"regexp"
)

var digitsNum = regexp.MustCompile(`^\d+$`)

func IsNumber(s string) bool {
	return digitsNum.MatchString(s)
}
