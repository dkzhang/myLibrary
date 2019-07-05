package myUtils

import "strings"

func TrimBlank(str string) string {
	return strings.Trim(str, " \r\n\v\f")
}
