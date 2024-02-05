package store

import (
	"strings"
)

func trim(value string) string {
	return strings.Trim(value, " \n\r")
}
