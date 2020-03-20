package latinize

import (
	"strings"
)

func String(str string) string {
	var result = []string{}

	for _, letter := range str {
		val, ok := charactersMap[string(letter)]
		if !ok {
			result = append(result, string(letter))
			continue
		}

		result = append(result, val)
	}

	return strings.Join(result, "")
}
