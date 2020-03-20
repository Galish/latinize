package latinize

import "strings"

func String(input string) string {
	r := strings.NewReplacer(characters...)
	return r.Replace(input)
}
