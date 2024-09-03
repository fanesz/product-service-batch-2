package helper

import (
	"fmt"
	"strings"
)

func Placeholders(n int, startsFrom int) string {
	if n == 0 {
		return ""
	}

	placeholders := make([]string, n)
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+startsFrom)
	}

	return strings.Join(placeholders, ", ")
}
