package latinize

import (
	"strings"
	"unicode"
)

func String(str string) string {
	var result = []string{}

	for _, strRune := range str {
		symbol := string(strRune)

		if !unicode.IsLetter(strRune) {
			result = append(result, symbol)
			continue
		}

		val, ok := charactersMap[symbol]
		if ok {
			result = append(result, val)
			continue
		}

		result = append(result, symbol)
	}

	return strings.Join(result, "")
}
